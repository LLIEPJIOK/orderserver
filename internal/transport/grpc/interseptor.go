package grpc

import (
	"context"

	"gitlab.crja72.ru/gospec/students/223640-nphne-et6ofbhg-course-1195/pkg/logger"
	"google.golang.org/grpc"
)

func LoggerInterceptor(l logger.Logger) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp any, err error) {
		l.Infof(ctx, "request started method=%s", info.FullMethod)
		ctx = logger.CtxWithLogger(ctx, l)

		return handler(ctx, req)
	}
}
