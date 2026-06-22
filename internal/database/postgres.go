package database

import (
	"context"
	"db-apeiron/internal/config"
	"db-apeiron/internal/logger"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Postgres struct {
	Pool *pgxpool.Pool
}

func NewPostgres(ctx context.Context, cfg config.DatabaseConfig) (*Postgres, error) {
	dbLogger := logger.WithComponent("postgres")

	pool, err := NewPool(ctx, cfg)
	if err != nil {
		return nil, err
	}

	db := &Postgres{
		Pool: pool,
	}

	dbLogger.Info().
		Msg("postgres initialized")

	return db, nil
}

func (p *Postgres) Close() {
	if p.Pool == nil {
		return
	}

	p.Pool.Close()
}
