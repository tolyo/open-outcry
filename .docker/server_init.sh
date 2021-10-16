#!/bin/bash
# Docker entrypoint script.

# Wait until Postgres is ready
while ! pg_isready -q -h $POSTGRES_HOST -p $POSTGRES_PORT -U $POSTGRES_USER
do
  echo "$(date) - waiting for database to start"
  sleep 2
done

if [[ -z `psql -Atqc "\\list $POSTGRES_TEST_DB"` ]]; then
  echo "Database $POSTGRES_TEST_DB does not exist. Creating..."
  createdb -E UTF8 $POSTGRES_TEST_DB -l en_US.UTF-8 -T template0
  echo "Database $POSTGRES_TEST_DB created."
fi

make compile
make db-update
make run-dev
