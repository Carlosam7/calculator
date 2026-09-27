package routes

import (
	"net/http"

	"github.com/Carlosam7/calculator/internal/interfaces/http/handlers"
)

func SetupRoutes(
	healthHandler *handlers.HealthHandler,
	calculatorHandler *handlers.CalculatorHandler,
) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/health", healthHandler.Check)
	mux.HandleFunc("POST /api/calculator/evaluate", calculatorHandler.Evaluate)

	return mux
}
