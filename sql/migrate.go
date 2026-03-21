package sql

import (
	"embed"
	"open-outcry/pkg/db"

	log "github.com/sirupsen/logrus"

	"github.com/pressly/goose/v3"
)

const generatedMigrationDir = "generated"

//go:embed generated/*.sql
var embedMigrations embed.FS

func configureGoose() {
	goose.SetBaseFS(embedMigrations)
	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}
}

func MigrateUp() error {
	log.Info("Migrate up")
	configureGoose()
	if err := goose.Up(db.Instance().DB, generatedMigrationDir); err != nil {
		panic(err)
	}

	return nil
}

func MigrateDown() error {
	configureGoose()
	if err := goose.DownTo(db.Instance().DB, generatedMigrationDir, 0); err != nil {
		panic(err)
	}

	return nil
}
