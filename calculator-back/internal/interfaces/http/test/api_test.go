package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"sort"
	"strings"
	"testing"

	appcalculator "github.com/Carlosam7/calculator/internal/application/calculator"
	"github.com/Carlosam7/calculator/internal/interfaces/http/handlers"
	"github.com/Carlosam7/calculator/internal/interfaces/http/routes"
)

const (
	healthPath      = "/api/health"
	evaluatePath    = "/api/calculator/evaluate"
	contentTypeJSON = "application/json"
)

type healthResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type evaluateResponse struct {
	Expression string  `json:"expression"`
	Result     float64 `json:"result"`
}

type errorResponse struct {
	Error struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

func newTestHandler() http.Handler {
	healthHandler := handlers.NewHealthHandler()
	calculatorHandler := handlers.NewCalculatorHandler(appcalculator.NewService())

	return routes.SetupRoutes(healthHandler, calculatorHandler)
}

func doRequest(t *testing.T, method string, path string, body string) *httptest.ResponseRecorder {
	t.Helper()

	request := httptest.NewRequest(method, path, strings.NewReader(body))
	request.Header.Set("Content-Type", contentTypeJSON)

	recorder := httptest.NewRecorder()
	newTestHandler().ServeHTTP(recorder, request)

	return recorder
}

func assertStatus(t *testing.T, recorder *httptest.ResponseRecorder, want int) {
	t.Helper()

	if recorder.Code != want {
		t.Errorf(
			"expected status %d, got %d (body: %s)",
			want,
			recorder.Code,
			recorder.Body.String(),
		)
	}
}

func assertContentType(t *testing.T, recorder *httptest.ResponseRecorder) {
	t.Helper()

	got := recorder.Header().Get("Content-Type")
	if got != contentTypeJSON {
		t.Errorf("expected Content-Type %q, got %q", contentTypeJSON, got)
	}
}

func assertJSONKeys(t *testing.T, recorder *httptest.ResponseRecorder, want ...string) {
	t.Helper()

	var decoded map[string]any
	if err := json.Unmarshal(recorder.Body.Bytes(), &decoded); err != nil {
		t.Fatalf("expected a JSON object body, got error: %v (body: %s)", err, recorder.Body.String())
	}

	got := make([]string, 0, len(decoded))
	for key := range decoded {
		got = append(got, key)
	}
	sort.Strings(got)

	sort.Strings(want)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("expected JSON keys %v, got %v (body: %s)", want, got, recorder.Body.String())
	}
}

func assertErrorCode(t *testing.T, recorder *httptest.ResponseRecorder, want string) {
	t.Helper()

	var response errorResponse
	if err := json.Unmarshal(recorder.Body.Bytes(), &response); err != nil {
		t.Fatalf("expected a JSON error body, got error: %v (body: %s)", err, recorder.Body.String())
	}

	if response.Error.Code != want {
		t.Errorf("expected error code %q, got %q (body: %s)", want, response.Error.Code, recorder.Body.String())
	}

	if response.Error.Message == "" {
		t.Errorf("expected a non-empty error message, got %q (body: %s)", response.Error.Message, recorder.Body.String())
	}
}

func TestHealth(t *testing.T) {
	recorder := doRequest(t, http.MethodGet, healthPath, "")

	assertStatus(t, recorder, http.StatusOK)
	assertContentType(t, recorder)
	assertJSONKeys(t, recorder, "status", "message")

	var response healthResponse
	if err := json.Unmarshal(recorder.Body.Bytes(), &response); err != nil {
		t.Fatalf("expected a JSON body, got error: %v (body: %s)", err, recorder.Body.String())
	}

	if response.Status != "ok" {
		t.Errorf("expected status %q, got %q", "ok", response.Status)
	}
}

func TestEvaluate(t *testing.T) {
	tests := []struct {
		name       string
		expression string
		expected   float64
	}{
		{
			name:       "operator precedence",
			expression: "4 + 5 * 10",
			expected:   54,
		},
		{
			name:       "parentheses",
			expression: "(4 + 5) * 10",
			expected:   90,
		},
		{
			name:       "square root with arithmetic",
			expression: "sqrt(25) + 2",
			expected:   7,
		},
		{
			name:       "right associative power",
			expression: "2 ^ 3 ^ 2",
			expected:   512,
		},
		{
			name:       "addition",
			expression: "4 + 5",
			expected:   9,
		},
		{
			name:       "modulo",
			expression: "10 % 3",
			expected:   1,
		},
		{
			name:       "negative exponent",
			expression: "2 ^ -2",
			expected:   0.25,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, err := json.Marshal(appcalculator.EvaluateExpressionRequest{
				Expression: tt.expression,
			})
			if err != nil {
				t.Fatalf("unexpected error building the request: %v", err)
			}

			recorder := doRequest(t, http.MethodPost, evaluatePath, string(body))

			assertStatus(t, recorder, http.StatusOK)
			assertContentType(t, recorder)
			assertJSONKeys(t, recorder, "expression", "result")

			var response evaluateResponse
			if err := json.Unmarshal(recorder.Body.Bytes(), &response); err != nil {
				t.Fatalf("expected a JSON body, got error: %v (body: %s)", err, recorder.Body.String())
			}

			if response.Expression != tt.expression {
				t.Errorf("expected expression %q, got %q", tt.expression, response.Expression)
			}

			if response.Result != tt.expected {
				t.Errorf("expected result %v, got %v", tt.expected, response.Result)
			}
		})
	}
}

func TestEvaluateDivisionByZero(t *testing.T) {
	tests := []string{
		"10 / 0",
		"10 / (5 - 5)",
		"10 % 0",
	}

	for _, expression := range tests {
		t.Run(expression, func(t *testing.T) {
			body, err := json.Marshal(appcalculator.EvaluateExpressionRequest{
				Expression: expression,
			})
			if err != nil {
				t.Fatalf("unexpected error building the request: %v", err)
			}

			recorder := doRequest(t, http.MethodPost, evaluatePath, string(body))

			assertStatus(t, recorder, http.StatusBadRequest)
			assertContentType(t, recorder)
			assertJSONKeys(t, recorder, "error")
			assertErrorCode(t, recorder, "DIVISION_BY_ZERO")
		})
	}
}

func TestEvaluateInvalidExpression(t *testing.T) {
	tests := []struct {
		name       string
		expression string
		code       string
	}{
		{
			name:       "unexpected operator",
			expression: "4 + * 5",
			code:       "INVALID_EXPRESSION",
		},
		{
			name:       "missing operand",
			expression: "4 +",
			code:       "INVALID_EXPRESSION",
		},
		{
			name:       "missing operator",
			expression: "4 5",
			code:       "INVALID_EXPRESSION",
		},
		{
			name:       "unknown function",
			expression: "unknown(25)",
			code:       "INVALID_EXPRESSION",
		},
		{
			name:       "invalid number",
			expression: "10.5.2",
			code:       "INVALID_EXPRESSION",
		},
		{
			name:       "missing closing parenthesis",
			expression: "(4 + 5",
			code:       "MISSING_PAREN",
		},
		{
			name:       "empty expression",
			expression: "",
			code:       "EMPTY_EXPRESSION",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, err := json.Marshal(appcalculator.EvaluateExpressionRequest{
				Expression: tt.expression,
			})
			if err != nil {
				t.Fatalf("unexpected error building the request: %v", err)
			}

			recorder := doRequest(t, http.MethodPost, evaluatePath, string(body))

			assertStatus(t, recorder, http.StatusBadRequest)
			assertContentType(t, recorder)
			assertJSONKeys(t, recorder, "error")
			assertErrorCode(t, recorder, tt.code)
		})
	}
}

func TestEvaluateInvalidJSONBody(t *testing.T) {
	tests := []struct {
		name string
		body string
	}{
		{
			name: "not json at all",
			body: "not json",
		},
		{
			name: "truncated object",
			body: `{"expression": "4 + 5"`,
		},
		{
			name: "empty body",
			body: "",
		},
		{
			name: "json array instead of object",
			body: `["4 + 5"]`,
		},
		{
			name: "expression with wrong type",
			body: `{"expression": 10}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			recorder := doRequest(t, http.MethodPost, evaluatePath, tt.body)

			assertStatus(t, recorder, http.StatusBadRequest)
			assertContentType(t, recorder)
			assertJSONKeys(t, recorder, "error")
			assertErrorCode(t, recorder, "INVALID_REQUEST")
		})
	}
}

func TestEvaluateEmptyExpressionField(t *testing.T) {
	tests := []struct {
		name string
		body string
	}{
		{
			name: "field spelled differently",
			body: `{"expr": "4 + 5"}`,
		},
		{
			name: "expression is null",
			body: `{"expression": null}`,
		},
		{
			name: "expression is blank",
			body: `{"expression": "   "}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			recorder := doRequest(t, http.MethodPost, evaluatePath, tt.body)

			assertStatus(t, recorder, http.StatusBadRequest)
			assertContentType(t, recorder)
			assertJSONKeys(t, recorder, "error")
			assertErrorCode(t, recorder, "EMPTY_EXPRESSION")
		})
	}
}
