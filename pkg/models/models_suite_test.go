package models

import (
	"open-outcry/pkg/conf"
	"open-outcry/pkg/db"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ModelsTestSuite struct {
	suite.Suite
}

func TestModelSuite(t *testing.T) {
	conf.LoadTestConfig()
	suite.Run(t, &ModelsTestSuite{})
}

func (suite *ModelsTestSuite) SetupTest() {
	err := db.SetupInstance()
	if err != nil {
		panic(err)
	}
	db.MigrateUp()
}

func (suite *ModelsTestSuite) TearDownTest() {
	db.MigrateDown()
	db.Instance().Close()
}
