# Calculator

A web-based calculator application built with **React + TypeScript** on the frontend and **Go** on the backend.

The application is designed with a separated frontend/backend architecture and can be run either locally for development or using Docker Compose.

## Technologies

### Frontend

- React
- TypeScript
- Vite
- Tailwind CSS
- pnpm

### Backend

- Go 1.27+
- REST API

### Infrastructure

- Docker
- Docker Compose
- Nginx

## Project Structure

```text
calculator/
├── calculator-front/
│   ├── src/
│   │   ├── components/
│   │   ├── App.tsx
│   │   └── ...
│   ├── Dockerfile
│   ├── nginx.conf
│   ├── package.json
│   └── ...
│
├── calculator-back/
│   ├── cmd/
│   │   └── server/
│   ├── ...
│   ├── Dockerfile
│   ├── go.mod
│   └── go.sum
│
├── docker-compose.yml
└── README.md
```

# Setup

## Prerequisites

Make sure the following tools are installed:

- Node.js 22+
- pnpm
- Go 1.27+
- Docker
- Docker Compose

## Clone the repository

```bash
git clone <repository-url>
cd calculator
```

## Frontend setup

Navigate to the frontend directory:

```bash
cd calculator-front
```

Install dependencies:

```bash
pnpm install
```

Create a `.env` file if necessary:

```env
VITE_API_BASE_URL=http://localhost:8080
```

Start the development server:

```bash
pnpm dev
```

The frontend will be available at:

```text
http://localhost:5173
```

## Backend setup

Open another terminal and navigate to the backend:

```bash
cd calculator-back
```

Download dependencies:

```bash
go mod download
```

Run the server:

```bash
go run ./cmd/server
```

The backend will be available at:

```text
http://localhost:8080
```

# Running with Docker

The entire application can also be started using Docker Compose.

From the project root:

```bash
docker compose up --build
```

The application will be available at:

```text
http://localhost:3000
```

The backend will be available at:

```text
http://localhost:8080
```

To run the application in the background:

```bash
docker compose up --build -d
```

To stop the containers:

```bash
docker compose down
```

To check the running services:

```bash
docker compose ps
```

To view logs:

```bash
docker compose logs -f
```

# Environment Variables

The following environment variables can be configured through a `.env` file:

```env
BACKEND_PORT=8080
FRONTEND_PORT=3000
VITE_API_BASE_URL=http://localhost:8080
```

The default values are used when the variables are not explicitly defined.

> `VITE_API_BASE_URL` must use an address accessible from the user's browser. The Docker service name `backend` should not be used here because the API request is made by the browser, not directly by the frontend container.

# API

The backend exposes a REST API for calculator operations.

## Calculate an expression

### Request

```http
POST /calculate
Content-Type: application/json
```

### Request body

```json
{
  "expression": "5 + 3 * 2"
}
```

### Example using cURL

```bash
curl -X POST http://localhost:8080/calculate \
  -H "Content-Type: application/json" \
  -d '{"expression":"5 + 3 * 2"}'
```

### Example response

```json
{
  "result": 11
}
```

The frontend uses this API to evaluate the expression entered by the user.

> The exact endpoint and request/response format should match the implementation exposed by the backend.

# Design Decisions and Assumptions

## Frontend and backend separation

The application separates the user interface from the calculation logic.

The frontend is responsible for:

- Rendering the calculator interface.
- Handling user interactions.
- Building the mathematical expression.
- Sending expressions to the backend.
- Displaying the result.

The backend is responsible for:

- Receiving mathematical expressions.
- Evaluating expressions.
- Returning the calculated result.
- Handling invalid expressions and calculation errors.

This separation keeps the UI independent from the calculation implementation.

## Expression-based calculation

The calculator builds the expression as the user interacts with the interface.

For example:

```text
5
5 +
5 + 3
5 + 3 ×
5 + 3 × 2
```

When the user presses `=`, the complete expression is sent to the backend for evaluation.

## API-based calculation

The frontend does not perform the final calculation itself. Instead, it delegates expression evaluation to the backend through the REST API.

This provides a clear separation of responsibilities and allows the calculation logic to be tested independently from the UI.

## REST API

REST was selected as the communication mechanism because the application has a relatively simple request/response interaction between the frontend and backend.

The frontend sends an expression and the backend returns the calculated result.

## Containerization

Docker Compose is used to run both services together:

```text
                    Docker Compose
                         │
             ┌───────────┴───────────┐
             │                       │
             ▼                       ▼
        Frontend                 Backend
        React/Vite                Go API
        Nginx                     Port 8080
        Port 3000
```

The frontend is built using Node.js and served using Nginx in the production container.

The backend is compiled using Go and runs using a lightweight Distroless image.

## Browser-to-backend communication

When running with Docker, the browser accesses:

```text
http://localhost:3000
```

and API requests are sent to:

```text
http://localhost:8080
```

The Docker service name `backend` is intended for container-to-container communication and is therefore not used as the browser-facing API URL.

## Error handling

Invalid expressions should return an appropriate HTTP error response from the backend rather than causing the application to crash.

The frontend should display an appropriate error state when the API request fails or when the backend rejects an expression.

# Development Commands

## Frontend

Install dependencies:

```bash
pnpm install
```

Run development server:

```bash
pnpm dev
```

Build for production:

```bash
pnpm build
```

Preview production build:

```bash
pnpm preview
```

Run tests:

```bash
pnpm test
```

## Backend

Download dependencies:

```bash
go mod download
```

Run the server:

```bash
go run ./cmd/server
```

Build the application:

```bash
go build -o bin/server ./cmd/server
```

Run Go tests:

```bash
go test ./...
```

# Testing

The frontend uses:

- Vitest
- React Testing Library
- Testing Library User Event

Frontend tests cover component behavior and user interactions.

The backend uses Go's built-in testing framework.

Run frontend tests:

```bash
pnpm test
```

Run backend tests:

```bash
go test ./...
```

# Project Status

- [x] Calculator UI
- [x] React + TypeScript
- [x] Tailwind CSS
- [x] Go backend
- [x] REST API
- [x] Docker configuration
- [x] Docker Compose
- [x] Nginx configuration
- [x] Complete frontend/backend integration
- [x] Complete expression evaluation
- [x] Comprehensive backend unit tests
- [x] Comprehensive frontend unit tests
- [ ] CI/CD

# License

This project is intended for academic and/or personal use.
