package utils

func GetCount(tableName string) int {
	return db.QueryVal(utils.Format("SELECT COUNT(*) FROM {{.}}", tableName))
}
