package services

import (
	"open-outcry/pkg/conf"
	"open-outcry/pkg/db"
	"open-outcry/pkg/models/app_entity"
	"open-outcry/pkg/models/instrument_account"
	"open-outcry/pkg/utils"
	"open-outcry/sql"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ServiceTestSuite struct {
	suite.Suite
	appEntity1         appentity.AppEntityId
	instrumentAccount1 instrumentaccount.InstrumentAccountId
	appEntity2         appentity.AppEntityId
	instrumentAccount2 instrumentaccount.InstrumentAccountId
}

func TestServiceTestSuite(t *testing.T) {
	conf.LoadTestConfig()
	suite.Run(t, &ServiceTestSuite{})
}

func (assert *ServiceTestSuite) SetupSuite() {
	if err := db.SetupInstance(); err != nil {
		panic(err)
	}
	_ = sql.MigrateUp()
}

func (assert *ServiceTestSuite) SetupTest() {
	assert.appEntity1, assert.instrumentAccount1 = Acc("test")
	assert.appEntity2, assert.instrumentAccount2 = Acc("test2")
}

func (assert *ServiceTestSuite) TearDownTest() {
	utils.Each([]string{"stop_order",
		"instrument_account_ledger_entry",
		"instrument_account_transfer",
		"trade",
		"trade_order",
		"price_level",
		"transfer_ledger_entry",
		"transfer",
		"currency_account WHERE app_entity_id != 1",
		"instrument_account",
		"app_entity WHERE pub_id != 'MASTER'",
	}, db.DeleteAll)
}

func (assert *ServiceTestSuite) TearDownAllSuite() {
	_ = sql.MigrateDown()
	_ = db.Instance().Close()
}
