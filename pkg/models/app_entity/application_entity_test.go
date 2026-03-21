package appentity

import (
	"testing"

	"open-outcry/pkg/db"
	"open-outcry/pkg/models/testutil"

	"github.com/stretchr/testify/assert"
)

func TestMasterEntity(t *testing.T) {
	teardown := testutil.Setup(t)
	defer teardown()

	assert.Equal(t, 1, db.GetCount("app_entity"))
	assert.Equal(t, AppEntityId(Master), FindAppEntityExternalId("MASTER"))
}

func TestGetAppEntities(t *testing.T) {
	teardown := testutil.Setup(t)
	defer teardown()

	entities := GetAppEntities()
	assert.GreaterOrEqual(t, len(entities), 1)

	found := false
	for _, e := range entities {
		if e.Type == Master {
			found = true
			assert.Equal(t, AppEntityExternalId("MASTER"), e.ExternalId)
		}
	}
	assert.True(t, found, "MASTER entity should exist")
}

func TestGetAppEntity(t *testing.T) {
	teardown := testutil.Setup(t)
	defer teardown()

	masterPubID := FindAppEntityExternalId("MASTER")
	entity := GetAppEntity(masterPubID)
	assert.NotNil(t, entity)
	assert.Equal(t, masterPubID, entity.Id)
	assert.Equal(t, Master, entity.Type)
	assert.Equal(t, AppEntityExternalId("MASTER"), entity.ExternalId)
}

func TestGetAppEntityNotFound(t *testing.T) {
	teardown := testutil.Setup(t)
	defer teardown()

	entity := GetAppEntity("nonexistent-id")
	assert.Nil(t, entity)
}
