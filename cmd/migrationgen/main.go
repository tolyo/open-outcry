package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	manifest "open-outcry/sql/manifest"
)

const (
	repoRoot     = "."
	generatedDir = "sql/generated"
)

func main() {
	if err := generate(); err != nil {
		fmt.Fprintf(os.Stderr, "migration generation failed: %v\n", err)
		os.Exit(1)
	}
}

func generate() error {
	if err := os.RemoveAll(generatedDir); err != nil {
		return err
	}
	if err := os.MkdirAll(generatedDir, 0o755); err != nil {
		return err
	}

	sources := manifest.MigrationSources()
	for index, source := range sources {
		if err := writeGeneratedMigration(index, source); err != nil {
			return err
		}
	}

	fmt.Printf("generated %d migrations in %s\n", len(sources), generatedDir)
	return nil
}

func writeGeneratedMigration(index int, source string) error {
	cleanSource := filepath.Clean(filepath.FromSlash(source))
	if cleanSource == "." || cleanSource == ".." || strings.HasPrefix(cleanSource, ".."+string(filepath.Separator)) {
		return fmt.Errorf("invalid migration source path: %s", source)
	}
	if filepath.Ext(cleanSource) != ".sql" {
		return fmt.Errorf("migration source must be a .sql file: %s", source)
	}

	sourcePath := filepath.Join(repoRoot, cleanSource)
	contents, err := os.ReadFile(sourcePath)
	if err != nil {
		return fmt.Errorf("read %s: %w", sourcePath, err)
	}

	targetPath := filepath.Join(generatedDir, manifest.GeneratedMigrationName(index, source))
	if err := os.WriteFile(targetPath, contents, 0o644); err != nil {
		return fmt.Errorf("write %s: %w", targetPath, err)
	}

	return nil
}
