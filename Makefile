
.DEFAULT_GOAL := help
.PHONY: help

setup:
	go install golang.org/x/tools/cmd/goimports@latest
	go install honnef.co/go/tools/cmd/staticcheck@latest
	go install github.com/pressly/goose/v3/cmd/goose@latest
	npm i
	go get ./...

build: ## Installs and compiles dependencies
	go build -v ./...

run: ## Start dev mode
	go run main.go

test:
	go test ./... -v -cover -p 1

lint:
	go fmt ./...
	goimports -l -w .
	staticcheck ./...
	go vet ./...

include ./pkg/conf/dev.env
DBDSN:="host=$(POSTGRES_HOST) user=$(POSTGRES_USER) password=$(POSTGRES_PASSWORD) dbname=$(POSTGRES_DB) port=$(POSTGRES_PORT) sslmode=disable"
MIGRATE_OPTIONS=-allow-missing -dir="./pkg/db/migrations"

db-update: ## Migrate down on database
	goose -v $(MIGRATE_OPTIONS) postgres $(DBDSN) up

db-downgrade: ## Migrate up on database
	echo "$(MIGRATE_OPTIONS)"
	goose -v $(MIGRATE_OPTIONS) postgres $(DBDSN) reset

db-rebuild: ## Reset the database
	$(MAKE) db-downgrade
	$(MAKE) db-update

build-api: ## Build OpenAPI
	node node_modules/swagger-cli/swagger-cli.js bundle --dereference \
 		-o pkg/static/docs/api/openapi.json \
 		-t json \
 		-r openapi/openapi.yaml

generate-api: ## Generate server bindings
	npx @openapitools/openapi-generator-cli generate \
		-i openapi/openapi.yaml \
		-g go-server \
		-o pkg/rest \
		--additional-properties=packageName=api \
		--additional-properties=sourceFolder=api \
		--additional-properties=outputAsLibrary=true
	$(MAKE) lint

help:
	grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) \
	| sed -n 's/^\(.*\): \(.*\)##\(.*\)/\1\3/p'