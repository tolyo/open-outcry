package services

import (
	"open-outcry/pkg/db"
)

func CreateAppEntity(id AppEntityExternalId) AppEntityId {
	var newId string
	db.Instance().QueryRow("SELECT create_client($1)", id).Scan(&newId)
	return AppEntityId(newId)
}
