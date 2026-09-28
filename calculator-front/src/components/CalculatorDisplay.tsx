interface CalculatorDisplayProps {
  expression: string;
  result: number | null;
  error: string | null;
  isLoading: boolean;
}

export function CalculatorDisplay({ expression, result, error, isLoading }: CalculatorDisplayProps) {
  return (
    <div className="flex min-h-36 flex-col items-end justify-end gap-2 px-2 pb-8 pt-10 text-right">
      <p aria-label="Expresión" className="w-full break-all text-4xl text-white sm:text-5xl">
        {expression || '0'}
      </p>
      {error ? (
        <p className="w-full wrap-break text-sm text-red-400 sm:text-base" role="alert">
          {error}
        </p>
      ) : (
        <p aria-label="Resultado" className="w-full break-all text-xl text-white/70 sm:text-2xl">
          {isLoading ? '…' : result !== null ? result : '\u00A0'}
        </p>
      )}
    </div>
  );
}
