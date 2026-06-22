package database

import (
	"context"
	"db-apeiron/internal/logger"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	migrationsDir = "./migrations"
	bootstrapDir  = "./bootstrap"
)

func RunMigrations(ctx context.Context, pool *pgxpool.Pool) error {
	migrationLogger := logger.WithComponent("migration")

	migrationLogger.Info().
		Msg("starting database migrations")

	if err := ensureMigrationTable(ctx, pool); err != nil {
		return err
	}

	files, err := loadSQLFiles(migrationsDir)
	if err != nil {
		return err
	}

	for _, file := range files {
		applied, err := isMigrationApplied(ctx, pool, file)
		if err != nil {
			return err
		}

		if applied {
			migrationLogger.Info().
				Str("file", file).
				Msg("migration skipped")
			continue
		}

		if err := executeMigration(ctx, pool, file); err != nil {
			return fmt.Errorf("migration failed (%s): %w", file, err)
		}

		migrationLogger.Info().
			Str("file", file).
			Msg("migration applied")
	}

	migrationLogger.Info().
		Msg("migrations completed successfully")

	return nil
}

func RunSeeds(ctx context.Context, pool *pgxpool.Pool) error {
	seedLogger := logger.WithComponent("seed")

	seedLogger.Info().
		Msg("starting database seeds")

	files, err := loadSQLFiles(bootstrapDir)
	if err != nil {
		return err
	}

	for _, file := range files {
		if err := executeSQLFile(ctx, pool, file); err != nil {
			return fmt.Errorf("seed failed (%s): %w", file, err)
		}

		seedLogger.Info().
			Str("file", file).
			Msg("seed applied")
	}

	seedLogger.Info().
		Msg("seeds completed successfully")

	return nil
}

func ensureMigrationTable(ctx context.Context, pool *pgxpool.Pool) error {
	_, err := pool.Exec(ctx, `
		CREATE SCHEMA IF NOT EXISTS apeiron;

		CREATE TABLE IF NOT EXISTS apeiron.schema_migrations (
			file_name TEXT PRIMARY KEY,
			applied_at TIMESTAMP NOT NULL DEFAULT NOW()
		);
	`)

	return err
}

func isMigrationApplied(ctx context.Context, pool *pgxpool.Pool, file string) (bool, error) {
	fileName := filepath.Base(file)

	var exists bool

	err := pool.QueryRow(ctx, `
		SELECT EXISTS (
			SELECT 1
			FROM apeiron.schema_migrations
			WHERE file_name = $1
		)
	`, fileName).Scan(&exists)

	return exists, err
}

func executeMigration(ctx context.Context, pool *pgxpool.Pool, file string) error {
	sqlBytes, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	sqlContent := strings.TrimSpace(string(sqlBytes))
	if sqlContent == "" {
		return markMigrationApplied(ctx, pool, file)
	}

	tx, err := pool.Begin(ctx)
	if err != nil {
		return err
	}

	defer func() {
		_ = tx.Rollback(ctx)
	}()

	if _, err := tx.Exec(ctx, sqlContent); err != nil {
		return err
	}

	fileName := filepath.Base(file)

	if _, err := tx.Exec(ctx, `
		INSERT INTO apeiron.schema_migrations (file_name)
		VALUES ($1)
		ON CONFLICT (file_name) DO NOTHING
	`, fileName); err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func markMigrationApplied(ctx context.Context, pool *pgxpool.Pool, file string) error {
	fileName := filepath.Base(file)

	_, err := pool.Exec(ctx, `
		INSERT INTO apeiron.schema_migrations (file_name)
		VALUES ($1)
		ON CONFLICT (file_name) DO NOTHING
	`, fileName)

	return err
}

func executeSQLFile(ctx context.Context, pool *pgxpool.Pool, file string) error {
	sqlBytes, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	sqlContent := strings.TrimSpace(string(sqlBytes))
	if sqlContent == "" {
		return nil
	}

	tx, err := pool.Begin(ctx)
	if err != nil {
		return err
	}

	defer func() {
		_ = tx.Rollback(ctx)
	}()

	if _, err := tx.Exec(ctx, sqlContent); err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func loadSQLFiles(dir string) ([]string, error) {
	var files []string

	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		if strings.HasSuffix(strings.ToLower(path), ".sql") {
			files = append(files, path)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	sort.Strings(files)

	return files, nil
}
