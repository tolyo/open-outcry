package testutil

// Package testutil provides a shared integration-test harness for model packages.
// The model tests all run against the same Postgres database, so setup and
// teardown are centralized here and access is serialized with a file lock to
// avoid concurrent migration and data-reset races across packages.

import (
	"os"
	"path/filepath"
	"syscall"
	"testing"

	"open-outcry/pkg/conf"
	"open-outcry/pkg/db"
	"open-outcry/sql"
)

const lockPath = "/tmp/open-outcry-model-tests.lock"

func Setup(tb testing.TB) func() {
	tb.Helper()

	lockFile, err := os.OpenFile(filepath.Clean(lockPath), os.O_CREATE|os.O_RDONLY, 0o644)
	if err != nil {
		tb.Fatalf("open test lock: %v", err)
	}
	if err := syscall.Flock(int(lockFile.Fd()), syscall.LOCK_EX); err != nil {
		_ = lockFile.Close()
		tb.Fatalf("lock test db: %v", err)
	}

	conf.LoadTestConfig()
	if err := db.SetupInstance(); err != nil {
		_ = syscall.Flock(int(lockFile.Fd()), syscall.LOCK_UN)
		_ = lockFile.Close()
		tb.Fatalf("setup db: %v", err)
	}
	sql.MigrateUp()

	return func() {
		sql.MigrateDown()
		if db.Instance() != nil {
			_ = db.Instance().Close()
		}
		_ = syscall.Flock(int(lockFile.Fd()), syscall.LOCK_UN)
		_ = lockFile.Close()
	}
}
