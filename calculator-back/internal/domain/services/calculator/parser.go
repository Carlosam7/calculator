package calculator

import (
	"fmt"
	"math"
	"strconv"
)

type Parser struct {
	tokens []Token
	pos    int
}

func newParser(tokens []Token) *Parser {
	return &Parser{
		tokens: tokens,
	}
}

func (p *Parser) current() Token {
	return p.tokens[p.pos]
}

func (p *Parser) advance() {
	if p.pos < len(p.tokens)-1 {
		p.pos++
	}
}

func (p *Parser) parseExpression() (float64, error) {
	result, err := p.parseTerm()
	if err != nil {
		return 0, err
	}

	for {
		switch p.current().Type {
		case TokenPlus:
			p.advance()

			right, err := p.parseTerm()
			if err != nil {
				return 0, err
			}

			result += right

		case TokenMinus:
			p.advance()

			right, err := p.parseTerm()
			if err != nil {
				return 0, err
			}

			result -= right

		default:
			return result, nil
		}
	}
}

func (p *Parser) parseTerm() (float64, error) {
	result, err := p.parseFactor()
	if err != nil {
		return 0, err
	}

	for {
		switch p.current().Type {
		case TokenMultiply:
			p.advance()

			right, err := p.parseFactor()
			if err != nil {
				return 0, err
			}

			result *= right

		case TokenDivide:
			p.advance()

			right, err := p.parseFactor()
			if err != nil {
				return 0, err
			}

			if right == 0 {
				return 0, fmt.Errorf("division by zero")
			}

			result /= right

		case TokenModulo:
			p.advance()

			right, err := p.parseFactor()
			if err != nil {
				return 0, err
			}

			if right == 0 {
				return 0, fmt.Errorf("division by zero")
			}

			result = math.Mod(result, right)

		default:
			return result, nil
		}
	}
}

func (p *Parser) parseFactor() (float64, error) {
	if p.current().Type == TokenPlus {
		p.advance()
		return p.parseFactor()
	}

	if p.current().Type == TokenMinus {
		p.advance()

		result, err := p.parseFactor()
		if err != nil {
			return 0, err
		}

		return -result, nil
	}

	return p.parsePower()
}

func (p *Parser) parsePower() (float64, error) {
	left, err := p.parsePrimary()
	if err != nil {
		return 0, err
	}

	if p.current().Type != TokenPower {
		return left, nil
	}

	p.advance()

	right, err := p.parseFactor()
	if err != nil {
		return 0, err
	}

	return math.Pow(left, right), nil
}

func (p *Parser) parsePrimary() (float64, error) {
	token := p.current()

	switch token.Type {
	case TokenNumber:
		result, err := strconv.ParseFloat(token.Value, 64)
		if err != nil {
			return 0, fmt.Errorf(
				"%w: invalid number %q",
				token.Value,
			)
		}

		p.advance()
		return result, nil

	case TokenSqrt:
		p.advance()

		if p.current().Type != TokenLeftParen {
			return 0, fmt.Errorf(
				"%w: expected opening parenthesis after sqrt",
			)
		}

		p.advance()

		result, err := p.parseExpression()
		if err != nil {
			return 0, err
		}

		if result < 0 {
			return 0, fmt.Errorf(
				"%w: cannot calculate square root of a negative number",
			)
		}

		if p.current().Type != TokenRightParen {
			return 0, fmt.Errorf(
				"%w: expected closing parenthesis after sqrt",
			)
		}

		p.advance()

		return math.Sqrt(result), nil

	case TokenLeftParen:
		p.advance()

		result, err := p.parseExpression()
		if err != nil {
			return 0, err
		}

		if p.current().Type != TokenRightParen {
			return 0, fmt.Errorf(
				"%w: expected closing parenthesis",
			)
		}

		p.advance()
		return result, nil

	default:
		return 0, fmt.Errorf(
			"%w: unexpected token %q",
			token.Value,
		)
	}
}

func Parse(expression string) (float64, error) {
	tokens, err := Tokenize(expression)
	if err != nil {
		return 0, fmt.Errorf("%w: %v", err)
	}

	parser := newParser(tokens)

	result, err := parser.parseExpression()
	if err != nil {
		return 0, err
	}

	if parser.current().Type != TokenEOF {
		return 0, fmt.Errorf(
			"%w: unexpected token %q",
			parser.current().Value,
		)
	}
	return result, nil
}
