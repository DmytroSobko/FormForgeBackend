package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/DmytroSobko/FormForgeBackend/internal/configs"
)

type IntensitiesConfigHandler struct {
	Intensities *configs.IntensitiesEnvelope
}

func NewIntensitiesConfigHandler(
	intensitiesConfig *configs.IntensitiesEnvelope,
) *IntensitiesConfigHandler {
	return &IntensitiesConfigHandler{
		Intensities: intensitiesConfig,
	}
}

func (h *IntensitiesConfigHandler) GetIntensitiesConfig(
	w http.ResponseWriter,
	r *http.Request,
) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(h.Intensities)
}
