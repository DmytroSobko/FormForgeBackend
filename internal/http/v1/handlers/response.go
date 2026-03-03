package handlers

import (
	"encoding/json"
	"net/http"
)

// WriteJSON writes a JSON response with the given HTTP status code.
func WriteJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(v); err != nil {
		// If JSON encoding fails, fall back to a plain 500 error
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}

func DecodeJSON(r *http.Request, v any) error {
	return json.NewDecoder(r.Body).Decode(v)
}
