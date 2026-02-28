package routes

import (
	"net/http"

	"github.com/DmytroSobko/FormForgeBackend/internal/app"
)

func registerSimulationRoutes(mux *http.ServeMux, deps app.Dependencies) {

	//mux.HandleFunc("/simulate/week", NewSimulationHandler(deps.Engine).SimulateWeek)
}
