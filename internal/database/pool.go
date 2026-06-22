package database

import (
	"context"
	"db-apeiron/internal/config"
	"db-apeiron/internal/logger"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPool(
	ctx context.Context,
	cfg config.DatabaseConfig,
) (*pgxpool.Pool, error) {

	connectionString := buildConnectionString(cfg)

	dbLogger := logger.WithComponent("database")

	dbLogger.Info().
		Str("host", cfg.Host).
		Str("port", cfg.Port).
		Str("database", cfg.Name).
		Msg("connecting to postgres")

	poolConfig, err := pgxpool.ParseConfig(connectionString)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to parse postgres config: %w",
			err,
		)
	}

	applyPoolConfig(poolConfig, cfg)

	var pool *pgxpool.Pool

	for attempt := 1; attempt <= cfg.ConnectRetries; attempt++ {

		pool, err = pgxpool.NewWithConfig(
			ctx,
			poolConfig,
		)

		if err == nil {
			err = pool.Ping(ctx)
		}

		if err == nil {
			dbLogger.Info().
				Int("attempt", attempt).
				Msg("postgres connection established")

			return pool, nil
		}

		dbLogger.Warn().
			Int("attempt", attempt).
			Err(err).
			Msg("failed to connect postgres")

		time.Sleep(cfg.ConnectRetryDelay)
	}

	return nil, fmt.Errorf(
		"could not connect to postgres after %d attempts",
		cfg.ConnectRetries,
	)
}

func buildConnectionString(
	cfg config.DatabaseConfig,
) string {

	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.Name,
		cfg.SSLMode,
	)
}

func applyPoolConfig(
	poolConfig *pgxpool.Config,
	cfg config.DatabaseConfig,
) {

	poolConfig.MaxConns = cfg.MaxConns
	poolConfig.MinConns = cfg.MinConns
	poolConfig.MaxConnLifetime = cfg.MaxConnLifetime
	poolConfig.MaxConnIdleTime = cfg.MaxConnIdleTime

	poolConfig.HealthCheckPeriod = 1 * time.Minute
}
