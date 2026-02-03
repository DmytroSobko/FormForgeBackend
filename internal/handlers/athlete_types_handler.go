package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/DmytroSobko/FormForgeBackend/internal/configs"
)

type AthleteTypesConfigHandler struct {
	AthleteTypes *configs.AthleteTypesEnvelope
}

func NewAthleteTypesConfigHandler(
	athleteTypesConfig *configs.AthleteTypesEnvelope,
) *AthleteTypesConfigHandler {
	return &AthleteTypesConfigHandler{
		AthleteTypes: athleteTypesConfig,
	}
}

func (h *AthleteTypesConfigHandler) GetAthleteTypesConfig(
	w http.ResponseWriter,
	r *http.Request,
) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(h.AthleteTypes)
}
