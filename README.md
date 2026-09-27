# Calculator

A web-based calculator application built with **React + TypeScript** on the frontend and **Go** on the backend. The project is designed to run both locally for development and through Docker Compose for a containerized environment.

## Technologies

### Frontend

- React
- TypeScript
- Vite
- Tailwind CSS
- pnpm
- Nginx

### Backend

- Go 1.27+
- REST API
- Docker
- Distroless

### Infrastructure

- Docker
- Docker Compose

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
│   ├── pnpm-lock.yaml
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

## Requirements

To run the project locally, make sure you have the following installed:

- Node.js 22+
- pnpm
- Go 1.27+
- Docker
- Docker Compose

## Running with Docker

The recommended way to run the complete application is using Docker Compose.

From the project root:

```bash
docker compose up --build
```

The services will be available at:

| Service  | URL                   |
| -------- | --------------------- |
| Frontend | http://localhost:3000 |
| Backend  | http://localhost:8080 |

To run the services in the background:

```bash
docker compose up --build -d
```

To stop the services:

```bash
docker compose down
```

To check the status of the containers:

```bash
docker compose ps
```

To view the logs:

```bash
docker compose logs -f
```

To view only the backend logs:

```bash
docker compose logs -f backend
```

To view only the frontend logs:

```bash
docker compose logs -f frontend
```

## Environment Variables

Docker Compose supports configuring the application ports and backend URL through environment variables.

Default values:

```env
BACKEND_PORT=8080
FRONTEND_PORT=3000
VITE_API_BASE_URL=http://localhost:8080
```

These variables can also be defined in a `.env` file:

```env
BACKEND_PORT=8080
FRONTEND_PORT=3000
VITE_API_BASE_URL=http://localhost:8080
```

> `VITE_API_BASE_URL` must point to a URL accessible from the user's browser. Therefore, `http://backend:8080` should not be used for this variable.

## Frontend Development

Navigate to the frontend directory:

```bash
cd calculator-front
```

Install dependencies:

```bash
pnpm install
```

Start the development server:

```bash
pnpm dev
```

By default, Vite will be available at:

```text
http://localhos
```
