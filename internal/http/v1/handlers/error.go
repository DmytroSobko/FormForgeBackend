package handlers

import "net/http"

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

const (
	ErrInvalidRequest = "invalid_request"
	ErrInvalidType    = "invalid_type"
	ErrInternal       = "internal_error"
	ErrNotFound       = "not_found"
)

func WriteError(w http.ResponseWriter, status int, code string, message string) {
	WriteJSON(w, status, ErrorResponse{
		Error:   code,
		Message: message,
	})
}
