package matching

import (
	"open-outcry/pkg/db"
)

func (assert *ServiceTestSuite) TestGtc() {
	// given

	entity := assert.appEntity1

	// when: given a new order
	tradeOrder, _ := ProcessTradeOrder(assert.instrumentAccount1, "BTC_EUR", "LIMIT", Sell, 1.0, 10.0, "GTC")

	// then: it remains in the book until cancelled
	assert.Equal(1, db.QueryVal[int]("SELECT COUNT(*) FROM book_order"))
	assert.Equal(Open, GetTradeOrder(tradeOrder).Status)
	assert.Equal(10.0, FindCurrencyAccountByAppEntityIdAndCurrencyName(entity, "BTC").AmountReserved)
}

func (assert *ServiceTestSuite) TestFok() {
	// given
	entity := assert.appEntity1

	// when: given a new order that cannot be filled
	tradeOrder, _ := ProcessTradeOrder(assert.instrumentAccount1, "BTC_EUR", "LIMIT", Sell, 1.0, 1.0, FOK)

	// then: it is rejected
	assert.Equal(0, db.QueryVal[int]("SELECT COUNT(*) FROM book_order"))
	assert.Equal(Rejected, GetTradeOrder(tradeOrder).Status)
	assert.Equal(0.0, GetVolumeAtPrice("BTC_EUR", Sell, 1.0))

	assert.Equal(0.0, FindCurrencyAccountByAppEntityIdAndCurrencyName(entity, "BTC").AmountReserved)

	//when: given a new order that cannot be filled even when other orders present
	ProcessTradeOrder(assert.instrumentAccount2, "BTC_EUR", "LIMIT", Buy, 1.0, 1.0, "GTC")

	tradeOrder, _ = ProcessTradeOrder(assert.instrumentAccount1, "BTC_EUR", "LIMIT", Sell, 1.0, 2.0, FOK)

	//then: it is rejected
	assert.Equal(Rejected, GetTradeOrder(tradeOrder).Status)

	assert.Equal(0.0, FindCurrencyAccountByAppEntityIdAndCurrencyName(entity, "BTC").AmountReserved)

	//when: added another market order that can fill
	ProcessTradeOrder(assert.instrumentAccount2, "BTC_EUR", Market, Buy, 0.0, 2.0, "GTC")

	tradeOrder, _ = ProcessTradeOrder(assert.instrumentAccount1, "BTC_EUR", "LIMIT", Sell, 1.0, 2.0, FOK)

	//then: it is not reject
	assert.Equal(Filled, GetTradeOrder(tradeOrder).Status)
	assert.Equal(0.0, GetVolumeAtPrice("BTC_EUR", Sell, 1.0))

	assert.Equal(0.0, FindCurrencyAccountByAppEntityIdAndCurrencyName(entity, "BTC").AmountReserved)
}

func (assert *ServiceTestSuite) TestIoc() {
	// given
	entity := assert.appEntity1

	// when: given a new order that cannot be filled
	tradeOrder, _ := ProcessTradeOrder(assert.instrumentAccount1, "BTC_EUR", "LIMIT", Sell, 1.0, 1.0, IOC)

	//	then: it is rejected
	assert.Equal(0, db.QueryVal[int]("SELECT COUNT(*) FROM book_order"))
	assert.Equal(Rejected, GetTradeOrder(tradeOrder).Status)
	assert.Equal(0.0, GetVolumeAtPrice("BTC_EUR", Sell, 1.0))

	assert.Equal(0.0, FindCurrencyAccountByAppEntityIdAndCurrencyName(entity, "BTC").AmountReserved)

	//when: given a new order that can only be partially filled by a standing order in the order book
	ProcessTradeOrder(assert.instrumentAccount2, "BTC_EUR", "LIMIT", Buy, 1.0, 1, "GTC")
	tradeOrder, _ = ProcessTradeOrder(assert.instrumentAccount1, "BTC_EUR", "LIMIT", Sell, 1.0, 2, IOC)

	// then: it is partially rejected
	assert.Equal(1, GetTradeCount())
	assert.Equal(0, db.QueryVal[int]("SELECT COUNT(*) FROM book_order"))
	assert.Equal(PartiallyRejected, GetTradeOrder(tradeOrder).Status)
	assert.Equal(0.0, GetVolumeAtPrice("BTC_EUR", Sell, 1.0))

	assert.Equal(0.0, FindCurrencyAccountByAppEntityIdAndCurrencyName(entity, "BTC").AmountReserved)
}
