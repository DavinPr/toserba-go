#!/bin/bash

# Database configuration
DB_HOST="localhost"
DB_PORT="5432"
DB_USER="postgres"
DB_PASSWORD="postgres"
DB_NAME="toserba"

# Terminate all connections to the database
echo "Terminating connections to database '$DB_NAME'..."
PGPASSWORD=$DB_PASSWORD psql -U $DB_USER -h $DB_HOST -p $DB_PORT -d postgres -c "SELECT pg_terminate_backend(pg_stat_activity.pid) FROM pg_stat_activity WHERE pg_stat_activity.datname = '$DB_NAME' AND pid <> pg_backend_pid();"

# Drop the database
echo "Dropping database '$DB_NAME'..."
PGPASSWORD=$DB_PASSWORD psql -U $DB_USER -h $DB_HOST -p $DB_PORT -c "DROP DATABASE IF EXISTS $DB_NAME;"

# Check if the database was dropped successfully
if [ $? -eq 0 ]; then
  echo "Database '$DB_NAME' dropped successfully."
else
  echo "Failed to drop database '$DB_NAME'."
  exit 1
fi
