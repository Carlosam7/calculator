# Development Prompts

This file contains the main prompts used during the development of the Calculator project.

## 1. Backend Implementation

> We are going to build the backend of a calculator application for a technical test for a Software Engineer Intern position Required stack:

- Go 1.27+
- REST API
- JSON
- No database
- No heavy web framework; use net/http from the standard library
- Idiomatic Go code
- Clean and simple architecture
- Unit tests
- CORS to allow the local frontend

The frontend must NOT evaluate mathematical expressions on its own.

The API must expose:

GET /api/v1/health

Response:
{
“status”: “ok”
}

POST /api/v1/calculator/evaluate

Request:
{
“expression”: “4 + 5 _ 10”
}
Response:
{
“expression”: “4 + 5 _ 10”,
“result”: 54
}

Errors:
HTTP 400

{
“error”: {
“code”: “DIVISION_BY_ZERO”,
“message”: “Cannot divide by zero”
}
}

or:

{
“error”: {
“code”: “INVALID_EXPRESSION”,
“message”: “Invalid mathematical expression”
}
}

The implementation will support complete expressions, not isolated operations.

Don’t write any code yet. I’m just giving you some context.

> Implement only the lexer/tokenizer for the calculator domain. You must define: type TokenType int with the following tokens:
> TokenNumber
> TokenPlus
> TokenMinus
> TokenMultiply
> TokenDivide
> TokenPower
> TokenModulo
> TokenLeftParen
> TokenRightParen
> TokenSqrt
> TokenEOF

> Implement only the mathematical expression parser.
> You must use the tokens created by the lexer.
> I want a top-down recursive parser.
> The grammar should be:
> expression := term ((“+” | “-”) term)_
> term := factor ((“_” | “/” | “%”) factor)\*
> factor := (“+” | “-”) factor

        | power

power := primary
| primary “^” factor
primary := number
| “sqrt” “(” expression “)”
| “(” expression “)”

The hierarchy must produce the following order of operations:

1. Parentheses / primary
2. Exponentiation
3. Unary operator
4. Multiplication, division, and modulo
5. Addition and subtraction

> Include input validation and appropriate error handling.

> Add unit tests covering valid expressions, invalid expressions, and edge cases.

## 2. Frontend Implementation

> Implement the view calculator from this design: https://www.figma.com/design/gbn4KMl8xiC4ebmQ9taeAU/Untitled?node-id=0-1&t=6T2yXo2HnbmmGQW2-1 using React and TypeScript.
>
> The display must show the expression as the user builds it instead of replacing it with zero after pressing an operator.
>
> For example:
>
> ```text
> 5
> 5 +
> 5 + 3
> 5 + 3 ×
> 5 + 3 × 2
> ```

## 3. Frontend Unit Tests

> Add unit and component tests for the calculator frontend using Vitest and React Testing Library.
>
> Test:
>
> - Calculator rendering.
> - Number input.
> - Operators.
> - Decimal input.
> - Clear functionality.
> - Expression building.
> - Calculation behavior.
> - API/service behavior.
> - Calculator history behavior.
> - User interactions.
>
> Existing behavior must remain unchanged.
>
> Run the complete test suite and verify that all tests pass.

## 5. Backend Unit Tests

> Add comprehensive unit tests for the Go calculator backend.
>
> Cover:
>
> - Basic arithmetic operations.
> - Operator precedence.
> - Decimal values.
> - Invalid expressions.
> - Division by zero.
> - Edge cases.
> - HTTP/API behavior where applicable.
>
> Use Go's standard testing framework.
>
> Run:
>
> ```bash
> go test ./...
> ```
>
> and make sure the test suite passes.

## 6. Dockerization

> Dockerize the calculator application.
>
> The project contains:
>
> - A React/Vite frontend.
> - A Go backend.
>
> Create separate Dockerfiles for the frontend and backend and a Docker Compose configuration to run both services together.
>
> Requirements:
>
> - Use multi-stage builds.
> - Use Node.js for the frontend build.
> - Serve the production frontend using Nginx.
> - Compile the Go backend into a production-ready binary.
> - Use a lightweight runtime image for the backend.
> - Expose the frontend on port 3000 and the backend on port 8080.
> - Configure the frontend to communicate with the backend through an environment variable.
>
> Make sure the configuration works with:
>
> ```bash
> docker compose up --build
> ```

## 7. Testing and Verification

> Review the complete Calculator project and verify that the frontend, backend, API integration, tests, and Docker configuration work together correctly.
>
> Run the relevant test and build commands:
>
> ```bash
> pnpm test
> pnpm build
> go test ./...
> go build ./...
> docker compose build
> ```
>
> Identify any failures and fix only the issues related to the current implementation.

## 8. Documentation

> Create a professional README for the Calculator project.
>
> The README must include:
>
> - Project overview.
> - Technologies used.
> - Project structure.
> - Setup instructions.
> - How to run the frontend.
> - How to run the backend.
> - How to run the complete application with Docker Compose.
> - Environment variables.
> - Examples of REST API calls.
> - Design decisions and assumptions.
> - Testing instructions.
> - Current project status.
