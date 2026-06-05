#!/bin/bash

# WFA Meetup API - M1 Mac Setup Script
# This script automates the setup process for M1 Mac developers

set -e

echo "🚀 WFA Meetup API - M1 Mac Setup"
echo "=================================="
echo ""

# Check if running on Mac
if [[ ! "$OSTYPE" == "darwin"* ]]; then
    echo "⚠️  This script is intended for macOS. For other systems, see M1_SETUP.md"
    exit 1
fi

# Check for Apple Silicon
if [[ $(uname -m) != "arm64" ]]; then
    echo "⚠️  This system doesn't appear to be Apple Silicon (M1/M2/M3)"
    echo "Run: uname -m"
    exit 1
fi

echo "✓ Running on macOS Apple Silicon (ARM64)"
echo ""

# Check for Homebrew
if ! command -v brew &> /dev/null; then
    echo "📦 Installing Homebrew..."
    /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
    echo "✓ Homebrew installed"
else
    echo "✓ Homebrew found"
fi

# Check for Go
if ! command -v go &> /dev/null; then
    echo "📦 Installing Go..."
    brew install go
    echo "✓ Go installed"
else
    GO_VERSION=$(go version | awk '{print $3}')
    echo "✓ Go found: $GO_VERSION"
fi

# Check for Docker
if ! command -v docker &> /dev/null; then
    echo "⚠️  Docker Desktop for Mac not found"
    echo "Please install from: https://www.docker.com/products/docker-desktop"
    echo "After installation, ensure Docker Desktop is running"
else
    echo "✓ Docker found"
fi

# Check Go version
GO_VERSION=$(go version | grep -oE '[0-9]+\.[0-9]+' | head -1)
echo "Go version: $GO_VERSION"

echo ""
echo "📥 Installing Go dependencies..."
go mod download
go mod tidy
echo "✓ Dependencies installed"

echo ""
echo "🔧 Installing development tools..."

# Install golangci-lint
if ! command -v golangci-lint &> /dev/null; then
    echo "  Installing golangci-lint..."
    go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
    echo "  ✓ golangci-lint installed"
else
    echo "  ✓ golangci-lint found"
fi

# Install air for hot reload
if ! command -v air &> /dev/null; then
    echo "  Installing air (hot reload)..."
    go install github.com/cosmtrek/air@latest
    echo "  ✓ air installed"
else
    echo "  ✓ air found"
fi

# Install goimports
if ! command -v goimports &> /dev/null; then
    echo "  Installing goimports..."
    go install golang.org/x/tools/cmd/goimports@latest
    echo "  ✓ goimports installed"
else
    echo "  ✓ goimports found"
fi

echo ""
echo "✅ Setup complete!"
echo ""
echo "Next steps:"
echo "1. Start Docker Desktop for Mac"
echo "2. Run: make docker-up"
echo "3. Run: make run"
echo ""
echo "For development with hot reload:"
echo "  make dev"
echo ""
echo "For more information, see M1_SETUP.md"
