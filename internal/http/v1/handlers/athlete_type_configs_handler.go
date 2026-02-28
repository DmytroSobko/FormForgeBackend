package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/DmytroSobko/FormForgeBackend/internal/athlete"
	"github.com/DmytroSobko/FormForgeBackend/internal/http/v1/dto"
	"github.com/DmytroSobko/FormForgeBackend/internal/http/v1/mappers"
)

type AthleteTypeConfigsHandler struct {
	types []athlete.AthleteTypeConfig
}

func NewAthleteTypeConfigsHandler(
	types []athlete.AthleteTypeConfig,
) *AthleteTypeConfigsHandler {
	return &AthleteTypeConfigsHandler{
		types: types,
	}
}

func (h *AthleteTypeConfigsHandler) GetAthleteTypeConfigs(
	w http.ResponseWriter,
	r *http.Request,
) {
	athleteTypes := make([]dto.AthleteTypeConfig, 0, len(h.types))

	for _, t := range h.types {
		athleteTypes = append(athleteTypes, mappers.ToAthleteTypeConfig(t))
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(dto.AthleteTypeConfigsResponse{
		AthleteTypes: athleteTypes,
	})
}
