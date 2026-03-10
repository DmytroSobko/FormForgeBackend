package handlers

import (
	"net/http"

	"github.com/DmytroSobko/FormForgeBackend/internal/apperror"
	"github.com/DmytroSobko/FormForgeBackend/internal/athlete"
	"github.com/DmytroSobko/FormForgeBackend/internal/http/pagination"
	"github.com/DmytroSobko/FormForgeBackend/internal/http/v1/dto"
	"github.com/DmytroSobko/FormForgeBackend/internal/http/v1/mappers"
	"github.com/DmytroSobko/FormForgeBackend/internal/logging"
)

type AthleteHandler struct {
	service athlete.AthleteService
}

func NewAthleteHandler(service athlete.AthleteService) *AthleteHandler {
	return &AthleteHandler{service: service}
}

func (h *AthleteHandler) HandleAthletes(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.createAthlete(w, r)

	case http.MethodGet:
		h.getAthletes(w, r)

	default:
		WriteAppError(r.Context(), w, apperror.MethodNotAllowed("Method not allowed"))
	}
}

func (h *AthleteHandler) createAthlete(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	logger := logging.FromContext(r.Context())

	logger.Info("create athlete request started")

	var req dto.CreateAthleteRequest

	if err := DecodeJSON(r, &req); err != nil {
		WriteAppError(r.Context(), w, apperror.InvalidRequest(err.Error()))
		return
	}

	if !req.Type.IsValid() {
		WriteAppError(r.Context(), w, apperror.Validation("Invalid athlete type"))
		return
	}

	a, err := h.service.CreateAthlete(r.Context(), req.Type, req.Name)
	if err != nil {
		WriteAppError(r.Context(), w, apperror.Internal(err.Error()))
		return
	}

	logger.Info(
		"athlete created successfully",
		"id", a.ID,
		"name", a.Name,
		"type", a.Type,
	)

	WriteJSON(w, http.StatusCreated, mappers.ToAthleteResponse(a))
}

func (h *AthleteHandler) getAthletes(w http.ResponseWriter, r *http.Request) {
	logger := logging.FromContext(r.Context())

	logger.Info("get athletes request started")
	pagination := pagination.ParsePagination(r)

	athletes, err := h.service.GetAthletes(
		r.Context(),
		pagination.Limit,
		pagination.Offset,
	)

	if err != nil {
		WriteAppError(r.Context(), w, apperror.Internal(err.Error()))
		return
	}

	resp := dto.AthletesResponse{
		Athletes: mappers.ToAthleteResponses(athletes),
	}

	WriteJSON(w, http.StatusOK, resp)
}
