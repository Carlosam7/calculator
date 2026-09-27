import type { ApiErrorResponse, EvaluateExpressionResponse } from '../types/calculator';

const API_BASE_URL = import.meta.env.VITE_API_BASE_URL ?? 'http://localhost:8080';

/** Error de la API: conserva el `code` del envelope del backend además del mensaje. */
export class CalculatorApiError extends Error {
  readonly code: string;

  constructor(code: string, message: string) {
    super(message);
    this.name = 'CalculatorApiError';
    this.code = code;
  }
}

/**
 * Evalúa una expresión llamando a POST /api/calculator/evaluate del backend en Go.
 * Contrato (ver README del backend):
 *   request  -> { "expression": "4 + 5 * 10" }
 *   response -> { "expression": "4 + 5 * 10", "result": 54 }
 *   error    -> { "error": { "code": "DIVISION_BY_ZERO", "message": "..." } }  (HTTP 400/500)
 */
export async function evaluateExpression(expression: string): Promise<EvaluateExpressionResponse> {
  let response: Response;
  try {
    response = await fetch(`${API_BASE_URL}/api/calculator/evaluate`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ expression }),
    });
  } catch {
    throw new CalculatorApiError('NETWORK_ERROR', 'No se pudo conectar con la API.');
  }

  const data = await response.json().catch(() => null);

  if (!response.ok) {
    const errorBody = data as ApiErrorResponse | null;
    throw new CalculatorApiError(
      errorBody?.error?.code ?? 'INTERNAL_ERROR',
      errorBody?.error?.message ?? 'Ocurrió un error al calcular.',
    );
  }

  return data as EvaluateExpressionResponse;
}