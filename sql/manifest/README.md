# Migration Manifest

`manifest.go` is the single source of truth for migration ordering.

## How It Works

1. SQL files live next to the Go code that owns them under `pkg/...`.
2. `sql/manifest/manifest.go` lists those source files in dependency order.
3. `go run ./cmd/migrationgen -out <dir>` reads that list and writes numbered Goose files into the provided temp directory.
4. `make test`, `make db-up`, and `make db-down` generate that temp directory immediately before execution and remove it afterward.
5. The in-app migration entrypoint in `sql/migrate.go` also generates temp numbered migrations on demand when no migration directory is provided.

## Ordering Rule

Keep manifest entries in execution order, not alphabetic order. File location under `pkg/` does not affect migration order.
