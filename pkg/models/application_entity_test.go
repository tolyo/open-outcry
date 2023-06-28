package models

import "open-outcry/pkg/utils"

func (assert *ModelsTestSuite) TestMasterEntity() {
	// expect master entity to exist
	assert.Equal(1, utils.GetCount("app_entity"))
	assert.Equal(AppEntityId(Master), FindAppEntityExternalId("MASTER"))
}
