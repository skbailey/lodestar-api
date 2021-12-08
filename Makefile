.PHONY: migrate.new migrate.up migrate.down

migration name = "new_migration"
ifdef migration_name
	migration_name_option = migration_name
endif

migrate.new:
	@migrate create -ext sql -dir migrations -seq $(migration_name_option)

ssm_db_path := "/dev/lodestar/db/"
db_params := $(shell aws ssm get-parameters-by-path --path ${ssm_db_path} | jq -c 'reduce .Parameters[] as $$o ({}; .[$$o.Name] = $$o.Value)')

host = $(shell jq -r '."/dev/lodestar/db/host"' <<< '${db_params}')
port = $(shell jq -r '."/dev/lodestar/db/port"' <<< '${db_params}')
name = $(shell jq -r '."/dev/lodestar/db/name"' <<< '${db_params}')
password = $(shell jq -r '."/dev/lodestar/db/password"' <<< '${db_params}')
username = $(shell jq -r '."/dev/lodestar/db/username"' <<< '${db_params}')

db_url = "postgres://${username}:${password}@${host}:${port}/${name}?sslmode=disable"

migrate.up:
	@migrate -path migrations -database $(db_url) up

migrate.down:
	@migrate -path migrations -database $(db_url) down
