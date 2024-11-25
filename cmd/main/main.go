package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/LLIEPJIOK/orderserver/internal/config"
	"github.com/LLIEPJIOK/orderserver/internal/repository"
	"github.com/LLIEPJIOK/orderserver/internal/service"
	"github.com/LLIEPJIOK/orderserver/internal/transport/grpc"
	"github.com/LLIEPJIOK/orderserver/pkg/database/postgres"
	"github.com/LLIEPJIOK/orderserver/pkg/database/redis"
	"github.com/LLIEPJIOK/orderserver/pkg/logger"
)

const (
	serviceName = "lyceum"
	loggerLvl   = slog.LevelInfo
)

func main() {
	ctx := context.Background()

	serviceLogger := logger.New(os.Stdout, loggerLvl, serviceName)
	ctx = logger.CtxWithLogger(ctx, serviceLogger)

	cfg, err := config.New()
	if err != nil {
		serviceLogger.Fatalf(ctx, "failed to get config: %s", err)
	}

	db, err := postgres.NewPostgres(cfg.PostgresConfig)
	if err != nil {
		serviceLogger.Fatalf(ctx, "failed to init database: %s", err)
	}

	cache, err := redis.New(cfg.RedisConfig)
	if err != nil {
		serviceLogger.Fatalf(ctx, "failed to init cache: %s", err)
	}

	repo := repository.NewOrderRepository(db, cache)
	ordService := service.NewOrderService(repo)

	grpcServer, err := grpc.New(ctx, cfg.GRPCServerPort, cfg.RESTServerPort, ordService)
	if err != nil {
		serviceLogger.Fatalf(ctx, "failed to prepare server: %s", err)
	}

	graceCh := make(chan os.Signal, 1)
	signal.Notify(graceCh, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := grpcServer.Start(ctx); err != nil {
			serviceLogger.Errorf(ctx, "failed to start server: %s", err)
		}
	}()

	<-graceCh

	if err := grpcServer.Stop(ctx); err != nil {
		serviceLogger.Errorf(ctx, "failed to stop server: %s", err)
	}

	if err := db.Close(); err != nil {
		serviceLogger.Errorf(ctx, "failed to close db: %s", err)
	}

	serviceLogger.Infof(ctx, "Server stopped")
}
