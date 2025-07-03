# Makefile for microservice
.PHONY: build run test clean migrate-up migrate-down docker-build docker-run help

# Variables
BINARY_NAME=microservice
MAIN_PATH=cmd/main.go
BUILD_DIR=bin

# Default target
help:
	@echo "Available commands:"
	@echo "  build          - Build the application"
	@echo "  run            - Run the application"
	@echo "  test           - Run tests"
	@echo "  test-coverage  - Run tests with coverage"
	@echo "  clean          - Clean build artifacts"
	@echo "  migrate-up     - Run database migrations up"
	@echo "  migrate-down   - Run database migrations down"
	@echo "  docker-build   - Build Docker image"
	@echo "  docker-run     - Run with Docker Compose"
	@echo "  docker-stop    - Stop Docker containers"
	@echo "  dev            - Run in development mode"
	@echo "  lint           - Run linter"
	@echo "  fmt            - Format code"
	@echo "  deps           - Install dependencies"

# Build the application
build:
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_PATH)
	@echo "Build completed: $(BUILD_DIR)/$(BINARY_NAME)"

# Run the application
run:
	@echo "Running $(BINARY_NAME)..."
	go run $(MAIN_PATH)

# Run tests
test:
	@echo "Running tests..."
	go test ./...

# Run tests with coverage
test-coverage:
	@echo "Running tests with coverage..."
	go test -cover ./...
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	rm -rf $(BUILD_DIR)/
	rm -f coverage.out coverage.html
	@echo "Clean completed"

# Database migrations
migrate-up:
	@echo "Running database migrations up..."
	migrate -path internal/migrations -database "postgres://postgres:password@localhost:5432/microservice?sslmode=disable" up

migrate-down:
	@echo "Running database migrations down..."
	migrate -path internal/migrations -database "postgres://postgres:password@localhost:5432/microservice?sslmode=disable" down

# Docker commands
docker-build:
	@echo "Building Docker image..."
	docker build -t $(BINARY_NAME):latest .

docker-run:
	@echo "Running with Docker Compose..."
	docker-compose up --build

docker-stop:
	@echo "Stopping Docker containers..."
	docker-compose down

# Development mode with hot reload
dev:
	@echo "Starting development server..."
	air

# Lint code
lint:
	@echo "Running linter..."
	golangci-lint run

# Format code
fmt:
	@echo "Formatting code..."
	go fmt ./...

# Install dependencies
deps:
	@echo "Installing dependencies..."
	go mod tidy
	go mod download

# Generate mocks (if using mockery)
mocks:
	@echo "Generating mocks..."
	mockery --all --output=internal/mocks

# Run specific tests
test-unit:
	@echo "Running unit tests..."
	go test ./tests/Unit/...

test-feature:
	@echo "Running feature tests..."
	go test ./tests/Feature/...

# Database operations
db-reset: migrate-down migrate-up
	@echo "Database reset completed"

# Production build
build-prod:
	@echo "Building for production..."
	@mkdir -p $(BUILD_DIR)
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_PATH)
	@echo "Production build completed"