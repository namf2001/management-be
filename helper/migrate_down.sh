#!/bin/bash

# Load environment variables from .env file
export $(grep -v '^#' .env | xargs)

# Check if schema_migrations table exists
migrations_exist=$(PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -p $DB_POST -U $DB_USER -d $DB_NAME -t -c "SELECT EXISTS (SELECT FROM information_schema.tables WHERE table_schema = 'public' AND table_name = 'schema_migrations');")

if [ "$migrations_exist" = " t" ]; then
    # Check if there are any rows in the schema_migrations table
    row_count=$(PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -p $DB_POST -U $DB_USER -d $DB_NAME -t -c "SELECT COUNT(*) FROM schema_migrations;")
    
    # Trim spaces from row_count
    row_count_trimmed=$(echo "$row_count" | tr -d '[:space:]')
    
    echo "Row count in schema_migrations table: '$row_count_trimmed'"
    
    if [ "$row_count_trimmed" = "0" ]; then
        echo "Schema_migrations table exists but has no rows. Nothing to roll back."
    else
        # Check if version is greater than 0
        version=$(PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -p $DB_POST -U $DB_USER -d $DB_NAME -t -c "SELECT version FROM schema_migrations LIMIT 1;")
        
        # Trim spaces from version
        version_trimmed=$(echo "$version" | tr -d '[:space:]')
        
        echo "Detected version: '$version', Trimmed version: '$version_trimmed'"
        
        if [ "$version_trimmed" = "0" ]; then
            echo "Current version is 0. Nothing to roll back."
        else
            echo "Rolling back migration from version $version_trimmed..."
            migrate -path=./data/migrations -database "postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_POST}/${DB_NAME}?sslmode=disable" down 1
        fi
    fi
else
    echo "Schema_migrations table does not exist. Nothing to roll back."
fi