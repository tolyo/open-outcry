package utils

import (
	"context"
	"log"
	"open-outcry/pkg/db"
)

func GetCount(tableName string) int {
	var count int
	rows := db.Instance().QueryRow(Format("SELECT COUNT(*) FROM {{.}}", tableName))
	rows.Scan(&count)
	return count
}

func DeteleAll(tableName string) error {
	_, err := db.Instance().ExecContext(context.Background(), Format("DELETE FROM {{.}}", tableName))
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
