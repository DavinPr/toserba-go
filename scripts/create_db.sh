#!/bin/bash

# Database configuration
DB_HOST="localhost"
DB_PORT="5432"
DB_USER="postgres"
DB_PASSWORD="postgres"
DB_NAME="toserba"

# Create the database
echo "Creating database '$DB_NAME'..."
PGPASSWORD=$DB_PASSWORD psql -U $DB_USER -h $DB_HOST -p $DB_PORT -c "CREATE DATABASE $DB_NAME;"

# Check if the database was created successfully
if [ $? -eq 0 ]; then
  echo "Database '$DB_NAME' created successfully."
else
  echo "Failed to create database '$DB_NAME'."
  exit 1
fi
