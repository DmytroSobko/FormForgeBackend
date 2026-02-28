package routes

import (
	"net/http"

	"github.com/DmytroSobko/FormForgeBackend/internal/app"
	"github.com/DmytroSobko/FormForgeBackend/internal/http/v1/handlers"
)

func registerAthleteRoutes(mux *http.ServeMux, deps app.Dependencies) {

	mux.HandleFunc("/athletes",
		handlers.NewAthleteHandler(deps.AthleteService).HandleAthletes)
}
