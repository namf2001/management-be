#!/bin/bash

# Load environment variables from .env file
export $(grep -v '^#' .env | xargs)

# Connect to the database and check the migration version
PGPASSWORD=$BLUEPRINT_DB_PASSWORD psql -h $BLUEPRINT_DB_HOST -p $BLUEPRINT_DB_PORT -U $BLUEPRINT_DB_USERNAME -d $BLUEPRINT_DB_DATABASE -c "SELECT * FROM schema_migrations;"