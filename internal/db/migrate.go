package db

import (
	"github.com/DmytroSobko/FormForgeBackend/internal/logging"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
)

func RunMigrations(pool *pgxpool.Pool) error {
	dbURL := pool.Config().ConnString()

	m, err := migrate.New("file:///app/migrations", dbURL)
	if err != nil {
		return err
	}

	defer m.Close()

	if err := m.Up(); err != nil {
		if err == migrate.ErrNoChange {
			logging.Logger.Info("no migrations to apply")
			return nil
		}
		return err
	}

	logging.Logger.Info("migrations applied successfully")
	return nil
}
