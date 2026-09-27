import type { HistoryEntry } from '../types/calculator';

interface HistoryProps {
  entries: HistoryEntry[];
  onClear: () => void;
  onSelect: (entry: HistoryEntry) => void;
}

export function History({ entries, onClear, onSelect }: HistoryProps) {
  if (entries.length === 0) {
    return (
      <div className="flex flex-1 items-center justify-center px-8 text-center text-white/50">
        <p>Aún no hay cálculos. Los resultados confirmados con "=" aparecerán aquí.</p>
      </div>
    );
  }

  return (
    <div className="flex flex-1 flex-col overflow-hidden px-6 sm:px-8">
      <ul className="flex-1 space-y-1 overflow-y-auto py-4">
        {entries.map((entry) => (
          <li key={entry.id}>
            <button
              type="button"
              onClick={() => onSelect(entry)}
              className="flex w-full flex-col items-end rounded-xl px-3 py-2 text-right transition-colors hover:bg-white/5"
            >
              <span className="break-all text-sm text-white/60">{entry.expression}</span>
              <span className="break-all text-xl text-white">{entry.result}</span>
            </button>
          </li>
        ))}
      </ul>
      <button
        type="button"
        onClick={onClear}
        className="mb-6 self-end rounded-full px-4 py-2 text-sm text-white/60 transition-colors hover:text-amber-accent"
      >
        Borrar historial
      </button>
    </div>
  );
}
