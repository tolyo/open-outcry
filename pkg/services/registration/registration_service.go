package registration

import (
	"open-outcry/pkg/db"
	appentity "open-outcry/pkg/models/app_entity"
)

func CreateAppEntity(id appentity.AppEntityExternalId) appentity.AppEntityId {
	var newId string
	db.Instance().QueryRow("SELECT create_client($1)", id).Scan(&newId)
	return appentity.AppEntityId(newId)
}
