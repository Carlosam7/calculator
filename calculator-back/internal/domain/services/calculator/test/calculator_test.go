package test

import (
	"errors"
	"testing"

	"github.com/Carlosam7/calculator/internal/domain/services/calculator"
)

func TestEvaluate(t *testing.T) {
	tests := []struct {
		name       string
		expression string
		expected   float64
	}{
		{
			name:       "addition",
			expression: "4 + 5",
			expected:   9,
		},
		{
			name:       "subtraction",
			expression: "10 - 3",
			expected:   7,
		},
		{
			name:       "multiplication",
			expression: "4 * 5",
			expected:   20,
		},
		{
			name:       "division",
			expression: "20 / 4",
			expected:   5,
		},
		{
			name:       "operator precedence",
			expression: "4 + 5 * 10",
			expected:   54,
		},
		{
			name:       "parentheses",
			expression: "(4 + 5) * 10",
			expected:   90,
		},
		{
			name:       "nested parentheses",
			expression: "(10 / 4) * ((4 + 8) / 2)",
			expected:   15,
		},
		{
			name:       "decimal numbers",
			expression: "0.5 * 4",
			expected:   2,
		},
		{
			name:       "negative number",
			expression: "-5 + 10",
			expected:   5,
		},
		{
			name:       "power",
			expression: "2 ^ 3",
			expected:   8,
		},
		{
			name:       "power precedence",
			expression: "2 + 3 ^ 2",
			expected:   11,
		},
		{
			name:       "power with multiplication",
			expression: "2 ^ 3 * 2",
			expected:   16,
		},
		{
			name:       "right associative power",
			expression: "2 ^ 3 ^ 2",
			expected:   512,
		},
		{
			name:       "modulo",
			expression: "10 % 3",
			expected:   1,
		},
		{
			name:       "modulo with expression",
			expression: "(10 + 5) % 4",
			expected:   3,
		},
		{
			name:       "square root",
			expression: "sqrt(25)",
			expected:   5,
		},
		{
			name:       "square root expression",
			expression: "sqrt(9 + 7)",
			expected:   4,
		},
		{
			name:       "square root with arithmetic",
			expression: "sqrt(16) + 2",
			expected:   6,
		},
		{
			name:       "negative power",
			expression: "-2 ^ 2",
			expected:   -4,
		},
		{
			name:       "parenthesized negative power",
			expression: "(-2) ^ 2",
			expected:   4,
		},
		{
			name:       "negative exponent",
			expression: "2 ^ -2",
			expected:   0.25,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := calculator.Evaluate(tt.expression)

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if got != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, got)
			}
		})
	}
}

func TestEvaluateInvalidExpression(t *testing.T) {
	tests := []struct {
		name          string
		expression    string
		expectedError error
	}{
		{
			name:          "empty expression",
			expression:    "",
			expectedError: calculator.ErrEmptyExpression,
		},
		{
			name:          "whitespace expression",
			expression:    "   ",
			expectedError: calculator.ErrEmptyExpression,
		},
		{
			name:          "missing expression after operator",
			expression:    "4 +",
			expectedError: calculator.ErrInvalidExpression,
		},
		{
			name:          "unexpected operator",
			expression:    "4 + * 5",
			expectedError: calculator.ErrInvalidExpression,
		},
		{
			name:          "missing operator",
			expression:    "4 5",
			expectedError: calculator.ErrInvalidExpression,
		},
		{
			name:          "missing closing parenthesis",
			expression:    "(4 + 5",
			expectedError: calculator.ErrMissingParen,
		},
		{
			name:          "unexpected closing parenthesis",
			expression:    "4 + 5)",
			expectedError: calculator.ErrInvalidExpression,
		},
		{
			name:          "empty parentheses",
			expression:    "()",
			expectedError: calculator.ErrInvalidExpression,
		},
		{
			name:          "invalid identifier",
			expression:    "hello + 5",
			expectedError: calculator.ErrInvalidExpression,
		},
		{
			name:          "invalid number",
			expression:    "10.5.2",
			expectedError: calculator.ErrInvalidExpression,
		},
		{
			name:          "sqrt without argument",
			expression:    "sqrt",
			expectedError: calculator.ErrMissingParen,
		},
		{
			name:          "sqrt with empty parentheses",
			expression:    "sqrt()",
			expectedError: calculator.ErrInvalidExpression,
		},
		{
			name:          "sqrt of negative number",
			expression:    "sqrt(-4)",
			expectedError: calculator.ErrInvalidExpression,
		},
		{
			name:          "missing closing parenthesis in sqrt",
			expression:    "sqrt(25",
			expectedError: calculator.ErrMissingParen,
		},
		{
			name:          "sqrt without parentheses",
			expression:    "sqrt 25",
			expectedError: calculator.ErrMissingParen,
		},
		{
			name:          "unknown function",
			expression:    "unknown(25)",
			expectedError: calculator.ErrInvalidExpression,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := calculator.Evaluate(tt.expression)

			if err == nil {
				t.Fatalf("expected error %v, got nil", tt.expectedError)
			}

			if !errors.Is(err, tt.expectedError) {
				t.Fatalf(
					"expected error %v, got %v",
					tt.expectedError,
					err,
				)
			}
		})
	}
}

func TestEvaluateDivisionByZero(t *testing.T) {
	tests := []string{
		"10 / 0",
		"10 / (5 - 5)",
		"(10 + 5) / 0",
		"10 % 0",
		"10 % (5 - 5)",
	}

	for _, expression := range tests {
		t.Run(expression, func(t *testing.T) {
			_, err := calculator.Evaluate(expression)

			if !errors.Is(err, calculator.ErrDivisionByZero) {
				t.Fatalf(
					"expected ErrDivisionByZero, got %v",
					err,
				)
			}
		})
	}
}
