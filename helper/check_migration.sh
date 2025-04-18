#!/bin/bash

# Load environment variables from .env file
export $(grep -v '^#' .env | xargs)

# Connect to the database and check the migration version
PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -p $DB_POST -U $DB_USER -d $DB_NAME -c "SELECT * FROM schema_migrations;"