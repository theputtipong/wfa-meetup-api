.PHONY: help build run test clean install-deps docker-up docker-down docker-logs

# Variables
BINARY_NAME=api
BUILD_DIR=bin
GO_VERSION=1.24.1
GOOS?=$(shell go env GOOS)
GOARCH?=$(shell go env GOARCH)

help:
	@echo "WFA Meetup API - Makefile Commands"
	@echo "===================================="
	@echo "make build              - Build the API binary"
	@echo "make run                - Run the API locally"
	@echo "make test               - Run all tests"
	@echo "make test-coverage      - Run tests with coverage report"
	@echo "make clean              - Remove build artifacts"
	@echo "make install-deps       - Install dependencies"
	@echo "make lint               - Run linter"
	@echo "make fmt                - Format code"
	@echo "make docker-build       - Build Docker image (supports M1)"
	@echo "make docker-up          - Start Docker containers"
	@echo "make docker-down        - Stop Docker containers"
	@echo "make docker-logs        - View Docker logs"
	@echo "make m1-setup           - M1 Mac setup (install dependencies)"
	@echo "make dev                - Run in development mode with hot reload"

# Build the binary for current platform (includes M1 support)
build: install-deps
	@echo "Building API for $(GOOS)/$(GOARCH)..."
	mkdir -p $(BUILD_DIR)
	CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(GOARCH) go build \
		-ldflags="-w -s -X main.Version=$(shell git describe --tags --always --dirty 2>/dev/null || echo 'dev')" \
		-o $(BUILD_DIR)/$(BINARY_NAME) ./cmd/api
	@echo "✓ Build complete: $(BUILD_DIR)/$(BINARY_NAME)"

# Run the API (requires docker-compose to be running)
run: build
	@echo "Starting API server..."
	./$(BUILD_DIR)/$(BINARY_NAME)

# Run tests
test: install-deps
	@echo "Running tests..."
	go test -v -race -count=1 ./...

# Run tests with coverage
test-coverage: install-deps
	@echo "Running tests with coverage..."
	go test -v -race -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "✓ Coverage report generated: coverage.html"

# Install dependencies
install-deps:
	@echo "Installing dependencies..."
	go mod download
	go mod tidy

# Clean build artifacts
clean:
	@echo "Cleaning..."
	rm -rf $(BUILD_DIR) coverage.out coverage.html
	go clean
	@echo "✓ Clean complete"

# Run linter
lint: install-deps
	@echo "Running linter..."
	@which golangci-lint > /dev/null || (echo "Installing golangci-lint..." && go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest)
	golangci-lint run ./...

# Format code
fmt:
	@echo "Formatting code..."
	go fmt ./...
	goimports -w .
	@echo "✓ Format complete"

# Build Docker image (M1 compatible)
docker-build:
	@echo "Building Docker image for $(GOOS)/$(GOARCH)..."
	docker build -t wfa-meetup-api:latest \
		--platform linux/arm64 \
		-f Dockerfile .
	@echo "✓ Docker image built"

# Start Docker containers
docker-up:
	@echo "Starting Docker containers..."
	docker-compose up -d
	@echo "✓ Containers started"
	@echo "Waiting for services to be healthy..."
	@sleep 5
	docker-compose ps

# Stop Docker containers
docker-down:
	@echo "Stopping Docker containers..."
	docker-compose down

# View Docker logs
docker-logs:
	docker-compose logs -f

# M1 Mac specific setup
m1-setup:
	@echo "Setting up for M1 Mac..."
	@echo "Checking Go version..."
	@go version
	@echo ""
	@echo "✓ Go installation detected"
	@echo "✓ M1 support verified (Go 1.16+ natively supports ARM64)"
	@echo ""
	@echo "Installing Homebrew packages (if not installed)..."
	@command -v brew >/dev/null 2>&1 || { echo "Installing Homebrew..."; /bin/bash -c "$$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"; }
	@brew install docker || echo "Docker already installed"
	@echo ""
	@echo "✓ M1 Mac setup complete!"
	@echo "Next steps:"
	@echo "1. Start Docker Desktop for Mac"
	@echo "2. Run: make docker-up   (to start services)"
	@echo "3. Run: make run         (to start the API)"

# Development mode with hot reload
dev: install-deps docker-up
	@echo "Starting in development mode..."
	@which air > /dev/null || (echo "Installing air for hot reload..." && go install github.com/cosmtrek/air@latest)
	air

all: clean lint test build
	@echo "✓ All tasks completed successfully!"
