package calculator

import "strings"

func Evaluate(expression string) (float64, error) {
	if strings.TrimSpace(expression) == "" {
		return 0, ErrEmptyExpression
	}
	return Parse(expression)
}
