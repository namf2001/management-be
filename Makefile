.PHONY: build-local-go-image api-setup api-run api-down api-pg-migrate-up api-pg-migrate-down api-gen-models api-go-generate api-gen-mocks pg all build run docker-run docker-down test itest clean watch

DOCKER_BIN := docker
PROJECT_NAME := management-football
DOCKER_COMPOSE_BIN := docker-compose
DOCKER_BIN := docker

COMPOSE := PROJECT_NAME=${PROJECT_NAME} ${DOCKER_COMPOSE_BIN} -f build/docker-compose.base.yaml -f build/docker-compose.local.yaml
API_COMPOSE = ${COMPOSE} run --name ${PROJECT_NAME}-api-local --rm --service-ports -w /api api

build-local-go-image:
	${DOCKER_BIN} build -f build/local.go.Dockerfile -t ${PROJECT_NAME}-api-local:latest .
	-${DOCKER_BIN} images -q -f "dangling=true" | xargs ${DOCKER_BIN} rmi -f
api-setup: pg api-pg-migrate-up
	sleep 5
	${DOCKER_BIN} image inspect ${PROJECT_NAME}-api-local:latest >/dev/null 2>&1 || make build-local-go-image
api-run:
	${API_COMPOSE} sh -c "go run -mod=vendor cmd/api/main.go server -c configs/.env"
api-down:
	${COMPOSE} down --remove-orphans
api-pg-migrate-up:
	${COMPOSE} run --rm pg-migrate -path=/api-migrations -database="postgres://${PROJECT_NAME}:${PROJECT_NAME}@pg:5432/${PROJECT_NAME}?sslmode=disable" up
api-pg-migrate-down:
	${COMPOSE} run --rm pg-migrate -path=/api-migrations -database="postgres://${PROJECT_NAME}:${PROJECT_NAME}@pg:5432/${PROJECT_NAME}?sslmode=disable" drop
api-gen-models:
		${API_COMPOSE} sh -c 'cd ./internal/repository && go run ariga.io/entimport/cmd/entimport -dsn "postgres://${PROJECT_NAME}:@pg:5432/${PROJECT_NAME}?sslmode=disable" && go run entgo.io/ent/cmd/ent generate --feature sql/execquery ./ent/schema'
api-go-generate:
	${API_COMPOSE} sh -c "go generate ./..."
api-gen-mocks:
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
