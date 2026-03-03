package handlers

import (
	"net/http"

	"github.com/DmytroSobko/FormForgeBackend/internal/http/v1/dto"
	"github.com/DmytroSobko/FormForgeBackend/internal/http/v1/mappers"
	"github.com/DmytroSobko/FormForgeBackend/internal/simulation"
)

type IntensitiesHandler struct {
	intensities []simulation.Intensity
}

func NewIntensitiesHandler(intensities []simulation.Intensity) *IntensitiesHandler {
	return &IntensitiesHandler{intensities: intensities}
}

func (h *IntensitiesHandler) HandleIntensities(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.getIntensities(w, r)
	default:
		WriteError(w, http.StatusMethodNotAllowed, ErrInvalidRequest, "method not allowed")
	}
}

func (h *IntensitiesHandler) getIntensities(w http.ResponseWriter, _ *http.Request) {
	configs := make([]dto.IntensityConfig, len(h.intensities))

	for i, intensity := range h.intensities {
		configs[i] = mappers.ToIntensityConfig(intensity)
	}

	response := dto.IntensityConfigsResponse{
		Intensities: configs,
	}

	WriteJSON(w, http.StatusOK, response)
}
