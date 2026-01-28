package http

import (
	"log"
	"net/http"
	"runtime/debug"
)

// RecoveryMiddleware prevents panics from crashing the server
func RecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf(
					"PANIC recovered: %v\n%s",
					err,
					debug.Stack(),
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
