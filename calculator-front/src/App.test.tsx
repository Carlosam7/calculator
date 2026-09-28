import { render, screen } from '@testing-library/react';
import { describe, expect, it } from 'vitest';
import App from './App';

describe('App', () => {
  it('monta la calculadora', () => {
    render(<App />);

    expect(screen.getByLabelText('Expresión')).toBeInTheDocument();
    expect(screen.getByRole('tab', { name: 'Calculator' })).toBeInTheDocument();
    expect(screen.getByRole('tab', { name: 'History' })).toBeInTheDocument();
  });
});
