import type { CalculatorKey } from '../types/calculator';

interface CalculatorKeypadProps {
  onKeyPress: (key: CalculatorKey) => void;
  disabled?: boolean;
}

type Variant = 'digit' | 'operator' | 'muted' | 'accent' | 'pill';

const VARIANT_CLASSES: Record<Variant, string> = {
  digit: 'border border-amber-accent/30 text-white',
  operator: 'bg-amber-accent/20 text-white/70',
  muted: 'bg-amber-accent/20 text-white',
  accent: 'bg-amber-accent text-white',
  pill: 'bg-white/10 text-amber-accent',
};

interface KeypadButtonProps {
  label: string;
  variant: Variant;
  onClick: () => void;
  ariaLabel: string;
  span?: 1 | 2;
  disabled?: boolean;
}

function KeypadButton({ label, variant, onClick, ariaLabel, span = 1, disabled }: KeypadButtonProps) {
  return (
    <button
      type="button"
      onClick={onClick}
      disabled={disabled}
      aria-label={ariaLabel}
      className={[
        'flex items-center justify-center text-xl transition-colors duration-150 sm:text-2xl',
        'hover:brightness-110 active:scale-95 disabled:pointer-events-none disabled:opacity-40',
        span === 2 ? 'col-span-2 rounded-full py-3' : 'aspect-square rounded-full',
        VARIANT_CLASSES[variant],
      ].join(' ')}
    >
      {label}
    </button>
  );
}

export function CalculatorKeypad({ onKeyPress, disabled }: CalculatorKeypadProps) {
  return (
    <div className="grid grid-cols-4 gap-3 px-6 pb-6 sm:gap-4 sm:px-8 sm:pb-8">
      <KeypadButton label="C" variant="pill" span={2} ariaLabel="Limpiar" onClick={() => onKeyPress({ type: 'clear' })} />
      <KeypadButton label="←" variant="pill" span={2} ariaLabel="Borrar" onClick={() => onKeyPress({ type: 'backspace' })} />

      <KeypadButton label="√" variant="muted" ariaLabel="Raíz cuadrada" onClick={() => onKeyPress({ type: 'sqrt' })} />
      <KeypadButton
        label="xʸ"
        variant="muted"
        ariaLabel="Potencia"
        onClick={() => onKeyPress({ type: 'operator', value: '^' })}
      />
      <KeypadButton
        label="x"
        variant="operator"
        ariaLabel="Multiplicar"
        onClick={() => onKeyPress({ type: 'operator', value: '*' })}
      />
      <KeypadButton
        label="/"
        variant="operator"
        ariaLabel="Dividir"
        onClick={() => onKeyPress({ type: 'operator', value: '/' })}
      />

      <KeypadButton label="7" variant="digit" ariaLabel="7" onClick={() => onKeyPress({ type: 'digit', value: '7' })} />
      <KeypadButton label="8" variant="digit" ariaLabel="8" onClick={() => onKeyPress({ type: 'digit', value: '8' })} />
      <KeypadButton label="9" variant="digit" ariaLabel="9" onClick={() => onKeyPress({ type: 'digit', value: '9' })} />
      <KeypadButton
        label="−"
        variant="operator"
        ariaLabel="Restar"
        onClick={() => onKeyPress({ type: 'operator', value: '-' })}
      />

      <KeypadButton label="4" variant="digit" ariaLabel="4" onClick={() => onKeyPress({ type: 'digit', value: '4' })} />
      <KeypadButton label="5" variant="digit" ariaLabel="5" onClick={() => onKeyPress({ type: 'digit', value: '5' })} />
      <KeypadButton label="6" variant="digit" ariaLabel="6" onClick={() => onKeyPress({ type: 'digit', value: '6' })} />
      <KeypadButton
        label="+"
        variant="operator"
        ariaLabel="Sumar"
        onClick={() => onKeyPress({ type: 'operator', value: '+' })}
      />

      <KeypadButton label="1" variant="digit" ariaLabel="1" onClick={() => onKeyPress({ type: 'digit', value: '1' })} />
      <KeypadButton label="2" variant="digit" ariaLabel="2" onClick={() => onKeyPress({ type: 'digit', value: '2' })} />
      <KeypadButton label="3" variant="digit" ariaLabel="3" onClick={() => onKeyPress({ type: 'digit', value: '3' })} />
      <KeypadButton
        label="%"
        variant="operator"
        ariaLabel="Módulo"
        onClick={() => onKeyPress({ type: 'operator', value: '%' })}
      />

      <div className="flex aspect-square gap-1.5">
        <button
          type="button"
          onClick={() => onKeyPress({ type: 'parenthesis', value: '(' })}
          aria-label="Paréntesis que abre"
          className={[
            'h-full flex-1 rounded-full text-lg transition-colors duration-150 sm:text-xl',
            'hover:brightness-110 active:scale-95',
            VARIANT_CLASSES.muted,
          ].join(' ')}
        >
          (
        </button>
        <button
          type="button"
          onClick={() => onKeyPress({ type: 'parenthesis', value: ')' })}
          aria-label="Paréntesis que cierra"
          className={[
            'h-full flex-1 rounded-full text-lg transition-colors duration-150 sm:text-xl',
            'hover:brightness-110 active:scale-95',
            VARIANT_CLASSES.muted,
          ].join(' ')}
        >
          )
        </button>
      </div>
      <KeypadButton label="0" variant="muted" ariaLabel="0" onClick={() => onKeyPress({ type: 'digit', value: '0' })} />
      <KeypadButton label="." variant="muted" ariaLabel="Punto decimal" onClick={() => onKeyPress({ type: 'decimal' })} />
      <KeypadButton
        label="="
        variant="accent"
        ariaLabel="Calcular"
        disabled={disabled}
        onClick={() => onKeyPress({ type: 'equals' })}
      />
    </div>
  );
}
