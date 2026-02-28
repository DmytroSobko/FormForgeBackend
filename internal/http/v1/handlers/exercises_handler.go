package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/DmytroSobko/FormForgeBackend/internal/http/v1/dto"
	"github.com/DmytroSobko/FormForgeBackend/internal/http/v1/mappers"
	"github.com/DmytroSobko/FormForgeBackend/internal/simulation"
)

type ExercisesHandler struct {
	exercises []simulation.Exercise
}

func NewExercisesHandler(
	exercises []simulation.Exercise,
) *ExercisesHandler {
	return &ExercisesHandler{
		exercises: exercises,
	}
}

func (h *ExercisesHandler) GetExercises(
	w http.ResponseWriter,
	r *http.Request,
) {

	exercises := make([]dto.ExerciseConfig, 0, len(h.exercises))

	for _, e := range h.exercises {
		exercises = append(exercises, mappers.ToExerciseConfig(e))
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(dto.ExerciseConfigsResponse{
		Exercises: exercises,
	})
}
