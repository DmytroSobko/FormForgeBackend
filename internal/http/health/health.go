package health

import (
	"net/http"

	"github.com/DmytroSobko/FormForgeBackend/internal/db"
)

func RegisterHealthRoutes(mux *http.ServeMux, database *db.DB) {
	health := NewHealthHandler(database)

	mux.HandleFunc("/health/live", health.Live)
	mux.HandleFunc("/health/ready", health.Ready)
}
