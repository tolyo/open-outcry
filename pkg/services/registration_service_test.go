package services

import (
	"open-outcry/pkg/db"

	"github.com/google/uuid"
)

func (assert *ServiceTestSuite) TestCreateAppEntity() {
	// when
	count := db.GetCount("app_entity")
	res := CreateAppEntity("test")

	// then
	assert.NotNil(res)
	_, err := uuid.Parse(string(res))
	assert.Nil(err)
	assert.Equal(count+1, db.GetCount("app_entity"))

	// when
	res2 := CreateAppEntity("test")

	// then
	assert.NotEqual(res, res2)
	assert.Equal(count+2, db.GetCount("app_entity"))
}
