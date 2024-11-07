package grpc

import (
	"context"
	"fmt"
	"log"
	"net"

	client "gitlab.crja72.ru/gospec/students/223640-nphne-et6ofbhg-course-1195/pkg/api/order"
	"google.golang.org/grpc"
)

type Server struct {
	grpcServer *grpc.Server
	listener   net.Listener
}

func New(_ context.Context, port int, srv Service) (*Server, error) {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		return nil, fmt.Errorf("failed to listen: %w", err)
	}

	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	client.RegisterOrderServiceServer(grpcServer, NewOrderService(srv))

	return &Server{
		grpcServer: grpcServer,
		listener:   lis,
	}, nil
}

func (s *Server) Start() error {
	log.Printf("starting gRPC server at port=%d", s.listener.Addr().(*net.TCPAddr).Port)
	return s.grpcServer.Serve(s.listener)
}
