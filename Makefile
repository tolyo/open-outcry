
.DEFAULT_GOAL := help
.PHONY: help

setup: ## Installs and compiles dependencies
	@go install github.com/pressly/goose/v3/cmd/goose@latest 

run: ## Start dev mode
	@go run main.go

DB_DSN:=$$(yq e '.DB_DSN' /pkg/conf/dev.yaml)
MIGRATE_OPTIONS=-dir="/pkg/db/migrations"

db-update: ## Migrate down on database
	@goose -v $(MIGRATE_OPTIONS) postgres "$(DB_DSN)" up

db-downgrade: ## Migrate up on database
	@goose -v $(MIGRATE_OPTIONS) postgres "$(DB_DSN)" down

db-rebuild: ## Reset the database
	$(MAKE) db-downgrade 
	$(MAKE) db-update

build-api: ## Build OpenAPI
	node node_modules/swagger-cli/swagger-cli.js bundle -o static/docs/api/openapi.json -t json -r lib/web/openapi/openapi.yaml

help:
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) \
	| sed -n 's/^\(.*\): \(.*\)##\(.*\)/\1\3/p' 