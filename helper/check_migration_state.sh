#!/bin/bash

# Load environment variables from .env file
export $(grep -v '^#' .env | xargs)

# Connect to the database and check the schema_migrations table
echo "Checking schema_migrations table..."
PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -p $DB_POST -U $DB_USER -d $DB_NAME -c "SELECT * FROM schema_migrations;"

# Check if the users table exists
echo "Checking if users table exists..."
PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -p $DB_POST -U $DB_USER -d $DB_NAME -c "SELECT EXISTS (SELECT FROM information_schema.tables WHERE table_schema = 'public' AND table_name = 'schema_migrations');"