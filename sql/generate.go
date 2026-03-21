package sql

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	manifest "open-outcry/sql/manifest"
)

func GenerateMigrations(outputDir string) (int, error) {
	if err := os.RemoveAll(outputDir); err != nil {
		return 0, err
	}
	if err := os.MkdirAll(outputDir, 0o755); err != nil {
		return 0, err
	}

	sources := manifest.MigrationSources()
	for index, source := range sources {
		if err := writeGeneratedMigration(outputDir, index, source); err != nil {
			return 0, err
		}
	}

	return len(sources), nil
}

func CreateTempMigrations() (string, func(), error) {
	dir, err := os.MkdirTemp("", "open-outcry-migrations-*")
	if err != nil {
		return "", nil, err
	}

	cleanup := func() {
		_ = os.RemoveAll(dir)
	}

	if _, err := GenerateMigrations(dir); err != nil {
		cleanup()
		return "", nil, err
	}

	return dir, cleanup, nil
}

func writeGeneratedMigration(outputDir string, index int, source string) error {
	cleanSource := filepath.Clean(filepath.FromSlash(source))
	if cleanSource == "." || cleanSource == ".." || strings.HasPrefix(cleanSource, ".."+string(filepath.Separator)) {
		return fmt.Errorf("invalid migration source path: %s", source)
	}
	if filepath.Ext(cleanSource) != ".sql" {
		return fmt.Errorf("migration source must be a .sql file: %s", source)
	}

	sourcePath := filepath.Join(repoRoot(), cleanSource)
	contents, err := os.ReadFile(sourcePath)
	if err != nil {
		return fmt.Errorf("read %s: %w", sourcePath, err)
	}

	targetPath := filepath.Join(outputDir, manifest.GeneratedMigrationName(index, source))
	if err := os.WriteFile(targetPath, contents, 0o644); err != nil {
		return fmt.Errorf("write %s: %w", targetPath, err)
	}

	return nil
}

func repoRoot() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return "."
	}
	return filepath.Dir(filepath.Dir(filename))
}
