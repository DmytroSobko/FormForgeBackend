package db

import (
	"os"

	"github.com/DmytroSobko/FormForgeBackend/internal/logging"
)

type DBConfig struct {
	Port        string
	DatabaseURL string
}

func LoadDBConfig() DBConfig {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		logging.Logger.Error("environment variable missing", "name", "DATABASE_URL")
		os.Exit(1)
	}

	return DBConfig{
		Port:        port,
		DatabaseURL: dbURL,
	}
}
