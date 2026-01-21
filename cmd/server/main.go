package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/DmytroSobko/FormForgeBackend/internal/config"
	"github.com/DmytroSobko/FormForgeBackend/internal/db"
	apphttp "github.com/DmytroSobko/FormForgeBackend/internal/http"
)

func main() {
	cfg := config.Load()
	database := db.Connect(cfg.DatabaseURL)

	router := apphttp.NewRouter(database)

	server := &http.Server{
		Addr:    cfg.Port,
		Handler: router,
	}

	// Run server in goroutine
	go func() {
		log.Println("Server starting on", cfg.Port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen error: %v", err)
		}
	}()

	// Listen for shutdown signals
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop // wait for signal

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown failed: %v", err)
	}

	log.Println("Server stopped gracefully")

	defer database.Close()
}
