package main

import (
	"context"
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
	"github.com/DmytroSobko/FormForgeBackend/internal/logging"
	"github.com/DmytroSobko/FormForgeBackend/internal/simulation"
)

func main() {

	simCfg, simulationVersion, err := configs.LoadSimulationConfig("configs/simulation.v1.json")
	if err != nil {
		logging.Logger.Error("failed to load simulation config", "error", err)
	}

	exercises, exercisesVersion, err := configs.LoadExercises("configs/exercises.v1.json")
	if err != nil {
		logging.Logger.Error("failed to load exercises config", "error", err)
	}

	intensities, intensitiesVersion, err := configs.LoadIntensities("configs/intensities.v1.json")
	if err != nil {
		logging.Logger.Error("failed to load intensities config", "error", err)
	}

	athleteTypes, athleteTypesVersion, err := configs.LoadAthleteTypes("configs/athlete_types.v1.json")
	if err != nil {
		logging.Logger.Error("failed to load athlete_types config", "error", err)
	}

	logging.Logger.Info(
		"configs loaded",
		"simulation_version", simulationVersion,
		"exercises_version", exercisesVersion,
		"intensities_version", intensitiesVersion,
		"athlete_types_version", athleteTypesVersion,
	)

	// -------------------------
	// Initialize DB
	// -------------------------

	cfg := db.LoadDBConfig()
	logging.Logger.Info("connecting to database...")
	dbConn := db.Connect(cfg.DatabaseURL)

	if dbConn == nil {
		logging.Logger.Error("failed to connect to database")
		os.Exit(1)
	}

	defer dbConn.Close()

	logging.Logger.Info("running migrations...")
	if err := db.RunMigrations(dbConn.Pool()); err != nil {
		logging.Logger.Error("migration failed", "error", err)
		os.Exit(1)
	}

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

	athleteRepo := athlete.NewPostgresRepository(dbConn)

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
		DB:             dbConn,
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
		logging.Logger.Info("server starting", "port", cfg.Port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logging.Logger.Error("server listen failed", "error", err)
		}
	}()

	// -------------------------
	// Graceful shutdown
	// -------------------------

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop

	logging.Logger.Info("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logging.Logger.Error("server shutdown failed", "error", err)
	}

	logging.Logger.Info("server stopped gracefully")
}
