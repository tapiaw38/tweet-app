#!/bin/sh

# Wait for PostgreSQL to be ready
until pg_isready --host=tweet-postgres-db --port=5432 --username=postgres --dbname=tweet-db
do
  echo "Waiting for PostgreSQL to be ready..."
  sleep 2
done

# Both PostgreSQL and RabbitMQ are now ready, so run the Go application
echo "========== Starting Go application =========="
exec go run cmd/api/main.go --host 0.0.0.0