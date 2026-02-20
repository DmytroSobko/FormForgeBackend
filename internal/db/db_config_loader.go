package db

import (
	"log"
	"os"
)

type DBConfig struct {
	Port        string
	DatabaseURL string
}

func LoadDBConfig() DBConfig {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT is not set")
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
