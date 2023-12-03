package db

import (
	"embed"

	"github.com/pressly/goose/v3"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func MigrateUp() error {
	goose.SetBaseFS(embedMigrations)
	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}
	if err := goose.Up(Instance().DB, "migrations"); err != nil {
		return err
	}

	return nil
}

func MigrateDown() error {
	goose.SetBaseFS(embedMigrations)
	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}
	if err := goose.DownTo(Instance().DB, "migrations", 0); err != nil {
		return err
	}

	return nil
}
