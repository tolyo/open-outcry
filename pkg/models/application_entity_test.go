package models

import (
	"open-outcry/pkg/db"
)

func (assert *ModelsTestSuite) TestMasterEntity() {
	// expect master entity to exist
	assert.Equal(1, db.GetCount("app_entity"))
	assert.Equal(AppEntityId(Master), FindAppEntityExternalId("MASTER"))
}
