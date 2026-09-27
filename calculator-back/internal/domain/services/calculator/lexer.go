package calculator

import (
	"fmt"
	"regexp"
	"unicode"
)

var numberRegex = regexp.MustCompile(`^[0-9]+(\.[0-9]+)?$`)

func Tokenize(expression string) ([]Token, error) {
	var tokens []Token

	for i := 0; i < len(expression); {
		char := rune(expression[i])

		// Ignorar espacios.
		if unicode.IsSpace(char) {
			i++
			continue
		}

		if unicode.IsLetter(char) {
			start := i

			for i < len(expression) && unicode.IsLetter(rune(expression[i])) {
				i++
			}
			value := expression[start:i]
			if value != "sqrt" {
				return nil, fmt.Errorf("invalid identifier: %s", value)
			}

			tokens = append(tokens, Token{
				Type:  TokenSqrt,
				Value: value,
			})
			continue
		}

		switch char {
		case '+':
			tokens = append(tokens, Token{Type: TokenPlus, Value: string(char)})
			i++

		case '-':
			tokens = append(tokens, Token{Type: TokenMinus, Value: string(char)})
			i++

		case '*':
			tokens = append(tokens, Token{Type: TokenMultiply, Value: string(char)})
			i++

		case '/':
			tokens = append(tokens, Token{Type: TokenDivide, Value: string(char)})
			i++

		case '(':
			tokens = append(tokens, Token{Type: TokenLeftParen, Value: string(char)})
			i++

		case ')':
			tokens = append(tokens, Token{Type: TokenRightParen, Value: string(char)})
			i++

		case '^':
			tokens = append(tokens, Token{Type: TokenPower, Value: string(char)})
			i++

		case '%':
			tokens = append(tokens, Token{Type: TokenModulo, Value: string(char)})
			i++

		default:
			if unicode.IsDigit(char) || char == '.' {
				start := i
				for i < len(expression) {
					current := rune(expression[i])
					if unicode.IsDigit(current) || current == '.' {
						i++
						continue
					}
					break
				}
				value := expression[start:i]
				if !numberRegex.MatchString(value) {
					return nil, fmt.Errorf("invalid number: %s", value)
				}

				tokens = append(tokens, Token{
					Type:  TokenNumber,
					Value: value,
				})
				continue
			}
			return nil, fmt.Errorf("invalid character: %q", char)
		}
	}

	tokens = append(tokens, Token{
		Type: TokenEOF,
	})
	return tokens, nil
}
