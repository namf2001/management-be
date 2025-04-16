#!/bin/bash

# Load environment variables from .env file
export $(grep -v '^#' .env | xargs)

# Connect to the database and check the schema_migrations table
echo "Checking schema_migrations table..."
PGPASSWORD=$BLUEPRINT_DB_PASSWORD psql -h $BLUEPRINT_DB_HOST -p $BLUEPRINT_DB_PORT -U $BLUEPRINT_DB_USERNAME -d $BLUEPRINT_DB_DATABASE -c "SELECT * FROM schema_migrations;"

# Check if the users table exists
echo "Checking if users table exists..."
PGPASSWORD=$BLUEPRINT_DB_PASSWORD psql -h $BLUEPRINT_DB_HOST -p $BLUEPRINT_DB_PORT -U $BLUEPRINT_DB_USERNAME -d $BLUEPRINT_DB_DATABASE -c "SELECT EXISTS (SELECT FROM information_schema.tables WHERE table_schema = 'public' AND table_name = 'schema_migrations');"