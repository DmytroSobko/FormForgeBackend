package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/DmytroSobko/FormForgeBackend/internal/config"
)

type ConfigHandler struct {
	Simulation *config.SimulationConfigEnvelope
}

func NewConfigHandler(
	simConfig *config.SimulationConfigEnvelope,
) *ConfigHandler {
	return &ConfigHandler{
		Simulation: simConfig,
	}
}

func (h *ConfigHandler) GetSimulationConfig(
	w http.ResponseWriter,
	r *http.Request,
) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(h.Simulation)
}
