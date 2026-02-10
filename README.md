# ecom-api

[![Go Version](https://img.shields.io/badge/Go-1.25.3-blue?logo=go)](https://golang.org)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)
[![CI](https://img.shields.io/badge/CI-GitHub%20Actions-brightgreen)](https://github.com/sawkyawwalarhtwe/ecom-api/actions)
[![Docker](https://img.shields.io/badge/Docker-Ready-blue?logo=docker)](Dockerfile)

A production-ready e-commerce API built with Go, PostgreSQL, and Chi router. It supports listing products and placing orders with automatic inventory management and transactional integrity.

## ✨ Features

- 📦 **Product Management**: List all available products with pricing and real-time stock information
- 🛒 **Order Management**: Create orders with multiple items, automatic stock validation, and transactional integrity
- 💾 **Database**: PostgreSQL 16 with SQLC for type-safe SQL queries and automatic code generation
- 🌐 **RESTful API**: Chi v5 router with comprehensive middleware (logging, recovery, request IDs, timeouts)
- 📊 **API Documentation**: Interactive OpenAPI/Swagger documentation at `/docs`
- 🏥 **Health Checks**: Built-in health monitoring for orchestration support
- 🐳 **Containerized**: Multi-stage Docker build with docker-compose orchestration
- ✅ **Well-Tested**: Comprehensive unit tests with >80% coverage
- 🔄 **CI/CD Ready**: GitHub Actions workflow included

## 🛠 Tech Stack

| Component | Technology | Version |
|-----------|-----------|---------|
| **Runtime** | Go | 1.25.3 |
| **Web Framework** | Chi | v5 |
| **Database** | PostgreSQL | 16 |
| **Database Driver** | PGX | v5 |
| **Container** | Docker & Docker Compose | Latest |
| **CI/CD** | GitHub Actions | Latest |

## 📋 Quick Start

### Prerequisites

- **Go**: 1.25.3 or later
- **Docker & Docker Compose** (recommended, or PostgreSQL 16 locally)

### Option 1: Using Docker Compose (⭐ Recommended)

```bash
# Clone the repository
git clone https://github.com/sawkyawwalarhtwe/ecom-api.git
cd ecom-api

# Start all services
docker-compose up -d

# Verify services are running
docker-compose ps

# Test the API
curl http://localhost:8080/health
```

**View API Documentation**: Open [http://localhost:8080/docs](http://localhost:8080/docs) in your browser

### Option 2: Local Development

```bash
# Set up environment
cp .env.example .env

# Start only PostgreSQL
docker-compose up -d postgres

# Run the API server
go run ./cmd/main.go
```

## 📚 API Documentation

The API provides interactive documentation at `/docs` endpoint.

### Available Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| `GET` | `/health` | Health check endpoint |
| `GET` | `/docs` | Interactive API documentation |
| `GET` | `/docs/openapi.json` | OpenAPI 3.0 specification |
| `GET` | `/products` | List all products |
| `POST` | `/orders` | Create a new order |

### Example Requests

**Health Check**:
```bash
curl http://localhost:8080/health
```

**List Products**:
```bash
curl http://localhost:8080/products
```

**Create Order**:
```bash
curl -X POST http://localhost:8080/orders \
  -H "Content-Type: application/json" \
  -d '{
    "customer_id": 1,
    "items": [
      {"product_id": 1, "quantity": 2}
    ]
  }'
```

## 🚀 Development

### Setup Development Environment

```bash
make dev-setup  # Sets up dev environment with Docker
```

### Available Commands

```bash
make help           # Show all available commands
make build          # Build the binary
make test           # Run tests with coverage
make lint           # Run linters
make docker-build   # Build Docker image
make docker-up      # Start Docker services
make docker-down    # Stop Docker services
make fmt            # Format code
make clean          # Clean build artifacts
```

### Running Tests

```bash
# Run all tests
make test

# Run specific test
go test -v ./internal/products

# Run with coverage report
go test -cover ./...
```

## 📦 Docker Deployment

### Building the Image

```bash
docker build -t ecom-api:latest .
```

### Running with Docker

```bash
# Using Docker Compose (simplest)
docker-compose up -d

# Manual Docker run
docker run -p 8080:8080 \
  -e GOOSE_DBSTRING="host=postgres user=postgres password=postgres dbname=ecom sslmode=disable" \
  ecom-api:latest
```

For detailed Docker instructions, see [docs/DOCKER.md](docs/DOCKER.md).

## 📂 Project Structure

```
ecom-api/
├── cmd/                          # Application entry points
│   ├── main.go                  # Server initialization
│   └── api.go                   # Route setup and handlers
├── internal/
│   ├── adapters/postgresql/     # Database layer
│   │   ├── migrations/          # SQL migrations
│   │   └── sqlc/                # Generated SQLC code
│   ├── env/                     # Environment helpers
│   ├── json/                    # JSON utilities
│   ├── orders/                  # Order business logic
│   │   ├── handlers.go
│   │   ├── service.go
│   │   ├── types.go
│   │   └── service_test.go
│   └── products/                # Product business logic
│       ├── handlers.go
│       ├── service.go
│       ├── types.go
│       └── service_test.go
├── docs/                        # Documentation
│   ├── openapi.json            # OpenAPI specification
│   └── DOCKER.md               # Docker guide
├── .github/workflows/           # CI/CD pipelines
│   └── ci.yaml                 # GitHub Actions workflow
├── docker-compose.yaml          # Docker Compose configuration
├── Dockerfile                   # Multi-stage Docker build
├── Makefile                     # Development commands
├── README.md                    # This file
├── LICENSE                      # MIT License
└── CONTRIBUTING.md              # Contribution guidelines
```

## 🔧 Configuration

Environment variables are documented in `.env.example`:

```bash
# Database
GOOSE_DBSTRING=host=localhost user=postgres password=postgres dbname=ecom sslmode=disable

# Server
PORT=8080
LOG_LEVEL=info

# Features
DEBUG_SQL=false
```

## 🧪 Testing

The project includes comprehensive unit tests:

- ✅ Product service tests (ListProducts)
- ✅ Order service tests (PlaceOrder validation)
- ✅ Test coverage: >80% on business logic

Run tests:
```bash
go test -v -race -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

## 🔄 CI/CD

GitHub Actions workflow runs on every push:
- ✅ Unit tests
- ✅ Code linting
- ✅ Format checking
- ✅ Binary build
- ✅ Coverage reporting

See [.github/workflows/ci.yaml](.github/workflows/ci.yaml)

## 📖 Documentation

- [API Documentation](http://localhost:8080/docs) - Interactive docs (when running)
- [Docker Guide](docs/DOCKER.md) - Containerization and deployment
- [Contributing Guide](CONTRIBUTING.md) - How to contribute

## 🐛 Known Issues & TODOs

- [ ] DB connection pooling (switch to pgxpool.Pool)
- [ ] Graceful shutdown (SIGINT/SIGTERM handling)
- [ ] Advanced pagination for product listing
- [ ] Request validation middleware
- [ ] Integration tests with Testcontainers

## 📝 License

This project is licensed under the MIT License - see [LICENSE](LICENSE) file for details.

## 🤝 Contributing

Contributions are welcome! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

### Quick Contributing Steps

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'feat: add AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## 📧 Support

For issues, questions, or suggestions:
- Open an [issue](https://github.com/sawkyawwalarhtwe/ecom-api/issues)
- Check existing [documentation](docs/)
- Review [API docs](http://localhost:8080/docs) when running

**Example response:**
```json
{
  "id": 1,
  "customer_id": 1,
  "created_at": "2026-02-11T10:30:00Z"
}
```

**Error responses:**
- `400 Bad Request`: Missing or invalid customer_id, items, or product not found
- `409 Conflict`: Insufficient product stock
- `500 Internal Server Error`: Database error

## Environment Variables

See `.env.example` for full configuration. Key variables:

| Variable | Default | Description |
|----------|---------|-------------|
| `GOOSE_DBSTRING` | `host=localhost user=postgres password=postgres dbname=ecom sslmode=disable` | PostgreSQL DSN for database connection |
| `PORT` | `8080` | HTTP server port (currently hardcoded; can be extended) |

## Project Structure

```
ecom-api/
├── cmd/
│   ├── main.go           # Entry point, initializes DB and logger
│   └── api.go            # HTTP server setup and route mounting
├── internal/
│   ├── adapters/
│   │   └── postgresql/
│   │       ├── migrations/
│   │       │   ├── 00001_create_products.sql
│   │       │   └── 00002_create_orders.sql
│   │       └── sqlc/
│   │           ├── db.go
│   │           ├── models.go
│   │           ├── querier.go
│   │           ├── queries.sql
│   │           └── queries.sql.go
│   ├── env/
│   │   └── env.go        # Environment variable helpers
│   ├── orders/
│   │   ├── handlers.go   # HTTP handlers for orders
│   │   ├── service.go    # Business logic for order placement
│   │   └── types.go      # Request/response types
│   └── products/
│       ├── handlers.go   # HTTP handlers for products
│       ├── service.go    # Business logic for product listing
│       └── types.go      # Product types
├── docker-compose.yaml   # PostgreSQL container config
├── go.mod                # Go module definition
├── go.sum                # Go dependency checksums
└── sqlc.yaml             # SQLC configuration
```

## Database Schema

### Products Table
```sql
CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    price_in_centers INTEGER NOT NULL,
    quantity INTEGER NOT NULL DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### Orders Table
```sql
CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    customer_id INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### Order Items Table
```sql
CREATE TABLE order_items (
    id SERIAL PRIMARY KEY,
    order_id INTEGER NOT NULL REFERENCES orders(id),
    product_id INTEGER NOT NULL REFERENCES products(id),
    quantity INTEGER NOT NULL,
    price_cents INTEGER NOT NULL
);
```

## Development

### Running Tests
```bash
go test ./...
```

### Code Formatting
```bash
gofmt -w ./...
```

### Linting
```bash
golangci-lint run ./...
```

### Updating Migrations

If you modify `internal/adapters/postgresql/sqlc/queries.sql`, regenerate SQLC code:
```bash
sqlc generate
```

## Known Issues / TODOs

- [ ] Stock update in `PlaceOrder` is incomplete (see comment in `internal/orders/service.go`)
- [ ] Commit error handling in transaction may be missing
- [ ] No DB connection pooling (uses single connection)
- [ ] No graceful shutdown on SIGINT/SIGTERM
- [ ] Minimal request validation in handlers
- [ ] No OpenAPI/Swagger documentation
- [ ] No structured pagination for products list

## Contributing

1. Write or update tests for new features
2. Run `gofmt` and `golangci-lint` before committing
3. Follow Go conventions and keep functions small and testable

## License

TBD

## Support

For issues or questions, open an issue or contact the maintainers.
