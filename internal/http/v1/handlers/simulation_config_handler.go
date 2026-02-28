package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/DmytroSobko/FormForgeBackend/internal/http/v1/mappers"
	"github.com/DmytroSobko/FormForgeBackend/internal/simulation"
)

type SimulationConfigHandler struct {
	cfg simulation.Config
}

func NewSimulationConfigHandler(
	cfg simulation.Config,
) *SimulationConfigHandler {
	return &SimulationConfigHandler{
		cfg: cfg,
	}
}

func (h *SimulationConfigHandler) GetSimulationConfig(
	w http.ResponseWriter,
	r *http.Request,
) {
	resp := mappers.ToSimulationConfigResponse(h.cfg)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}
