package routes

import (
	"net/http"

	"github.com/DmytroSobko/FormForgeBackend/internal/app"
	"github.com/DmytroSobko/FormForgeBackend/internal/http/v1/handlers"
)

func registerConfigRoutes(mux *http.ServeMux, deps app.Dependencies) {

	mux.HandleFunc("/config/simulation",
		handlers.NewSimulationConfigHandler(deps.SimConfig).GetSimulationConfig)

	mux.HandleFunc("/config/athleteTypes",
		handlers.NewAthleteTypeConfigsHandler(deps.AthleteTypes).GetAthleteTypeConfigs)

	mux.HandleFunc("/config/exercises",
		handlers.NewExercisesHandler(deps.Exercises).GetExercises)

	mux.HandleFunc("/config/intensities",
		handlers.NewIntensitiesHandler(deps.Intensities).GetIntensities)
}
