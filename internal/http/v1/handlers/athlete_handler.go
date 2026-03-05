package handlers

import (
	"log"
	"net/http"

	"github.com/DmytroSobko/FormForgeBackend/internal/apperror"
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
		WriteAppError(w, apperror.MethodNotAllowed("Method not allowed"))
	}
}

func (h *AthleteHandler) createAthlete(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	log.Printf("createAthlete started")

	var req dto.CreateAthleteRequest

	if err := DecodeJSON(r, &req); err != nil {
		WriteAppError(w, apperror.InvalidRequest(err.Error()))
		return
	}

	if !req.Type.IsValid() {
		WriteAppError(w, apperror.Validation("Invalid athlete type"))
		return
	}

	a, err := h.service.CreateAthlete(r.Context(), req.Type, req.Name)
	if err != nil {
		WriteAppError(w, apperror.Internal(err.Error()))
		return
	}

	WriteJSON(w, http.StatusCreated, mappers.ToAthleteResponse(a))
}
