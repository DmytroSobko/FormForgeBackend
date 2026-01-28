package config

import (
	"log"
	"os"
)

type DBConfig struct {
	Port        string
	DatabaseURL string
}

func Load() DBConfig {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	return DBConfig{
		Port:        port,
		DatabaseURL: dbURL,
	}
}
