package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/DmytroSobko/FormForgeBackend/internal/configs"
)

type ExercisesConfigHandler struct {
	Exercises *configs.ExercisesEnvelope
}

func NewExercisesConfigHandler(
	exercisesConfig *configs.ExercisesEnvelope,
) *ExercisesConfigHandler {
	return &ExercisesConfigHandler{
		Exercises: exercisesConfig,
	}
}

func (h *ExercisesConfigHandler) GetExercisesConfig(
	w http.ResponseWriter,
	r *http.Request,
) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(h.Exercises)
}
