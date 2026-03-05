package handlers

import (
	"net/http"

	"github.com/DmytroSobko/FormForgeBackend/internal/apperror"
	"github.com/DmytroSobko/FormForgeBackend/internal/http/v1/dto"
	"github.com/DmytroSobko/FormForgeBackend/internal/http/v1/mappers"
	"github.com/DmytroSobko/FormForgeBackend/internal/simulation"
)

type IntensityConfigsHandler struct {
	intensities []simulation.Intensity
}

func NewIntensitiesHandler(intensities []simulation.Intensity) *IntensityConfigsHandler {
	return &IntensityConfigsHandler{intensities: intensities}
}

func (h *IntensityConfigsHandler) HandleIntensityConfigs(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.getIntensityConfigs(w, r)

	default:
		w.Header().Set("Allow", http.MethodGet)

		WriteAppError(w, apperror.MethodNotAllowed("Method not allowed"))
	}
}

func (h *IntensityConfigsHandler) getIntensityConfigs(w http.ResponseWriter, _ *http.Request) {
	configs := make([]dto.IntensityConfig, len(h.intensities))

	for i, intensity := range h.intensities {
		configs[i] = mappers.ToIntensityConfig(intensity)
	}

	response := dto.IntensityConfigsResponse{
		Intensities: configs,
	}

	WriteJSON(w, http.StatusOK, response)
}
