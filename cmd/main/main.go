package main

import (
	"context"
	"log"

	"gitlab.crja72.ru/gospec/students/223640-nphne-et6ofbhg-course-1195/internal/config"
	"gitlab.crja72.ru/gospec/students/223640-nphne-et6ofbhg-course-1195/internal/repository"
	"gitlab.crja72.ru/gospec/students/223640-nphne-et6ofbhg-course-1195/internal/service"
	"gitlab.crja72.ru/gospec/students/223640-nphne-et6ofbhg-course-1195/internal/transport/grpc"
)

func main() {
	ctx := context.Background()

	cfg, err := config.New()
	if err != nil {
		log.Fatalf("failed to get config: %s", err)
	}

	repo := repository.NewOrderRepository()
	ordService := service.NewOrderService(repo)
	grpcServer, err := grpc.New(ctx, cfg.GRPCServerPort, ordService)
	if err != nil {
		log.Fatalf("failed to prepare server: %s", err)
	}

	if err := grpcServer.Start(); err != nil {
		log.Fatalf("failed to start server: %s", err)
	}
}
