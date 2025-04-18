#!/bin/bash

# Load environment variables from .env file
export $(grep -v '^#' .env | xargs)

# Connect to the database and drop the users table if it exists
PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -p $DB_POST -U $DB_USER -d $DB_NAME -c "DROP TABLE IF EXISTS users;"

# Also drop the schema_migrations table to reset the migration state
PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -p $DB_POST -U $DB_USER -d $DB_NAME -c "DROP TABLE IF EXISTS schema_migrations;"