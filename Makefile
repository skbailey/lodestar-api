.PHONY: migrate.new migrate.up migrate.down

SHELL := /bin/bash

environment = dev
ifdef env
	environment = $(env)
endif

migration_name = new_migration
ifdef migration_name
	migration_name_option = migration_name
endif

migrate.new:
	@migrate create -ext sql -dir migrations -seq $(migration_name_option)

ssm_db_path = /${environment}/lodestar/db/
ssm_results = $(shell aws ssm get-parameters-by-path --region us-east-1 --path ${ssm_db_path})
db_params = $(shell echo '${ssm_results}' | jq -c 'reduce .Parameters[] as $$o ({}; .[$$o.Name] = $$o.Value)')

host = $(shell jq -r '."/${environment}/lodestar/db/host"' <<< '${db_params}')
port = $(shell jq -r '."/${environment}/lodestar/db/port"' <<< '${db_params}')
name = $(shell jq -r '."/${environment}/lodestar/db/name"' <<< '${db_params}')
password = $(shell jq -r '."/${environment}/lodestar/db/password"' <<< '${db_params}')
username = $(shell jq -r '."/${environment}/lodestar/db/username"' <<< '${db_params}')

db_url = "postgres://${username}:${password}@${host}:${port}/${name}?sslmode=disable"

migrate.up:
	@migrate -path migrations -database $(db_url) up

migrate.down:
	@migrate -path migrations -database $(db_url) down
