package db

import "log"

type DB struct {
	URL string
}

func Connect(databaseURL string) *DB {
	log.Println("Connecting to database:", databaseURL)
	return &DB{URL: databaseURL}
}

func (d *DB) IsHealthy() bool {
	// Stub for now
	// Later: ping database here
	return true
}
