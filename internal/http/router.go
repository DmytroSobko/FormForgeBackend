package http

import (
	"net/http"

	"github.com/DmytroSobko/FormForgeBackend/internal/app"
	"github.com/DmytroSobko/FormForgeBackend/internal/http/v1/routes"
)

func NewRouter(deps app.Dependencies) http.Handler {
	rootMux := http.NewServeMux()

	registerHealthRoutes(rootMux, deps.DB)

	routes.RegisterRoutes(rootMux, deps)

	handler := LoggingMiddleware(rootMux)
	handler = RecoveryMiddleware(handler)

	return handler
}
