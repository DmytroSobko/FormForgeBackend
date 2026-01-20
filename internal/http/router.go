package http

import (
	"net/http"

	"github.com/DmytroSobko/FormForgeBackend/internal/db"
	"github.com/DmytroSobko/FormForgeBackend/internal/handlers"
)

func NewRouter(database *db.DB) http.Handler {
	mux := http.NewServeMux()

	health := handlers.NewHealthHandler(database)

	mux.HandleFunc("/health/live", health.Live)
	mux.HandleFunc("/health/ready", health.Ready)

	return mux
}
