package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/DmytroSobko/FormForgeBackend/internal/http/v1/dto"
	"github.com/DmytroSobko/FormForgeBackend/internal/http/v1/mappers"
	"github.com/DmytroSobko/FormForgeBackend/internal/simulation"
)

type IntensitiesHandler struct {
	intensities []simulation.Intensity
}

func NewIntensitiesHandler(
	intensities []simulation.Intensity,
) *IntensitiesHandler {
	return &IntensitiesHandler{
		intensities: intensities,
	}
}

func (h *IntensitiesHandler) GetIntensities(
	w http.ResponseWriter,
	r *http.Request,
) {
	intensities := make([]dto.IntensityConfig, 0, len(h.intensities))

	for _, i := range h.intensities {
		intensities = append(intensities, mappers.ToIntensityConfig(i))
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(dto.IntensityConfigsResponse{
		Intensities: intensities,
	})
}
