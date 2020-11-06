package main

import (
	"context"
	apipb "github.com/autom8ter/graphik/api"
	"github.com/autom8ter/graphik/helpers"
	"github.com/autom8ter/graphik/logger"
	"github.com/autom8ter/graphik/runtime"
	"github.com/autom8ter/graphik/service/private"
	"github.com/autom8ter/machine"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/pflag"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func init() {
	pflag.CommandLine.StringVar(&configFile, "config", "./graphik.json", "path to json config")
}

var (
	configFile string
	cfg        = &apipb.Config{
		Http: &apipb.HTTPConfig{},
		Grpc: &apipb.GRPCConfig{},
		Raft: &apipb.RaftConfig{},
	}
)

func main() {
	pflag.Parse()
	f, err := os.Open(configFile)
	if err == nil {
		if err := helpers.JSONDecode(f, cfg); err != nil {
			logger.Error("failed to decode config", zap.Error(err))
			return
		}
		f.Close()
	}

	cfg.SetDefaults()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(interrupt)
	runt, err := runtime.New(ctx, cfg)
	if err != nil {
		logger.Error("failed to create runtime", zap.Error(err))
		return
	}
	router := http.NewServeMux()
	router.Handle("/metrics", promhttp.Handler())
	router.HandleFunc("/debug/pprof/", pprof.Index)
	router.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	router.HandleFunc("/debug/pprof/profile", pprof.Profile)
	router.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	router.HandleFunc("/debug/pprof/trace", pprof.Trace)

	server := &http.Server{
		Handler: router,
	}

	runt.Go(func(routine machine.Routine) {
		lis, err := net.Listen("tcp", cfg.GetHttp().GetBind())
		if err != nil {
			logger.Error("failed to create http server listener", zap.Error(err))
			return
		}
		defer lis.Close()
		logger.Info("starting http server",
			zap.String("address", lis.Addr().String()),
		)
		if err := server.Serve(lis); err != nil && err != http.ErrServerClosed {
			logger.Error("http server failure", zap.Error(err))
		}
	})

	gserver := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			grpc_prometheus.UnaryServerInterceptor,
			grpc_zap.UnaryServerInterceptor(logger.Logger()),
			grpc_auth.UnaryServerInterceptor(runt.JWTMiddleware()),
			grpc_recovery.UnaryServerInterceptor(),
		),
		grpc.ChainStreamInterceptor(
			grpc_prometheus.StreamServerInterceptor,
			grpc_zap.StreamServerInterceptor(logger.Logger()),
			grpc_auth.StreamServerInterceptor(runt.JWTMiddleware()),
			grpc_recovery.StreamServerInterceptor(),
		),
	)

	privateService := private.NewService(runt)
	apipb.RegisterPrivateServiceServer(gserver, privateService)
	grpc_prometheus.Register(gserver)

	runt.Go(func(routine machine.Routine) {
		lis, err := net.Listen("tcp", cfg.GetGrpc().GetBind())
		if err != nil {
			logger.Error("failed to create gRPC server listener", zap.Error(err))
			return
		}
		defer lis.Close()
		logger.Info("starting gRPC server",
			zap.String("address", lis.Addr().String()),
		)
		if err := gserver.Serve(lis); err != nil && err != http.ErrServerClosed {
			logger.Error("gRPC server failure", zap.Error(err))
		}
	})
	select {
	case <-interrupt:
		runt.Cancel()
		break
	case <-ctx.Done():
		runt.Cancel()
		break
	}

	logger.Warn("shutdown signal received")
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownCancel()

	_ = server.Shutdown(shutdownCtx)
	stopped := make(chan struct{})
	go func() {
		gserver.GracefulStop()
		close(stopped)
	}()

	t := time.NewTimer(5 * time.Second)

	select {
	case <-t.C:
		gserver.Stop()
	case <-stopped:
		t.Stop()
	}
	_ = runt.Close()
	logger.Info("shutdown successful")
	runt.Wait()
}
