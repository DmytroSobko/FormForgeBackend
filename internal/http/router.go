package http

import (
	"net/http"

	"github.com/DmytroSobko/FormForgeBackend/internal/app"
	"github.com/DmytroSobko/FormForgeBackend/internal/http/health"
	"github.com/DmytroSobko/FormForgeBackend/internal/http/middleware"
	"github.com/DmytroSobko/FormForgeBackend/internal/http/v1/routes"
)

func NewRouter(deps app.Dependencies) http.Handler {
	rootMux := http.NewServeMux()

	health.RegisterHealthRoutes(rootMux, deps.DB)

	routes.RegisterRoutes(rootMux, deps)

	handler := middleware.Chain(
		rootMux,
		middleware.RecoveryMiddleware,
		middleware.RequestIDMiddleware,
		middleware.LoggingMiddleware,
	)

	return handler
}
