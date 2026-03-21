package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	migration "open-outcry/sql"
)

func main() {
	var outDir string
	flag.StringVar(&outDir, "out", "", "output directory for generated goose migrations")
	flag.Parse()

	if outDir == "" {
		fmt.Fprintln(os.Stderr, "migration generation failed: -out is required")
		os.Exit(2)
	}

	absOutDir, err := filepath.Abs(outDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "migration generation failed: resolve output directory: %v\n", err)
		os.Exit(1)
	}

	count, err := migration.GenerateMigrations(absOutDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "migration generation failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("generated %d migrations in %s\n", count, absOutDir)
}
