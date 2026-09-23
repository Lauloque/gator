CONNECTION_STRING = "postgres://postgres:postgres@localhost:5432/gator"

reset:
	cd sql/schema && goose postgres $(CONNECTION_STRING) down && goose postgres $(CONNECTION_STRING) up

up:
	goose postgres $(CONNECTION_STRING) up

down:
	goose postgres $(CONNECTION_STRING) down
