# Go Template

Go API template using Echo with JWT authentication, Swagger docs, and automated tests.

## Stack

- Go 1.22
- Echo
- swaggo/echo-swagger
- Testify

## Endpoints

- GET /docs
- GET /v1/public
- GET /v1/customer
- POST /v1/auth/login
- GET /v1/private

## JWT Contract

Login returns:
- Body: {"token":"..."}
- Header: Authorization: Bearer <token>
- Header: X-JWT-Token: <token>

## Commands

- make install
- make install-dev
- make run
- make test
- make test-coverage
- make docker-build
- make docker-test
- make docker-curl-test

## Environment Variables

- PORT=8000
- JWT_SECRET=your-super-secret-jwt-key-at-least-32-characters-long-for-hs256
- JWT_ALGORITHM=HS256
- JWT_EXPIRATION=3600

## Project Structure

```text
.
├── main.go
├── Makefile
├── go.mod
├── docker/
│   ├── build.Dockerfile
│   └── test.Dockerfile
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
│       ├── test_with_curl.sh
│       └── test_with_curl.ps1
└── scripts/
    ├── ubuntu/docker-curl-test.sh
    └── windows/docker-curl-test.ps1
```

## Architecture

```text
HTTP -> API Routes -> JWT Middleware -> Services -> Repositories -> Domain
```

## Testing

- Unit tests validate service behavior.
- E2E tests validate endpoint behavior.
- Docker curl tests validate runtime image + auth contract.
