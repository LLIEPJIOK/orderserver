package logger

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
)

const (
	LoggerKey   = "logger"
	RequestID   = "requestID"
	ServiceName = "serviceName"
)

type Logger interface {
	Infof(ctx context.Context, msg string, args ...any)
	Errorf(ctx context.Context, msg string, args ...any)
	Fatalf(ctx context.Context, msg string, args ...any)
}

type logger struct {
	logger      *slog.Logger
	serviceName string
}

func New(w io.Writer, lvl slog.Level, serviceName string) Logger {
	return &logger{
		logger: slog.New(slog.NewJSONHandler(w, &slog.HandlerOptions{
			Level: lvl,
		})),
		serviceName: serviceName,
	}
}

func (l *logger) Infof(ctx context.Context, msg string, args ...any) {
	attrs := make([]any, 0)
	attrs = append(attrs, slog.String("serviceName", l.serviceName))

	reqID := ctx.Value(RequestID)
	if reqID != nil {
		attrs = append(attrs, slog.String(RequestID, reqID.(string)))
	}

	l.logger.InfoContext(ctx, fmt.Sprintf(msg, args...), attrs...)
}

func (l *logger) Errorf(ctx context.Context, msg string, args ...any) {
	attrs := make([]any, 0)
	attrs = append(attrs, slog.String(ServiceName, l.serviceName))

	reqID := ctx.Value(RequestID)
	if reqID != nil {
		attrs = append(attrs, slog.String(RequestID, reqID.(string)))
	}

	l.logger.ErrorContext(ctx, fmt.Sprintf(msg, args...), attrs...)
}

func (l *logger) Fatalf(ctx context.Context, msg string, args ...any) {
	l.Errorf(ctx, msg, args...)
	os.Exit(1)
}

func GetLoggerFromCtx(ctx context.Context) Logger {
	return ctx.Value(LoggerKey).(Logger)
}

func CtxWithLogger(ctx context.Context, lg Logger) context.Context {
	return context.WithValue(ctx, LoggerKey, lg)
}
