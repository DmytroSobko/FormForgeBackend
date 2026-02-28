package routes

import (
	"net/http"

	"github.com/DmytroSobko/FormForgeBackend/internal/app"
)

func RegisterRoutes(rootMux *http.ServeMux, deps app.Dependencies) {
	v1Mux := http.NewServeMux()

	registerConfigRoutes(v1Mux, deps)
	registerSimulationRoutes(v1Mux, deps)
	registerAthleteRoutes(v1Mux, deps)

	rootMux.Handle("/api/v1/", http.StripPrefix("/api/v1", v1Mux))
}
