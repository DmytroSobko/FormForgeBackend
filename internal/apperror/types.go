package apperror

import "net/http"

func Validation(message string) *AppError {
	return New("validation_error", message, http.StatusBadRequest)
}

func InvalidRequest(message string) *AppError {
	return New("invalid_request", message, http.StatusBadRequest)
}

func Unauthorized(message string) *AppError {
	return New("unauthorized", message, http.StatusUnauthorized)
}

func Forbidden(message string) *AppError {
	return New("forbidden", message, http.StatusForbidden)
}

func NotFound(message string) *AppError {
	return New("not_found", message, http.StatusNotFound)
}

func MethodNotAllowed(message string) *AppError {
	return New("method_not_allowed", message, http.StatusMethodNotAllowed)
}

func Conflict(message string) *AppError {
	return New("conflict", message, http.StatusConflict)
}

func Internal(message string) *AppError {
	return New("internal_error", message, http.StatusInternalServerError)
}

func ServiceUnavailable(message string) *AppError {
	return New("service_unavailable", message, http.StatusServiceUnavailable)
}
