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

func (suite *ServiceTestSuite) SetupSuite() {
	err := db.SetupInstance()
	if err != nil {
		panic(err)
	}
	sql.MigrateUp()
}

func (suite *ServiceTestSuite) SetupTest() {
	suite.appEntity1, suite.tradingAccount1 = Acc("test")
	suite.appEntity2, suite.tradingAccount2 = Acc("test2")
}

func (suite *ServiceTestSuite) TearDownTest() {
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

func (suite *ServiceTestSuite) TearDownAllSuite() {
	sql.MigrateDown()
	db.Instance().Close()
}
