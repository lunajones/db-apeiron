package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"db-apeiron/internal/bootstrap"
	"db-apeiron/internal/config"
	"db-apeiron/internal/database"
	"db-apeiron/internal/logger"
)

func main() {
	// ==================================================
	// CONTEXT ROOT
	// ==================================================
	ctx := context.Background()

	// ==================================================
	// LOAD CONFIG
	// ==================================================
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	// ==================================================
	// INIT LOGGER
	// ==================================================
	logger.Initialize(cfg.Logger)

	log := logger.WithComponent("main")

	log.Info().
		Str("app", cfg.App.Name).
		Str("env", cfg.App.Environment).
		Msg("starting server")

	// ==================================================
	// INIT DATABASE
	// ==================================================
	db, err := database.NewPostgres(ctx, cfg.DB)
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("failed to connect database")
	}

	// ==================================================
	// RUN MIGRATIONS
	// ==================================================
	if err := database.RunMigrations(ctx, db.Pool); err != nil {
		log.Fatal().
			Err(err).
			Msg("migration failed")
	}

	// ==================================================
	// RUN SEEDS
	// ==================================================
	if err := database.RunSeeds(ctx, db.Pool); err != nil {
		log.Fatal().
			Err(err).
			Msg("seed failed")
	}

	log.Info().
		Msg("database initialized successfully")

	// ==================================================
	// INIT APPLICATION
	// ==================================================
	app := bootstrap.NewApp(ctx, cfg, db)

	if err := app.Start(ctx); err != nil {
		log.Fatal().
			Err(err).
			Msg("failed to start application")
	}

	// ==================================================
	// SHUTDOWN HANDLER
	// ==================================================
	waitForShutdown(ctx, app)
}

func waitForShutdown(ctx context.Context, app *bootstrap.App) {
	shutdownSignal := make(chan os.Signal, 1)

	signal.Notify(
		shutdownSignal,
		syscall.SIGINT,
		syscall.SIGTERM,
	)

	<-shutdownSignal

	ctxShutdown, cancel := context.WithTimeout(
		ctx,
		10*time.Second,
	)
	defer cancel()

	log := logger.WithComponent("shutdown")

	log.Info().Msg("shutting down gracefully")

	done := make(chan struct{})

	go func() {
		defer close(done)

		if app != nil {
			if err := app.Shutdown(ctxShutdown); err != nil {
				log.Error().
					Err(err).
					Msg("application shutdown failed")
			}
		}
	}()

	select {
	case <-done:
		log.Info().Msg("shutdown complete")
	case <-ctxShutdown.Done():
		log.Error().
			Err(ctxShutdown.Err()).
			Msg("shutdown timeout")
	}
}
