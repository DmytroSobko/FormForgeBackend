package handlers

import (
	"net/http"

	"github.com/DmytroSobko/FormForgeBackend/internal/apperror"
	"github.com/DmytroSobko/FormForgeBackend/internal/http/v1/mappers"
	"github.com/DmytroSobko/FormForgeBackend/internal/simulation"
)

type SimulationConfigHandler struct {
	cfg simulation.Config
}

func NewSimulationConfigHandler(cfg simulation.Config) *SimulationConfigHandler {
	return &SimulationConfigHandler{cfg: cfg}
}

func (h *SimulationConfigHandler) HandleSimulationConfig(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.getSimulationConfig(w, r)

	default:
		w.Header().Set("Allow", http.MethodGet)

		WriteAppError(r.Context(), w, apperror.MethodNotAllowed("Method not allowed"))
	}
}

func (h *SimulationConfigHandler) getSimulationConfig(w http.ResponseWriter, _ *http.Request) {
	response := mappers.ToSimulationConfigResponse(h.cfg)

	WriteJSON(w, http.StatusOK, response)
}
