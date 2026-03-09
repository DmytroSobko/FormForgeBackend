package db

import (
	"log"

	"github.com/DmytroSobko/FormForgeBackend/internal/logging"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrations(dbURL string) {
	m, err := migrate.New("file:///app/migrations", dbURL)

	if err != nil {
		log.Fatal(err)
	}

	defer m.Close()

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}

	logging.Logger.Info("migrations applied successfully")
}
