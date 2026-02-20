package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/DmytroSobko/FormForgeBackend/internal/models"
	"github.com/DmytroSobko/FormForgeBackend/internal/simulation"
)

type SimulationHandler struct {
	engine *simulation.Engine
}

func NewSimulationHandler(
	engine *simulation.Engine,
) *SimulationHandler {
	return &SimulationHandler{
		engine: engine,
	}
}

func (h *SimulationHandler) SimulateWeek(
	w http.ResponseWriter,
	r *http.Request,
) {
	var req struct {
		Athlete models.Athlete      `json:"athlete"`
		Plan    models.TrainingPlan `json:"plan"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	result, err := h.engine.SimulateWeek(req.Athlete, req.Plan)

	if err != nil {
		http.Error(w, "simulation failed", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
