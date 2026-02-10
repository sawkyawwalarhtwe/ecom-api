# Docker & Containerization Guide

## Overview

This project includes:
- **Dockerfile**: Multi-stage build that compiles the Go binary and runs it in a minimal Alpine Linux container
- **docker-compose.yaml**: Orchestrates PostgreSQL and the API service with networking and health checks

## Building the Docker Image

### Build locally:
```bash
docker build -t ecom-api:latest .
```

### Build with a specific tag:
```bash
docker build -t ecom-api:v1.0.0 .
```

## Running with Docker Compose

### Start all services:
```bash
docker-compose up -d
```

This starts:
- **PostgreSQL 16** on `localhost:5432`
- **ecom-api** on `localhost:8080`
- Both services on a shared network (`ecom-network`)

### View logs:
```bash
# All services
docker-compose logs -f

# Specific service
docker-compose logs -f api
docker-compose logs -f postgres
```

### Stop services:
```bash
docker-compose down
```

### Stop and remove volumes:
```bash
docker-compose down -v
```

## Verifying the Setup

### Check if services are running:
```bash
docker-compose ps
```

### Test the API:
```bash
# Health check
curl http://localhost:8080/health

# API documentation
curl http://localhost:8080/docs

# List products
curl http://localhost:8080/products

# Place an order
curl -X POST http://localhost:8080/orders \
  -H "Content-Type: application/json" \
  -d '{
    "customer_id": 1,
    "items": [{"product_id": 1, "quantity": 2}]
  }'
```

### Check database connection:
```bash
docker-compose exec postgres psql -U postgres -d ecom -c "SELECT * FROM products;"
```

## Image Details

### Multi-Stage Build Explanation:

**Stage 1 (Builder):**
- Uses `golang:1.25.3-alpine` to compile the application
- Downloads dependencies
- Builds a static binary with `CGO_ENABLED=0` for portability

**Stage 2 (Runtime):**
- Uses minimal `alpine:3.19` base image (~5MB)
- Copies only the compiled binary from Stage 1
- Creates a non-root `app` user for security
- Includes health check configuration

### Final Image Size:
Typical size: **~30-40 MB** (minimal for a Go application)

## Production Deployment

### Using a container registry:

```bash
# Tag for registry
docker tag ecom-api:latest myregistry.azurecr.io/ecom-api:latest

# Push to registry
docker push myregistry.azurecr.io/ecom-api:latest
```

### Running without docker-compose:

```bash
# Start PostgreSQL
docker run -d \
  --name postgres \
  -e POSTGRES_PASSWORD=postgres \
  -e POSTGRES_DB=ecom \
  -p 5432:5432 \
  postgres:16-alpine

# Start API (ensure DB is ready first)
docker run -d \
  --name ecom-api \
  -e GOOSE_DBSTRING="host=postgres user=postgres password=postgres dbname=ecom sslmode=disable" \
  -p 8080:8080 \
  --link postgres:postgres \
  ecom-api:latest
```

## Environment Variables in Docker

The following environment variables can be set in `docker-compose.yaml`:

| Variable | Default | Description |
|----------|---------|-------------|
| `GOOSE_DBSTRING` | `host=postgres user=postgres password=postgres dbname=ecom sslmode=disable` | PostgreSQL connection string |
| `PORT` | `8080` | API server port |
| `LOG_LEVEL` | `info` | Logging level (debug, info, warn, error) |

## Health Checks

Both services include health checks:

- **PostgreSQL**: Checks if `pg_isready` responds
- **ecom-api**: Checks the `/health` endpoint

This ensures services only start when dependencies are ready.

## Troubleshooting

### Container won't start:
```bash
# Check logs
docker-compose logs api

# Inspect the image
docker inspect ecom-api:latest
```

### Database connection refused:
```bash
# Verify PostgreSQL is healthy
docker-compose ps
# Status should show "healthy" for postgres

# Check network connectivity
docker-compose exec api ping postgres
```

### Port already in use:
Edit `docker-compose.yaml` to change port mappings (e.g., `8081:8080`)

### Rebuild image without cache:
```bash
docker-compose build --no-cache
```

## Development vs Production

### Development:
```bash
docker-compose up -d
# Services auto-restart, no security constraints
```

### Production Checklist:
- Use secrets management (not environment variables in compose)
- Set `restart: unless-stopped` (already configured)
- Run non-root user (already configured)
- Configure resource limits:
  ```yaml
  resources:
    limits:
      cpus: '0.5'
      memory: 512M
  ```
- Use read-only filesystem where possible
- Configure log rotation
- Monitor health checks

## CI/CD Integration

### Building in CI/CD:
```bash
# In GitHub Actions, GitLab CI, etc.
docker build -t ecom-api:${{ github.sha }} .
docker push registry.example.com/ecom-api:${{ github.sha }}
```

The included `.github/workflows/ci.yaml` can be extended to build and push Docker images.
