# Makefile

.PHONY: help build run test lint fmt clean docker-build docker-up docker-down docs

help: ## Show this help message
	@echo "Available commands:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  make %-20s %s\n", $$1, $$2}'

build: ## Build the binary
	@echo "Building ecom-api..."
	@go build -o ecom-api ./cmd

run: ## Run the API server locally
	@echo "Starting API server..."
	@go run ./cmd/main.go

test: ## Run all tests
	@echo "Running tests..."
	@go test -v -race -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report: coverage.html"

lint: ## Run linters
	@echo "Running linters..."
	@go vet ./...
	@go fmt ./...
	@which golangci-lint > /dev/null && golangci-lint run ./... || echo "golangci-lint not installed"

fmt: ## Format code
	@echo "Formatting code..."
	@gofmt -s -w .
	@goimports -w .

clean: ## Clean build artifacts
	@echo "Cleaning..."
	@rm -f ecom-api
	@rm -f coverage.out coverage.html
	@go clean

docker-build: ## Build Docker image
	@echo "Building Docker image..."
	@docker build -t ecom-api:latest .

docker-up: ## Start services with Docker Compose
	@echo "Starting services..."
	@docker-compose up -d
	@echo "Services started. Check status with: docker-compose ps"

docker-down: ## Stop Docker Compose services
	@echo "Stopping services..."
	@docker-compose down

docker-logs: ## View Docker logs
	@docker-compose logs -f

docker-clean: ## Remove Docker containers and volumes
	@echo "Removing containers and volumes..."
	@docker-compose down -v

dev-setup: ## Setup development environment
	@echo "Setting up development environment..."
	@cp .env.example .env
	@go mod download
	@docker-compose up -d
	@echo "Development environment ready!"

dev-start: ## Start development setup
	@docker-compose up -d

dev-stop: ## Stop development setup
	@docker-compose down

db-shell: ## Connect to PostgreSQL shell
	@docker-compose exec postgres psql -U postgres -d ecom

db-reset: ## Reset database (WARNING: deletes all data)
	@echo "WARNING: This will delete all data!"
	@read -p "Are you sure? (y/N) " -n 1 -r; \
	echo; \
	if [[ $$REPLY =~ ^[Yy]$$ ]]; then \
		docker-compose exec postgres psql -U postgres -d ecom -c "DROP TABLE IF EXISTS order_items CASCADE; DROP TABLE IF EXISTS orders CASCADE; DROP TABLE IF EXISTS products CASCADE;"; \
		echo "Database reset complete."; \
	fi

api-docs: ## Open API documentation
	@echo "Opening API documentation at http://localhost:8080/docs"
	@open http://localhost:8080/docs || xdg-open http://localhost:8080/docs || start http://localhost:8080/docs

deps: ## Download dependencies
	@echo "Downloading dependencies..."
	@go mod download
	@go mod tidy

sqlc-gen: ## Generate SQLC code
	@echo "Generating SQLC code..."
	@sqlc generate

all: clean lint test build ## Build, test, and lint everything
