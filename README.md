# Project management-be

One Paragraph of project description goes here

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

## Docker-based Development Environment

This project uses Docker for development to ensure a consistent environment across all developers' machines. The following Makefile commands are available for working with the Docker-based development environment:

### Build the API Docker Image

```bash
make build-local-go-image
```

This command builds the Docker image for the API service using the Dockerfile in the `build` directory.

### Setup the API Environment

```bash
make api-setup
```

This command sets up the API environment by starting the PostgreSQL database, running migrations, and building the API Docker image if it doesn't exist.

### Run the API

```bash
make api-run
```

This command runs the API service in the Docker container.

### Stop the API Environment

```bash
make api-down
```

This command stops and removes all containers, networks, and volumes defined in the Docker Compose files.

### Database Management

Start the PostgreSQL database:
```bash
make pg
```

Run database migrations:
```bash
make api-pg-migrate-up
```

Rollback database migrations:
```bash
make api-pg-migrate-down
```

Seed the database with fake data:
```bash
make db-seed
```

### Code Generation

Generate models from the database:
```bash
make api-gen-models
```

Generate Go code:
```bash
make api-go-generate
```

Generate mocks for testing:
```bash
make api-gen-mocks
```

### Workflow Steps

1. Create or modify migration files in `data/migrations/`
2. Run `make api-setup` to set up the API environment
3. Run `make api-pg-migrate-up` to apply the migrations to the database
4. (Optional) Run `make db-seed` to seed the database with fake data for development and testing
5. Run `make api-gen-models` to generate models from the database
6. Run `make api-go-generate` to generate Go code
7. Run `make api-run` to run the API service
