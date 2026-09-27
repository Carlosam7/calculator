# Calculator API — Project Structure

REST calculator API written in Go with the standard library. No database, no
framework. Evaluates complete mathematical expressions with operator
precedence, parentheses, powers, modulo and square roots.

## Directory structure

```
calculator-back/
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── application/
│   │   └── calculator/
│   │       ├── dto.go
│   │       ├── service.go
│   │       └── test/
│   │           └── (service-level tests, not yet written)
│   ├── domain/
│   │   └── services/
│   │       └── calculator/
│   │           ├── calculator.go
│   │           ├── errors.go
│   │           ├── lexer.go
│   │           ├── parser.go
│   │           ├── token.go
│   │           └── test/
│   │               ├── calculator_test.go
│   │               ├── lexer_test.go
│   │               └── parser_test.go
│   └── interfaces/
│       └── http/
│           ├── handlers/
│           │   ├── calculator_handler.go
│           │   ├── error.go
│           │   └── health_handler.go
│           ├── routes/
│           │   └── routes.go
│           └── test/
│               └── api_test.go
├── go.mod
├── go.sum
└── README.md
```

## Purpose of each folder

### `cmd/server`

The only executable. Reads configuration, wires the layers together, starts the
HTTP server and owns the process lifecycle. Contains no business logic — if a
rule about how to calculate something appears here, it is in the wrong place.

`main.go` builds the handler chain: CORS → routes → handlers → service → domain.

### `internal/domain/services/calculator`

The heart of the application and the only layer that knows what a mathematical
expression is. It has no knowledge of HTTP, JSON or status codes, which is what
makes it testable without a server.

| File | Responsibility |
| --- | --- |
| `token.go` | Defines the token types produced by the lexer (`TokenNumber`, `TokenPlus`, …, `TokenEOF`). |
| `lexer.go` | `Tokenize` converts a raw string into a flat token list. Rejects unknown characters, malformed numbers and identifiers other than `sqrt`. |
| `parser.go` | Recursive-descent parser. One function per precedence level (`parseExpression` → `parseTerm` → `parseFactor` → `parsePower` → `parsePrimary`) so precedence and associativity fall out of the call structure. Returns the numeric result directly. |
| `calculator.go` | `Evaluate`, the public entry point: validates that the expression is not blank, then delegates to the parser. |
| `errors.go` | The four sentinel errors (`ErrInvalidExpression`, `ErrDivisionByZero`, `ErrMissingParen`, `ErrEmptyExpression`) that the HTTP layer translates into response codes. |

### `internal/application/calculator`

The use-case layer. Owns the request and response shapes and the operations the
application exposes.

| File | Responsibility |
| --- | --- |
| `dto.go` | The JSON contract: `EvaluateExpressionRequest`, `EvaluateExpressionResponse`. |
| `service.go` | `Service.Evaluate`, the single use case exposed to the outside world. |

### `internal/interfaces/http`

Everything that speaks HTTP. Translates JSON into a service call and the result
back into JSON. It contains no calculation logic.

| File | Responsibility |
| --- | --- |
| `routes/routes.go` | Registers the URL patterns on a `ServeMux` and returns the handler. |
| `handlers/health_handler.go` | `GET /api/health` → `{"status":"ok","message":"…"}`. |
| `handlers/calculator_handler.go` | `POST /api/calculator/evaluate` → decodes the body, calls the service, maps domain errors to status codes. |
| `handlers/error.go` | Single place that writes the `{"error":{"code","message"}}` envelope. |

## API

| Method | Path | Body | Response |
| --- | --- | --- | --- |
| `GET` | `/api/health` | — | `{"status":"ok","message":"Hello, welcome to your calculator."}` |
| `POST` | `/api/calculator/evaluate` | `{"expression":"4 + 5 * 10"}` | `{"expression":"4 + 5 * 10","result":54}` |

Errors always use the same envelope with HTTP 400:

```json
{"error":{"code":"DIVISION_BY_ZERO","message":"Cannot divide by zero"}}
```

| Code | Cause |
| --- | --- |
| `INVALID_EXPRESSION` | Malformed expression, unknown character, unknown function. |
| `DIVISION_BY_ZERO` | `/` or `%` by zero. |
| `EMPTY_EXPRESSION` | Blank or missing `expression` field. |
| `MISSING_PAREN` | Unbalanced or missing parentheses. |
| `INVALID_REQUEST` | Body is not valid JSON. |
| `INTERNAL_ERROR` | Unexpected server-side failure (HTTP 500). |

## Supported syntax

- Numbers: integers and decimals (`10`, `0.5`)
- Operators: `+` `-` `*` `/` `%` `^`
- Parentheses, nested to any depth
- Function: `sqrt(x)`
- Precedence: `+ -` < `* / %` < unary `-` < `^`; `^` is right-associative

```json
{"expression":"(10 / 4) * ((4 + 8) / 2)"}   // 15
{"expression":"2 ^ 3 ^ 2"}                   // 512  (right associative)
{"expression":"-2 ^ 2"}                      // -4   (unary binds looser than ^)
{"expression":"(-2) ^ 2"}                    // 4
```

## Running

```bash
go run ./cmd/server
```

Listens on `:8080`. CORS is enabled for all origins, so a frontend on
`http://localhost:5173` can call the API directly.

## Tests

```bash
go test ./...
go test -race -coverpkg=./... ./...
```

Tests live in a `test/` subpackage next to the code they exercise and use only
the standard library (`testing`, `net/http/httptest`, `encoding/json`). The HTTP
tests build the real handler chain and drive it with `httptest`, so no server
is started. Because the tests live outside the packages under test, coverage
needs `-coverpkg=./...` to be reported (currently 85.9%).
