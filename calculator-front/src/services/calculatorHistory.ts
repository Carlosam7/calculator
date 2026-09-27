import type { HistoryEntry } from '../types/calculator';

const STORAGE_KEY = 'calculator.history';
const MAX_ENTRIES = 50;

function readAll(): HistoryEntry[] {
  try {
    const raw = localStorage.getItem(STORAGE_KEY);
    return raw ? (JSON.parse(raw) as HistoryEntry[]) : [];
  } catch {
    return [];
  }
}

export function getHistory(): HistoryEntry[] {
  return readAll();
}

export function addToHistory(expression: string, result: number): HistoryEntry[] {
  const entry: HistoryEntry = {
    id: crypto.randomUUID(),
    expression,
    result,
    evaluatedAt: Date.now(),
  };
  const updated = [entry, ...readAll()].slice(0, MAX_ENTRIES);
  localStorage.setItem(STORAGE_KEY, JSON.stringify(updated));
  return updated;
}

export function clearHistory(): HistoryEntry[] {
  localStorage.removeItem(STORAGE_KEY);
  return [];
}
