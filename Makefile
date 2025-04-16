# Simple Makefile for a Go project

# Build the application
all: build test

build:
	@echo "Building..."


	@go build -o main cmd/api/main.go

# Run the application
run:
	@go run cmd/api/main.go

# Install golang-migrate tool
install-migrate:
	@if command -v migrate > /dev/null; then \
		echo "golang-migrate is already installed"; \
	else \
		echo "Installing golang-migrate..."; \
		go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest; \
	fi

# Run database migrations
migrate-up:
	@echo "Running migrations..."
	@if [ ! -f .env ]; then \
		echo "Error: .env file not found"; \
		exit 1; \
	fi
	@echo "Checking migration state..."
	@if [ -f ./fix_migration_version.sh ]; then \
		./fix_migration_version.sh; \
	else \
		echo "Warning: fix_migration_version.sh not found. Skipping migration state check."; \
	fi
	@export $$(grep -v '^#' .env | xargs) && \
	if command -v migrate > /dev/null; then \
		migrate -path=./data/migrations -database "postgres://$${BLUEPRINT_DB_USERNAME}:$${BLUEPRINT_DB_PASSWORD}@$${BLUEPRINT_DB_HOST}:$${BLUEPRINT_DB_PORT}/$${BLUEPRINT_DB_DATABASE}?sslmode=disable" up; \
	else \
		read -p "golang-migrate is not installed. Do you want to install it? [Y/n] " choice; \
		if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
			go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest; \
			migrate -path=./data/migrations -database "postgres://$${BLUEPRINT_DB_USERNAME}:$${BLUEPRINT_DB_PASSWORD}@$${BLUEPRINT_DB_HOST}:$${BLUEPRINT_DB_PORT}/$${BLUEPRINT_DB_DATABASE}?sslmode=disable" up; \
		else \
			echo "You chose not to install golang-migrate. Exiting..."; \
			exit 1; \
		fi; \
	fi

# Rollback database migrations
migrate-down:
	@echo "Rolling back migrations..."
	@if [ ! -f .env ]; then \
		echo "Error: .env file not found"; \
		exit 1; \
	fi
	@if [ ! -f ./migrate_down.sh ]; then \
		echo "Error: migrate_down.sh not found"; \
		exit 1; \
	fi
	@chmod +x ./migrate_down.sh
	@./migrate_down.sh

# Force migration version (useful for fixing dirty database state)
migrate-force:
	@echo "Forcing migration version..."
	@if [ ! -f .env ]; then \
		echo "Error: .env file not found"; \
		exit 1; \
	fi
	@export $$(grep -v '^#' .env | xargs) && \
	if command -v migrate > /dev/null; then \
		migrate -path=./data/migrations -database "postgres://$${BLUEPRINT_DB_USERNAME}:$${BLUEPRINT_DB_PASSWORD}@$${BLUEPRINT_DB_HOST}:$${BLUEPRINT_DB_PORT}/$${BLUEPRINT_DB_DATABASE}?sslmode=disable" force 0; \
	else \
		read -p "golang-migrate is not installed. Do you want to install it? [Y/n] " choice; \
		if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
			go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest; \
			migrate -path=./data/migrations -database "postgres://$${BLUEPRINT_DB_USERNAME}:$${BLUEPRINT_DB_PASSWORD}@$${BLUEPRINT_DB_HOST}:$${BLUEPRINT_DB_PORT}/$${BLUEPRINT_DB_DATABASE}?sslmode=disable" force 0; \
		else \
			echo "You chose not to install golang-migrate. Exiting..."; \
			exit 1; \
		fi; \
	fi

# Generate ent schema from database
ent-import:
	@echo "Generating ent schema from database..."
	@if [ ! -f .env ]; then \
		echo "Error: .env file not found"; \
		exit 1; \
	fi
	@export $$(grep -v '^#' .env | xargs) && \
	go run -mod=mod ariga.io/entimport/cmd/entimport -dsn "postgres://$${BLUEPRINT_DB_USERNAME}:$${BLUEPRINT_DB_PASSWORD}@$${BLUEPRINT_DB_HOST}:$${BLUEPRINT_DB_PORT}/$${BLUEPRINT_DB_DATABASE}?sslmode=disable" -tables users -schema-path ./internal/repository/ent/schema

# Generate ent code from schema
ent-generate:
	@echo "Generating ent code from schema..."
	@go generate ./internal/repository/ent

# Complete workflow: run migrations, import schema, generate code
workflow:
	@echo "Running complete workflow..."
	@make install-migrate
	@make migrate-up
	@make ent-import
	@make ent-generate
	@echo "Workflow completed successfully!"
# Create DB container
docker-run:
	@if docker compose up --build 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker-compose up --build; \
	fi

# Shutdown DB container
docker-down:
	@if docker compose down 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker-compose down; \
	fi

# Test the application
test:
	@echo "Testing..."
	@go test ./... -v
# Integrations Tests for the application
itest:
	@echo "Running integration tests..."
	@go test ./internal/database -v

# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f main

# Live Reload
watch:
	@if command -v air > /dev/null; then \
            air; \
            echo "Watching...";\
        else \
            read -p "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
            if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
                go install github.com/air-verse/air@latest; \
                air; \
                echo "Watching...";\
            else \
                echo "You chose not to install air. Exiting..."; \
                exit 1; \
            fi; \
        fi

.PHONY: all build run test clean watch docker-run docker-down itest migrate-up migrate-down migrate-force ent-import ent-generate workflow install-migrate
