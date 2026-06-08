.PHONY: help build run test clean install-deps docker-up docker-down docker-logs

# Variables
BINARY_NAME=api
BUILD_DIR=bin
GO_VERSION=1.24.1
GOOS?=$(shell go env GOOS)
GOARCH?=$(shell go env GOARCH)

ifeq ($(OS),Windows_NT)
EXE = .exe
RUN_CMD = .\\$(BUILD_DIR)\\$(BINARY_NAME)$(EXE)
BIN_PATH = $(BUILD_DIR)/$(BINARY_NAME)$(EXE)
ENV_SET = set CGO_ENABLED=0&& set GOOS=$(GOOS)&& set GOARCH=$(GOARCH)&&
MKDIR = if not exist $(BUILD_DIR) mkdir $(BUILD_DIR)
else
EXE =
RUN_CMD = ./$(BUILD_DIR)/$(BINARY_NAME)
BIN_PATH = $(BUILD_DIR)/$(BINARY_NAME)
ENV_SET = CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(GOARCH)
MKDIR = mkdir -p $(BUILD_DIR)
endif

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
	@$(MKDIR)
	@$(ENV_SET) go build \
		-ldflags="-w -s -X main.Version=$(shell git describe --tags --always --dirty 2>/dev/null || echo 'dev')" \
		-o $(BIN_PATH) ./cmd/api
	@echo "Build complete: $(BIN_PATH)"

# Run the API (requires docker-compose to be running)
run: build
	@echo "Starting API server..."
	@$(RUN_CMD)

# Run tests
test: install-deps
	@echo "Running tests..."
	go test -v -race -count=1 ./...

# Run tests with coverage
test-coverage: install-deps
	@echo "Running tests with coverage..."
	go test -v -race -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Install dependencies
install-deps:
	@echo "Installing dependencies..."
	@echo "Installing Go dependencies..."
	go mod download
	go mod tidy
	@echo "Dependencies installed"

# Clean build artifacts
clean:
	@echo "Cleaning..."
ifeq ($(OS),Windows_NT)
	@if exist $(BUILD_DIR) rmdir /s /q $(BUILD_DIR)
	@if exist coverage.out del /f /q coverage.out
	@if exist coverage.html del /f /q coverage.html
else
	@rm -rf $(BUILD_DIR) coverage.out coverage.html
endif
	go clean
	@echo "Clean complete"

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
	@echo "Format complete"

# Build Docker image (M1 compatible)
docker-build:
	@echo "Building Docker image for $(GOOS)/$(GOARCH)..."
	docker build -t wfa-meetup-api:latest \
		--platform linux/arm64 \
		-f Dockerfile .
	@echo "Docker image built"

# Cross-platform wait command for Windows and Unix
ifeq ($(OS),Windows_NT)
WAIT_COMMAND = timeout /t 5 /nobreak >nul
else
WAIT_COMMAND = sleep 5
endif

# Start Docker containers
docker-up:
	@echo "Starting Docker containers..."
	docker-compose up -d
	@echo "Containers started"
	@echo "Waiting for services to be healthy..."
	@$(WAIT_COMMAND)
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
	@echo "Go installation detected"
	@echo "M1 support verified (Go 1.16+ natively supports ARM64)"
	@echo ""
	@echo "Installing Homebrew packages (if not installed)..."
	@command -v brew >/dev/null 2>&1 || { echo "Installing Homebrew..."; /bin/bash -c "$$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"; }
	@brew install docker || echo "Docker already installed"
	@echo ""
	@echo "M1 Mac setup complete!"
	@echo "Next steps:"
	@echo "1. Start Docker (Docker Desktop or 'colima start')"
	@echo "1. Start Docker Desktop for Mac"
	@echo "2. Run: make docker-up   (to start services)"
	@echo "3. Run: make run         (to start the API)"

# Development mode with hot reload
dev: install-deps docker-up
	@echo "Starting in development mode..."
	@which air > /dev/null || (echo "Installing air for hot reload..." && go install github.com/cosmtrek/air@latest)
	air

all: clean lint test build
	@echo "All tasks completed successfully!"
