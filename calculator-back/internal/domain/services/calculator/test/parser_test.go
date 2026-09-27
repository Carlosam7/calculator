package test

import (
	"testing"

	"github.com/Carlosam7/calculator/internal/domain/services/calculator"
)

func TestParseValidExpressions(t *testing.T) {
	expressions := []string{
		"10",
		"10 + 5",
		"10 - 5",
		"10 * 5",
		"10 / 5",
		"4 + 5 * 10",
		"(4 + 5) * 10",
		"(10 / 4) * ((4 + 8) / 2)",
		"-5",
		"+5",
		"4 * -2",
		"((10 + 5) * 2) / 5",
	}

	for _, expression := range expressions {
		t.Run(expression, func(t *testing.T) {
			if _, err := calculator.Parse(expression); err != nil {
				t.Fatalf(
					"expected expression %q to be valid, got error: %v",
					expression,
					err,
				)
			}
		})
	}
}

func TestParseInvalidExpressions(t *testing.T) {
	expressions := []string{
		"4 +",
		"+",
		"4 *",
		"4 + * 5",
		"4 5",
		"(4 + 5",
		"4 + 5)",
		"()",
		"(4 + )",
		"((4 + 5)",
		"4 / / 2",
	}

	for _, expression := range expressions {
		t.Run(expression, func(t *testing.T) {
			if _, err := calculator.Parse(expression); err == nil {
				t.Fatalf(
					"expected expression %q to be invalid",
					expression,
				)
			}
		})
	}
}

func TestParseResults(t *testing.T) {
	tests := []struct {
		expression string
		expected   float64
	}{
		{
			expression: "10",
			expected:   10,
		},
		{
			expression: "4 + 5",
			expected:   9,
		},
		{
			expression: "10 - 3",
			expected:   7,
		},
		{
			expression: "4 * 5",
			expected:   20,
		},
		{
			expression: "20 / 4",
			expected:   5,
		},
		{
			expression: "4 + 5 * 10",
			expected:   54,
		},
		{
			expression: "(4 + 5) * 10",
			expected:   90,
		},
		{
			expression: "(10 / 4) * ((4 + 8) / 2)",
			expected:   15,
		},
		{
			expression: "-5 + 10",
			expected:   5,
		},
		{
			expression: "4 * -2",
			expected:   -8,
		},
		{
			expression: "10.5 * 2",
			expected:   21,
		},
	}

	for _, tt := range tests {
		t.Run(tt.expression, func(t *testing.T) {
			got, err := calculator.Parse(tt.expression)

			if err != nil {
				t.Fatalf(
					"unexpected error for %q: %v",
					tt.expression,
					err,
				)
			}

			if got != tt.expected {
				t.Errorf(
					"expected %v, got %v",
					tt.expected,
					got,
				)
			}
		})
	}
}

func TestParseDivisionByZero(t *testing.T) {
	expressions := []string{
		"10 / 0",
		"10 / (5 - 5)",
		"(10 + 5) / 0",
	}

	for _, expression := range expressions {
		t.Run(expression, func(t *testing.T) {
			_, err := calculator.Parse(expression)

			if err == nil {
				t.Fatalf(
					"expected division by zero error for %q",
					expression,
				)
			}
		})
	}
}
