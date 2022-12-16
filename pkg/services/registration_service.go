package services

import (
	"open-outcry/pkg/db"
	"open-outcry/pkg/models"
)

func CreateAppEntity(id models.AppEntityExternalId) models.AppEntityId {
	var newId string
	db.Instance().QueryRow("SELECT create_client($1)", id).Scan(&newId)
	return models.AppEntityId(newId)
}
