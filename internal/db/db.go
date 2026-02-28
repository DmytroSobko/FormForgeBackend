package db

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
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
				log.Println("Connected to PostgreSQL")
				return &DB{pool: pool}
			}
		}

		cancel()

		log.Printf("DB not ready (attempt %d/%d), retrying...", i, maxRetries)
		time.Sleep(2 * time.Second)
	}

	log.Fatalf("Unable to connect to DB after %d attempts: %v", maxRetries, err)
	return nil
}

// Close gracefully shuts down the connection pool.
func (d *DB) Close() {
	d.pool.Close()
}

// IsHealthy checks database connectivity (used by health checks).
func (d *DB) IsHealthy() bool {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	return d.pool.Ping(ctx) == nil
}

// Exec executes a query without returning rows.
func (d *DB) Exec(ctx context.Context, query string, args ...any) error {
	_, err := d.pool.Exec(ctx, query, args...)
	return err
}

// Query executes a query that returns multiple rows.
func (d *DB) Query(ctx context.Context, query string, args ...any) (pgx.Rows, error) {
	return d.pool.Query(ctx, query, args...)
}

// QueryRow executes a query that returns a single row.
func (d *DB) QueryRow(ctx context.Context, query string, args ...any) pgx.Row {
	return d.pool.QueryRow(ctx, query, args...)
}
