# calculator-front

React + TypeScript + Vite frontend for the calculator API. Talks to the Go backend
on `http://localhost:8080` by default (override with `VITE_API_BASE_URL`).

## Scripts

| Command | Purpose |
| --- | --- |
| `npm run dev` | Vite dev server on `http://localhost:5173` |
| `npm run build` | Typecheck (`tsc -b`) and build to `dist/` |
| `npm run preview` | Serve the production build locally |
| `npm run lint` | Oxlint |
| `npm test` | Run the test suite once |
| `npm run test:watch` | Re-run tests on change |
| `npm run test:coverage` | Test suite with a coverage report |

## Testing

Tests use [Vitest](https://vitest.dev) with Testing Library and `jsdom`. No test
runner was configured before; the config lives in the `test` block of
`vite.config.ts` and the shared setup in `src/test/setup.ts`.

| File | Covers |
| --- | --- |
| `src/App.test.tsx` | Mounts the component tree. |
| `src/components/Calculator.test.tsx` | The keypad state machine: expression building, `=` behaviour, continuing from a result, clear/backspace, loading and error states, history tab. |
| `src/services/calculator.test.ts` | The API client: request shape, error-envelope mapping, network failures, malformed responses. |
| `src/services/calculatorHistory.test.ts` | `localStorage` persistence, the 50-entry cap and corrupt data. |

Components are queried through accessible names (`aria-label`, `role`) rather
than test ids, so the tests double as an accessibility check.

### Known failing tests

`src/services/calculator.test.ts` contains two cases marked `it.fails`. They
document a real defect: when the backend returns `200 OK` with an empty body —
which it currently does when a result overflows to `+Inf` — `evaluateExpression`
resolves to `null` instead of throwing, and the component then throws
`TypeError: Cannot read properties of null`.

They pass as long as the bug exists and will turn red once it is fixed, which is
the signal to change them back to `it` and fix the service.

## React + TypeScript + Vite

This project is based on the Vite React template.

Currently, two official plugins are available:

- [@vitejs/plugin-react](https://github.com/vitejs/vite-plugin-react/blob/main/packages/plugin-react) uses [Oxc](https://oxc.rs)
- [@vitejs/plugin-react-swc](https://github.com/vitejs/vite-plugin-react/blob/main/packages/plugin-react-swc) uses [SWC](https://swc.rs/)

## React Compiler

The React Compiler is not enabled on this template because of its impact on dev & build performances. To add it, see [this documentation](https://react.dev/learn/react-compiler/installation).

## Expanding the Oxlint configuration

If you are developing a production application, we recommend enabling type-aware lint rules by installing `oxlint-tsgolint` and editing `.oxlintrc.json`:

```json
{
  "$schema": "./node_modules/oxlint/configuration_schema.json",
  "plugins": ["react", "typescript", "oxc"],
  "options": {
    "typeAware": true
  },
  "rules": {
    "react/rules-of-hooks": "error",
    "react/only-export-components": ["warn", { "allowConstantExport": true }]
  }
}
```

See the [Oxlint rules documentation](https://oxc.rs/docs/guide/usage/linter/rules) for the full list of rules and categories.
