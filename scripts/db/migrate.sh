#!/bin/bash

# Bash script to manage database migrations
# Usage:
# ./migrate.sh up - Apply all migrations
# ./migrate.sh down - Rollback the last migration
# ./migrate.sh create NAME - Create a new migration with the given name

set -e

# Set variables
MIGRATIONS_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)/db/migrations"
DB_HOST=${DB_HOST:-localhost}
DB_PORT=${DB_PORT:-5432}
DB_USER=${DB_USER:-secura}
DB_PASSWORD=${DB_PASSWORD:-securapassword}
DB_NAME=${DB_NAME:-secura}

# Function to create a new migration
create_migration() {
  local name=$1
  if [ -z "$name" ]; then
    echo "Error: Migration name is required"
    exit 1
  fi

  # Get the latest migration number
  latest_num=$(ls -1 $MIGRATIONS_DIR/*.sql 2>/dev/null | grep -oE '[0-9]+' | sort -n | tail -1)
  
  # Default to 1 if no migrations yet
  if [ -z "$latest_num" ]; then
    latest_num=0
  fi
  
  # Calculate the next number
  next_num=$((latest_num + 1))
  
  # Format with leading zeros
  next_num_padded=$(printf "%03d" $next_num)
  
  # Create the new migration file
  new_file="$MIGRATIONS_DIR/${next_num_padded}_${name}.sql"
  
  cat > "$new_file" << EOF
-- Migration: ${next_num_padded}_${name}

-- Up migration

-- Down migration

EOF

  echo "Created new migration: $new_file"
}

# Function to run migrations
run_migrations() {
  local direction=$1
  
  # Ensure PSQL is available
  if ! command -v psql >/dev/null 2>&1; then
    echo "Error: psql command not found"
    exit 1
  fi
  
  # Set the connection string
  PGPASSWORD=$DB_PASSWORD
  export PGPASSWORD
  
  # Database connection parameters
  conn_params="-h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME"
  
  if [ "$direction" = "up" ]; then
    # Create migrations table if it doesn't exist
    psql $conn_params -c "
      CREATE TABLE IF NOT EXISTS migrations (
        id SERIAL PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
        applied_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
      );
    "
    
    # Get already applied migrations
    applied_migrations=$(psql $conn_params -t -c "SELECT name FROM migrations;" | tr -d ' ')
    
    # Apply each migration that hasn't been applied yet
    for migration in $(ls -1 $MIGRATIONS_DIR/*.sql | sort); do
      filename=$(basename "$migration")
      if ! echo "$applied_migrations" | grep -q "$filename"; then
        echo "Applying migration: $filename"
        
        # Start a transaction
        psql $conn_params -c "BEGIN;"
        
        # Apply the migration - only the up part
        sed -n '/-- Up migration/,/-- Down migration/p' "$migration" | 
          grep -v "-- Down migration" |
          psql $conn_params -f -
        
        # Record the migration
        psql $conn_params -c "INSERT INTO migrations (name) VALUES ('$filename');"
        
        # Commit the transaction
        psql $conn_params -c "COMMIT;"
        
        echo "Migration applied: $filename"
      else
        echo "Migration already applied: $filename"
      fi
    done
  elif [ "$direction" = "down" ]; then
    # Get the last applied migration
    last_migration=$(psql $conn_params -t -c "SELECT name FROM migrations ORDER BY id DESC LIMIT 1;" | tr -d ' ')
    
    if [ -n "$last_migration" ]; then
      echo "Rolling back migration: $last_migration"
      
      # Start a transaction
      psql $conn_params -c "BEGIN;"
      
      # Apply the down migration
      sed -n '/-- Down migration/,/$/p' "$MIGRATIONS_DIR/$last_migration" |
        psql $conn_params -f -
      
      # Remove the migration record
      psql $conn_params -c "DELETE FROM migrations WHERE name = '$last_migration';"
      
      # Commit the transaction
      psql $conn_params -c "COMMIT;"
      
      echo "Migration rolled back: $last_migration"
    else
      echo "No migrations to roll back"
    fi
  else
    echo "Error: Invalid direction. Use 'up' or 'down'."
    exit 1
  fi
}

# Main execution
case "$1" in
  up)
    run_migrations up
    ;;
  down)
    run_migrations down
    ;;
  create)
    create_migration "$2"
    ;;
  *)
    echo "Usage:"
    echo "  $0 up - Apply all migrations"
    echo "  $0 down - Rollback the last migration"
    echo "  $0 create NAME - Create a new migration with the given name"
    exit 1
    ;;
esac 