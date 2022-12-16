package services

import (
	"github.com/stretchr/testify/suite"
	"open-outcry/pkg/conf"
	"open-outcry/pkg/db"
	"testing"
)

type ServiceTestSuite struct {
	suite.Suite
}

func TestServiceTestSuite(t *testing.T) {
	conf.LoadTestConfig()
	suite.Run(t, &ServiceTestSuite{})
}

func (suite *ServiceTestSuite) SetupTest() {
	db.SetupInstance()
	db.MigrateUp()
}

func (suite *ServiceTestSuite) TearDownTest() {
	db.MigrateDown()
	db.Instance().Close()
}
