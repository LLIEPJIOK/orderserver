package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

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

	grpcServer, err := grpc.New(ctx, cfg.GRPCServerPort, cfg.RESTServerPort, ordService)
	if err != nil {
		log.Fatalf("failed to prepare server: %s", err)
	}

	graceCh := make(chan os.Signal, 1)
	signal.Notify(graceCh, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := grpcServer.Start(ctx); err != nil {
			log.Printf("failed to start server: %s", err)
		}
	}()

	<-graceCh

	if err := grpcServer.Stop(ctx); err != nil {
		log.Printf("failed to stop server: %s", err)
		os.Exit(1)
	}

	log.Println("Server stopped")
}
