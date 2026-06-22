package database

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	ErrTxNotFound = errors.New("transaction not found in context")
)

type ctxKeyTx struct{}

type txManager struct {
	pool *pgxpool.Pool
}

func NewTxManager(pool *pgxpool.Pool) TxManager {
	return &txManager{
		pool: pool,
	}
}

func (t *txManager) WithTx(ctx context.Context, fn func(ctx context.Context) error) error {
	tx, err := t.pool.BeginTx(ctx, pgx.TxOptions{
		IsoLevel: pgx.ReadCommitted,
	})
	if err != nil {
		return err
	}

	defer func() {
		if r := recover(); r != nil {
			_ = tx.Rollback(ctx)
			panic(r)
		}
	}()

	txCtx := context.WithValue(ctx, ctxKeyTx{}, tx)

	if err := fn(txCtx); err != nil {
		_ = tx.Rollback(ctx)
		return err
	}

	return tx.Commit(ctx)
}

func GetTx(ctx context.Context) (pgx.Tx, error) {
	tx, ok := ctx.Value(ctxKeyTx{}).(pgx.Tx)
	if !ok || tx == nil {
		return nil, ErrTxNotFound
	}
	return tx, nil
}

func (t *txManager) Exec(ctx context.Context, sql string, args ...any) (pgconn.CommandTag, error) {
	tx, err := GetTx(ctx)
	if err == nil {
		return tx.Exec(ctx, sql, args...)
	}

	return t.pool.Exec(ctx, sql, args...)
}

func (t *txManager) Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error) {
	tx, err := GetTx(ctx)
	if err == nil {
		return tx.Query(ctx, sql, args...)
	}

	return t.pool.Query(ctx, sql, args...)
}

func (t *txManager) QueryRow(ctx context.Context, sql string, args ...any) pgx.Row {
	tx, err := GetTx(ctx)
	if err == nil {
		return tx.QueryRow(ctx, sql, args...)
	}

	return t.pool.QueryRow(ctx, sql, args...)
}
