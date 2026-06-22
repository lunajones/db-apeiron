package interceptors

import (
	"context"
	"time"

	"db-apeiron/internal/logger"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func LoggingInterceptor() grpc.UnaryServerInterceptor {
	log := logger.WithComponent("grpc")

	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (any, error) {

		start := time.Now()

		resp, err := handler(ctx, req)

		duration := time.Since(start)

		st, _ := status.FromError(err)
		code := st.Code().String()

		log.Info().
			Str("method", info.FullMethod).
			Str("status", code).
			Dur("duration", duration).
			Err(err).
			Msg("grpc request")

		return resp, err
	}
}
