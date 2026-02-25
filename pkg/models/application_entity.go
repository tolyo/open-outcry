package models

import "open-outcry/pkg/db"

// AppEntityId `app_entity.pub_id` db reference
type AppEntityId string

// AppEntityExternalId `app_entity.pub_id` db reference
type AppEntityExternalId string

// AppEntityType Type of application entity
type AppEntityType string

const (
	Client AppEntityType = "CLIENT"
	Master AppEntityType = "MASTER"
)

// AppEntity Application entity is any generic entity capable of being an actor in financial transaction
type AppEntity struct {
	Id         AppEntityId         `db:"pub_id"`
	Type       AppEntityType       `db:"type"`
	ExternalId AppEntityExternalId `db:"external_id"`
}

func FindAppEntityExternalId(id AppEntityExternalId) AppEntityId {
	res := db.QueryVal[string]("SELECT pub_id FROM app_entity WHERE external_id = $1", id)
	return AppEntityId(res)
}

func GetAppEntities() []AppEntity {
	return db.QueryList[AppEntity](`SELECT pub_id, type, external_id FROM app_entity`)
}

func GetAppEntity(id AppEntityId) *AppEntity {
	var entity AppEntity
	err := db.Instance().QueryRow(
		`SELECT pub_id, type, external_id FROM app_entity WHERE pub_id = $1`, id,
	).Scan(&entity.Id, &entity.Type, &entity.ExternalId)
	if err != nil {
		return nil
	}
	return &entity
}
