# Go Template Project

A small Golang project demonstrating development and testing best practices.

## Project Structure

```
go-template/
├── cmd/
│   └── api/
│       └── main.go          # Main application entry point
├── internal/
│   ├── calculator/
│   │   ├── calculator.go    # Calculator business logic
│   │   └── calculator_test.go # Unit tests
│   ├── user/
│   │   ├── user.go          # User model and operations
│   │   └── user_test.go     # User tests
│   └── server/
│       ├── server.go        # HTTP server
│       └── server_test.go   # Server integration tests
├── pkg/
│   └── utils/
│       ├── strings.go       # Utility functions
│       └── strings_test.go  # Utility tests
├── test/
│   └── integration_test.go  # Integration tests
├── go.mod                   # Go module definition
├── go.sum                   # Go module checksums
├── Makefile                 # Build automation
└── README.md               # This file
```

## Features Demonstrated

- **Go Modules**: Dependency management
- **Package Organization**: Clean separation of concerns
- **Unit Testing**: Individual function testing
- **Integration Testing**: End-to-end testing
- **HTTP Server**: Simple REST API
- **Error Handling**: Proper error patterns
- **Documentation**: Code comments and README

## Getting Started

### Prerequisites

- Go 1.21 or later

### Installation

1. Clone or download this project
2. Navigate to the project directory
3. Install dependencies:

```bash
go mod tidy
```

### Running the Application

```bash
# Run the main application
go run cmd/api/main.go

# Or use the Makefile
make run
```

### Testing

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run only unit tests
go test ./internal/... ./pkg/...

# Run only integration tests
go test ./test/...

# Or use the Makefile
make test
make test-coverage
```

### Building

```bash
# Build the binary
go build -o bin/api cmd/api/main.go

# Or use the Makefile
make build
```

## API Endpoints

The application provides a simple REST API:

- `GET /health` - Health check
- `GET /users/{id}` - Get user by ID
- `POST /calculate` - Perform calculations

## Examples

### Health Check
```bash
curl http://localhost:8080/health
```

### Get User
```bash
curl http://localhost:8080/users/1
```

### Calculate
```bash
curl -X POST http://localhost:8080/calculate \
  -H "Content-Type: application/json" \
  -d '{"operation": "add", "a": 5, "b": 3}'
```

## Testing Philosophy

This project demonstrates several testing approaches:

1. **Unit Tests**: Test individual functions in isolation
2. **Integration Tests**: Test components working together
3. **Table-Driven Tests**: Test multiple scenarios efficiently
4. **Mocking**: Test dependencies in isolation
5. **Benchmarks**: Performance testing

## Best Practices Demonstrated

- Clear package structure
- Proper error handling
- Comprehensive testing
- Code documentation
- Build automation
- Dependency management
