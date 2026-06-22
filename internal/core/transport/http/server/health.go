package core_http_server

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	core_postgres_pool "github.com/Mertvyki/book-shop/internal/core/repository/postgres/pool"
)

type HealthHandler struct {
	pool core_postgres_pool.Pool
}

func NewHealthHandler(pool core_postgres_pool.Pool) *HealthHandler {
	return &HealthHandler{pool: pool}
}

type healthResponse struct {
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
	Database  string `json:"database"`
}

func (h *HealthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	dbStatus := "ok"
	if err := h.pool.Ping(ctx); err != nil {
		dbStatus = "error: " + err.Error()
	}

	statusCode := http.StatusOK
	overallStatus := "ok"
	if dbStatus != "ok" {
		statusCode = http.StatusServiceUnavailable
		overallStatus = "degraded"
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(healthResponse{
		Status:    overallStatus,
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Database:  dbStatus,
	})
}
