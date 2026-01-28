package db

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	Pool *pgxpool.Pool
}

func Connect(databaseURL string) *DB {
	const maxRetries = 10

	var pool *pgxpool.Pool
	var err error

	for i := 1; i <= maxRetries; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		pool, err = pgxpool.New(ctx, databaseURL)
		if err == nil {
			if err = pool.Ping(ctx); err == nil {
				log.Println("Connected to PostgreSQL")
				return &DB{Pool: pool}
			}
		}

		log.Printf("DB not ready (attempt %d/%d), retrying...", i, maxRetries)
		time.Sleep(2 * time.Second)
	}

	log.Fatalf("Unable to connect to DB after %d attempts: %v", maxRetries, err)
	return nil
}

func (d *DB) Close() {
	d.Pool.Close()
}

func (d *DB) IsHealthy() bool {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	return d.Pool.Ping(ctx) == nil
}
