.PHONY: help build run test clean install deps fmt lint docker-build docker-run

# Variables
BINARY_NAME := autogen-backend-v2-mcp
CONFIG_PATH := config-files/config.yaml
GO_VERSION := 1.23.5
DOCKER_IMAGE := autogen-v2-mcp
DOCKER_TAG := latest

# Default target
.DEFAULT_GOAL := help

## help: Show this help message
help:
	@echo "Usage:"
	@echo "  make [target]"
	@echo ""
	@echo "Available targets:"
	@echo "  build          - Build the binary"
	@echo "  run            - Build and run the application"
	@echo "  run-dev        - Run the application in development mode (go run)"
	@echo "  test           - Run tests"
	@echo "  test-coverage  - Run tests with coverage report"
	@echo "  clean          - Remove build artifacts"
	@echo "  install        - Install dependencies"
	@echo "  deps           - Download and update dependencies"
	@echo "  fmt            - Format code"
	@echo "  lint           - Run linter (requires golangci-lint)"
	@echo "  vet            - Run go vet"
	@echo "  check          - Run fmt, vet, and lint"
	@echo "  docker-build   - Build Docker image"
	@echo "  docker-run     - Build and run Docker container"
	@echo "  docker-clean   - Remove Docker image"
	@echo "  all            - Clean, deps, fmt, vet, build, and test"

## build: Build the binary
build:
	@echo "Building $(BINARY_NAME)..."
	@go build -o $(BINARY_NAME) .
	@echo "Build complete: $(BINARY_NAME)"

## run: Run the application
run: build
	@echo "Running $(BINARY_NAME)..."
	@./$(BINARY_NAME) -config=$(CONFIG_PATH)

## run-dev: Run the application in development mode (using go run)
run-dev:
	@echo "Running in development mode..."
	@go run . -config=$(CONFIG_PATH)

## test: Run tests
test:
	@echo "Running tests..."
	@go test -v ./...

## test-coverage: Run tests with coverage
test-coverage:
	@echo "Running tests with coverage..."
	@go test -v -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

## clean: Remove build artifacts
clean:
	@echo "Cleaning..."
	@rm -f $(BINARY_NAME)
	@rm -f nodeops
	@rm -f coverage.out coverage.html
	@go clean
	@echo "Clean complete"

## install: Install dependencies
install: deps

## deps: Download dependencies
deps:
	@echo "Downloading dependencies..."
	@go mod download
	@go mod tidy
	@echo "Dependencies updated"

## fmt: Format code
fmt:
	@echo "Formatting code..."
	@go fmt ./...
	@echo "Format complete"

## lint: Run linter (requires golangci-lint)
lint:
	@echo "Running linter..."
	@if command -v golangci-lint > /dev/null; then \
		golangci-lint run; \
	else \
		echo "golangci-lint not installed. Install with: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"; \
	fi

## vet: Run go vet
vet:
	@echo "Running go vet..."
	@go vet ./...
	@echo "Vet complete"

## check: Run fmt, vet, and lint
check: fmt vet lint

## docker-build: Build Docker image
docker-build:
	@echo "Building Docker image..."
	@docker build -t $(DOCKER_IMAGE):$(DOCKER_TAG) .
	@echo "Docker image built: $(DOCKER_IMAGE):$(DOCKER_TAG)"

## docker-run: Run Docker container
docker-run: docker-build
	@echo "Running Docker container..."
	@docker run -p 8080:8080 \
		-v $(PWD)/config-files:/app/config-files \
		$(DOCKER_IMAGE):$(DOCKER_TAG)

## docker-clean: Remove Docker image
docker-clean:
	@echo "Removing Docker image..."
	@docker rmi $(DOCKER_IMAGE):$(DOCKER_TAG) || true
	@echo "Docker image removed"

## all: Build, test, and check
all: clean deps fmt vet build test

