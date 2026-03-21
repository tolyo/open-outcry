include api/api.mk
include demo/demo.mk

.DEFAULT_GOAL := help
.PHONY: help ensure-goimports ensure-staticcheck

TEST_PACKAGES := $(shell go list -f '{{if or .TestGoFiles .XTestGoFiles}}{{.ImportPath}}{{end}}' ./... | sed '/^$$/d')
GOBIN := $(shell go env GOPATH)/bin
GOIMPORTS_BIN := $(GOBIN)/goimports
STATICCHECK_BIN := $(GOBIN)/staticcheck
GO_VERSION := $(shell go env GOVERSION)
MIGRATION_DIR_ENV := OPEN_OUTCRY_MIGRATION_DIR
ENV ?= DEV

setup:
	go install golang.org/x/tools/cmd/goimports@latest
	GOTOOLCHAIN=$(GO_VERSION) go install honnef.co/go/tools/cmd/staticcheck@latest
	go install github.com/pressly/goose/v3/cmd/goose@latest
	go get ./...
	(cd api && npm i)

build: ## Installs and compiles dependencies
	go build -v ./...

run: ## Start dev mode
	make db-up
	go run main.go

test:
	tmpdir=$$(mktemp -d); \
	trap 'rm -rf "$$tmpdir"' EXIT; \
	ENV=DEV go run ./cmd/migrationgen -out "$$tmpdir"; \
	ENV=DEV $(MIGRATION_DIR_ENV)="$$tmpdir" go test $(TEST_PACKAGES) -v -cover -p 1

ensure-goimports:
	@if [ ! -x "$(GOIMPORTS_BIN)" ]; then \
		go install golang.org/x/tools/cmd/goimports@latest; \
	fi

ensure-staticcheck:
	@if [ ! -x "$(STATICCHECK_BIN)" ]; then \
		GOTOOLCHAIN=$(GO_VERSION) go install honnef.co/go/tools/cmd/staticcheck@latest; \
	elif [ "$$(go version -m "$(STATICCHECK_BIN)" 2>/dev/null | sed -n '1s/.*: //p')" != "$(GO_VERSION)" ]; then \
		GOTOOLCHAIN=$(GO_VERSION) go install honnef.co/go/tools/cmd/staticcheck@latest; \
	fi

lint: ensure-goimports ensure-staticcheck
	go fmt ./...
	$(GOIMPORTS_BIN) -l -w .
	$(STATICCHECK_BIN) ./...
	go vet ./...

include ./pkg/conf/dev.env
DB_DSN:="host=$(POSTGRES_HOST) user=$(POSTGRES_USER) password=$(POSTGRES_PASSWORD) dbname=$(POSTGRES_DB) port=$(POSTGRES_PORT) sslmode=disable"
MIGRATE_OPTIONS=-allow-missing

db-up: ## Migrate database up
	tmpdir=$$(mktemp -d); \
	trap 'rm -rf "$$tmpdir"' EXIT; \
	ENV=$(ENV) go run ./cmd/migrationgen -out "$$tmpdir"; \
	goose -v $(MIGRATE_OPTIONS) -dir="$$tmpdir" postgres $(DB_DSN) up

db-down: ## Reset database migrations
	tmpdir=$$(mktemp -d); \
	trap 'rm -rf "$$tmpdir"' EXIT; \
	ENV=$(ENV) go run ./cmd/migrationgen -out "$$tmpdir"; \
	goose -v $(MIGRATE_OPTIONS) -dir="$$tmpdir" postgres $(DB_DSN) reset

db-rebuild: ## Reset the database
	make db-down
	make db-up

help:
	grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) \
	| sed -n 's/^\(.*\): \(.*\)##\(.*\)/\1\3/p'
