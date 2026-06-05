# M1 Mac Quick Reference

This project is fully optimized for Apple Silicon (M1/M2/M3) Macs. Use these commands for a smooth development experience.

## 🚀 Quick Start (5 minutes)

```bash
# 1. One-time M1 setup (installs dependencies)
make m1-setup

# 2. Start all services (PostgreSQL, Redis, MongoDB, RabbitMQ)
make docker-up

# 3. Build and run the API (in another terminal)
make run
```

API runs at: http://localhost:3000

## 📝 Common Commands

| Command | Purpose |
|---------|---------|
| `make dev` | Development with hot reload |
| `make test` | Run all tests |
| `make lint` | Check code quality |
| `make docker-logs` | View service logs |
| `make clean` | Clean build artifacts |

## 🐛 Troubleshooting

**Docker won't start:**
- Ensure Docker Desktop for Mac is running
- Check: `docker ps`

**Port already in use:**
```bash
lsof -i :5432  # Check port 5432
kill -9 <PID>  # Kill the process
```

**Need more help?**
See [M1_SETUP.md](M1_SETUP.md) for detailed documentation.

---

**Optimized for M1/M2/M3 native performance** ⚡
