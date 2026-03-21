# Migration Manifest

The migration order is defined explicitly in [manifest.go](./manifest.go).

## How it works

1. Source SQL files live next to the Go packages that use them, under `pkg/...`.
2. `MigrationSources()` returns those source files in dependency order.
3. `go run ./cmd/migrationgen` reads that list and writes numbered Goose files into `sql/generated/`.
4. `make db-up` and the in-app migration entrypoint both execute `sql/generated/`, so they share the same ordering.

## Rule for changes

When you add, remove, or reorder migrations, update `migrationSources` in `manifest.go`.
Directory layout and file names under `pkg/` do not affect execution order by themselves.
