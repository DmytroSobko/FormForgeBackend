package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/DmytroSobko/FormForgeBackend/internal/app"
	"github.com/DmytroSobko/FormForgeBackend/internal/athlete"
	"github.com/DmytroSobko/FormForgeBackend/internal/configs"
	"github.com/DmytroSobko/FormForgeBackend/internal/db"
	httpRouter "github.com/DmytroSobko/FormForgeBackend/internal/http"
	"github.com/DmytroSobko/FormForgeBackend/internal/simulation"
)

func main() {
	// -------------------------
	// Load simulation config (domain)
	// -------------------------

	simCfg, simulationVersion, err := configs.LoadSimulationConfig("configs/simulation.v1.json")
	if err != nil {
		log.Fatalf("failed to load simulation config: %v", err)
	}

	exercises, exercisesVersion, err := configs.LoadExercises("configs/exercises.v1.json")
	if err != nil {
		log.Fatalf("failed to load exercises config: %v", err)
	}

	intensities, intensitiesVersion, err := configs.LoadIntensities("configs/intensities.v1.json")
	if err != nil {
		log.Fatalf("failed to load intensities config: %v", err)
	}

	athleteTypes, athleteTypesVersion, err := configs.LoadAthleteTypes("configs/athlete_types.v1.json")
	if err != nil {
		log.Fatalf("failed to load athlete_types config: %v", err)
	}

	log.Printf("Loaded configs: simulation=%s exercises=%s intensities=%s athleteTypes=%s",
		simulationVersion,
		exercisesVersion,
		intensitiesVersion,
		athleteTypesVersion,
	)

	// -------------------------
	// Initialize DB
	// -------------------------

	cfg := db.LoadDBConfig()
	database := db.Connect(cfg.DatabaseURL)
	defer database.Close()

	// -------------------------
	// Initialize simulation engine
	// -------------------------

	simEngine := simulation.NewEngine(
		simCfg,
		intensities,
	)

	// -------------------------
	// Initialize repository
	// -------------------------

	athleteRepo := athlete.NewPostgresRepository(database)

	// -------------------------
	// Initialize athlete service
	// -------------------------

	athleteService := athlete.NewService(
		athleteRepo,
		athleteTypes,
	)

	// -------------------------
	// Build application dependencies
	// -------------------------

	deps := app.Dependencies{
		DB:             database,
		Engine:         simEngine,
		Exercises:      exercises,
		Intensities:    intensities,
		AthleteTypes:   athleteTypes,
		SimConfig:      simCfg,
		AthleteService: athleteService,
	}

	// -------------------------
	// Build router
	// -------------------------

	router := httpRouter.NewRouter(deps)

	server := &http.Server{
		Addr:         cfg.Port,
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// -------------------------
	// Start server
	// -------------------------

	go func() {
		log.Println("Server starting on", cfg.Port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen error: %v", err)
		}
	}()

	// -------------------------
	// Graceful shutdown
	// -------------------------

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("server shutdown failed: %v", err)
	}

	log.Println("Server stopped gracefully")
}
