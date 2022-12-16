package services

import (
	"open-outcry/pkg/db"
	"open-outcry/pkg/models"
)

func (assert *ServiceTestSuite) TestGtc() {
	// given
	tradingAccount := Acc()
	entity := GetAppEntityId()

	// when: given a new order
	tradeOrder := ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "SELL", 1.0, 10.0, "GTC")

	// then: it remains in the book until cancelled
	assert.Equal(1, db.QueryVal[int]("SELECT COUNT(*) FROM book_order"))
	assert.Equal(models.Open, models.GetTradeOrder(tradeOrder).Status)
	assert.Equal(10.0, models.FindPaymentAccountByAppEntityIdAndCurrencyName(entity, "BTC").AmountReserved)
}

func (assert *ServiceTestSuite) TestFok() {
	// given
	tradingAccount := Acc()
	//tradingAccount2 := Acc2()
	//entity := GetAppEntityId()

	// when: given a new order that cannot be filled
	tradeOrder := ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "SELL", 1.0, 1.0, models.FOK)

	// then: it is rejected
	assert.Equal(0, db.QueryVal[int]("SELECT COUNT(*) FROM book_order"))
	assert.Equal(models.Rejected, models.GetTradeOrder(tradeOrder).Status)
	assert.Equal(0.0, GetVolumeAtPrice("BTC_EUR", "SELL", 1.0))
	//
	//    assert FindPaymentAccountByAppEntityIdAndCurrencyName(entity, "BTC").AmountReserved
	//           |> Decimal.to_float() == 0
	//
	// when: given a new order that cannot be filled even when other orders present
	//    ProcessTradeOrder(tradingAccount2, "BTC_EUR", "LIMIT", "BUY", 1.0, 1.0, "GTC")
	//
	//    tradeOrder :=
	//      ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "SELL", 1.0, 2.0, :FOK)
	//
	// then: it is rejected
	//    assert TradeOrder.get(tradeOrder).status == :REJECTED
	//
	//    assert FindPaymentAccountByAppEntityIdAndCurrencyName(entity, "BTC").AmountReserved
	//           |> Decimal.to_float() == 0
	//
	// when: added another market order that can fill
	//    ProcessTradeOrder(tradingAccount2, "BTC_EUR", models.Market, "BUY", 2.0, "GTC")
	//
	//    tradeOrder :=
	//      ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "SELL", 1.0, 2.0, :FOK)
	//
	// then: it is not reject
	//    assert TradeOrder.get(tradeOrder).status == :FILLED
	//    assert GetVolumeAtPrice("BTC_EUR", "SELL", 1.0) == 0
	//
	// assert FindPaymentAccountByAppEntityIdAndCurrencyName(entity, "BTC").AmountReserved
	//        |> Decimal.to_float() == 0
}

//
//  test "IOC" {
// given
//    tradingAccount = Acc()
//    tradingAccount2 = Acc2()
//    entity = GetAppEntityId()
//
// when: given a new order that cannot be filled
//    tradeOrder :=
//      ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "SELL", 1.0, 1.0, :IOC)
//
// then: it is rejected
//    assert db.QueryVal("SELECT COUNT(*) FROM book_order") == 0
//    assert TradeOrder.get(tradeOrder).status == :REJECTED
//    assert GetVolumeAtPrice("BTC_EUR", "SELL", 1.0) == 0
//
//    assert FindPaymentAccountByAppEntityIdAndCurrencyName(entity, "BTC").AmountReserved
//           |> Decimal.to_float() == 0
//
// when: given a new order that can only be partially filled by a standing order in the order book
//    ProcessTradeOrder(tradingAccount2, "BTC_EUR", "LIMIT", "BUY", 1.0, 1, "GTC")
//    tradeOrder := ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "SELL", 1.0, 2, :IOC)
//
// then: it is partially rejected
//    assert MatchingServiceTestHelpers.get_trade_count() == 1
//    assert db.QueryVal("SELECT COUNT(*) FROM book_order") == 0
//    assert TradeOrder.get(tradeOrder).status == :PARTIALLY_REJECTED
//    assert GetVolumeAtPrice("BTC_EUR", "SELL", 1.0) == 0
//
//    assert FindPaymentAccountByAppEntityIdAndCurrencyName(entity, "BTC").AmountReserved
//           |> Decimal.to_float() == 0
//  }
//}
