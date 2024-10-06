package db

import (
	"context"
	"fmt"
	"open-outcry/pkg/conf"
	"reflect"

	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

var db *sqlx.DB

// SetupInstance - Initializes the database
func SetupInstance() error {
	var err error
	db, err = sqlx.Connect("postgres", conf.Get().DBDsn)
	if err != nil {
		return err
	}
	return nil
}

// Instance gets database instance, initialized via SetupInstance func
func Instance() *sqlx.DB {
	return db
}

func QueryVal[T interface{}](query string, args ...any) T {
	var val T
	kind := reflect.ValueOf(val)
	if kind.Kind() == reflect.Struct {
		err := db.Get(val, query, args...)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		db.QueryRow(query, args...).Scan(&val)
	}
	return val
}

func QueryList[T interface{}](query string, args ...any) []T {
	var val T
	kind := reflect.ValueOf(val)
	if kind.Kind() == reflect.Struct {
		res := make([]T, 0)
		err := db.Select(&res, query, args...)
		if err != nil {
			log.Fatal(err)
		}
		return res
	} else {
		res := make([]T, 0)

		rows, err := Instance().Query(query, args...)
		if err != nil {
			log.Fatal(err)
		}
		for rows.Next() {
			var item T
			rows.Scan(&item)
			res = append(res, item)
		}
		return res
	}

}

func GetCount(tableName string) int {
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s", tableName)
	return QueryVal[int](query)
}

func DeleteAll(tableName string) {
	query := fmt.Sprintf("DELETE FROM %s", tableName)
	_, err := Instance().ExecContext(context.Background(),
		query,
	)
	if err != nil {
		log.Fatal(err)
	}
}
