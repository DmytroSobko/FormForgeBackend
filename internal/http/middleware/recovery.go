package middleware

import (
	"net/http"
	"runtime/debug"

	"github.com/DmytroSobko/FormForgeBackend/internal/logging"
)

func RecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		defer func() {
			if err := recover(); err != nil {

				logger := logging.FromContext(r.Context())

				logger.Error(
					"panic recovered",
					"error", err,
					"stack", string(debug.Stack()),
				)

				http.Error(
					w,
					http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError,
				)
			}
		}()

		next.ServeHTTP(w, r)
	})
}
