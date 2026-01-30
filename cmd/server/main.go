package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/DmytroSobko/FormForgeBackend/internal/configs"
	"github.com/DmytroSobko/FormForgeBackend/internal/db"
	"github.com/DmytroSobko/FormForgeBackend/internal/simulation"

	httpRouter "github.com/DmytroSobko/FormForgeBackend/internal/http"
)

func main() {
	configPath := "config/simulation.v1.json"

	simConfig, err := configs.LoadSimulationConfig(configPath)
	if err != nil {
		log.Fatalf(
			"failed to load simulation config: %v",
			err,
		)
	}

	simEngine := simulation.NewEngine(
		&simConfig.Simulation,
	)

	log.Printf(
		"Loaded simulation config version %s",
		simConfig.Version,
	)

	cfg := configs.LoadDBConfig()
	database := db.Connect(cfg.DatabaseURL)

	router := httpRouter.NewRouter(
		database,
		simConfig,
		simEngine,
	)

	server := &http.Server{
		Addr:         cfg.Port,
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
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
