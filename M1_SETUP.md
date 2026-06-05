# WFA Meetup API - M1 Mac Compatibility Guide

This project is now fully optimized for **Mac M1/M2/M3** (Apple Silicon) development. All components have been tested and verified to work seamlessly on ARM64 architecture.

## ✅ M1 Compatibility Features

- ✓ Go 1.24+ with native ARM64 support
- ✓ Multi-stage Docker builds supporting ARM64
- ✓ Docker Compose with explicit ARM64 platform specifications
- ✓ Health checks for all services
- ✓ Makefile for easy M1-native builds
- ✓ GitHub Actions testing on both x64 and ARM64

## 🚀 Quick Start on M1 Mac

### Prerequisites

1. **Install Homebrew** (if not already installed):
   ```bash
   /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
   ```

2. **Install Go 1.24+**:
   ```bash
   brew install go
   ```
   Verify: `go version`

3. **Install Docker Desktop for Mac**:
   - Download from [Docker's official website](https://www.docker.com/products/docker-desktop)
   - Ensure Docker Desktop is running before starting the API

### Setup & Run

1. **M1-specific setup** (one-time):
   ```bash
   make m1-setup
   ```

2. **Start all services**:
   ```bash
   make docker-up
   ```
   Wait for the health checks to pass:
   ```bash
   docker-compose ps
   ```

3. **Build and run the API**:
   ```bash
   make run
   ```

The API will start on `http://localhost:3000`

## 📚 Available Commands

### Development
```bash
make dev                # Run with hot reload (requires air)
make run                # Build and run
make build              # Build binary only
make test               # Run tests
make test-coverage      # Run tests with coverage report
make lint               # Run linter
make fmt                # Format code
```

### Docker Management
```bash
make docker-build       # Build Docker image for M1
make docker-up          # Start services
make docker-down        # Stop services
make docker-logs        # View logs
```

### Maintenance
```bash
make install-deps       # Install Go dependencies
make clean              # Clean build artifacts
make all                # Run full pipeline (clean, lint, test, build)
```

## 🔧 Configuration

### Environment Variables

Create a `.env` file in the project root (if needed):

```env
# PostgreSQL
DB_HOST=localhost
DB_PORT=5432
DB_USER=user
DB_PASSWORD=password
DB_NAME=wfa_meetup

# Redis
REDIS_URL=localhost:6379

# Google Maps
GOOGLE_MAPS_API_KEY=your_api_key_here
```

### Service Ports

| Service | Port | URL |
|---------|------|-----|
| API | 3000 | http://localhost:3000 |
| PostgreSQL | 5432 | localhost:5432 |
| MongoDB | 27017 | localhost:27017 |
| Redis | 6379 | localhost:6379 |
| RabbitMQ AMQP | 5672 | localhost:5672 |
| RabbitMQ UI | 15672 | http://localhost:15672 |

## 🏗️ Architecture

### Multi-stage Docker Build
The Dockerfile uses multi-stage builds optimized for M1:
- **Builder stage**: Compiles Go binary for ARM64
- **Final stage**: Minimal Alpine Linux image for runtime

This results in:
- ✓ Fast builds on M1 (native compilation)
- ✓ Small image size
- ✓ No CGO dependencies needed

### Platform Specifications
All Docker services explicitly specify `platform: linux/arm64` in docker-compose.yml to ensure native M1 performance.

## 🧪 Testing on M1

Run comprehensive tests:
```bash
make test               # Unit tests with race detection
make test-coverage      # Generate HTML coverage report
```

The GitHub Actions workflow automatically tests:
- macOS latest (M1 native)
- Ubuntu latest (x64 for compatibility)

## 📊 Development Workflow

### Hot Reload Development
```bash
make dev
```
This will:
1. Start all Docker services
2. Install `air` for hot reload (if not present)
3. Watch for file changes and automatically rebuild

### Building for Different Architectures

**Build for M1/ARM64** (default on M1 Mac):
```bash
make build
```

**Build for Intel/AMD64** (cross-compile):
```bash
GOOS=darwin GOARCH=amd64 make build
```

**Build for Linux/ARM64** (Docker):
```bash
make docker-build
```

## 🐛 Troubleshooting

### Docker Services Won't Start
```bash
# Check Docker Desktop is running
docker ps

# Check logs
make docker-logs

# Restart services
make docker-down
make docker-up
```

### Build Issues
```bash
# Clean and rebuild
make clean
make build

# Verify Go installation
go version
```

### Port Already in Use
If a port is already in use:
```bash
# Find process using port (example: port 5432)
lsof -i :5432

# Kill process
kill -9 <PID>

# Or change port in docker-compose.yml and .env
```

### Redis/Database Connection Issues
```bash
# Wait for health checks to pass
docker-compose ps

# Check service logs
docker-compose logs postgres
docker-compose logs redis
```

## 📝 Notes for M1 Users

1. **Native Performance**: This project is optimized to run natively on M1/M2/M3 without any translation layers like Rosetta.

2. **Docker Support**: Docker Desktop for Mac with Apple Silicon provides excellent performance with native ARM64 image support.

3. **Go Support**: Go 1.16+ has first-class ARM64 support, making Go projects naturally fast on M1.

4. **No CGO**: The build uses `CGO_ENABLED=0` to avoid C dependencies that might have ARM64 compatibility issues.

## 🔗 Resources

- [Go on ARM64 Documentation](https://golang.org/doc/install/source#requirements)
- [Docker Desktop for Mac with Apple Silicon](https://docs.docker.com/desktop/mac/apple-silicon/)
- [Fiber Web Framework Docs](https://docs.gofiber.io/)
- [GORM Documentation](https://gorm.io/)

## 🤝 Contributing

When contributing, please ensure:
1. Code works on both x64 and ARM64
2. Docker images build for both architectures
3. Tests pass: `make test`
4. Code is formatted: `make fmt`

---

**Happy coding on M1! 🚀**
