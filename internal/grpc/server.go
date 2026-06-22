package grpc

import (
	"context"
	"fmt"
	"net"
	"time"

	"db-apeiron/internal/config"
	"db-apeiron/internal/logger"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	grpcServer *grpc.Server
	port       string
}

func NewServer(cfg config.GRPCConfig) *Server {
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(loggingInterceptor()),
	)

	server := &Server{
		grpcServer: grpcServer,
		port:       cfg.Port,
	}

	server.registerSystemServices()

	return server
}

func (s *Server) registerSystemServices() {
	healthServer := health.NewServer()
	healthServer.SetServingStatus("", grpc_health_v1.HealthCheckResponse_SERVING)

	grpc_health_v1.RegisterHealthServer(s.grpcServer, healthServer)

	// Dev only.
	reflection.Register(s.grpcServer)
}

func (s *Server) Register(registerFn func(server *grpc.Server)) {
	registerFn(s.grpcServer)
}

func (s *Server) Start() error {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", s.port))
	if err != nil {
		return fmt.Errorf("failed to listen grpc port: %w", err)
	}

	log := logger.WithComponent("grpc")

	log.Info().
		Str("port", s.port).
		Msg("grpc server starting")

	if err := s.grpcServer.Serve(listener); err != nil {
		return fmt.Errorf("grpc server stopped: %w", err)
	}

	return nil
}

func (s *Server) Stop(ctx context.Context) {
	log := logger.WithComponent("grpc")

	log.Info().Msg("gracefully stopping grpc server")

	stopped := make(chan struct{})

	go func() {
		defer close(stopped)
		s.grpcServer.GracefulStop()
	}()

	select {
	case <-stopped:
		log.Info().Msg("grpc server stopped")
	case <-ctx.Done():
		log.Warn().
			Err(ctx.Err()).
			Msg("grpc shutdown timeout reached")

		s.grpcServer.Stop()
	}
}

func loggingInterceptor() grpc.UnaryServerInterceptor {
	log := logger.WithComponent("grpc")

	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (any, error) {
		start := time.Now()

		resp, err := handler(ctx, req)

		log.Info().
			Str("method", info.FullMethod).
			Dur("duration", time.Since(start)).
			Err(err).
			Msg("request handled")

		return resp, err
	}
}
