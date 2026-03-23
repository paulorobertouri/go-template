.PHONY: install install-dev run test format docker-build docker-test docker-curl-test clean help

install:
	bash ./scripts/ubuntu/install.sh

install-dev: install
	bash ./scripts/ubuntu/install-dev.sh

run:
	bash ./scripts/ubuntu/run.sh

test:
	bash ./scripts/ubuntu/test.sh

test-coverage:
	go test -count=1 -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

format:
	bash ./scripts/ubuntu/format.sh

docker-build:
	docker build -f docker/build.Dockerfile -t go-template:latest .

docker-test:
	docker build -f docker/test.Dockerfile -t go-template-test:latest .
	docker run --rm go-template-test:latest

docker-curl-test:
	bash ./scripts/ubuntu/docker-curl-test.sh

clean:
	go clean
	rm -rf coverage.out coverage.html

help:
	@echo "Available targets:"
	@echo "  make install        - Download dependencies via Ubuntu script"
	@echo "  make install-dev    - Install dev tools via Ubuntu script"
	@echo "  make run            - Run the application via Ubuntu script"
	@echo "  make test           - Run tests via Ubuntu script"
	@echo "  make test-coverage  - Run tests with coverage report"
	@echo "  make docker-build   - Build Docker image"
	@echo "  make docker-test    - Build and run tests in Docker"
	@echo "  make docker-curl-test - Run curl integration tests in Docker"
	@echo "  make format         - Format source code
	@echo "  make clean          - Clean build artifacts""
