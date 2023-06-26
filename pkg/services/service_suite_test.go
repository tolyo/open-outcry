package services

import (
	"open-outcry/pkg/conf"
	"open-outcry/pkg/db"
	"open-outcry/pkg/utils"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ServiceTestSuite struct {
	suite.Suite
}

func TestServiceTestSuite(t *testing.T) {
	conf.LoadTestConfig()
	suite.Run(t, &ServiceTestSuite{})
}

func (suite *ServiceTestSuite) SetupSuite() {
	err := db.SetupInstance()
	if err != nil {
		panic(err)
	}
	db.MigrateUp()
}

func (suite *ServiceTestSuite) TearDownTest() {
	// remove all trades
	utils.DeteleAll("stop_order")
	utils.DeteleAll("trading_account_transfer")
	utils.DeteleAll("trade")
	utils.DeteleAll("trade_order")
	utils.DeteleAll("price_level")
	utils.DeteleAll("payment")
	utils.DeteleAll("payment_account WHERE app_entity_id != 1")
	utils.DeteleAll("trading_account")
	utils.DeteleAll("app_entity WHERE pub_id != 'MASTER'")
	// remove all users
}

func (suite *ServiceTestSuite) TearDownAllSuite() {
	db.MigrateDown()
	db.Instance().Close()
}
