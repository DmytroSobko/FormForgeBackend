package http

import (
	"net/http"

	"github.com/DmytroSobko/FormForgeBackend/internal/configs"
	"github.com/DmytroSobko/FormForgeBackend/internal/db"
	"github.com/DmytroSobko/FormForgeBackend/internal/handlers"
	"github.com/DmytroSobko/FormForgeBackend/internal/simulation"
)

func NewRouter(
	database *db.DB,
	simConfig *configs.SimulationConfigEnvelope,
	athleteTypesConfig *configs.AthleteTypesEnvelope,
	exercisesConfig *configs.ExercisesEnvelope,
	intensitiesConfig *configs.IntensitiesEnvelope,
	engine *simulation.Engine,
) http.Handler {

	mux := http.NewServeMux()

	health := handlers.NewHealthHandler(database)

	mux.HandleFunc("/health/live", health.Live)
	mux.HandleFunc("/health/ready", health.Ready)

	simulationConfigHandler := handlers.NewSimulationConfigHandler(simConfig)
	mux.HandleFunc("/api/config/simulation", simulationConfigHandler.GetSimulationConfig)

	athleteTypesConfigHandler := handlers.NewAthleteTypesConfigHandler(athleteTypesConfig)
	mux.HandleFunc("/api/config/athleteTypes", athleteTypesConfigHandler.GetAthleteTypesConfig)

	exercisesConfigHandler := handlers.NewExercisesConfigHandler(exercisesConfig)
	mux.HandleFunc("/api/config/exercises", exercisesConfigHandler.GetExercisesConfig)

	intensitiesConfigHandler := handlers.NewIntensitiesConfigHandler(intensitiesConfig)
	mux.HandleFunc("/api/config/intensities", intensitiesConfigHandler.GetIntensitiesConfig)

	simulationHandler := handlers.NewSimulationHandler(engine)
	mux.HandleFunc("/api/simulate/week", simulationHandler.SimulateWeek)

	handler := LoggingMiddleware(mux)
	handler = RecoveryMiddleware(handler)
	return handler
}
