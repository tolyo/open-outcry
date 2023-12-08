package services

import (
	"open-outcry/pkg/conf"
	"open-outcry/pkg/db"
	"open-outcry/pkg/models"
	"open-outcry/pkg/utils"
	"open-outcry/sql"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ServiceTestSuite struct {
	suite.Suite
	appEntity1      models.AppEntityId
	tradingAccount1 models.TradingAccountId // seller
	appEntity2      models.AppEntityId
	tradingAccount2 models.TradingAccountId // buyer
}

func TestServiceTestSuite(t *testing.T) {
	conf.LoadTestConfig()
	suite.Run(t, &ServiceTestSuite{})
}

func (assert *ServiceTestSuite) SetupSuite() {
	err := db.SetupInstance()
	if err != nil {
		panic(err)
	}
	sql.MigrateUp()
}

func (assert *ServiceTestSuite) SetupTest() {
	assert.appEntity1, assert.tradingAccount1 = Acc("test")
	assert.appEntity2, assert.tradingAccount2 = Acc("test2")
}

func (assert *ServiceTestSuite) TearDownTest() {
	utils.Each([]string{"stop_order",
		"trading_account_transfer",
		"trade",
		"trade_order",
		"price_level",
		"payment",
		"payment_account WHERE app_entity_id != 1",
		"trading_account",
		"app_entity WHERE pub_id != 'MASTER'",
	}, db.DeleteAll)
}

func (assert *ServiceTestSuite) TearDownAllSuite() {
	sql.MigrateDown()
	db.Instance().Close()
}
