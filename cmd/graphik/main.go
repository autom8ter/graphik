package main

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/autom8ter/machine"
	"github.com/graphikDB/graphik/database"
	"github.com/graphikDB/graphik/discover/k8s"
	"github.com/graphikDB/graphik/gen/grpc/go"
	"github.com/graphikDB/graphik/gql"
	"github.com/graphikDB/graphik/helpers"
	"github.com/graphikDB/graphik/logger"
	"github.com/graphikDB/raft"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/cors"
	"github.com/soheilhy/cmux"
	"github.com/spf13/pflag"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/http/pprof"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

var global = &apipb.Flags{}

func init() {
	godotenv.Load()
	pflag.CommandLine.BoolVar(&global.Debug, "debug", helpers.BoolEnvOr("GRAPHIK_DEBUG", false), "enable debug logs (env: GRAPHIK_DEBUG)")
	pflag.CommandLine.StringVar(&global.RaftSecret, "raft-secret", os.Getenv("GRAPHIK_RAFT_SECRET"), "raft cluster secret (so only authorized nodes may join cluster) (env: GRAPHIK_RAFT_SECRET)")
	pflag.CommandLine.StringVar(&global.JoinRaft, "join-raft", os.Getenv("GRAPHIK_JOIN_RAFT"), "join raft cluster at target address (env: GRAPHIK_JOIN_RAFT)")
	pflag.CommandLine.StringVar(&global.RaftPeerId, "raft-peer-id", os.Getenv("GRAPHIK_RAFT_PEER_ID"), "raft peer ID - one will be generated if not set (env: GRAPHIK_RAFT_PEER_ID)")
	pflag.CommandLine.Int64Var(&global.ListenPort, "listen-port", int64(helpers.IntEnvOr("GRAPHIK_LISTEN_PORT", 7820)), "serve gRPC & graphQL on this port (env: GRAPHIK_LISTEN_PORT)")
	pflag.CommandLine.StringVar(&global.StoragePath, "storage", helpers.EnvOr("GRAPHIK_STORAGE_PATH", "/tmp/graphik"), "persistant storage path (env: GRAPHIK_STORAGE_PATH)")
	pflag.CommandLine.StringVar(&global.OpenIdDiscovery, "open-id", helpers.EnvOr("GRAPHIK_OPEN_ID", ""), "open id connect discovery uri ex: https://accounts.google.com/.well-known/openid-configuration (env: GRAPHIK_OPEN_ID) (required)")
	pflag.CommandLine.BoolVar(&global.Metrics, "metrics", helpers.BoolEnvOr("GRAPHIK_METRICS", true), "enable prometheus & pprof metrics (emv: GRAPHIK_METRICS = true)")
	pflag.CommandLine.StringSliceVar(&global.AllowHeaders, "allow-headers", helpers.StringSliceEnvOr("GRAPHIK_ALLOW_HEADERS", []string{"*"}), "cors allow headers (env: GRAPHIK_ALLOW_HEADERS)")
	pflag.CommandLine.StringSliceVar(&global.AllowOrigins, "allow-origins", helpers.StringSliceEnvOr("GRAPHIK_ALLOW_ORIGINS", []string{"*"}), "cors allow origins (env: GRAPHIK_ALLOW_ORIGINS)")
	pflag.CommandLine.StringSliceVar(&global.AllowMethods, "allow-methods", helpers.StringSliceEnvOr("GRAPHIK_ALLOW_METHODS", []string{"HEAD", "GET", "POST", "PUT", "PATCH", "DELETE"}), "cors allow methods (env: GRAPHIK_ALLOW_METHODS)")
	pflag.CommandLine.StringSliceVar(&global.RootUsers, "root-users", helpers.StringSliceEnvOr("GRAPHIK_ROOT_USERS", nil), "a list of email addresses that bypass registered authorizers (env: GRAPHIK_ROOT_USERS)  (required)")
	pflag.CommandLine.StringVar(&global.TlsCert, "tls-cert", helpers.EnvOr("GRAPHIK_TLS_CERT", ""), "path to tls certificate (env: GRAPHIK_TLS_CERT)")
	pflag.CommandLine.StringVar(&global.TlsKey, "tls-key", helpers.EnvOr("GRAPHIK_TLS_KEY", ""), "path to tls key (env: GRAPHIK_TLS_KEY)")
	pflag.CommandLine.BoolVar(&global.RequireRequestAuthorizers, "require-request-authorizers", helpers.BoolEnvOr("GRAPHIK_REQUIRE_REQUEST_AUTHORIZERS", false), "require request authorizers for all methods/endpoints (env: GRAPHIK_REQUIRE_REQUEST_AUTHORIZERS)")
	pflag.CommandLine.BoolVar(&global.RequireResponseAuthorizers, "require-response-authorizers", helpers.BoolEnvOr("GRAPHIK_REQUIRE_RESPONSE_AUTHORIZERS", false), "require request authorizers for all methods/endpoints (env: GRAPHIK_REQUIRE_RESPONSE_AUTHORIZERS)")
	pflag.CommandLine.StringVar(&global.PlaygroundClientId, "playground-client-id", helpers.EnvOr("GRAPHIK_PLAYGROUND_CLIENT_ID", ""), "playground oauth client id (env: GRAPHIK_PLAYGROUND_CLIENT_ID)")
	pflag.CommandLine.StringVar(&global.PlaygroundClientSecret, "playground-client-secret", helpers.EnvOr("GRAPHIK_PLAYGROUND_CLIENT_SECRET", ""), "playground oauth client secret (env: GRAPHIK_PLAYGROUND_CLIENT_SECRET)")
	pflag.CommandLine.StringVar(&global.PlaygroundRedirect, "playground-redirect", helpers.EnvOr("GRAPHIK_PLAYGROUND_REDIRECT", ""), "playground oauth redirect (env: GRAPHIK_PLAYGROUND_REDIRECT)")
	pflag.CommandLine.StringVar(&global.Environment, "environment", helpers.EnvOr("GRAPHIK_ENVIRONMENT", ""), "deployment environment (k8s) (env: GRAPHIK_ENVIRONMENT)")
	pflag.Parse()
}

const (
	k8sEnv    = "k8s"
	leaderPod = "graphik-0"
)

func main() {
	run(context.Background(), global)
}

func run(ctx context.Context, cfg *apipb.Flags) {
	lgger := logger.New(cfg.Debug)
	if cfg.OpenIdDiscovery == "" {
		lgger.Error("empty open-id connect discovery --open-id", zap.String("usage", pflag.CommandLine.Lookup("open-id").Usage))
		return
	}
	if len(cfg.GetRootUsers()) == 0 {
		lgger.Error("zero root users", zap.String("usage", pflag.CommandLine.Lookup("root-users").Usage))
		return
	}
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(interrupt)
	var localRaftAddr = fmt.Sprintf("localhost:%v", global.ListenPort+1)
	if global.Environment != "" {
		switch global.Environment {
		case k8sEnv:
			var (
				podname   = os.Getenv("POD_NAME")
				namespace = os.Getenv("POD_NAMESPACE")
				podIp     = os.Getenv("POD_IP")
			)
			if podname == "" {
				lgger.Error("expected POD_NAME environmental variable set in k8s environment")
				return
			}
			if podIp == "" {
				lgger.Error("expected POD_IP environmental variable set in k8s environment")
				return
			}
			prvider, err := k8s.NewInClusterProvider(namespace)
			if err != nil {
				lgger.Error("failed to get incluster k8s provider", zap.Error(err))
				return
			}
			pods, err := prvider.Pods(ctx)
			if err != nil {
				lgger.Error("failed to get pods", zap.Error(err))
				return
			}
			leaderIp := pods[leaderPod]
			lgger.Info("joining k8s cluster",
				zap.String("namespace", namespace),
				zap.String("podip", podIp),
				zap.String("podname", podname),
				zap.String("leaderIp", leaderIp),
				zap.Any("discovery", pods),
			)
			global.RaftPeerId = podname
			global.RaftAdvertise = fmt.Sprintf("%s:%v", podIp, global.ListenPort+1)
			if podname != leaderPod {
				global.JoinRaft = fmt.Sprintf("%s:%v", leaderIp, global.ListenPort)
				localRaftAddr = fmt.Sprintf("%s:%v", podIp, global.ListenPort+1)
			}
		default:
			lgger.Error("unsupported environment", zap.String("env", global.Environment))
			return
		}
	}
	m := machine.New(ctx)

	adminLis, err := net.Listen("tcp", fmt.Sprintf(":%v", global.ListenPort+1))
	if err != nil {
		lgger.Error("metrics server failure", zap.Error(err))
		return
	}
	adminMux := cmux.New(adminLis)
	defer adminLis.Close()
	raftLis := adminMux.Match(cmux.Any())
	lgger.Info("starting raft listener", zap.String("address", raftLis.Addr().String()))
	defer raftLis.Close()
	var metricServer *http.Server
	{
		router := http.NewServeMux()
		router.Handle("/metrics", promhttp.Handler())
		router.HandleFunc("/debug/pprof/", pprof.Index)
		router.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
		router.HandleFunc("/debug/pprof/profile", pprof.Profile)
		router.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
		router.HandleFunc("/debug/pprof/trace", pprof.Trace)
		metricServer = &http.Server{Handler: router}
	}

	m.Go(func(routine machine.Routine) {
		hmatcher := adminMux.Match(cmux.HTTP1())
		defer hmatcher.Close()
		lgger.Info("starting metrics/admin server", zap.String("address", hmatcher.Addr().String()))
		if err := metricServer.Serve(hmatcher); err != nil && err != http.ErrServerClosed {
			lgger.Error("metrics server failure", zap.Error(err))
		}
	})
	m.Go(func(routine machine.Routine) {
		if err := adminMux.Serve(); err != nil && !strings.Contains(err.Error(), "closed network connection") {
			lgger.Error("listener mux error", zap.Error(err))
		}
	})
	g, err := database.NewGraph(
		global.OpenIdDiscovery,
		database.WithRaftSecret(global.RaftSecret),
		database.WithLogger(lgger),
		database.WithMachine(m),
		database.WithRootUsers(global.RootUsers),
		database.WithRequireRequestAuthorizers(global.RequireRequestAuthorizers),
		database.WithRequireResponseAuthorizers(global.RequireResponseAuthorizers),
	)
	if err != nil {
		lgger.Error("failed to create graph", zap.Error(err))
		return
	}
	defer g.Close()
	var (
		lis       net.Listener
		advertise net.Addr
	)
	var tlsConfig *tls.Config
	if global.TlsCert != "" && global.TlsKey != "" {
		cer, err := tls.LoadX509KeyPair(global.TlsCert, global.TlsKey)
		if err != nil {
			log.Println(err)
			return
		}
		tlsConfig = &tls.Config{Certificates: []tls.Certificate{cer}}
	}
	if tlsConfig != nil {
		lis, err = tls.Listen("tcp", fmt.Sprintf(":%v", global.ListenPort), tlsConfig)
		if err != nil {
			lgger.Error("failed to create tls server listener", zap.Error(err))
			return
		}
	} else {
		lis, err = net.Listen("tcp", fmt.Sprintf(":%v", global.ListenPort))
		if err != nil {
			lgger.Error("failed to create server listener", zap.Error(err))
			return
		}
	}
	defer lis.Close()

	ropts := []raft.Opt{
		raft.WithIsLeader(global.JoinRaft == ""),
		raft.WithRaftDir(fmt.Sprintf("%s/raft", global.StoragePath)),
		raft.WithPeerID(global.RaftPeerId),
		raft.WithMaxPool(5),
		raft.WithRestoreSnapshotOnRestart(false),
	}
	if global.RaftAdvertise != "" {
		advertise, err = net.ResolveTCPAddr("tcp", global.RaftAdvertise)
		if err != nil {
			lgger.Error("failed to resolve raft advertise addr", zap.Error(err))
			return
		}
		ropts = append(ropts, raft.WithAdvertiseAddr(advertise))
	}

	var config *oauth2.Config
	if global.PlaygroundClientId != "" {
		resp, err := http.DefaultClient.Get(global.OpenIdDiscovery)
		if err != nil {
			lgger.Error("failed to get oidc", zap.Error(err))
			return
		}
		defer resp.Body.Close()
		var openID = map[string]interface{}{}
		bits, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			lgger.Error("failed to get oidc", zap.Error(err))
			return
		}
		if err := json.Unmarshal(bits, &openID); err != nil {
			lgger.Error("failed to get oidc", zap.Error(err))
			return
		}
		config = &oauth2.Config{
			ClientID:     global.PlaygroundClientId,
			ClientSecret: global.PlaygroundClientSecret,
			Endpoint: oauth2.Endpoint{
				AuthURL:  openID["authorization_endpoint"].(string),
				TokenURL: openID["token_endpoint"].(string),
			},
			RedirectURL: global.PlaygroundRedirect,
			Scopes:      []string{"openid", "email", "profile"},
		}
	}
	self := fmt.Sprintf("localhost:%v", global.ListenPort)
	conn, err := grpc.DialContext(ctx, self, grpc.WithInsecure())
	if err != nil {
		lgger.Error("failed to setup graphql endpoint", zap.Error(err))
		return
	}
	defer conn.Close()
	resolver := gql.NewResolver(apipb.NewDatabaseServiceClient(conn), cors.New(cors.Options{
		AllowedOrigins: global.AllowOrigins,
		AllowedMethods: global.AllowMethods,
		AllowedHeaders: global.AllowHeaders,
	}), config, lgger)
	mux := http.NewServeMux()
	mux.Handle("/", resolver.QueryHandler())

	if config != nil {
		mux.Handle("/playground", resolver.Playground())
		mux.Handle("/playground/callback", resolver.PlaygroundCallback("/playground"))
	}
	httpServer := &http.Server{
		Handler: mux,
	}
	{
		rft, err := raft.NewRaft(g.RaftFSM(), raftLis, ropts...)
		if err != nil {
			lgger.Error("failed to create raft", zap.Error(err))
			return
		}
		g.SetRaft(rft)
	}
	apiMux := cmux.New(lis)

	m.Go(func(routine machine.Routine) {
		hmatcher := apiMux.Match(cmux.HTTP1())
		defer hmatcher.Close()
		lgger.Info("starting http server",
			zap.String("address", hmatcher.Addr().String()),
		)
		if err := httpServer.Serve(hmatcher); err != nil && err != http.ErrServerClosed {
			lgger.Error("http server failure", zap.Error(err))
		}
	})
	gserver := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			grpc_prometheus.UnaryServerInterceptor,
			grpc_zap.UnaryServerInterceptor(lgger.Zap()),
			grpc_validator.UnaryServerInterceptor(),
			g.UnaryInterceptor(),
			grpc_recovery.UnaryServerInterceptor(),
		),
		grpc.ChainStreamInterceptor(
			grpc_prometheus.StreamServerInterceptor,
			grpc_zap.StreamServerInterceptor(lgger.Zap()),
			grpc_validator.StreamServerInterceptor(),
			g.StreamInterceptor(),
			grpc_recovery.StreamServerInterceptor(),
		),
	)
	apipb.RegisterDatabaseServiceServer(gserver, g)
	apipb.RegisterRaftServiceServer(gserver, g)
	reflection.Register(gserver)

	grpc_prometheus.Register(gserver)
	m.Go(func(routine machine.Routine) {
		gmatcher := apiMux.Match(cmux.HTTP2())
		defer gmatcher.Close()
		lgger.Info("starting gRPC server", zap.String("address", lis.Addr().String()))
		if err := gserver.Serve(gmatcher); err != nil && err != http.ErrServerClosed && !strings.Contains(err.Error(), "mux: listener closed") {
			lgger.Error("gRPC server failure", zap.Error(err))
		}
	})
	m.Go(func(routine machine.Routine) {
		if err := apiMux.Serve(); err != nil && !strings.Contains(err.Error(), "closed network connection") {
			lgger.Error("listener mux error", zap.Error(err))
		}
	})

	if global.JoinRaft != "" {
		lgger.Info("joining raft cluster",
			zap.String("joinAddr", global.JoinRaft),
			zap.String("localAddr", localRaftAddr),
		)
		if err := join(ctx, global.JoinRaft, localRaftAddr, g, lgger); err != nil {
			lgger.Error(err.Error())
		}
	}
	select {
	case <-interrupt:
		m.Cancel()
		break
	case <-ctx.Done():
		m.Cancel()
		break
	}

	lgger.Warn("shutdown signal received")
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownCancel()

	_ = httpServer.Shutdown(shutdownCtx)
	if metricServer != nil {
		_ = metricServer.Shutdown(shutdownCtx)
	}
	stopped := make(chan struct{})
	go func() {
		gserver.GracefulStop()
		close(stopped)
	}()

	t := time.NewTimer(10 * time.Second)
	select {
	case <-t.C:
		gserver.Stop()
	case <-stopped:
		t.Stop()
	}
	m.Wait()
	lgger.Info("shutdown successful")
}

func join(ctx context.Context, joinAddr string, localAddr string, g *database.Graph, lgger *logger.Logger) error {
	leaderConn, err := grpc.DialContext(ctx, joinAddr, grpc.WithInsecure())
	if err != nil {
		return errors.Wrap(err, "failed to join raft")
	}
	defer leaderConn.Close()
	rclient := apipb.NewRaftServiceClient(leaderConn)
	for x := 0; x < 5; x++ {
		_, err := rclient.JoinCluster(metadata.AppendToOutgoingContext(ctx, "x-graphik-raft-secret", global.RaftSecret), &apipb.Peer{
			NodeId: g.Raft().PeerID(),
			Addr:   localAddr,
		})
		if err != nil {
			lgger.Error("failed to join cluster - retrying", zap.Error(err),
				zap.Int("attempt", x+1),
				zap.String("joinAddr", joinAddr),
				zap.String("localAddr", localAddr),
			)
			continue
		} else {
			break
		}
	}
	return nil
}
