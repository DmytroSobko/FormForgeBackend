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
	engine *simulation.Engine,
) http.Handler {

	mux := http.NewServeMux()

	health := handlers.NewHealthHandler(database)

	mux.HandleFunc("/health/live", health.Live)
	mux.HandleFunc("/health/ready", health.Ready)

	configHandler := handlers.NewConfigHandler(simConfig)

	mux.HandleFunc(
		"/api/config/simulation",
		configHandler.GetSimulationConfig,
	)

	simHandler := handlers.NewSimulationHandler(engine)
	mux.HandleFunc(
		"/api/simulate/week",
		simHandler.SimulateWeek,
	)

	handler := LoggingMiddleware(mux)
	handler = RecoveryMiddleware(handler)
	return handler
}
