#!/bin/bash

# Load environment variables from .env file
export $(grep -v '^#' .env | xargs)

# Connect to the database and drop the users table if it exists
PGPASSWORD=$BLUEPRINT_DB_PASSWORD psql -h $BLUEPRINT_DB_HOST -p $BLUEPRINT_DB_PORT -U $BLUEPRINT_DB_USERNAME -d $BLUEPRINT_DB_DATABASE -c "DROP TABLE IF EXISTS users;"

# Also drop the schema_migrations table to reset the migration state
PGPASSWORD=$BLUEPRINT_DB_PASSWORD psql -h $BLUEPRINT_DB_HOST -p $BLUEPRINT_DB_PORT -U $BLUEPRINT_DB_USERNAME -d $BLUEPRINT_DB_DATABASE -c "DROP TABLE IF EXISTS schema_migrations;"