# Go Template Project

A simple and clean Golang REST API template demonstrating basic CRUD operations, calculator functionality, and modern API development practices.

## Project Structure

```text
go-template/
├── bin/
│   └── api                      # Compiled binary (generated)
├── cmd/
│   └── api/
│       └── main.go              # Main application entry point with graceful shutdown
├── docs/                        # Swagger documentation (auto-generated)
│   ├── docs.go                  # Swagger docs package
│   ├── swagger.json             # OpenAPI JSON specification
│   └── swagger.yaml             # OpenAPI YAML specification
├── internal/
│   ├── common/                  # Common framework for all handlers
│   │   ├── handler.go           # Base handler with common methods
│   │   ├── params.go            # URL parameter extraction helpers
│   │   ├── response.go          # Standardized JSON response helpers
│   │   └── router.go            # Route registration helpers
│   ├── calculator/
│   │   ├── calculator.go        # Calculator business logic
│   │   ├── calculator_test.go   # Calculator unit tests
│   │   └── routes.go            # Calculator HTTP routes and handlers
│   ├── user/
│   │   ├── user.go              # User service and in-memory storage
│   │   ├── user_test.go         # User service unit tests
│   │   └── routes.go            # User CRUD HTTP routes and handlers
│   ├── greeting/                # Example endpoint implementation
│   │   ├── handler.go           # Greeting handlers
│   │   └── handler_test.go      # Greeting handler tests
│   └── server/
│       └── server.go            # HTTP server setup and route configuration
├── go.mod                       # Go module definition
├── go.sum                       # Go module checksums
├── Makefile                     # Build automation and development tasks
└── README.md                    # This file
```

## Features

- **Calculator API**: Basic arithmetic operations (add, subtract, multiply, divide) with validation
- **User Management API**: Full CRUD operations for user management with in-memory storage
- **Greeting API**: Simple greeting endpoints demonstrating easy endpoint creation patterns
- **Swagger Documentation**: Auto-generated OpenAPI documentation with interactive UI
- **Clean Architecture**: Modular, maintainable code structure with reusable components
- **Common Framework**: Simplified handler patterns for rapid endpoint development
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
