import { beforeEach, describe, expect, it } from 'vitest';
import { addToHistory, clearHistory, getHistory } from './calculatorHistory';

const STORAGE_KEY = 'calculator.history';

describe('calculatorHistory', () => {
  beforeEach(() => {
    localStorage.clear();
  });

  it('devuelve una lista vacía cuando no hay historial', () => {
    expect(getHistory()).toEqual([]);
  });

  it('añade la entrada más reciente al principio', () => {
    addToHistory('1+1', 2);
    const history = addToHistory('4+5*10', 54);

    expect(history).toHaveLength(2);
    expect(history[0]).toMatchObject({ expression: '4+5*10', result: 54 });
    expect(history[1]).toMatchObject({ expression: '1+1', result: 2 });
  });

  it('persiste en localStorage', () => {
    addToHistory('4+5*10', 54);

    const stored = JSON.parse(localStorage.getItem(STORAGE_KEY) ?? '[]');
    expect(stored).toHaveLength(1);
    expect(stored[0]).toMatchObject({ expression: '4+5*10', result: 54 });
  });

  it('genera un id único por entrada', () => {
    addToHistory('1+1', 2);
    addToHistory('1+1', 2);

    const [first, second] = getHistory();
    expect(first.id).not.toBe(second.id);
  });

  it('conserva como máximo 50 entradas', () => {
    for (let i = 0; i < 55; i += 1) {
      addToHistory(`${i}+0`, i);
    }

    const history = getHistory();
    expect(history).toHaveLength(50);
    expect(history[0].result).toBe(54);
    expect(history[49].result).toBe(5);
  });

  it('borra todo el historial', () => {
    addToHistory('1+1', 2);
    clearHistory();

    expect(getHistory()).toEqual([]);
    expect(localStorage.getItem(STORAGE_KEY)).toBeNull();
  });

  it('devuelve una lista vacía si el dato almacenado está corrupto', () => {
    localStorage.setItem(STORAGE_KEY, '{esto no es json');

    expect(getHistory()).toEqual([]);
  });

  it('acepta decimales y valores negativos en el historial', () => {
    addToHistory('2^-2', 0.25);
    addToHistory('-2^2', -4);

    const history = getHistory();
    expect(history[0].result).toBe(-4);
    expect(history[1].result).toBe(0.25);
  });
});
