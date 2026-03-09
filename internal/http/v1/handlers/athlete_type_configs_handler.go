package handlers

import (
	"net/http"

	"github.com/DmytroSobko/FormForgeBackend/internal/apperror"
	"github.com/DmytroSobko/FormForgeBackend/internal/athlete"
	"github.com/DmytroSobko/FormForgeBackend/internal/http/v1/dto"
	"github.com/DmytroSobko/FormForgeBackend/internal/http/v1/mappers"
)

type AthleteTypeConfigsHandler struct {
	types []athlete.AthleteTypeConfig
}

func NewAthleteTypeConfigsHandler(types []athlete.AthleteTypeConfig) *AthleteTypeConfigsHandler {
	return &AthleteTypeConfigsHandler{types: types}
}

func (h *AthleteTypeConfigsHandler) HandleAthleteTypeConfigs(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.getAthleteTypeConfigs(w, r)

	default:
		w.Header().Set("Allow", http.MethodGet)

		WriteAppError(r.Context(), w, apperror.MethodNotAllowed("Method not allowed"))
	}
}

func (h *AthleteTypeConfigsHandler) getAthleteTypeConfigs(w http.ResponseWriter, _ *http.Request) {
	configs := make([]dto.AthleteTypeConfig, len(h.types))

	for i, t := range h.types {
		configs[i] = mappers.ToAthleteTypeConfig(t)
	}

	response := dto.AthleteTypeConfigsResponse{
		AthleteTypes: configs,
	}

	WriteJSON(w, http.StatusOK, response)
}
