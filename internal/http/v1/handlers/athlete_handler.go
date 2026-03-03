package handlers

import (
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

func (h *AthleteHandler) HandleAthletes(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.createAthlete(w, r)
	default:
		WriteError(w, http.StatusMethodNotAllowed, ErrInvalidRequest, "method not allowed")
	}
}

func (h *AthleteHandler) createAthlete(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var req dto.CreateAthleteRequest
	if err := DecodeJSON(r, &req); err != nil {
		WriteError(w, http.StatusBadRequest, ErrInvalidRequest, "invalid request body")
		return
	}

	athleteType := athlete.AthleteType(req.Type)
	if !athleteType.IsValid() {
		WriteError(w, http.StatusBadRequest, ErrInvalidType, "invalid athlete type")
		return
	}

	a, err := h.service.CreateAthlete(r.Context(), athleteType, req.Name)
	if err != nil {
		WriteError(w, http.StatusBadRequest, ErrInvalidRequest, err.Error())
		return
	}

	WriteJSON(w, http.StatusCreated, mappers.ToAthleteResponse(a))
}
