package test

import (
	"testing"

	"github.com/Carlosam7/calculator/internal/domain/services/calculator"
)

func TestTokenize(t *testing.T) {
	tests := []struct {
		name       string
		expression string
		want       []calculator.Token
	}{
		{
			name:       "simple addition",
			expression: "10 + 5",
			want: []calculator.Token{
				{Type: calculator.TokenNumber, Value: "10"},
				{Type: calculator.TokenPlus, Value: "+"},
				{Type: calculator.TokenNumber, Value: "5"},
				{Type: calculator.TokenEOF},
			},
		},
		{
			name:       "parentheses and multiplication",
			expression: "(4 + 5) * 10",
			want: []calculator.Token{
				{Type: calculator.TokenLeftParen, Value: "("},
				{Type: calculator.TokenNumber, Value: "4"},
				{Type: calculator.TokenPlus, Value: "+"},
				{Type: calculator.TokenNumber, Value: "5"},
				{Type: calculator.TokenRightParen, Value: ")"},
				{Type: calculator.TokenMultiply, Value: "*"},
				{Type: calculator.TokenNumber, Value: "10"},
				{Type: calculator.TokenEOF},
			},
		},
		{
			name:       "nested expression",
			expression: "(10 / 4) * ((4 + 8) / 2)",
			want: []calculator.Token{
				{Type: calculator.TokenLeftParen, Value: "("},
				{Type: calculator.TokenNumber, Value: "10"},
				{Type: calculator.TokenDivide, Value: "/"},
				{Type: calculator.TokenNumber, Value: "4"},
				{Type: calculator.TokenRightParen, Value: ")"},
				{Type: calculator.TokenMultiply, Value: "*"},
				{Type: calculator.TokenLeftParen, Value: "("},
				{Type: calculator.TokenLeftParen, Value: "("},
				{Type: calculator.TokenNumber, Value: "4"},
				{Type: calculator.TokenPlus, Value: "+"},
				{Type: calculator.TokenNumber, Value: "8"},
				{Type: calculator.TokenRightParen, Value: ")"},
				{Type: calculator.TokenDivide, Value: "/"},
				{Type: calculator.TokenNumber, Value: "2"},
				{Type: calculator.TokenRightParen, Value: ")"},
				{Type: calculator.TokenEOF},
			},
		},
		{
			name:       "decimal numbers",
			expression: "10.5 * 2.25",
			want: []calculator.Token{
				{Type: calculator.TokenNumber, Value: "10.5"},
				{Type: calculator.TokenMultiply, Value: "*"},
				{Type: calculator.TokenNumber, Value: "2.25"},
				{Type: calculator.TokenEOF},
			},
		},
		{
			name:       "advanced_operations",
			expression: "sqrt(25) ^ 2 % 3",
			want: []calculator.Token{
				{Type: calculator.TokenSqrt, Value: "sqrt"},
				{Type: calculator.TokenLeftParen, Value: "("},
				{Type: calculator.TokenNumber, Value: "25"},
				{Type: calculator.TokenRightParen, Value: ")"},
				{Type: calculator.TokenPower, Value: "^"},
				{Type: calculator.TokenNumber, Value: "2"},
				{Type: calculator.TokenModulo, Value: "%"},
				{Type: calculator.TokenNumber, Value: "3"},
				{Type: calculator.TokenEOF},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := calculator.Tokenize(tt.expression)

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if len(got) != len(tt.want) {
				t.Fatalf(
					"expected %d tokens, got %d",
					len(tt.want),
					len(got),
				)
			}

			for i, token := range got {
				if token != tt.want[i] {
					t.Errorf(
						"token %d: expected %+v, got %+v",
						i,
						tt.want[i],
						token,
					)
				}
			}
		})
	}
}

func TestTokenizeInvalidExpression(t *testing.T) {
	tests := []struct {
		name       string
		expression string
	}{
		{
			name:       "invalid character",
			expression: "4 @ 2",
		},
		{
			name:       "invalid number",
			expression: "10.5.2",
		},
		{
			name:       "letters",
			expression: "hello + 5",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := calculator.Tokenize(tt.expression)

			if err == nil {
				t.Fatalf(
					"expected an error for expression %q",
					tt.expression,
				)
			}
		})
	}
}
