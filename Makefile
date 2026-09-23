CONNECTION_STRING = "postgres://postgres:postgres@localhost:5432/gator"

reset:
	sqlc generate
	cd sql/schema && goose postgres $(CONNECTION_STRING) down && goose postgres $(CONNECTION_STRING) up

up:
	sqlc generate
	cd sql/schema && goose postgres $(CONNECTION_STRING) up

down:
	cd sql/schema && goose postgres $(CONNECTION_STRING) down
