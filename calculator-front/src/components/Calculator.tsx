import { useState } from 'react';
import { CalculatorApiError, evaluateExpression } from '../services/calculator';
import type { CalculatorKey, HistoryEntry } from '../types/calculator';
import { CalculatorDisplay } from './CalculatorDisplay';
import { CalculatorKeypad } from './CalculatorKeypad';
import { addToHistory, clearHistory, getHistory } from '../services/calculatorHistory';
import { History } from './History';

type Tab = 'calculator' | 'history';

export function Calculator() {
  const [expression, setExpression] = useState('');
  const [result, setResult] = useState<number | null>(null);
  const [error, setError] = useState<string | null>(null);
  const [isLoading, setIsLoading] = useState(false);
  const [tab, setTab] = useState<Tab>('calculator');
  const [history, setHistory] = useState<HistoryEntry[]>(() => getHistory());
  
  const [justEvaluated, setJustEvaluated] = useState(false);

  async function handleKeyPress(key: CalculatorKey) {
    setError(null);

    if (key.type === 'clear') {
      setExpression('');
      setResult(null);
      setJustEvaluated(false);
      return;
    }

    if (key.type === 'backspace') {
      if (justEvaluated) {
        setExpression('');
        setResult(null);
        setJustEvaluated(false);
        return;
      }
      setExpression((prev) => prev.slice(0, -1));
      return;
    }

    if (key.type === 'equals') {
      if (expression.trim().length === 0) return;
      setIsLoading(true);
      try {
        const response = await evaluateExpression(expression);
        setResult(response.result);
        setHistory(addToHistory(response.expression, response.result));
        setJustEvaluated(true);
      } catch (err) {
        setError(err instanceof CalculatorApiError ? err.message : 'No se pudo calcular la expresión.');
      } finally {
        setIsLoading(false);
      }
      return;
    }

    // Cualquier otra tecla: si venimos de un "=", arranca expresión nueva
    // (a partir del resultado anterior si la tecla es un operador).
    let base = expression;
    if (justEvaluated) {
      base = key.type === 'operator' && result !== null ? String(result) : '';
      setResult(null);
      setJustEvaluated(false);
    }

    switch (key.type) {
      case 'digit':
        setExpression(base + key.value);
        return;
      case 'operator':
        setExpression(base.length === 0 ? base : base + key.value);
        return;
      case 'decimal':
        setExpression(base + '.');
        return;
      case 'sqrt':
        setExpression(base + 'sqrt(');
        return;
      case 'parenthesis': {
        const openParens = (base.match(/\(/g) ?? []).length;
        const closeParens = (base.match(/\)/g) ?? []).length;
        setExpression(base + (openParens > closeParens ? ')' : '('));
        return;
      }
    }
  }

  function handleSelectHistoryEntry(entry: HistoryEntry) {
    setExpression(String(entry.result));
    setResult(null);
    setError(null);
    setJustEvaluated(false);
    setTab('calculator');
  }

  function handleClearHistory() {
    setHistory(clearHistory());
  }

  return (
    <div className="flex h-screen w-full md:max-w-105 flex-col overflow-auto md:rounded-[25px] p-5 bg-calculator-surface shadow-2xl md:h-205">
      <div className="flex gap-8 border-b border-white/30 px-8 pt-6" role="tablist">
        <button
          type="button"
          role="tab"
          aria-selected={tab === 'calculator'}
          onClick={() => setTab('calculator')}
          className={[
            'pb-3 text-base transition-colors',
            tab === 'calculator' ? 'border-b-2 border-amber-accent text-white' : 'text-white/70 hover:text-white',
          ].join(' ')}
        >
          Calculator
        </button>
        <button
          type="button"
          role="tab"
          aria-selected={tab === 'history'}
          onClick={() => setTab('history')}
          className={[
            'pb-3 text-base transition-colors',
            tab === 'history' ? 'border-b-2 border-amber-accent text-white' : 'text-white/70 hover:text-white',
          ].join(' ')}
        >
          History
        </button>
      </div>

      {tab === 'calculator' ? (
        <>
          <CalculatorDisplay expression={expression} result={result} error={error} isLoading={isLoading} />
          <div className="mt-auto">
            <CalculatorKeypad onKeyPress={handleKeyPress} disabled={isLoading} />
          </div>
        </>
      ) : (
        <History entries={history} onClear={handleClearHistory} onSelect={handleSelectHistoryEntry} />
      )}
    </div>
  );
}
