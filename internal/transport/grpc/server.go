package grpc

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	client "github.com/LLIEPJIOK/orderserver/pkg/api/order"
	"github.com/LLIEPJIOK/orderserver/pkg/logger"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

type Server struct {
	grpcServer *grpc.Server
	restServer *http.Server
	listener   net.Listener
}

func New(ctx context.Context, grpcPort, restPort int, srv Service) (*Server, error) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		return nil, fmt.Errorf("failed to listen: %w", err)
	}

	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(LoggerInterceptor(logger.GetLoggerFromCtx(ctx))),
	}

	grpcServer := grpc.NewServer(opts...)
	client.RegisterOrderServiceServer(grpcServer, NewOrderService(srv))

	restSrv := runtime.NewServeMux()
	err = client.RegisterOrderServiceHandlerServer(ctx, restSrv, NewOrderService(srv))
	if err != nil {
		return nil, fmt.Errorf("failed to register service handler server: %w", err)
	}

	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", restPort),
		Handler: restSrv,
	}

	return &Server{
		grpcServer: grpcServer,
		restServer: httpServer,
		listener:   lis,
	}, nil
}

func (s *Server) Start(ctx context.Context) error {
	eg, ctx := errgroup.WithContext(ctx)
	serviceLogger := logger.GetLoggerFromCtx(ctx)

	eg.Go(func() error {
		serviceLogger.Infof(ctx, "starting gRPC server at address=%s", s.listener.Addr().String())

		if err := s.grpcServer.Serve(s.listener); err != nil {
			return fmt.Errorf("start grpc server: %w", err)
		}

		return nil
	})

	eg.Go(func() error {
		serviceLogger.Infof(ctx, "starting rest server at address=%q", s.restServer.Addr)

		if err := s.restServer.ListenAndServe(); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				return fmt.Errorf("start rest server: %w", err)
			}
		}

		return nil
	})

	if err := eg.Wait(); err != nil {
		return fmt.Errorf("starting server: %w", err)
	}

	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	s.grpcServer.GracefulStop()

	if err := s.restServer.Shutdown(ctx); err != nil {
		return fmt.Errorf("stop rest server: %w", err)
	}

	return nil
}
