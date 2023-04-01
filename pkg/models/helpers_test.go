package models

import (
	"github.com/stretchr/testify/suite"
	"open-outcry/pkg/conf"
	"open-outcry/pkg/db"
	"testing"
)

type ModelsTestSuite struct {
	suite.Suite
}

func TestServiceTestSuite(t *testing.T) {
	conf.LoadTestConfig()
	suite.Run(t, &ModelsTestSuite{})
}

func (suite *ModelsTestSuite) SetupTest() {
	db.SetupInstance()
	db.MigrateUp()
}

func (suite *ModelsTestSuite) TearDownTest() {
	db.MigrateDown()
	db.Instance().Close()
}
