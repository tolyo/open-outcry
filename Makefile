
.DEFAULT_GOAL := help
.PHONY: help

setup:
	go install golang.org/x/tools/cmd/goimports@latest
	go install honnef.co/go/tools/cmd/staticcheck@latest
	go install github.com/pressly/goose/v3/cmd/goose@latest
	go install github.com/cosmtrek/air@latest
	npm i
	go get ./...

build: ## Installs and compiles dependencies
	go build -v ./...

run: ## Start dev mode
	make db-up
	air main.go

test:
	go test ./... -v -cover -p 1

lint:
	go fmt ./...
	goimports -l -w .
	staticcheck ./...
	go vet ./...

include ./pkg/conf/dev.env
DBDSN:="host=$(POSTGRES_HOST) user=$(POSTGRES_USER) password=$(POSTGRES_PASSWORD) dbname=$(POSTGRES_DB) port=$(POSTGRES_PORT) sslmode=disable"
MIGRATE_OPTIONS=-allow-missing -dir="./sql"

db-up: ## Migrate down on database
	goose -v $(MIGRATE_OPTIONS) postgres $(DBDSN) up

db-down: ## Migrate up on database
	goose -v $(MIGRATE_OPTIONS) postgres $(DBDSN) reset

db-rebuild: ## Reset the database
	make db-down
	make db-up

OUTPUT_YAML:="./pkg/static/api.yaml"

bundle-api:
	npx @redocly/cli bundle \
		pkg/api/openapi.yaml \
		--output $(OUTPUT_YAML)

validate-api:
	make bundle-api
	npx @redocly/cli lint \
		$(OUTPUT_YAML) \
		--format=checkstyle

generate-api: ## Generate server bindings, move model files, fix imports
	make validate-api
	npx @openapitools/openapi-generator-cli generate \
		-i $(OUTPUT_YAML) \
		-g go-server \
		-o pkg/rest \
		--additional-properties=packageName=api \
		--additional-properties=sourceFolder=api \
		--additional-properties=outputAsLibrary=true
	make lint

help:
	grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) \
	| sed -n 's/^\(.*\): \(.*\)##\(.*\)/\1\3/p'
