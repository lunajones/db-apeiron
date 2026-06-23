package interceptors

import (
	"context"
	"runtime/debug"

	"db-apeiron/internal/logger"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func RecoveryInterceptor() grpc.UnaryServerInterceptor {
	log := logger.WithComponent("grpc-recovery")

	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp any, err error) {

		defer func() {
			if r := recover(); r != nil {

				log.Error().
					Str("method", info.FullMethod).
					Interface("panic", r).
					Str("stack", string(debug.Stack())).
					Msg("panic handled in grpc handler")

				resp = nil
				err = status.Errorf(
					codes.Internal,
					"internal server error",
				)
			}
		}()

		return handler(ctx, req)
	}
}
