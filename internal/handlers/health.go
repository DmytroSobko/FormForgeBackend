package handlers

import (
	"net/http"

	"github.com/DmytroSobko/FormForgeBackend/internal/db"
)

type HealthHandler struct {
	DB *db.DB
}

func NewHealthHandler(database *db.DB) *HealthHandler {
	return &HealthHandler{DB: database}
}

// Liveness probe
func (h *HealthHandler) Live(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("alive"))
}

// Readiness probe
func (h *HealthHandler) Ready(w http.ResponseWriter, r *http.Request) {
	if !h.DB.IsHealthy() {
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte("not ready"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ready"))
}