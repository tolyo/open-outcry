package models

import "open-outcry/pkg/db"

// `app_entity.pub_id` db reference
type AppEntityId string

// `app_entity.pub_id` db reference
type AppEntityExternalId string

// Type of application entity
type AppEntityType string

const (
	Client AppEntityType = "CLIENT"
	Master AppEntityType = "MASTER"
)

// Application entity is any generic enity capable of being an
// actor in financial transaction
type AppEntity struct {
	Id         AppEntityId
	Type       AppEntityType
	ExternalId AppEntityExternalId
}

func FindAppEntityExternalId(id AppEntityExternalId) AppEntityId {
	res := db.QueryVal[string]("SELECT pub_id FROM app_entity WHERE external_id = $1", id)
	return AppEntityId(res)
}
