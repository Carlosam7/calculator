package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	appcalculator "github.com/Carlosam7/calculator/internal/application/calculator"
	domaincalculator "github.com/Carlosam7/calculator/internal/domain/services/calculator"
)

type CalculatorHandler struct {
	service *appcalculator.Service
}

func NewCalculatorHandler(service *appcalculator.Service) *CalculatorHandler {
	return &CalculatorHandler{
		service: service,
	}
}

func (h *CalculatorHandler) Evaluate(w http.ResponseWriter, r *http.Request) {
	var request appcalculator.EvaluateExpressionRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		writeError(w, http.StatusBadRequest, "INVALID_REQUEST", "invalid JSON body")
		return
	}

	result, err := h.service.Evaluate(request.Expression)
	if err != nil {
		if errors.Is(err, domaincalculator.ErrDivisionByZero) {
			writeError(w, http.StatusBadRequest, "DIVISION_BY_ZERO", "Cannot divide by zero")
			return
		}

		if errors.Is(err, domaincalculator.ErrInvalidExpression) {
			writeError(w, http.StatusBadRequest, "INVALID_EXPRESSION", "Invalid mathematical expression")
			return
		}

		if errors.Is(err, domaincalculator.ErrEmptyExpression) {
			writeError(w, http.StatusBadRequest, "EMPTY_EXPRESSION", "Empty expression")
			return
		}

		if errors.Is(err, domaincalculator.ErrMissingParen) {
			writeError(w, http.StatusBadRequest, "MISSING_PAREN", "Missing parentheses")
			return
		}

		writeError(w, http.StatusInternalServerError, "INTERNAL_ERROR", "An unexpected error occurred")
		return

	}

	response := appcalculator.EvaluateExpressionResponse{
		Expression: request.Expression,
		Result:     result,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w).Encode(response)
}
