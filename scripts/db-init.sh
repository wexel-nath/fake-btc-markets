#!/bin/bash
set -e

# check for required vars
DB_HOST="${DB_HOST:?DB_HOST must be set}"
DB_NAME="${DB_NAME:?DB_NAME must be set}"
DB_PASS="${DB_PASS:?DB_PASS must be set}"
DB_PORT="${DB_PORT:?DB_PORT must be set}"
DB_USER="${DB_USER:?DB_USER must be set}"
POSTGRES_PASS="${POSTGRES_PASS:?POSTGRES_PASS must be set}"
POSTGRES_USER="${POSTGRES_USER:?POSTGRES_USER must be set}"

connect() {
	psql \
		-h "$DB_HOST" \
		-p "$DB_PORT" \
		-U "$DB_USER" \
		"$@"
}

run_query() {
	connect -d "$DB_NAME" -tAc "$1"
}

run_sql_file() {
	connect -d "$DB_NAME" -f "$1"
}

run_query_as_postgres() {
	psql \
		-h "$DB_HOST" \
		-p "$DB_PORT" \
		-U "$POSTGRES_USER" \
		-tAc "$@"
}

is_up() {
	run_query_as_postgres "SELECT 1" || return 1
}

wait_for_postgres() {
	echo "Waiting for postgres..."
	for i in {1..60}; do
		if [[ $(is_up) ]]; then
			echo "Postgres is ready"
			return
		fi
		echo "Postgres is not ready... waiting"
		sleep 5
	done
	echo "Postgres is not ready... exiting"
	exit 1
}

maybe_create_user() {
	echo "Creating user $DB_USER"

	if [[ $(run_query_as_postgres "SELECT 1 FROM pg_user WHERE usename = '$DB_USER';") -eq 1 ]]; then
		echo "User '$DB_USER' exists"
	else
		run_query_as_postgres "CREATE ROLE $DB_USER LOGIN PASSWORD '$DB_PASS';"
	fi
	echo
}

maybe_create_database() {
	echo "Creating database"

	if [[ $(run_query_as_postgres "SELECT 1 FROM pg_database WHERE datname = '$DB_NAME';") -eq 1 ]]; then
		echo "Database '$DB_NAME' exists"
	else
		run_query_as_postgres "CREATE DATABASE $DB_NAME;"
		run_query_as_postgres "GRANT ALL PRIVILEGES ON DATABASE $DB_NAME TO $DB_USER;"
	fi
	echo
}

maybe_create_schema() {
	echo "Creating database schema"
	for file in /sql/schema/*.sql; do
		if [[ "$file" =~ _(.+).sql ]]; then
			table_name=${BASH_REMATCH[1]}
		else
			continue
		fi

		echo "Running $file"
		if [[ $(run_query "SELECT 1 FROM pg_tables WHERE tablename = '$table_name';") -eq 1 ]]; then
			echo "Table '$table_name' exists"
		else
			run_sql_file "$file"
		fi
		echo
	done
}

maybe_run_updates() {
	echo "Running database migrations"
	for file in /sql/updates/*.sql; do
		if [[ $(run_query "SELECT 1 FROM update WHERE update_id = '$file';") -eq 1 ]]; then
			continue
		fi

		echo "Running $file"
		run_sql_file "$file"

		echo "Saving update"
		run_query "INSERT INTO update (update_id) VALUES ('$file');"
		echo
	done
}

# psql requires a password file
export PGPASSFILE='.pgpass'
echo "*:$DB_PORT:*:$POSTGRES_USER:$POSTGRES_PASS" > "$PGPASSFILE"
chmod 0600 "$PGPASSFILE"

wait_for_postgres

maybe_create_user
maybe_create_database

# create schema and updates as DB_USER:DB_PASS
echo "*:$DB_PORT:$DB_NAME:$DB_USER:$DB_PASS" > "$PGPASSFILE"

maybe_create_schema
maybe_run_updates

echo "Database initialization complete."

exit 0
