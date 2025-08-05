# Go Template Project

A simple and clean Golang REST API template demonstrating basic CRUD operations and calculator functionality.

## Project Structure

```text
go-template/
├── cmd/
│   └── api/
│       └── main.go              # Main application entry point
├── internal/
│   ├── calculator/
│   │   ├── calculator.go        # Simple calculator logic
│   │   ├── calculator_test.go   # Calculator tests
│   │   └── routes.go            # Calculator HTTP routes
│   ├── user/
│   │   ├── user.go              # User service and data
│   │   ├── user_test.go         # User tests
│   │   └── routes.go            # User HTTP routes
│   └── server/
│       └── server.go            # HTTP server setup
├── go.mod                       # Go module definition
├── go.sum                       # Go module checksums
├── Makefile                     # Build automation
└── README.md                    # This file
```

## Features

- **Calculator API**: Basic arithmetic operations (add, subtract, multiply, divide)
- **User Management API**: CRUD operations for user management
- **Clean Architecture**: Simple, easy-to-understand code structure
- **In-Memory Storage**: No external dependencies for data storage
- **Comprehensive Testing**: Unit tests for all components

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

# Or build and run
make build
./bin/api
```

### Testing

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...
```

## API Endpoints

### Health Check

- `GET /health` - Health check endpoint
- `GET /` - Welcome message

### Calculator Routes

- `GET /add/{a}/{b}` - Add two numbers
- `GET /subtract/{a}/{b}` - Subtract two numbers  
- `GET /multiply/{a}/{b}` - Multiply two numbers
- `GET /divide/{a}/{b}` - Divide two numbers

Example: `GET /add/5/3` returns `{"result": 8}`

### User Routes

- `GET /users` - Get all users
- `GET /users/{id}` - Get user by ID
- `POST /users` - Create a new user
- `PUT /users/{id}` - Update user by ID
- `DELETE /users/{id}` - Delete user by ID

User JSON format:

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

## Build

```bash
# Build the binary
go build -o bin/api cmd/api/main.go

# Or use the Makefile
make build
```

## License

This project is provided as a template for learning purposes.
