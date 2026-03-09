package db

import (
	"context"

	"github.com/jackc/pgx/v5"
)

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
