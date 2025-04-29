.PHONY: build-local-go-image api-setup api-run api-down api-pg-migrate-up api-pg-migrate-down api-gen-models api-go-generate api-gen-mocks pg all build run docker-run docker-down test itest clean watch db-seed workflow swagger-docs swagger-serve test-coverage api-drop-data

DOCKER_BIN := docker
PROJECT_NAME := management-football
DOCKER_COMPOSE_BIN := docker-compose
DOCKER_BIN := docker

COMPOSE := PROJECT_NAME=${PROJECT_NAME} ${DOCKER_COMPOSE_BIN} -f build/docker-compose.base.yaml -f build/docker-compose.local.yaml
API_COMPOSE = ${COMPOSE} run --name ${PROJECT_NAME}-api-local --rm --service-ports -w /api api

PORT_SWAGGER := 8080

build-local-go-image:
	${DOCKER_BIN} build -f build/local.go.Dockerfile -t ${PROJECT_NAME}-api-local:latest .
	-${DOCKER_BIN} images -q -f "dangling=true" | xargs ${DOCKER_BIN} rmi -f
api-setup: pg api-pg-migrate-up
	sleep 5
	${DOCKER_BIN} image inspect ${PROJECT_NAME}-api-local:latest >/dev/null 2>&1 || make build-local-go-image
api-run:
	-${DOCKER_BIN} rm -f ${PROJECT_NAME}-api-local 2>/dev/null || true
	${API_COMPOSE} sh -c "go run -mod=vendor cmd/api/main.go server -c configs/.env"
api-down:
	${COMPOSE} down --remove-orphans
api-pg-migrate-up:
	${COMPOSE} run --rm pg-migrate -path=/api-migrations -database="postgres://${PROJECT_NAME}:${PROJECT_NAME}@pg:5432/${PROJECT_NAME}?sslmode=disable" up
api-pg-migrate-down:
	${COMPOSE} run --rm pg-migrate -path=/api-migrations -database="postgres://${PROJECT_NAME}:${PROJECT_NAME}@pg:5432/${PROJECT_NAME}?sslmode=disable" drop
db-seed:
	@echo "Seeding database with fake data..."
	@go run data/seed/*.go
api-gen-models:
	-${DOCKER_BIN} rm -f ${PROJECT_NAME}-api-local 2>/dev/null || true
	${API_COMPOSE} sh -c 'cd ./internal/repository && rm -f ./ent/schema/*.go && entimport -dsn "postgres://${PROJECT_NAME}:${PROJECT_NAME}@pg:5432/${PROJECT_NAME}?sslmode=disable"'
api-go-generate:
	-${DOCKER_BIN} rm -f ${PROJECT_NAME}-api-local 2>/dev/null || true
	${API_COMPOSE} sh -c "go generate ./..."
api-gen-mocks:
	-${DOCKER_BIN} rm -f ${PROJECT_NAME}-mockery-local 2>/dev/null || true
	${COMPOSE} run --name ${PROJECT_NAME}-mockery-local --rm -w /api --entrypoint '' mockery /bin/sh -c "\
		mockery --dir internal/controller --all --recursive --inpackage && \
		mockery --dir internal/repository --all --recursive --inpackage"
pg:
	${COMPOSE} up -d pg

# Build the application
all: build test

build:
	@echo "Building..."
	@go build -o main cmd/api/main.go

# Run the application
run:
	@go run cmd/api/main.go

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
	@echo "Running tests..."
	@go test -v ./...

# Integrations Tests for the application
itest:
	@echo "Running integration tests..."
	@go test ./internal/database -v

# Clean the binary and generated files
clean:
	@echo "Cleaning..."
	@rm -f main
	@rm -rf docs/swagger/*
	@rm -f coverage.out coverage.html

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

# Workflow for setting up the entire project
workflow: api-setup api-gen-models api-go-generate db-seed
	@echo "Workflow completed successfully!"

# Generate Swagger documentation
swagger-docs:
	@echo "Generating Swagger documentation..."
	@~/go/bin/swag init -g swagger.go -o docs/swagger

# Serve Swagger documentation (requires server to be running)
swagger-serve:
	@echo "Open Swagger UI in browser..."
	@open http://localhost:$(PORT_SWAGGER)/swagger/index.html

# Run tests with coverage
test-coverage:
	@echo "Running tests with coverage..."
	@go test -v -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html
	@open coverage.html

# Drop all data from database while keeping structure
api-drop-data: pg
	@echo "Dropping all data from database..."
	@sleep 2
	${COMPOSE} run --rm pg psql "postgresql://${PROJECT_NAME}:${PROJECT_NAME}@pg:5432/${PROJECT_NAME}?sslmode=disable" -c "\
		TRUNCATE TABLE match_players CASCADE; \
		TRUNCATE TABLE player_statistics CASCADE; \
		TRUNCATE TABLE matches CASCADE; \
		TRUNCATE TABLE players CASCADE; \
		TRUNCATE TABLE teams CASCADE; \
		TRUNCATE TABLE departments CASCADE; \
		ALTER SEQUENCE departments_id_seq RESTART WITH 1; \
		ALTER SEQUENCE teams_id_seq RESTART WITH 1; \
		ALTER SEQUENCE players_id_seq RESTART WITH 1; \
		ALTER SEQUENCE matches_id_seq RESTART WITH 1; \
		ALTER SEQUENCE player_statistics_id_seq RESTART WITH 1; \
		ALTER SEQUENCE match_players_id_seq RESTART WITH 1;"