import { afterEach, beforeEach, describe, expect, it, vi } from 'vitest';
import { CalculatorApiError, evaluateExpression } from './calculator';

function jsonResponse(body: unknown, init: ResponseInit = {}): Response {
  return new Response(JSON.stringify(body), {
    status: 200,
    headers: { 'Content-Type': 'application/json' },
    ...init,
  });
}

describe('evaluateExpression', () => {
  const fetchMock = vi.fn<typeof fetch>();

  beforeEach(() => {
    fetchMock.mockReset();
    vi.stubGlobal('fetch', fetchMock);
  });

  afterEach(() => {
    vi.unstubAllGlobals();
  });

  it('posts the expression and returns the result', async () => {
    fetchMock.mockResolvedValue(
      jsonResponse({ expression: '4 + 5 * 10', result: 54 }),
    );

    const response = await evaluateExpression('4 + 5 * 10');

    expect(response).toEqual({ expression: '4 + 5 * 10', result: 54 });
    expect(fetchMock).toHaveBeenCalledTimes(1);

    const [url, init] = fetchMock.mock.calls[0];
    expect(url).toBe('http://localhost:8080/api/calculator/evaluate');
    expect(init?.method).toBe('POST');
    expect(init?.headers).toEqual({ 'Content-Type': 'application/json' });
    expect(init?.body).toBe(JSON.stringify({ expression: '4 + 5 * 10' }));
  });

  it.each([
    ['INVALID_EXPRESSION', 'Invalid mathematical expression'],
    ['DIVISION_BY_ZERO', 'Cannot divide by zero'],
    ['EMPTY_EXPRESSION', 'Empty expression'],
    ['MISSING_PAREN', 'Missing parentheses'],
  ])('maps the %s error envelope to a typed error', async (code, message) => {
    fetchMock.mockResolvedValue(
      jsonResponse({ error: { code, message } }, { status: 400 }),
    );

    const error = await evaluateExpression('10 / 0').catch((err: unknown) => err);

    expect(error).toBeInstanceOf(CalculatorApiError);
    expect(error).toMatchObject({ code, message });
  });

  it('reports a network failure as NETWORK_ERROR', async () => {
    fetchMock.mockRejectedValue(new TypeError('Failed to fetch'));

    const error = await evaluateExpression('1 + 1').catch((err: unknown) => err);

    expect(error).toBeInstanceOf(CalculatorApiError);
    expect(error).toMatchObject({ code: 'NETWORK_ERROR' });
  });

  it('falls back to defaults when the error body is not JSON', async () => {
    fetchMock.mockResolvedValue(new Response('gateway timeout', { status: 504 }));

    const error = await evaluateExpression('1 + 1').catch((err: unknown) => err);

    expect(error).toMatchObject({ code: 'INTERNAL_ERROR' });
  });

  it('falls back to a default message when the envelope has no message', async () => {
    fetchMock.mockResolvedValue(jsonResponse({ error: { code: 'INVALID_REQUEST' } }, { status: 400 }));

    const error = await evaluateExpression('1 + 1').catch((err: unknown) => err);

    expect(error).toMatchObject({ code: 'INVALID_REQUEST' });
  });

  // Regresión: el backend devolvía 200 con el cuerpo vacío cuando el resultado
  // se desborda a +Inf. `response.json()` falla, `.catch(() => null)` deja
  // `data === null` y la comprobación `response.ok` lo daba por bueno, así que
  // la función devolvía `null` y el componente reventaba al leer `.result`.
  //
  // Marcados como `it.fails` porque el defecto sigue presente: pasan mientras
  // el bug exista y se pondrán en rojo cuando se corrija `calculator.ts`, que es
  // la señal para convertirlos en `it` normales.
  it.fails('rejects a 200 response whose body is not decodable JSON', async () => {
    fetchMock.mockResolvedValue(
      new Response('', { status: 200, headers: { 'Content-Type': 'application/json' } }),
    );

    const error = await evaluateExpression('1 + 1').catch((err: unknown) => err);

    expect(error).toBeInstanceOf(CalculatorApiError);
    expect(error).not.toBeNull();
  });

  it.fails('rejects a 200 response whose body decodes to null', async () => {
    fetchMock.mockResolvedValue(jsonResponse(null));

    const error = await evaluateExpression('1 + 1').catch((err: unknown) => err);

    expect(error).toBeInstanceOf(CalculatorApiError);
  });
});
