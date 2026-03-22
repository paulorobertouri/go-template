# Go Template

A modern Go API template with JWT authentication, Swagger documentation, and comprehensive testing.

## Features

- **FastAPI-style API** with v1 endpoints
- **JWT Authentication** with configurable secret and algorithm
- **Swagger/OpenAPI Documentation** at `/docs`
- **Dependency Injection** with service/repository pattern
- **Docker Support** with multi-stage builds
- **Comprehensive Tests** (unit, e2e, docker integration)
- **Cross-platform Scripts** (Ubuntu `.sh` and Windows `.ps1`)

## API Routes

- `GET /docs` - Swagger API documentation
- `GET /v1/public` - Public endpoint
- `GET /v1/customer` - Get customer (public)
- `POST /v1/auth/login` - Login (returns JWT)
- `GET /v1/private` - Private endpoint (requires JWT)

## JWT Response Contract

Login endpoint returns:
```json
{
  "token": "eyJhbGc..."
}
```

Headers:
- `Authorization: Bearer <token>`
- `X-JWT-Token: <token>`

## Quick Start

### Local Development

```bash
# Install dependencies
make install
make install-dev

# Run the application
make run

# Run tests
make test
make test-coverage

# View coverage
open coverage.html
```

### Docker

```bash
# Build Docker image
make docker-build

# Run Docker image
docker run -p 8000:8000 go-template:latest

# Run tests in Docker
make docker-test

# Run docker+curl integration tests
make docker-curl-test
```

## Environment Variables

```
PORT=8000
JWT_SECRET=your-super-secret-jwt-key-at-least-32-characters-long-for-hs256
JWT_ALGORITHM=HS256
JWT_EXPIRATION=3600
```

## Project Structure

```
.
├── main.go              # Entry point
├── go.mod              # Module definition
├── Makefile            # Build tasks
├── .env                # Environment configuration
├── docker/
│   ├── build.Dockerfile      # Production build
│   └── test.Dockerfile       # Testing build
├── internal/
│   ├── api/
│   │   └── routes.go         # Route definitions
│   ├── services/
│   │   ├── auth_service.go   # JWT generation/validation
│   │   └── customer_service.go
│   ├── repositories/
│   │   └── customer_repo.go
│   ├── domain/
│   │   └── models.go         # Data models
│   ├── middleware/
│   │   └── auth.go           # JWT middleware
│   └── di/
│       └── providers.go      # Dependency injection
├── tests/
│   ├── unit/
│   │   └── services_test.go
│   ├── e2e/
│   │   └── api_test.go
│   └── docker/
│       ├── test_with_curl.sh
│       └── test_with_curl.ps1
└── scripts/
    ├── ubuntu/
    │   └── docker-curl-test.sh
    └── windows/
        └── docker-curl-test.ps1
```

## Architecture

```
HTTP Request
    ↓
API Layer (routes.go)
    ↓
Middleware (auth.go - JWT validation)
    ↓
Service Layer (services)
    ↓
Repository Layer (repositories)
    ↓
Domain Models (models.go)
```

## Testing

- **Unit Tests**: Service isolation tests
- **E2E Tests**: Full HTTP stack integration tests  
- **Docker Tests**: Container-based curl integration tests

Run all tests with `make test` or `make test-coverage`.

## License

MIT
- **Graceful Shutdown**: Proper server lifecycle management with signal handling
- **In-Memory Storage**: No external dependencies required for quick development and testing
- **Comprehensive Testing**: Unit tests for all business logic and handlers
- **Build Automation**: Makefile with common development tasks

## Getting Started

### Prerequisites

- Go 1.21 or later
- Git (for cloning the repository)

### Installation

1. Clone the repository:

```bash
git clone https://github.com/paulorobertouri/go-template.git
cd go-template
```

1. Install dependencies:

```bash
go mod tidy
```

### Running the Application

#### Development Mode

```bash
# Run directly with Go
go run cmd/api/main.go

# Or use the Makefile
make run
```

#### Production Build

```bash
# Build the binary
make build

# Run the compiled binary
./bin/api
```

The server will start on port 8080 and display:

- API endpoints available at: `http://localhost:8080`
- Swagger UI available at: `http://localhost:8080/swagger/`

### Testing

```bash
# Run all tests
make test

# Run tests with coverage report
make test-coverage

# Run tests for a specific package
go test ./internal/calculator/...
```

## API Documentation

This project includes comprehensive Swagger/OpenAPI documentation. Once the server is running, visit:

**Swagger UI**: <http://localhost:8080/swagger/>

The documentation includes:

- Interactive API explorer
- Request/response schemas
- Example payloads
- Error response formats

## API Endpoints

### Health and General

- `GET /health` - Health check endpoint
- `GET /` - Welcome message

### Calculator Routes

All calculator operations return results in the format: `{"data": {"result": <number>}}`

- `GET /add/{a}/{b}` - Add two numbers
- `GET /subtract/{a}/{b}` - Subtract two numbers  
- `GET /multiply/{a}/{b}` - Multiply two numbers
- `GET /divide/{a}/{b}` - Divide two numbers (returns error for division by zero)

**Examples:**

- `GET /add/5/3` → `{"data": {"result": 8}}`
- `GET /divide/10/2` → `{"data": {"result": 5}}`

### Greeting Routes

All greeting operations return messages in the format: `{"data": {"message": "<greeting>"}}`

- `GET /greeting/{name}` - Simple greeting
- `GET /greeting/formal/{name}` - Formal greeting

**Examples:**

- `GET /greeting/John` → `{"data": {"message": "Hello, John!"}}`
- `GET /greeting/formal/Alice` → `{"data": {"message": "Good day, Alice. It's a pleasure to meet you."}}`

### User Management Routes

Complete CRUD operations for user management:

- `GET /users` - Get all users
- `GET /users/{id}` - Get user by ID
- `POST /users` - Create a new user
- `PUT /users/{id}` - Update user by ID
- `DELETE /users/{id}` - Delete user by ID

**User JSON Schema:**

```json
{
  "id": 1,
  "name": "John Doe",
  "email": "john@example.com"
}
```

## Usage Examples

### Calculator Operations

```bash
# Add two numbers
curl http://localhost:8080/add/5/3

# Subtract two numbers
curl http://localhost:8080/subtract/10/4

# Multiply two numbers
curl http://localhost:8080/multiply/7/6

# Divide two numbers
curl http://localhost:8080/divide/15/3

# Handle division by zero (returns error)
curl http://localhost:8080/divide/10/0
```

### User Management

```bash
# Create a user
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{"name": "John Doe", "email": "john.doe@example.com"}'

# Get all users
curl http://localhost:8080/users

# Get user by ID
curl http://localhost:8080/users/1

# Update user
curl -X PUT http://localhost:8080/users/1 \
  -H "Content-Type: application/json" \
  -d '{"name": "Jane Doe", "email": "jane.doe@example.com"}'

# Delete user
curl -X DELETE http://localhost:8080/users/1
```

### Greeting Operations

```bash
# Simple greeting
curl http://localhost:8080/greeting/John

# Formal greeting
curl http://localhost:8080/greeting/formal/Alice
```

### Health Check

```bash
# Check API health
curl http://localhost:8080/health

# Get welcome message
curl http://localhost:8080/
```

## Architecture and Design

### Common Framework

This template includes a simplified framework that makes adding new endpoints very easy. The `internal/common` package provides:

- **BaseHandler**: Common functionality for all handlers (JSON responses, error handling)
- **Route Registration**: Declarative route setup with named routes
- **Parameter Extraction**: Easy URL parameter parsing with type safety
- **Standardized Responses**: Consistent JSON response format across all endpoints

### Adding New Endpoints

Creating new endpoints is straightforward. Here's a quick example:

```go
package myservice

import (
    "net/http"
    "github.com/gorilla/mux"
    "github.com/paulorobertouri/go-template/internal/common"
)

type Handler struct {
    common.BaseHandler
}

func NewHandler() *Handler {
    return &Handler{}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
    routes := common.RouteGroup{
        Prefix: "/my-api",
        Routes: []common.Route{
            common.NamedRoute("/{id}", "GET", "myservice.get", h.handleGet),
            common.NamedRoute("", "POST", "myservice.create", h.handleCreate),
        },
    }
    common.RegisterGroup(router, routes)
}

func (h *Handler) handleGet(w http.ResponseWriter, r *http.Request) {
    params := h.GetURLParams(r)
    id, exists := params.Int("id")
    if !exists {
        h.WriteError(w, http.StatusBadRequest, "Invalid ID")
        return
    }
    
    // Your business logic here
    result := map[string]interface{}{"id": id, "message": "Found"}
    h.WriteSuccess(w, result)
}
```

### Project Structure Benefits

- **Separation of Concerns**: Each package handles a specific domain
- **Testability**: Clear separation makes unit testing straightforward
- **Maintainability**: Common patterns reduce code duplication
- **Extensibility**: Easy to add new features following established patterns

## Development

### Available Make Commands

```bash
# Build the application
make build

# Run the application in development mode
make run

# Run tests
make test

# Run tests with coverage
make test-coverage

# Clean build artifacts
make clean

# View all available commands
make help
```

### Swagger Documentation Generation

The project uses Swaggo for automatic API documentation generation. To regenerate docs after adding new endpoints:

```bash
# Install swag if not already installed
go install github.com/swaggo/swag/cmd/swag@latest

# Generate documentation
swag init -g cmd/api/main.go -o ./docs
```

Make sure to add Swagger comments to your handlers following the format shown in existing handlers.

## Dependencies

### Core Dependencies

- **Gorilla Mux**: HTTP router and URL matcher
- **Swaggo**: Swagger documentation generation
- **Testify**: Testing toolkit for assertions and mocks

### Development Dependencies

- **Go 1.21+**: Modern Go features and performance improvements
- **Make**: Build automation

## Build and Deployment

### Building for Production

```bash
# Build with optimizations
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/api cmd/api/main.go

# Or use the Makefile
make build
```

### Docker Support

While not included in this template, the application is designed to be easily containerized:

```dockerfile
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o api cmd/api/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/api .
EXPOSE 8080
CMD ["./api"]
```

## Contributing

1. Fork the repository
1. Create a feature branch (`git checkout -b feature/amazing-feature`)
1. Make your changes
1. Add tests for new functionality
1. Ensure all tests pass (`make test`)
1. Commit your changes (`git commit -m 'Add some amazing feature'`)
1. Push to the branch (`git push origin feature/amazing-feature`)
1. Open a Pull Request

### Code Standards

- Follow Go conventions and best practices
- Add unit tests for new functionality
- Update documentation for API changes
- Use meaningful commit messages
- Ensure code passes all linting checks

## License

This project is provided as a template for learning and development purposes. Feel free to use it as a starting point for your own projects.

## Support

If you have questions or need help with this template:

1. Check the existing documentation
1. Look at the example implementations in the codebase
1. Review the test files for usage patterns
1. Open an issue for bugs or feature requests

This template is designed to be educational and easy to understand. The code is intentionally well-commented and structured to demonstrate Go best practices.
