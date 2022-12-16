package utils

import "open-outcry/pkg/db"

func GetCount(tableName string) int {
	var count int
	rows := db.Instance().QueryRow(Format("SELECT COUNT(*) FROM {{.}}", tableName))
	rows.Scan(&count)
	return count
}
