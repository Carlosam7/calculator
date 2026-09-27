package calculator

import "fmt"

type Code string

const (
	CodeInvalidExpression Code = "INVALID_EXPRESSION"
	CodeDivisionByZero    Code = "DIVISION_BY_ZERO"
	CodeMathError         Code = "MATH_ERROR"
	CodeInvalidRequest    Code = "INVALID_REQUEST"
)

type Error struct {
	Code    Code
	Message string
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

func newInvalidExpressionError() *Error {
	return &Error{Code: CodeInvalidExpression, Message: "Invalid mathematical expression"}
}

func newDivisionByZeroError() *Error {
	return &Error{Code: CodeDivisionByZero, Message: "Cannot divide by zero"}
}

func newMathError() *Error {
	return &Error{Code: CodeMathError, Message: "Expression cannot be evaluated to a real number"}
}

func NewInvalidRequestError() *Error {
	return &Error{Code: CodeInvalidRequest, Message: "Invalid request body"}
}
