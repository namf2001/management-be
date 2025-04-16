# Project management-be

One Paragraph of project description goes here

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

## MakeFile

Run build make command with tests
```bash
make all
```

Build the application
```bash
make build
```

Run the application
```bash
make run
```
Create DB container
```bash
make docker-run
```

Shutdown DB Container
```bash
make docker-down
```

DB Integrations Test:
```bash
make itest
```

Live reload the application:
```bash
make watch
```

Run the test suite:
```bash
make test
```

Clean up binary from the last build:
```bash
make clean
```

## Database Schema Management

This project uses a workflow for managing database schema changes and generating ent code from the database schema.

Install the golang-migrate tool (if not already installed):
```bash
make install-migrate
```

Run database migrations:
```bash
make migrate-up
```

Rollback database migrations:
```bash
make migrate-down
```

Force migration version (useful for fixing dirty database state):
```bash
make migrate-force
```

Generate ent schema from database:
```bash
make ent-import
```

Generate ent code from schema:
```bash
make ent-generate
```

Run the complete workflow (migrations, import schema, generate code):
```bash
make workflow
```

### Helper Scripts

The project includes several helper scripts to assist with database management:

- `fix_migration_version.sh`: Fixes issues with the schema_migrations table when it's in an inconsistent state. This script is automatically called by the `migrate-up` command.
- `migrate_down.sh`: Handles rolling back migrations safely, checking if there are any migrations to roll back before attempting to do so.
- `check_migration_state.sh`: Checks the current state of the schema_migrations table.
- `check_table.sh`: Checks if a specific table exists in the database.
- `drop_table.sh`: Drops specified tables from the database, useful for resetting the database state.

These scripts help ensure that the database is in a consistent state and that migrations can be applied and rolled back safely.

### Workflow Steps

1. Create or modify migration files in `data/migrations/`
2. Run `make docker-run` to start the database container
3. Run `make install-migrate` to install the golang-migrate tool (if not already installed)
4. Run `make migrate-up` to apply the migrations to the database
5. Run `make ent-import` to generate ent schema from the database
6. Run `make ent-generate` to generate ent code from the schema
7. Alternatively, run `make workflow` to execute steps 3-6 in sequence
