package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/autom8ter/graphik/config"
	"github.com/autom8ter/graphik/graph"
	"github.com/autom8ter/graphik/graph/generated"
	"github.com/autom8ter/graphik/jwks"
	"github.com/autom8ter/graphik/logger"
	"github.com/autom8ter/graphik/store"
	"github.com/autom8ter/machine"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/cors"
	"github.com/spf13/pflag"
	"go.uber.org/zap"
	"net"
	"net/http"
	"net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const version = "0.0.0"

func init() {
	pflag.CommandLine.IntVar(&cfg.Port, "port", 8080, "port to serve on")
	pflag.CommandLine.StringVar(&cfg.Raft.DBPath, "raft.path", "/tmp/graphik", "path to database folder")
	pflag.CommandLine.IntVar(&cfg.Raft.Bind, "raft.bind", 8081, "bind raft protocol to local port")
	pflag.CommandLine.StringVar(&cfg.Raft.Join, "raft.join", "", "join raft cluster leader")
	pflag.CommandLine.StringVar(&cfg.Raft.NodeID, "raft.id", "main", "unique raft node id")
	pflag.CommandLine.StringSliceVar(&cfg.JWKs, "jwks", nil, "remote json web key set(s)")
	pflag.CommandLine.StringSliceVar(&cfg.Cors.AllowedHeaders, "cors.headers", nil, "allowed cors headers")
	pflag.CommandLine.StringSliceVar(&cfg.Cors.AllowedMethods, "cors.methods", nil, "allowed cors methods")
	pflag.CommandLine.StringSliceVar(&cfg.Cors.AllowedOrigins, "cors.origins", nil, "allowed cors origins")
}

var (
	cfg = &config.Config{
		Raft: &config.Raft{},
	}
)

func main() {
	pflag.Parse()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(interrupt)

	mach := machine.New(ctx)
	router := mux.NewRouter()
	stor, err := store.New(
		store.WithLeader(cfg.Raft.Join == ""),
		store.WithID(cfg.Raft.NodeID),
		store.WithBindAddr(fmt.Sprintf("localhost:%v", cfg.Raft.Bind)),
		store.WithRaftDir(cfg.Raft.DBPath),
	)
	if err != nil {
		logger.Error("failed to create raft store", zap.Error(err))
		return
	}
	if cfg.Raft.Join != "" {
		if err := joinRaft(cfg.Raft.Join, fmt.Sprintf("localhost:%v", cfg.Raft.Bind), cfg.Raft.NodeID); err != nil {
			logger.Error("failed to join cluster", zap.Error(err))
			return
		}
	}
	resolver := graph.NewResolver(mach, stor)
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))
	router.Handle("/", playground.Handler("GraphQL playground", "/api/query"))
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   cfg.Cors.AllowedOrigins,
		AllowedMethods:   cfg.Cors.AllowedMethods,
		AllowedHeaders:   cfg.Cors.AllowedHeaders,
		AllowCredentials: true,
		Debug:            true,
	}).Handler)
	a, err := jwks.New(cfg.JWKs)
	if err != nil {
		logger.Error("failed to fetch jwks", zap.Error(err))
		return
	}
	mach.Go(func(routine machine.Routine) {
		logger.Info("refreshing jwks")
		if err := a.RefreshKeys(); err != nil {
			logger.Error("failed to refresh keys", zap.Error(err))
		}
	}, machine.GoWithMiddlewares(machine.Cron(time.NewTicker(1*time.Minute))))

	middleware := a.Middleware()
	router.Handle("/api/query", middleware(srv))
	router.Handle("/api/join", middleware(stor.Join())).Methods(http.MethodPost)
	router.Handle("/api/export", middleware(stor.Export())).Methods(http.MethodGet)
	router.Handle("/api/import", middleware(stor.Import())).Methods(http.MethodPost)
	router.Handle("/api/jwks", middleware(a.GetJWKS())).Methods(http.MethodGet)
	router.Handle("/api/jwks", middleware(a.PutJWKS())).Methods(http.MethodPut)

	router.Handle("/metrics", promhttp.Handler()).Methods(http.MethodGet)
	router.HandleFunc("/debug/pprof/", pprof.Index).Methods(http.MethodGet)
	router.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline).Methods(http.MethodGet)
	router.HandleFunc("/debug/pprof/profile", pprof.Profile).Methods(http.MethodGet)
	router.HandleFunc("/debug/pprof/symbol", pprof.Symbol).Methods(http.MethodGet)
	router.HandleFunc("/debug/pprof/trace", pprof.Trace).Methods(http.MethodGet)

	server := &http.Server{
		Handler: router,
	}
	mach.Go(func(routine machine.Routine) {
		lis, err := net.Listen("tcp", fmt.Sprintf(":%v", cfg.Port))
		if err != nil {
			logger.Error("failed to create server listener", zap.Error(err))
			return
		}
		defer lis.Close()
		logger.Info("starting graphql server",
			zap.String("address", lis.Addr().String()),
			zap.String("version", version),
		)
		if err := server.Serve(lis); err != nil && err != http.ErrServerClosed {
			logger.Error("server failure", zap.Error(err))
		}
	})
	select {
	case <-interrupt:
		mach.Cancel()
		break
	case <-ctx.Done():
		mach.Cancel()
		break
	}
	logger.Warn("shutdown signal received")
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownCancel()

	_ = server.Shutdown(shutdownCtx)
	_ = resolver.Close()
	logger.Info("shutdown successful")
	mach.Wait()
}

func joinRaft(joinAddr, raftAddr, nodeID string) error {
	b, err := json.Marshal(map[string]string{"addr": raftAddr, "id": nodeID})
	if err != nil {
		return err
	}
	resp, err := http.Post(fmt.Sprintf("http://%s/api/join", joinAddr), "application-type/json", bytes.NewReader(b))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}
