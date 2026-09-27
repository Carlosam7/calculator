export type ApiErrorCode =
  | 'INVALID_EXPRESSION'
  | 'DIVISION_BY_ZERO'
  | 'EMPTY_EXPRESSION'
  | 'MISSING_PAREN'
  | 'INVALID_REQUEST'
  | 'INTERNAL_ERROR'
  | 'NETWORK_ERROR';

export interface EvaluateExpressionResponse {
  expression: string;
  result: number;
}

export interface ApiErrorResponse {
  error: {
    code: ApiErrorCode;
    message: string;
  };
}

export interface HistoryEntry {
  id: string;
  expression: string;
  result: number;
  evaluatedAt: number;
}

export type CalculatorKey =
  | { type: 'digit'; value: string }
  | { type: 'operator'; value: '+' | '-' | '*' | '/' | '%' | '^' }
  | { type: 'sqrt' }
  | { type: 'decimal' }
  | { type: 'parenthesis' }
  | { type: 'clear' }
  | { type: 'backspace' }
  | { type: 'equals' };
