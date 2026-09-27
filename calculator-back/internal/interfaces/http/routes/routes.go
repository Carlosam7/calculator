package routes

import (
	"net/http"

	"github.com/Carlosam7/calculator/internal/interfaces/http/handlers"
)

func SetupRoutes(healthHandler *handlers.HealthHandler) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/v1/health", healthHandler.Check)

	return mux
}
