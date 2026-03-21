package matching

import (
	"open-outcry/pkg/conf"
	"open-outcry/pkg/db"
	appentity "open-outcry/pkg/models/app_entity"
	currencyaccount "open-outcry/pkg/models/currency_account"
	"open-outcry/pkg/models/instrument"
	instrumentaccount "open-outcry/pkg/models/instrument_account"
	order "open-outcry/pkg/models/order_book"
	tradeorder "open-outcry/pkg/models/trade_order"
	orderbook "open-outcry/pkg/services/order_book"
	"open-outcry/pkg/services/registration"
	"open-outcry/pkg/services/transfer"
	"open-outcry/pkg/utils"
	"open-outcry/sql"
	"testing"

	"github.com/stretchr/testify/suite"
)

type (
	AppEntityId         = appentity.AppEntityId
	AppEntityExternalId = appentity.AppEntityExternalId
	InstrumentAccountId = instrumentaccount.InstrumentAccountId
	InstrumentName      = instrument.InstrumentName
	CurrencyAccount     = currencyaccount.CurrencyAccount
	OrderBook           = order.OrderBook
	PriceVolume         = order.PriceVolume
	TradeOrder          = tradeorder.TradeOrder
	OrderPrice          = tradeorder.OrderPrice
	OrderSide           = tradeorder.OrderSide
	OrderType           = tradeorder.OrderType
	OrderTimeInForce    = tradeorder.OrderTimeInForce
	OrderStatus         = tradeorder.OrderStatus
)

const (
	Buy                = tradeorder.Buy
	Sell               = tradeorder.Sell
	Limit              = tradeorder.Limit
	Market             = tradeorder.Market
	StopLoss           = tradeorder.StopLoss
	StopLimit          = tradeorder.StopLimit
	GTC                = tradeorder.GTC
	IOC                = tradeorder.IOC
	FOK                = tradeorder.FOK
	Open               = tradeorder.Open
	PartiallyFilled    = tradeorder.PartiallyFilled
	Cancelled          = tradeorder.Cancelled
	PartiallyCancelled = tradeorder.PartiallyCancelled
	PartiallyRejected  = tradeorder.PartiallyRejected
	Filled             = tradeorder.Filled
	Rejected           = tradeorder.Rejected
)

var (
	CreateAppEntity                                 = registration.CreateAppEntity
	CreateCurrencyAccount                           = currencyaccount.CreateCurrencyAccount
	CreateTransferDeposit                           = transfer.CreateTransferDeposit
	FindCurrencyAccountByAppEntityIdAndCurrencyName = currencyaccount.FindCurrencyAccountByAppEntityIdAndCurrencyName
	FindInstrumentAccountByApplicationEntityId      = instrumentaccount.FindInstrumentAccountByApplicationEntityId
	GetTradeOrder                                   = tradeorder.GetTradeOrder
	GetVolumes                                      = orderbook.GetVolumes
	GetOrderBook                                    = orderbook.GetOrderBook
	GetVolumeAtPrice                                = orderbook.GetVolumeAtPrice
	GetAvailableLimitVolume                         = orderbook.GetAvailableLimitVolume
)

type ServiceTestSuite struct {
	suite.Suite
	appEntity1         AppEntityId
	instrumentAccount1 InstrumentAccountId
	appEntity2         AppEntityId
	instrumentAccount2 InstrumentAccountId
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
	utils.Each([]string{
		"stop_order",
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

type AppState struct {
	entity1         []CurrencyAccount
	entity2         []CurrencyAccount
	tradeCount      int
	orderBookStates OrderBook
}

type TestStep struct {
	initialState  AppState
	orders        []TradeOrder
	expectedState AppState
}

type MatchingServiceTestCase struct {
	steps []TestStep
}

func Acc(v string) (AppEntityId, InstrumentAccountId) {
	appEntityId := CreateAppEntity(AppEntityExternalId(v))
	CreateCurrencyAccount(appEntityId, "BTC")
	CreateTransferDeposit(appEntityId, 1000, "BTC", "Test", "Test")
	CreateTransferDeposit(appEntityId, 1000, "EUR", "Test", "Test")
	account := FindInstrumentAccountByApplicationEntityId(appEntityId)
	return appEntityId, account.Id
}
