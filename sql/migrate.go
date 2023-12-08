package sql

import (
	"embed"
	"open-outcry/pkg/db"

	"github.com/pressly/goose/v3"
)

//go:embed *.sql
var embedMigrations embed.FS

func MigrateUp() error {
	goose.SetBaseFS(embedMigrations)
	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}
	if err := goose.Up(db.Instance().DB, "/"); err != nil {
		return err
	}

	return nil
}

func MigrateDown() error {
	goose.SetBaseFS(embedMigrations)
	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}
	if err := goose.DownTo(db.Instance().DB, "/", 0); err != nil {
		return err
	}

	return nil
}
