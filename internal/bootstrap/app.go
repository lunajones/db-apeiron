package bootstrap

import (
	"context"

	apeironv1 "db-apeiron/gen/apeiron/v1"
	"db-apeiron/internal/config"
	"db-apeiron/internal/database"
	internalgrpc "db-apeiron/internal/grpc"
	"db-apeiron/internal/grpc/handlers"
	"db-apeiron/internal/logger"

	"google.golang.org/grpc"
)

type App struct {
	Config *config.Config
	DB     *database.Postgres
	Deps   *Dependencies

	GRPCServer *internalgrpc.Server
}

func NewApp(ctx context.Context, cfg *config.Config, db *database.Postgres) *App {
	deps := NewDependencies(db.Pool)

	_ = ctx

	return &App{
		Config: cfg,
		DB:     db,
		Deps:   deps,
	}
}

func (a *App) Start(ctx context.Context) error {
	log := logger.WithComponent("app")

	log.Info().
		Msg("application dependencies initialized")

	a.GRPCServer = internalgrpc.NewServer(a.Config.GRPC)

	a.GRPCServer.Register(func(server *grpc.Server) {
		apeironv1.RegisterCacheServiceServer(
			server,
			handlers.NewCacheHandler(a.Deps.CacheLoader),
		)
		apeironv1.RegisterCreatureDataServiceServer(
			server,
			handlers.NewCreatureDataHandler(a.Deps.Caches.Templates),
		)
		apeironv1.RegisterSkillDataServiceServer(
			server,
			handlers.NewSkillDataHandler(a.Deps.Caches.Skills),
		)
		apeironv1.RegisterProfileDataServiceServer(
			server,
			handlers.NewProfileDataHandler(a.Deps.Caches.Profiles),
		)
		apeironv1.RegisterWorldDataServiceServer(
			server,
			handlers.NewWorldDataHandler(a.Deps.Caches.World),
		)
	})

	go func() {
		if err := a.GRPCServer.Start(); err != nil {
			log.Error().
				Err(err).
				Msg("grpc server stopped with error")
		}
	}()

	_ = ctx

	log.Info().
		Msg("application started")

	return nil
}

func (a *App) Shutdown(ctx context.Context) error {
	log := logger.WithComponent("app")

	log.Info().
		Msg("stopping application")

	if a.GRPCServer != nil {
		a.GRPCServer.Stop(ctx)
	}

	if a.DB != nil {
		a.DB.Close()
	}

	log.Info().
		Msg("application stopped")

	return nil
}
