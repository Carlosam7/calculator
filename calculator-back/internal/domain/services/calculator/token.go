package calculator

type TokenType int

const (
	TokenNumber TokenType = iota
	TokenPlus
	TokenMinus
	TokenMultiply
	TokenDivide
	TokenPower
	TokenModulo
	TokenLeftParen
	TokenRightParen
	TokenSqrt
	TokenEOF
)

type Token struct {
	Type  TokenType
	Value string
}
