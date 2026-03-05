package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
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

func DecodeJSON(r *http.Request, dst any) error {
	if r.Body == nil {
		return errors.New("request body must not be empty")
	}

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(dst)
	if err != nil {

		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError

		switch {
		case errors.As(err, &syntaxError):
			return fmt.Errorf("malformed JSON at position %d", syntaxError.Offset)

		case errors.Is(err, io.ErrUnexpectedEOF):
			return errors.New("malformed JSON")

		case errors.As(err, &unmarshalTypeError):
			return fmt.Errorf(
				"invalid value for field '%s' (expected %s)",
				unmarshalTypeError.Field,
				unmarshalTypeError.Type,
			)

		case strings.HasPrefix(err.Error(), "json: unknown field "):
			field := strings.TrimPrefix(err.Error(), "json: unknown field ")
			return fmt.Errorf("unknown field %s", field)

		case errors.Is(err, io.EOF):
			return errors.New("request body must not be empty")

		default:
			return err
		}
	}

	// Ensure only one JSON object is sent
	if decoder.Decode(&struct{}{}) != io.EOF {
		return errors.New("request body must contain only a single JSON object")
	}

	return nil
}
