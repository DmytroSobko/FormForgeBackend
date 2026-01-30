package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/DmytroSobko/FormForgeBackend/internal/configs"
)

type ConfigHandler struct {
	Simulation *configs.SimulationConfigEnvelope
}

func NewConfigHandler(
	simConfig *configs.SimulationConfigEnvelope,
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
