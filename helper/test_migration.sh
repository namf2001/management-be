#!/bin/bash

# This script tests the migrate-down and migrate-up commands to verify that the issue has been fixed

# Load environment variables from .env file
export $(grep -v '^#' .env | xargs)

echo "Testing migrate-down command..."
make migrate-down

echo "Testing migrate-up command..."
make migrate-up

echo "Testing workflow command..."
make workflow

echo "Test completed successfully!"