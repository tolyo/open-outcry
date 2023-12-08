package sql

import (
	"embed"
	"open-outcry/pkg/db"

	log "github.com/sirupsen/logrus"

	"github.com/pressly/goose/v3"
)

//go:embed *.sql
var embedMigrations embed.FS

func MigrateUp() error {
	log.Info("Migrate up")
	goose.SetBaseFS(embedMigrations)
	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}
	if err := goose.Up(db.Instance().DB, "."); err != nil {
		panic(err)
	}

	return nil
}

func MigrateDown() error {
	goose.SetBaseFS(embedMigrations)
	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}
	if err := goose.DownTo(db.Instance().DB, ".", 0); err != nil {
		panic(err)
	}

	return nil
}
