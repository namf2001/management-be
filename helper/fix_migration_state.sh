#!/bin/bash

# Load environment variables from .env file
export $(grep -v '^#' .env | xargs)

# Drop the users table manually
echo "Dropping users table..."
PGPASSWORD=$BLUEPRINT_DB_PASSWORD psql -h $BLUEPRINT_DB_HOST -p $BLUEPRINT_DB_PORT -U $BLUEPRINT_DB_USERNAME -d $BLUEPRINT_DB_DATABASE -c "DROP TABLE IF EXISTS users;"

# Verify that the users table no longer exists
echo "Verifying users table no longer exists..."
PGPASSWORD=$BLUEPRINT_DB_PASSWORD psql -h $BLUEPRINT_DB_HOST -p $BLUEPRINT_DB_PORT -U $BLUEPRINT_DB_USERNAME -d $BLUEPRINT_DB_DATABASE -c "SELECT EXISTS (SELECT FROM information_schema.tables WHERE table_schema = 'public' AND table_name = 'users');"

# Check the schema_migrations table
echo "Checking schema_migrations table..."
PGPASSWORD=$BLUEPRINT_DB_PASSWORD psql -h $BLUEPRINT_DB_HOST -p $BLUEPRINT_DB_PORT -U $BLUEPRINT_DB_USERNAME -d $BLUEPRINT_DB_DATABASE -c "SELECT * FROM schema_migrations;"

echo "Migration state fixed!"