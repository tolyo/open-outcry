package db

import (
	"database/sql"
	"open-outcry/pkg/conf"

	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
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

func QueryList[T comparable](query string, args ...any) []T {
	rows, err := Instance().Query(query, args...)

	if err != nil {
		log.Fatal(err)
	}
	res := make([]T, 0)
	for rows.Next() {
		var item T
		rows.Scan(&item)
		res = append(res, item)
	}
	return res
}
