package handlers

import (
	"net/http"

	"github.com/DmytroSobko/FormForgeBackend/internal/http/v1/dto"
	"github.com/DmytroSobko/FormForgeBackend/internal/http/v1/mappers"
	"github.com/DmytroSobko/FormForgeBackend/internal/simulation"
)

type ExercisesHandler struct {
	exercises []simulation.Exercise
}

func NewExercisesHandler(exercises []simulation.Exercise) *ExercisesHandler {
	return &ExercisesHandler{exercises: exercises}
}

func (h *ExercisesHandler) HandleExercises(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.getExercises(w, r)
	default:
		WriteError(w, http.StatusMethodNotAllowed, ErrInvalidRequest, "method not allowed")
	}
}

func (h *ExercisesHandler) getExercises(w http.ResponseWriter, _ *http.Request) {
	configs := make([]dto.ExerciseConfig, len(h.exercises))

	for i, e := range h.exercises {
		configs[i] = mappers.ToExerciseConfig(e)
	}

	response := dto.ExerciseConfigsResponse{
		Exercises: configs,
	}

	WriteJSON(w, http.StatusOK, response)
}
