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
        echo "Schema_migrations table exists but has no rows. Dropping schema_migrations table to allow clean migration..."
        PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -p $DB_POST -U $DB_USER -d $DB_NAME -c "DROP TABLE IF EXISTS schema_migrations;"
        echo "Schema_migrations table dropped. You can now run migrations again."
    else
        # Check if version is 0
        version=$(PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -p $DB_POST -U $DB_USER -d $DB_NAME -t -c "SELECT version FROM schema_migrations LIMIT 1;")

        # Trim spaces from version
        version_trimmed=$(echo "$version" | tr -d '[:space:]')

        echo "Detected version: '$version', Trimmed version: '$version_trimmed'"

        if [ "$version_trimmed" = "0" ]; then
            echo "Detected version 0 in schema_migrations table. Dropping schema_migrations table to allow clean migration..."
            PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -p $DB_POST -U $DB_USER -d $DB_NAME -c "DROP TABLE IF EXISTS schema_migrations;"
            echo "Schema_migrations table dropped. You can now run migrations again."
        else
            echo "Schema_migrations table exists with version $version_trimmed. No action needed."
        fi
    fi
else
    echo "Schema_migrations table does not exist. No action needed."
fi
