package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/DmytroSobko/FormForgeBackend/internal/configs"
)

type SimulationConfigHandler struct {
	Simulation *configs.SimulationConfigEnvelope
}

func NewSimulationConfigHandler(
	simConfig *configs.SimulationConfigEnvelope,
) *SimulationConfigHandler {
	return &SimulationConfigHandler{
		Simulation: simConfig,
	}
}

func (h *SimulationConfigHandler) GetSimulationConfig(
	w http.ResponseWriter,
	r *http.Request,
) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(h.Simulation)
}
