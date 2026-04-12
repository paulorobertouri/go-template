# go-template

Go API template using Echo with JWT authentication, Swagger docs, and automated tests.

## Stack

- Go 1.22+
- Echo
- swaggo/echo-swagger
- golang-jwt/jwt
- Testify

## Endpoints

- `GET /health`
- `GET /docs`
- `GET /v1/public`
- `GET /v1/customer`
- `POST /v1/auth/login`
- `GET /v1/private`

## JWT Contract

Login returns:

- Body: `{"token":"..."}`
- Header: `Authorization: Bearer <token>`
- Header: `X-JWT-Token: <token>`

## Commands

- `make install` — download dependencies
- `make install-dev` — install dev tools
- `make run` — run the application
- `make test` — run tests
- `make test-coverage` — run tests with coverage report
- `make docker-build` — build Docker image
- `make docker-test` — run tests in Docker
- `make docker-curl-test` — run curl integration tests in Docker

## Environment Variables

- `PORT=8000`
- `JWT_SECRET=your-super-secret-jwt-key-at-least-32-characters-long-for-hs256`
- `JWT_ALGORITHM=HS256`
- `JWT_EXPIRATION=3600`

## Project Structure

```text
.
├── main.go
├── Makefile
├── go.mod
├── internal/
│   ├── api/routes.go
│   ├── di/providers.go
│   ├── domain/models.go
│   ├── middleware/auth.go
│   ├── repositories/customer_repo.go
│   └── services/
│       ├── auth_service.go
│       └── customer_service.go
├── tests/
│   ├── unit/services_test.go
│   ├── e2e/api_test.go
│   └── docker/
├── scripts/
│   ├── ubuntu/
│   └── windows/
└── docker/
    ├── build.Dockerfile
    └── test.Dockerfile
```

## Architecture

```text
HTTP -> API Routes -> JWT Middleware -> Services -> Repositories -> Domain
```

## Quality Gates

- Tests must pass.
- Coverage report available via `make test-coverage`.
- Docker image builds and tests pass in container.
