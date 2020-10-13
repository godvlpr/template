PATH_TO_MIGRATIONS=storage/db/migrations

POSTGRESQL_URL='postgres://$(TEMPLATE_DATABASE_USER):$(TEMPLATE_DATABASE_PASSWORD)@$(TEMPLATE_DATABASE_HOST):$(TEMPLATE_DATABASE_PORT)/$(TEMPLATE_DATABASE_NAME)?sslmode=$(TEMPLATE_DATABASE_SSL)'

generate-docs:
	swag init -g cmd/main.go

compose-up:
	docker-compose up &

create-migration:
	migrate create -ext sql -dir $(PATH_TO_MIGRATIONS) $(table_name)

migration-up:
	migrate -database ${POSTGRESQL_URL} -path $(PATH_TO_MIGRATIONS) up

migration-down:
	migrate -database ${POSTGRESQL_URL} -path $(PATH_TO_MIGRATIONS) down

format:
	go fmt github.com/godvlpr/template/...

build:
	go build ./cmd/main.go

run:
	go run ./cmd/main.go