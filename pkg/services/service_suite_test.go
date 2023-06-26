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
	utils.Each([]string{"stop_order",
		"trading_account_transfer",
		"trade",
		"trade_order",
		"price_level",
		"payment",
		"payment_account WHERE app_entity_id != 1",
		"trading_account",
		"app_entity WHERE pub_id != 'MASTER'",
	}, utils.DeleteAll)
}

func (suite *ServiceTestSuite) TearDownAllSuite() {
	db.MigrateDown()
	db.Instance().Close()
}
