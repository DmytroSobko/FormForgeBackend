package db

import (
	"context"
	"time"

	"github.com/DmytroSobko/FormForgeBackend/internal/logging"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	pool *pgxpool.Pool
}

// Connect initializes a pgx connection pool with retry logic.
func Connect(databaseURL string) *DB {
	const maxRetries = 10

	var pool *pgxpool.Pool
	var err error

	for i := 1; i <= maxRetries; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

		pool, err = pgxpool.New(ctx, databaseURL)
		if err == nil {
			if err = pool.Ping(ctx); err == nil {
				cancel()

				logging.Logger.Info(
					"connected to PostgreSQL",
				)

				return &DB{pool: pool}
			}

			pool.Close()
		}

		cancel()

		logging.Logger.Info(
			"database not ready, retrying connection",
			"attempt", i,
			"max_attempts", maxRetries,
			"error", err,
		)

		time.Sleep(2 * time.Second)
	}

	logging.Logger.Error(
		"failed to connect to PostgreSQL after retries",
		"max_attempts", maxRetries,
		"error", err,
	)

	return nil
}

// Close gracefully shuts down the connection pool.
func (d *DB) Close() {
	logging.Logger.Info("closing database connection pool")
	d.pool.Close()
}

// IsHealthy checks database connectivity (used by health checks).
func (d *DB) IsHealthy() bool {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err := d.pool.Ping(ctx)
	if err != nil {
		logging.Logger.Warn(
			"database health check failed",
			"error", err,
		)
		return false
	}

	return true
}
