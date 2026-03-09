package handlers

import (
	"context"
	"errors"
	"net/http"

	"github.com/DmytroSobko/FormForgeBackend/internal/apperror"
	"github.com/DmytroSobko/FormForgeBackend/internal/logging"
)

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

func WriteAppError(ctx context.Context, w http.ResponseWriter, err error) {

	logger := logging.FromContext(ctx)

	var appErr *apperror.AppError

	if errors.As(err, &appErr) {

		WriteJSON(w, appErr.StatusCode, ErrorResponse{
			Error:   appErr.Code,
			Message: appErr.Message,
		})

		return
	}

	logger.Error(
		"internal error",
		"error", err,
	)

	WriteJSON(w, http.StatusInternalServerError, ErrorResponse{
		Error:   "internal_error",
		Message: "Internal server error",
	})
}
