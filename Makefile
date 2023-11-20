
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
	@goose -v $(MIGRATE_OPTIONS) postgres $(DBDSN) up

db-downgrade: ## Migrate up on database
	@goose -v $(MIGRATE_OPTIONS) postgres $(DBDSN) reset

db-rebuild: ## Reset the database
	@make db-downgrade
	@make db-update

validate-api: ## Validate api
	npx @openapitools/openapi-generator-cli validate \
		-i pkg/openapi.yaml \
		--recommend

bundle-api:
	@npx @redocly/cli bundle \
		pkg/api/openapi.yaml \
		--output ./pkg/api.yaml \

validate-api:
	@npx @redocly/cli lint \
		./pkg/api.yaml \
		--format=checkstyle

generate-api: ## Generate server bindings, move model files, fix imports
	npx @openapitools/openapi-generator-cli generate \
		-i pkg/api.yaml \
		-g go-server \
		-o pkg/rest \
		--additional-properties=packageName=api \
		--additional-properties=sourceFolder=api \
		--additional-properties=outputAsLibrary=true
	@make lint

help:
	grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) \
	| sed -n 's/^\(.*\): \(.*\)##\(.*\)/\1\3/p'
