#!/bin/bash

# Load environment variables from .env file
export $(grep -v '^#' .env | xargs)

# Drop the users table manually
echo "Dropping users table..."
PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -p $DB_POST -U $DB_USER -d $DB_NAME -c "DROP TABLE IF EXISTS users;"

# Verify that the users table no longer exists
echo "Verifying users table no longer exists..."
PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -p $DB_POST -U $DB_USER -d $DB_NAME -c "SELECT EXISTS (SELECT FROM information_schema.tables WHERE table_schema = 'public' AND table_name = 'users');"

# Check the schema_migrations table
echo "Checking schema_migrations table..."
PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -p $DB_POST -U $DB_USER -d $DB_NAME -c "SELECT * FROM schema_migrations;"

echo "Migration state fixed!"