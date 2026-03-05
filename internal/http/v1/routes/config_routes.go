package routes

import (
	"net/http"

	"github.com/DmytroSobko/FormForgeBackend/internal/app"
	"github.com/DmytroSobko/FormForgeBackend/internal/http/v1/handlers"
)

func registerConfigRoutes(mux *http.ServeMux, deps app.Dependencies) {

	mux.HandleFunc("/config/simulation",
		handlers.NewSimulationConfigHandler(deps.SimConfig).HandleSimulationConfig)

	mux.HandleFunc("/config/athleteTypes",
		handlers.NewAthleteTypeConfigsHandler(deps.AthleteTypes).HandleAthleteTypeConfigs)

	mux.HandleFunc("/config/exercises",
		handlers.NewExercisesHandler(deps.Exercises).HandleExerciseConfigs)

	mux.HandleFunc("/config/intensities",
		handlers.NewIntensitiesHandler(deps.Intensities).HandleIntensityConfigs)
}
