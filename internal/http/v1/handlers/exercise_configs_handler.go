package handlers

import (
	"net/http"

	"github.com/DmytroSobko/FormForgeBackend/internal/apperror"
	"github.com/DmytroSobko/FormForgeBackend/internal/http/v1/dto"
	"github.com/DmytroSobko/FormForgeBackend/internal/http/v1/mappers"
	"github.com/DmytroSobko/FormForgeBackend/internal/simulation"
)

type ExerciseConfigsHandler struct {
	exercises []simulation.Exercise
}

func NewExercisesHandler(exercises []simulation.Exercise) *ExerciseConfigsHandler {
	return &ExerciseConfigsHandler{exercises: exercises}
}

func (h *ExerciseConfigsHandler) HandleExerciseConfigs(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.getExerciseConfigs(w, r)

	default:
		w.Header().Set("Allow", http.MethodGet)

		WriteAppError(w, apperror.MethodNotAllowed("Method not allowed"))
	}
}

func (h *ExerciseConfigsHandler) getExerciseConfigs(w http.ResponseWriter, _ *http.Request) {
	configs := make([]dto.ExerciseConfig, len(h.exercises))

	for i, e := range h.exercises {
		configs[i] = mappers.ToExerciseConfig(e)
	}

	response := dto.ExerciseConfigsResponse{
		Exercises: configs,
	}

	WriteJSON(w, http.StatusOK, response)
}
