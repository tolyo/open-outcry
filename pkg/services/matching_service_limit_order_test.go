package services

//funcmodule MatchingServiceLimitOrderTest {
//  use DataCase
//
//  import TestUtils
//
func (assert *ServiceTestSuite) TestProcessLimitSellOrderSave() {
// when: a limit order is sent to an empty matching unit
//    res = ProcessTradeOrder(Acc(), "BTC_EUR", "LIMIT", "SELL", 10, 100, "GTC")
//
// then: a matching unit should save the trade order on save order to the order book
//    assert.Equal(res != nil
//    assert.Equal(res == db.QueryVal("SELECT pub_id FROM tradeOrder WHERE pub_id = '#{res}'")
//    assert.Equal(GetSellBookOrderCount() == 1
//
//    assert.Equal(GetVolumes("BTC_EUR", "SELL") == [
//             {10, 100}
//           ]
}

func (assert *ServiceTestSuite) TestProcessLimitBuyOrderSave() {
// when: a limit order is sent to an empty matching unit
//    ProcessTradeOrder(Acc(), "BTC_EUR", "LIMIT", "BUY", 10, 100, "GTC")
//
// then: a matching unit should save the orderSELL
//    assert.Equal(GetBuyBookOrderCount() == 1
//
//    assert.Equal(GetVolumes("BTC_EUR", "BUY") == [
//             {10, 100}
//           ]
}

func (assert *ServiceTestSuite) TestProcessLimitNoMatchCaseIncomingBuy() {
// when: there is a SELL order in the book and a BUY limit order arrives that {es not cross
//    ProcessTradeOrder(Acc(), "BTC_EUR", "LIMIT", "SELL", 10, 100, "GTC")
//    ProcessTradeOrder(Acc2(), "BTC_EUR", "LIMIT", "BUY", 9, 100, "GTC")
//
// then: the book should have both orders and no trade should be generated
//    assert.Equal(GetSellBookOrderCount() == 1
//    assert.Equal(GetBuyBookOrderCount() == 1
//    assert.Equal(get_trade_count() == 0
//
//    assert.Equal(GetVolumes("BTC_EUR", "SELL") == [
//             {10, 100}
//           ]
//
//    assert.Equal(GetVolumes("BTC_EUR", "BUY") == [
//             {9, 100}
//           ]
}

func (assert *ServiceTestSuite) TestProcessLimitNoMatchCaseIncomingSell() {
// when: there is a BUY order in the book and a SELL limit order arrives that {es not cross
//    ProcessTradeOrder(Acc(), "BTC_EUR", "LIMIT", "BUY", 9, 100, "GTC")
//    ProcessTradeOrder(Acc2(), "BTC_EUR", "LIMIT", "SELL", 10, 100, "GTC")
//
// then: the book should have both orders and no trade should be generated
//    assert.Equal(GetSellBookOrderCount() == 1
//    assert.Equal(GetBuyBookOrderCount() == 1
//    assert.Equal(get_trade_count() == 0
//
//    assert.Equal(GetVolumes("BTC_EUR", "SELL") == [
//             {10, 100}
//           ]
//
//    assert.Equal(GetVolumes("BTC_EUR", "BUY") == [
//             {9, 100}
//           ]
}

func (assert *ServiceTestSuite) TestProcessLimitExactMatchIncomingBuy() {
// when: there is a SELL order in the book
//    ProcessTradeOrder(Acc(), "BTC_EUR", "LIMIT", "SELL", 10, 100, "GTC")
//
// then:
//    assert.Equal(GetSellBookOrderCount() == 1
//    assert.Equal(GetBuyBookOrderCount() == 0
//    assert.Equal(get_trade_count() == 0
//
//    assert.Equal(GetAvailableLimitVolume("SELL", 10) ==
//             100
//
//    assert.Equal(GetVolumes("BTC_EUR", "SELL") == [
//             {10, 100}
//           ]
//
// when: a BUY limit order arrives that crossed
//    ProcessTradeOrder(Acc2(), "BTC_EUR", "LIMIT", "BUY", 10, 100, "GTC")
//
// then: the book should have no orders and a single trade should be generated
//    assert.Equal(GetSellBookOrderCount() == 0
//    assert.Equal(GetBuyBookOrderCount() == 0
//    assert.Equal(get_trade_count() == 1
//    assert.Equal(GetAvailableLimitVolume("SELL", 10) == 0
//    assert.Equal(GetVolumes("BTC_EUR", "SELL") == []
//    assert.Equal(GetVolumes("BTC_EUR", "BUY") == []
}

func (assert *ServiceTestSuite) TestProcessLimitPartialMatchIncomingBuySingleTrade() {
//    ProcessTradeOrder(Acc(), "BTC_EUR", "LIMIT", "SELL", 10.00, 100.00, "GTC")
//
// then:
//    assert.Equal(GetSellBookOrderCount() == 1
//    assert.Equal(GetBuyBookOrderCount() == 0
//    assert.Equal(get_trade_count() == 0
//
//    assert.Equal(GetVolumes("BTC_EUR", "SELL") == [
//             {10.00, 100}
//           ]
//
// when: incoming buy order is only partially matched
//    ProcessTradeOrder(Acc2(), "BTC_EUR", "LIMIT", "BUY", 10.00, 50.00, "GTC")
//
// then: the book should have one sell order and a single trade should be generated
//    assert.Equal(GetSellBookOrderCount() == 1
//    assert.Equal(GetBuyBookOrderCount() == 0
//    assert.Equal(get_trade_count() == 1
//
//    assert.Equal(GetVolumes("BTC_EUR", "SELL") == [
//             {10.00, 50}
//           ]
//
//    assert.Equal(GetVolumes("BTC_EUR", "BUY") == []
}

func (assert *ServiceTestSuite) TestProcessLimitOverflowMatchIncomingBuySingleTrade() {
// when: there is a SELL order in the book
//    ProcessTradeOrder(Acc(), "BTC_EUR", "LIMIT", "SELL", 10.00, 10.00, "GTC")
//
// then:
//    assert.Equal(GetSellBookOrderCount() == 1
//    assert.Equal(GetBuyBookOrderCount() == 0
//    assert.Equal(get_trade_count() == 0
//    assert.Equal(GetAvailableLimitVolume("SELL", 10) == 10
//
// when: a BUY limit order arrives that crosses and is more that the book amount
//    ProcessTradeOrder(Acc2(), "BTC_EUR", "LIMIT", "BUY", 10.00, 15, "GTC")
//
// then: the book should be one buy order,  no sell orders and a one trade
//    assert.Equal(GetSellBookOrderCount() == 0
//    assert.Equal(GetBuyBookOrderCount() == 1
//    assert.Equal(get_trade_count() == 1
//    assert.Equal(GetVolumes("BTC_EUR", "SELL") == []
//
//    assert.Equal(GetVolumes("BTC_EUR", "BUY") == [
//             {10.00, 5}
//           ]
}

func (assert *ServiceTestSuite) TestProcessLimitPartialMatchIncomingSellSingleTrade() {
// when: there is a BUY order
//    ProcessTradeOrder(Acc(), "BTC_EUR", "LIMIT", "BUY", 10.00, 100.00, "GTC")
//
// then:
//    assert.Equal(GetSellBookOrderCount() == 0
//    assert.Equal(GetBuyBookOrderCount() == 1
//    assert.Equal(get_trade_count() == 0
//
//    assert.Equal(GetAvailableLimitVolume("BUY", 10) == 100
//
// when: incoming SELL order is only partially matched
//    ProcessTradeOrder(Acc2(), "BTC_EUR", "LIMIT", "SELL", 10.00, 50.00, "GTC")
//
// then: the book should have one BUY order and a 1 trade should be generated
//    assert.Equal(GetSellBookOrderCount() == 0
//    assert.Equal(GetBuyBookOrderCount() == 1
//    assert.Equal(get_trade_count() == 1
//    assert.Equal(GetAvailableLimitVolume("BUY", 10) == 50
//    assert.Equal(GetVolumes("BTC_EUR", "SELL") == []
//
//    assert.Equal(GetVolumes("BTC_EUR", "BUY") == [
//             {10.00, 50}
//           ]
}

func (assert *ServiceTestSuite) TestProcessLimitOverflowMatchIncomingSellSingleTrade() {
// when: there is a BUY order in the book
//    ProcessTradeOrder(Acc(), "BTC_EUR", "LIMIT", "BUY", 10.00, 100.00, "GTC")
//
// then:
//    assert.Equal(GetSellBookOrderCount() == 0
//    assert.Equal(GetBuyBookOrderCount() == 1
//    assert.Equal(get_trade_count() == 0
//    assert.Equal(GetAvailableLimitVolume("BUY", 10) == 100
//
//    # when: a SELL limit order arrives that crosses and is more that the book amount
//    ProcessTradeOrder(Acc2(), "BTC_EUR", "LIMIT", "SELL", 10.00, 150.00, "GTC")
//
//    # then: the book should be one SELL order,  no BUY orders and 1 trade
//    assert.Equal(GetSellBookOrderCount() == 1
//    assert.Equal(GetBuyBookOrderCount() == 0
//    assert.Equal(get_trade_count() == 1
//    assert.Equal(GetAvailableLimitVolume("SELL", 10) == 50
//
//    assert.Equal(GetVolumes("BTC_EUR", "SELL") == [
//             {10.00, 50}
//           ]
//
//    assert.Equal(GetVolumes("BTC_EUR", "BUY") == []
}

func (assert *ServiceTestSuite) TestProcessLimitExactMatchIncomingBuysMultipleTrades() {
//    # when: there is a SELL order in the book and 2 BUY limit order arrive that cross
//    ProcessTradeOrder(Acc(), "BTC_EUR", "LIMIT", "SELL", 10.00, 10.00, "GTC")
//
//    # then:
//    assert.Equal(GetSellBookOrderCount() == 1
//    assert.Equal(GetBuyBookOrderCount() == 0
//    assert.Equal(get_trade_count() == 0
//
//    assert.Equal(GetAvailableLimitVolume("SELL", 10) == 10
//
//    # when: incoming buy order that are partially matched
//    tradingAccount = Acc2()
//    ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "BUY", 10.00, 5.00, "GTC")
//    ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "BUY", 10.00, 5.00, "GTC")
//
//    # then: the book should have 2 trades should be generated
//    assert.Equal(GetSellBookOrderCount() == 0
//    assert.Equal(GetBuyBookOrderCount() == 0
//    assert.Equal(get_trade_count() == 2
//    assert.Equal(GetAvailableLimitVolume("SELL", 10) == 0
//    assert.Equal(GetVolumes("BTC_EUR", "SELL") == []
//    assert.Equal(GetVolumes("BTC_EUR", "BUY") == []
}

func (assert *ServiceTestSuite) TestProcessLimitExactMatchIncomingSellMultipleTrades() {
//    tradingAccount = Acc()
//    tradingAccount2 = Acc2()
//
//    # when: there is a BUY order in the book and 2 SELL limit order arrive that cross
//    ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "BUY", 10.00, 10.00, "GTC")
//
//    # then:
//    assert.Equal(GetSellBookOrderCount() == 0
//    assert.Equal(GetBuyBookOrderCount() == 1
//    assert.Equal(get_trade_count() == 0
//    assert.Equal(GetAvailableLimitVolume("BUY", 10) == 10
//
//    # when: incoming buy order is only partially matched
//    ProcessTradeOrder(tradingAccount2, "BTC_EUR", "LIMIT", "SELL", 10.00, 5.00, "GTC")
//    ProcessTradeOrder(tradingAccount2, "BTC_EUR", "LIMIT", "SELL", 10.00, 5.00, "GTC")
//
//    # then: the book should have 2 trades should be generated
//    assert.Equal(GetSellBookOrderCount() == 0
//    assert.Equal(GetBuyBookOrderCount() == 0
//    assert.Equal(get_trade_count() == 2
//    assert.Equal(GetVolumes("BTC_EUR", "SELL") == []
//    assert.Equal(GetVolumes("BTC_EUR", "BUY") == []
}

func (assert *ServiceTestSuite) TestProcessLimitPartialMatchMultipleBookSellsToMultipleTrades() {
//    # given:
//    tradingAccount = Acc()
//    tradingAccount2 = Acc2()
//
//    # when: there are multiple SELL orders in the book and BUY limit order arrive that cross but can fill only partially one of the orders
//    ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "SELL", 10.00, 50.00, "GTC")
//    ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "SELL", 10.00, 50.00, "GTC")
//
//    # then:
//    assert.Equal(GetSellBookOrderCount() == 2
//    assert.Equal(GetBuyBookOrderCount() == 0
//    assert.Equal(get_trade_count() == 0
//
//    assert.Equal(GetAvailableLimitVolume("SELL", 10) ==
//             100
//
//    # when: incoming buy order is only partially matched
//    ProcessTradeOrder(tradingAccount2, "BTC_EUR", "LIMIT", "BUY", 10.00, 75.00, "GTC")
//
//    # then: the book should have 2 trades should be generated
//    assert.Equal(GetSellBookOrderCount() == 1
//    assert.Equal(GetBuyBookOrderCount() == 0
//    assert.Equal(get_trade_count() == 2
//
//    assert.Equal(GetVolumes("BTC_EUR", "SELL") == [
//             {10.00, 25}
//           ]
//
//    assert.Equal(GetVolumes("BTC_EUR", "BUY") == []
}

func (assert *ServiceTestSuite) TestProcessLimitPartialMatchMultipleBookBuysToMultipleTrades() {
//    # given:
//    tradingAccount = Acc()
//    tradingAccount2 = Acc2()
//
//    # when: there are multiple BUY orders in the book and SELL limit order arrive that cross but can fill only partially one of the orders
//    ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "BUY", 10.00, 50.00, "GTC")
//    ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "BUY", 10.00, 50.00, "GTC")
//
//    # then:
//    assert.Equal(GetSellBookOrderCount() == 0
//    assert.Equal(GetBuyBookOrderCount() == 2
//    assert.Equal(get_trade_count() == 0
//    assert.Equal(GetAvailableLimitVolume("BUY", 10) == 100
//
//    # when: incoming SELL order is only partially matched
//    ProcessTradeOrder(tradingAccount2, "BTC_EUR", "LIMIT", "SELL", 10.00, 75.00, "GTC")
//
//    # then: the book should have 2 trades should be generated
//    assert.Equal(GetSellBookOrderCount() == 0
//    assert.Equal(GetBuyBookOrderCount() == 1
//    assert.Equal(get_trade_count() == 2
//    assert.Equal(GetVolumes("BTC_EUR", "SELL") == []
//
//    assert.Equal(GetVolumes("BTC_EUR", "BUY") == [
//             {10.00, 25}
//           ]
}

func (assert *ServiceTestSuite) TestProcessLimitExactMatchMultipleBookSellsToMultipleTrades() {
//    # given:
//    tradingAccount = Acc()
//    tradingAccount2 = Acc2()
//
//    # when:
//    ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "SELL", 10.00, 10.00, "GTC")
//    ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "SELL", 10.00, 10.00, "GTC")
//    ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "SELL", 10.00, 10.00, "GTC")
//    ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "SELL", 10.00, 10.00, "GTC")
//
//    # then:
//    assert.Equal(GetSellBookOrderCount() == 4
//    assert.Equal(GetBuyBookOrderCount() == 0
//    assert.Equal(get_trade_count() == 0
//    assert.Equal(GetAvailableLimitVolume("SELL", 10) == 40
//
//    # when: incoming buy order is only partially matched
//    ProcessTradeOrder(tradingAccount2, "BTC_EUR", "LIMIT", "BUY", 10.00, 40.00, "GTC")
//
//    # then: the book should have 2 trades should be generated
//    assert.Equal(GetSellBookOrderCount() == 0
//    assert.Equal(GetBuyBookOrderCount() == 0
//    assert.Equal(get_trade_count() == 4
//    assert.Equal(GetVolumes("BTC_EUR", "SELL") == []
//    assert.Equal(GetVolumes("BTC_EUR", "BUY") == []
}

func (assert *ServiceTestSuite) TestProcessLimitExactMatchMultipleBookBuyToMultipleTrades() {
//    # given:
//    tradingAccount = Acc()
//    tradingAccount2 = Acc2()
//
//    # when:
//    ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "BUY", 10.00, 5.00, "GTC")
//    ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "BUY", 10.00, 5.00, "GTC")
//    ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "BUY", 10.00, 5.00, "GTC")
//    ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "BUY", 10.00, 5.00, "GTC")
//
//    # then:
//    assert.Equal(GetSellBookOrderCount() == 0
//    assert.Equal(GetBuyBookOrderCount() == 4
//    assert.Equal(get_trade_count() == 0
//    assert.Equal(GetAvailableLimitVolume("BUY", 10) == 20
//
//    # when: incoming buy order is only partially matched
//    ProcessTradeOrder(tradingAccount2, "BTC_EUR", "LIMIT", "SELL", 10.00, 20.00, "GTC")
//
//    # then: the book should have 2 trades should be generated
//    assert.Equal(GetSellBookOrderCount() == 0
//    assert.Equal(GetBuyBookOrderCount() == 0
//    assert.Equal(get_trade_count() == 4
//    assert.Equal(GetVolumes("BTC_EUR", "SELL") == []
//    assert.Equal(GetVolumes("BTC_EUR", "BUY") == []
}

func (assert *ServiceTestSuite) TestProcessLimitIncompleteMatchMultipleBookSellsToMultipleTrades() {
//    # given:
//    tradingAccount = Acc()
//    tradingAccount2 = Acc2()
//
//    # when: there are multiple SELL orders in the book and BUY limit order arrive that cross but can be only partially filled
//    ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "SELL", 10.00, 5.00, "GTC")
//    ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "SELL", 10.00, 5.00, "GTC")
//
//    # then:
//    assert.Equal(GetSellBookOrderCount() == 2
//    assert.Equal(GetBuyBookOrderCount() == 0
//    assert.Equal(get_trade_count() == 0
//    assert.Equal(GetAvailableLimitVolume("SELL", 10) == 10
//
//    # when: incoming buy order is only partially matched
//    ProcessTradeOrder(tradingAccount2, "BTC_EUR", "LIMIT", "BUY", 10.00, 17, "GTC")
//
//    # then: the book should have 2 trades should be generated
//    assert.Equal(GetSellBookOrderCount() == 0
//    assert.Equal(GetBuyBookOrderCount() == 1
//    assert.Equal(get_trade_count() == 2
//    assert.Equal(GetVolumes("BTC_EUR", "SELL") == []
//
//    assert.Equal(GetVolumes("BTC_EUR", "BUY") == [
//             {10.00, 7}
//           ]
}

func (assert *ServiceTestSuite) TestProcessLimitIncompleteMatchMultipleBookBuysToMultipleTrades() {
//    # given:
//    tradingAccount = Acc()
//    tradingAccount2 = Acc2()
//
//    # when: there are multiple BUY orders in the book and SELL limit order arrive that cross but can be only partially filled
//    ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "BUY", 10.00, 50.00, "GTC")
//    ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "BUY", 10.00, 50.00, "GTC")
//
//    # then:
//    assert.Equal(GetSellBookOrderCount() == 0
//    assert.Equal(GetBuyBookOrderCount() == 2
//    assert.Equal(get_trade_count() == 0
//    assert.Equal(GetAvailableLimitVolume("BUY", 10) == 100
//
//    # when: incoming SELL order is only partially matched
//    ProcessTradeOrder(tradingAccount2, "BTC_EUR", "LIMIT", "SELL", 10.00, 175.00, "GTC")
//
//    # then: the book should have 2 trades should be generated
//    assert.Equal(GetSellBookOrderCount() == 1
//    assert.Equal(GetBuyBookOrderCount() == 0
//    assert.Equal(get_trade_count() == 2
//
//    assert.Equal(GetVolumes("BTC_EUR", "SELL") == [
//             {10.00, 75}
//           ]
//
//    assert.Equal(GetVolumes("BTC_EUR", "BUY") == []
}

func (assert *ServiceTestSuite) TestProcessLimitPartialMatchMultipleBookSellsToMultipleTradesMultiplePrices() {
//    # given:
//    tradingAccount = Acc()
//    tradingAccount2 = Acc2()
//
//    # when: there are multiple SELL orders in the book and BUY limit order arrive that cross but can fill only partially one of the orders
//    ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "SELL", 9.00, 50.00, "GTC")
//    ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "SELL", 10.00, 50.00, "GTC")
//
//    # then:
//    assert.Equal(GetSellBookOrderCount() == 2
//    assert.Equal(GetBuyBookOrderCount() == 0
//    assert.Equal(get_trade_count() == 0
//
//    # when: incoming buy order is only partially matched
//    ProcessTradeOrder(tradingAccount2, "BTC_EUR", "LIMIT", "BUY", 11.00, 75.00, "GTC")
//
//    # then: the book should have 2 trades should be generated
//    assert.Equal(GetSellBookOrderCount() == 1
//    assert.Equal(GetBuyBookOrderCount() == 0
//    assert.Equal(get_trade_count() == 2
//
//    assert.Equal(GetVolumes("BTC_EUR", "SELL") == [
//             {10.00, 25}
//           ]
//
//    assert.Equal(GetVolumes("BTC_EUR", "BUY") == []
}

func (assert *ServiceTestSuite) TestProcessLimitPartialMatchMultipleBookBuysToMultipleTradesMultiplePrices() {
//    # given:
//    tradingAccount = Acc()
//    tradingAccount2 = Acc2()
//
//    # when: there are multiple BUY orders in the book and SELL limit order arrive that cross but can fill only partially one of the orders
//    ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "BUY", 9.00, 50.00, "GTC")
//    ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "BUY", 10.00, 50.00, "GTC")
//
//    # then:
//    assert.Equal(GetSellBookOrderCount() == 0
//    assert.Equal(GetBuyBookOrderCount() == 2
//    assert.Equal(get_trade_count() == 0
//    assert.Equal(GetAvailableLimitVolume("BUY", 10) == 50
//    assert.Equal(GetAvailableLimitVolume("BUY", 9) == 100
//
//    # when: incoming SELL order is only partially matched
//    ProcessTradeOrder(tradingAccount2, "BTC_EUR", "LIMIT", "SELL", 8.00, 75.00, "GTC")
//
//    # then: the book should have 2 trades should be generated
//    assert.Equal(GetSellBookOrderCount() == 0
//    assert.Equal(GetBuyBookOrderCount() == 1
//    assert.Equal(get_trade_count() == 2
//    assert.Equal(GetAvailableLimitVolume("BUY", 9) == 25
//    assert.Equal(GetVolumes("BTC_EUR", "SELL") == []
//
//    assert.Equal(GetVolumes("BTC_EUR", "BUY") == [
//             {9.00, 25}
//           ]
}

func (assert *ServiceTestSuite) TestProcessLimitSelfTradePreventions() {
	//    # given:
	//    tradingAccount = Acc()
	//
	//    # when:
	//    ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "BUY", 10.00, 50.00, "GTC")
	//    assert.Equal(GetBuyBookOrderCount() == 1
	//
	//    # then:
	//    ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "SELL", 10.00, 50.00, "GTC")
	//    assert.Equal(GetSellBookOrderCount() == 1
	//    assert.Equal(get_trade_count() == 0
	//    assert.Equal(GetAvailableLimitVolume("BUY", 10) == 50
	//    assert.Equal(GetAvailableLimitVolume("SELL", 10) == 50
	//  }
}
