package models

// `application_entity.pub_id` db reference
type ApplicationEntityId string

// `application_entity.pub_id` db reference
type ApplicationEntityExternalId string

// Type of application entity
type ApplicationEntityType string

const (
	Client ApplicationEntityType = "CLIENT"
	Master ApplicationEntityType = "MASTER"
)

// Application entity is any generic enity capable of being an
// actor in financial transaction
type ApplicationEntity struct {
	Id         ApplicationEntityId
	Type       ApplicationEntityType
	ExternalId ApplicationEntityExternalId
}

func FindByExternalId(id ApplicationEntityExternalId) ApplicationEntityId {
	db.QueryVal("SELECT pub_id FROM application_entity WHERE external_id = $1", id)
}
