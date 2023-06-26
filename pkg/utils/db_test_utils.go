package utils

import (
	"context"
	"log"
	"open-outcry/pkg/db"
)

func GetCount(tableName string) int {
	var count int
	db.Instance().QueryRow("SELECT COUNT(*) FROM " + tableName).Scan(&count)
	return count
}

func DeleteAll(tableName string) {
	_, err := db.Instance().ExecContext(context.Background(), "DELETE FROM "+tableName)
	if err != nil {
		log.Fatal(err)
	}
}
