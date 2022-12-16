package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"open-outcry/pkg/conf"
)

var db *sql.DB

// SetupInstance - Initializes the database
func SetupInstance() error {
	var err error
	db, err = sql.Open("postgres", conf.Get().DBDsn)
	if err != nil {
		return err
	}
	return nil
}

// Instance gets database instance, initialized via SetupInstance func
func Instance() *sql.DB {
	return db
}

func QueryVal[T comparable](query string, args ...any) T {
	var val T
	db.QueryRow(query, args...).Scan(&val)
	return val
}
