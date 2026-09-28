import { render, screen, waitFor } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import type { UserEvent } from '@testing-library/user-event';
import { beforeEach, describe, expect, it, vi } from 'vitest';
import { CalculatorApiError, evaluateExpression } from '../services/calculator';
import { Calculator } from './Calculator';

vi.mock('../services/calculator', async (importOriginal) => {
  const actual = await importOriginal<typeof import('../services/calculator')>();
  return { ...actual, evaluateExpression: vi.fn() };
});

const evaluateMock = vi.mocked(evaluateExpression);

const expressionDisplay = () => screen.getByLabelText('Expresión');
const resultDisplay = () => screen.getByLabelText('Resultado');
const key = (name: string) => screen.getByRole('button', { name });

async function press(user: UserEvent, ...labels: string[]) {
  for (const label of labels) {
    await user.click(key(label));
  }
}

function resolved(expression: string, result: number) {
  evaluateMock.mockResolvedValue({ expression, result });
}

describe('Calculator', () => {
  beforeEach(() => {
    evaluateMock.mockReset();
  });

  it('muestra la expresión mientras se escribe', async () => {
    const user = userEvent.setup();
    render(<Calculator />);

    await press(user, '4', 'Sumar', '5', 'Multiplicar', '1', '0');

    expect(expressionDisplay()).toHaveTextContent('4+5*10');
    expect(evaluateMock).not.toHaveBeenCalled();
  });

  it('envía la expresión al backend y muestra el resultado', async () => {
    const user = userEvent.setup();
    resolved('4+5*10', 54);
    render(<Calculator />);

    await press(user, '4', 'Sumar', '5', 'Multiplicar', '1', '0', 'Calcular');

    await waitFor(() => expect(resultDisplay()).toHaveTextContent('54'));
    expect(evaluateMock).toHaveBeenCalledWith('4+5*10');
  });

  it('respeta la precedencia de operadores mediante la API', async () => {
    const user = userEvent.setup();
    resolved('(4+5)*10', 90);
    render(<Calculator />);

    await press(
      user,
      'Paréntesis que abre',
      '4',
      'Sumar',
      '5',
      'Paréntesis que cierra',
      'Multiplicar',
      '1',
      '0',
      'Calcular',
    );

    await waitFor(() => expect(resultDisplay()).toHaveTextContent('90'));
    expect(evaluateMock).toHaveBeenCalledWith('(4+5)*10');
  });

  it('añade sqrt( al pulsar la raíz cuadrada', async () => {
    const user = userEvent.setup();
    render(<Calculator />);

    await press(user, '2', '5', 'Raíz cuadrada');

    expect(expressionDisplay()).toHaveTextContent('25sqrt(');
  });

  it('ignora un operador al inicio de la expresión', async () => {
    const user = userEvent.setup();
    render(<Calculator />);

    await press(user, 'Sumar', 'Multiplicar');

    expect(expressionDisplay()).toHaveTextContent('0');
  });

  it('empieza una expresión nueva al pulsar un dígito después de "="', async () => {
    const user = userEvent.setup();
    resolved('4+5*10', 54);
    render(<Calculator />);

    await press(user, '4', 'Sumar', '5', 'Multiplicar', '1', '0', 'Calcular');
    await waitFor(() => expect(resultDisplay()).toHaveTextContent('54'));

    await press(user, '7');

    expect(expressionDisplay()).toHaveTextContent('7');
    expect(resultDisplay()).not.toHaveTextContent('54');
  });

  it('continúa desde el resultado al pulsar un operador después de "="', async () => {
    const user = userEvent.setup();
    resolved('4+5*10', 54);
    render(<Calculator />);

    await press(user, '4', 'Sumar', '5', 'Multiplicar', '1', '0', 'Calcular');
    await waitFor(() => expect(resultDisplay()).toHaveTextContent('54'));

    await press(user, 'Sumar', '6');

    expect(expressionDisplay()).toHaveTextContent('54+6');
  });

  it('añade paréntesis y punto decimal al pulsarlos', async () => {
    const user = userEvent.setup();
    render(<Calculator />);

    await press(user, '1', 'Punto decimal', '5', 'Paréntesis que abre', '2', 'Paréntesis que cierra');

    expect(expressionDisplay()).toHaveTextContent('1.5(2)');
  });

  it('empieza una expresión nueva al abrir un paréntesis después de "="', async () => {
    const user = userEvent.setup();
    resolved('4+5', 9);
    render(<Calculator />);

    await press(user, '4', 'Sumar', '5', 'Calcular');
    await waitFor(() => expect(resultDisplay()).toHaveTextContent('9'));

    await press(user, 'Paréntesis que abre');

    expect(expressionDisplay()).toHaveTextContent('(');
    expect(resultDisplay()).not.toHaveTextContent('9');
  });

  it('borra el último carácter con la tecla de retroceso', async () => {
    const user = userEvent.setup();
    render(<Calculator />);

    await press(user, '4', '5', '6', 'Borrar');

    expect(expressionDisplay()).toHaveTextContent('45');
  });

  it('la tecla de retroceso limpia todo justo después de "="', async () => {
    const user = userEvent.setup();
    resolved('4+5*10', 54);
    render(<Calculator />);

    await press(user, '4', 'Sumar', '5', 'Multiplicar', '1', '0', 'Calcular');
    await waitFor(() => expect(resultDisplay()).toHaveTextContent('54'));

    await press(user, 'Borrar');

    expect(expressionDisplay()).toHaveTextContent('0');
    expect(resultDisplay()).not.toHaveTextContent('54');
  });

  it('la tecla C limpia la expresión y el resultado', async () => {
    const user = userEvent.setup();
    resolved('4+5*10', 54);
    render(<Calculator />);

    await press(user, '4', 'Sumar', '5', 'Multiplicar', '1', '0', 'Calcular');
    await waitFor(() => expect(resultDisplay()).toHaveTextContent('54'));

    await press(user, 'Limpiar');

    expect(expressionDisplay()).toHaveTextContent('0');
    expect(resultDisplay()).not.toHaveTextContent('54');
  });

  it('no llama a la API si la expresión está vacía', async () => {
    const user = userEvent.setup();
    render(<Calculator />);

    await press(user, 'Calcular');

    expect(evaluateMock).not.toHaveBeenCalled();
  });

  it('muestra el mensaje de error devuelto por la API', async () => {
    const user = userEvent.setup();
    evaluateMock.mockRejectedValue(
      new CalculatorApiError('DIVISION_BY_ZERO', 'Cannot divide by zero'),
    );
    render(<Calculator />);

    await press(user, '1', '0', 'Dividir', '0', 'Calcular');

    const alert = await screen.findByRole('alert');
    expect(alert).toHaveTextContent('Cannot divide by zero');
  });

  it('muestra un mensaje genérico si el error no es de la API', async () => {
    const user = userEvent.setup();
    evaluateMock.mockRejectedValue(new Error('boom'));
    render(<Calculator />);

    await press(user, '1', 'Sumar', '1', 'Calcular');

    const alert = await screen.findByRole('alert');
    expect(alert).toHaveTextContent('No se pudo calcular la expresión.');
  });

  it('deshabilita el teclado mientras la petición está en vuelo', async () => {
    const user = userEvent.setup();
    let release: (value: { expression: string; result: number }) => void = () => {};
    evaluateMock.mockReturnValue(
      new Promise((resolve) => {
        release = resolve;
      }),
    );
    render(<Calculator />);

    await press(user, '4', 'Sumar', '5', 'Calcular');

    await waitFor(() => expect(key('Calcular')).toBeDisabled());
    expect(resultDisplay()).toHaveTextContent('…');

    release({ expression: '4+5', result: 9 });
    await waitFor(() => expect(resultDisplay()).toHaveTextContent('9'));
  });

  it('guarda el resultado en el historial y permite recuperarlo', async () => {
    const user = userEvent.setup();
    resolved('4+5*10', 54);
    render(<Calculator />);

    await press(user, '4', 'Sumar', '5', 'Multiplicar', '1', '0', 'Calcular');
    await waitFor(() => expect(resultDisplay()).toHaveTextContent('54'));

    await user.click(screen.getByRole('tab', { name: 'History' }));

    const entry = await screen.findByRole('button', { name: /54/ });
    expect(entry).toHaveTextContent('4+5*10');

    await user.click(entry);

    expect(screen.getByRole('tab', { name: 'Calculator' })).toHaveAttribute('aria-selected', 'true');
    expect(expressionDisplay()).toHaveTextContent('54');
  });

  it('muestra un estado vacío cuando el historial no tiene entradas', async () => {
    const user = userEvent.setup();
    render(<Calculator />);

    await user.click(screen.getByRole('tab', { name: 'History' }));

    expect(screen.getByText(/Aún no hay cálculos/)).toBeInTheDocument();
  });

  it('permite borrar todo el historial', async () => {
    const user = userEvent.setup();
    resolved('4+5', 9);
    render(<Calculator />);

    await press(user, '4', 'Sumar', '5', 'Calcular');
    await waitFor(() => expect(resultDisplay()).toHaveTextContent('9'));

    await user.click(screen.getByRole('tab', { name: 'History' }));
    await screen.findByRole('button', { name: /9/ });
    await user.click(screen.getByRole('button', { name: 'Borrar historial' }));

    expect(screen.getByText(/Aún no hay cálculos/)).toBeInTheDocument();
  });
});
