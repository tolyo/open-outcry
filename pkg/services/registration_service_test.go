package services

import (
	"github.com/google/uuid"
	"open-outcry/pkg/models"
	"open-outcry/pkg/utils"
)

func (assert *ServiceTestSuite) TestCreateAppEntity() {
	// when
	count := utils.GetCount("app_entity")
	res := CreateAppEntity(models.AppEntityExternalId("test"))

	// then
	assert.NotNil(res)
	_, err := uuid.Parse(string(res))
	assert.Nil(err)
	assert.Equal(count+1, utils.GetCount("app_entity"))

	// when
	res2 := CreateAppEntity(models.AppEntityExternalId("test"))

	// then
	assert.NotEqual(res, res2)
	assert.Equal(count+2, utils.GetCount("app_entity"))
}
