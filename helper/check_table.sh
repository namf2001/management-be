#!/bin/bash

# Load environment variables from .env file
export $(grep -v '^#' .env | xargs)

# Connect to the database and check if the users table exists
PGPASSWORD=$BLUEPRINT_DB_PASSWORD psql -h $BLUEPRINT_DB_HOST -p $BLUEPRINT_DB_PORT -U $BLUEPRINT_DB_USERNAME -d $BLUEPRINT_DB_DATABASE -c "SELECT EXISTS (SELECT FROM information_schema.tables WHERE table_schema = 'public' AND table_name = 'users');"