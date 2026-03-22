.PHONY: install install-dev run test docker-build docker-test docker-curl-test clean help

install:
	go mod tidy
	go mod download

install-dev: install
	go install github.com/cosmtrek/air@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

run:
	go run main.go

test:
	go test -count=1 -v ./...

test-coverage:
	go test -count=1 -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

docker-build:
	docker build -f docker/build.Dockerfile -t go-template:latest .

docker-test:
	docker build -f docker/test.Dockerfile -t go-template-test:latest .
	docker run --rm go-template-test:latest

docker-curl-test:
	bash tests/docker/test_with_curl.sh

clean:
	go clean
	rm -rf coverage.out coverage.html

help:
	@echo "Available targets:"
	@echo "  make install        - Download dependencies"
	@echo "  make install-dev    - Install dev tools"
	@echo "  make run            - Run the application"
	@echo "  make test           - Run tests"
	@echo "  make test-coverage  - Run tests with coverage report"
	@echo "  make docker-build   - Build Docker image"
	@echo "  make docker-test    - Build and run tests in Docker"
	@echo "  make docker-curl-test - Run curl integration tests in Docker"
	@echo "  make clean          - Clean build artifacts"
