package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/DmytroSobko/FormForgeBackend/internal/athlete"
	"github.com/DmytroSobko/FormForgeBackend/internal/http/v1/dto"
	"github.com/DmytroSobko/FormForgeBackend/internal/http/v1/mappers"
)

type AthleteHandler struct {
	service *athlete.Service
}

func NewAthleteHandler(service *athlete.Service) *AthleteHandler {
	return &AthleteHandler{service: service}
}

func (h *AthleteHandler) CreateAthlete(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var req dto.CreateAthleteRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	athleteType := athlete.AthleteType(req.Type)

	if !athleteType.IsValid() {
		http.Error(w, "invalid athlete type", http.StatusBadRequest)
		return
	}

	a, err := h.service.CreateAthlete(r.Context(), athleteType, req.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp := mappers.ToAthleteResponse(a)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(resp)
}

func (h *AthleteHandler) HandleAthletes(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.CreateAthlete(w, r)
	// case http.MethodGet:
	// 	h.GetAthletes(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
