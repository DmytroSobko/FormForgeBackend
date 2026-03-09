package middleware

import (
	"net/http"

	"github.com/DmytroSobko/FormForgeBackend/internal/logging"
	"github.com/google/uuid"
)

func RequestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		requestID := uuid.NewString()

		w.Header().Set("X-Request-ID", requestID)

		// attach request-scoped logger
		reqLogger := logging.Logger.With(
			"request_id", requestID,
			"method", r.Method,
			"path", r.URL.Path,
		)

		ctx := logging.WithLogger(r.Context(), reqLogger)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
