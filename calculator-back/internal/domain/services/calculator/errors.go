package calculator

import "errors"

var (
	ErrInvalidExpression = errors.New("invalid expression")
	ErrDivisionByZero    = errors.New("division by zero")
	ErrMissingParen      = errors.New("missing parenthesis")
	ErrEmptyExpression   = errors.New("empty expression")
)
