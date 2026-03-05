package handlers

import (
	"errors"
	"log"
	"net/http"

	"github.com/DmytroSobko/FormForgeBackend/internal/apperror"
)

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

func WriteAppError(w http.ResponseWriter, err error) {

	var appErr *apperror.AppError

	if errors.As(err, &appErr) {

		WriteJSON(w, appErr.StatusCode, ErrorResponse{
			Error:   appErr.Code,
			Message: appErr.Message,
		})

		return
	}

	// fallback for unexpected errors
	log.Printf("internal error: %v", err)

	WriteJSON(w, http.StatusInternalServerError, ErrorResponse{
		Error:   "internal_error",
		Message: "Internal server error",
	})
}
