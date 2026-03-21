package sql

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"open-outcry/pkg/db"

	log "github.com/sirupsen/logrus"

	"github.com/pressly/goose/v3"
)

const migrationDirEnv = "OPEN_OUTCRY_MIGRATION_DIR"

var (
	migrationDirMu        sync.Mutex
	activeMigrationDir    string
	cleanupMigrationDirFn func()
)

func configureGoose() {
	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}
}

func MigrateUp() error {
	log.Info("Migrate up")
	configureGoose()

	migrationDir, err := ensureMigrationDir()
	if err != nil {
		return err
	}
	if err := goose.Up(db.Instance().DB, migrationDir); err != nil {
		panic(err)
	}

	return nil
}

func MigrateDown() error {
	configureGoose()

	migrationDir, err := ensureMigrationDir()
	if err != nil {
		return err
	}
	if err := goose.DownTo(db.Instance().DB, migrationDir, 0); err != nil {
		panic(err)
	}

	releaseMigrationDir()
	return nil
}

func ensureMigrationDir() (string, error) {
	migrationDirMu.Lock()
	defer migrationDirMu.Unlock()

	if activeMigrationDir != "" {
		return activeMigrationDir, nil
	}

	if configuredDir := os.Getenv(migrationDirEnv); configuredDir != "" {
		activeMigrationDir = filepath.Clean(configuredDir)
		if _, err := os.Stat(activeMigrationDir); err != nil {
			return "", fmt.Errorf("migration dir %s: %w", activeMigrationDir, err)
		}
		cleanupMigrationDirFn = nil
		return activeMigrationDir, nil
	}

	generatedDir, cleanup, err := CreateTempMigrations()
	if err != nil {
		return "", err
	}
	activeMigrationDir = generatedDir
	cleanupMigrationDirFn = cleanup
	return activeMigrationDir, nil
}

func releaseMigrationDir() {
	migrationDirMu.Lock()
	defer migrationDirMu.Unlock()

	if cleanupMigrationDirFn != nil {
		cleanupMigrationDirFn()
	}
	activeMigrationDir = ""
	cleanupMigrationDirFn = nil
}
