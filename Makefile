include api/api.mk
include demo/demo.mk

.DEFAULT_GOAL := help
.PHONY: help

setup:
	go install golang.org/x/tools/cmd/goimports@latest
	go install honnef.co/go/tools/cmd/staticcheck@latest
	go install github.com/pressly/goose/v3/cmd/goose@latest
	go get ./...
	(cd api && npm i)

build: ## Installs and compiles dependencies
	go build -v ./...

run: ## Start dev mode
	make db-up
	go run main.go

test:
	go test ./... -v -cover -p 1

lint:
	go fmt ./...
	goimports -l -w .
	staticcheck ./...
	go vet ./...

include ./pkg/conf/dev.env
DB_DSN:="host=$(POSTGRES_HOST) user=$(POSTGRES_USER) password=$(POSTGRES_PASSWORD) dbname=$(POSTGRES_DB) port=$(POSTGRES_PORT) sslmode=disable"
MIGRATE_OPTIONS=-allow-missing -dir="./sql"

db-up: ## Migrate down on database
	goose -v $(MIGRATE_OPTIONS) postgres $(DB_DSN) up

db-down: ## Migrate up on database
	goose -v $(MIGRATE_OPTIONS) postgres $(DB_DSN) reset

db-rebuild: ## Reset the database
	make db-down
	make db-up

help:
	grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) \
	| sed -n 's/^\(.*\): \(.*\)##\(.*\)/\1\3/p'
