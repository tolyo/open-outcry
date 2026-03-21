# Migration Manifest

`manifest.go` is the single source of truth for migration ordering.

## How It Works

1. SQL files live next to the Go code that owns them under `pkg/...`.
2. `sql/manifest/manifest.go` lists the base migration files in dependency order.
3. `go run ./cmd/migrationgen -out <dir>` reads that list and writes numbered Goose files into the provided temp directory.
4. `make test`, `make db-up`, and `make db-down` generate that temp directory immediately before execution and remove it afterward.
5. The in-app migration entrypoint in `sql/migrate.go` also generates temp numbered migrations on demand when no migration directory is provided.

## Environment-Specific Seeds

`pkg/conf/seeds_dev.sql` is appended only when `ENV=DEV`.

- `make db-up` and `make db-down` default to `ENV=DEV`, so local development gets seed data.
- `make test` currently generates migrations with `ENV=DEV`, because the existing test suite depends on the baseline seeded data.
- If `ENV` is unset, migration generation defaults to `DEV`, matching local app startup and direct `go test` usage.
- If `ENV` is set to a non-listed value, no environment-specific seed files are added.

## Ordering Rule

Keep manifest entries in execution order, not alphabetic order. File location under `pkg/` does not affect migration order.
