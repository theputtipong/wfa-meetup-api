# Windows 11 Quick Reference

This project fully supports Windows 11 development. We use `winget` (Windows Package Manager) to streamline the setup process.

## 🚀 Quick Start (5 minutes)

Open **PowerShell** as Administrator (for first-time setup only) and run:

```powershell
# 1. One-time Windows setup (installs Go, Make, and tools)
.\scripts\windows-setup.ps1
```

_Note: You may need to run `Set-ExecutionPolicy -ExecutionPolicy RemoteSigned -Scope CurrentUser` to allow the script to execute._

Once installed, open a fresh terminal to make sure all environment variables are updated:

```bash
# 2. Start all services (PostgreSQL, Redis, MongoDB, RabbitMQ)
make docker-up

# 3. Build and run the API (in another terminal)
make run
```

API runs at: http://localhost:3000

## 📝 Common Commands

| Command            | Purpose                     |
| ------------------ | --------------------------- |
| `make dev`         | Development with hot reload |
| `make test`        | Run all tests               |
| `make lint`        | Check code quality          |
| `make docker-logs` | View service logs           |
| `make clean`       | Clean build artifacts       |

Ensure **Docker Desktop** is running before using any `make docker-*` commands.
