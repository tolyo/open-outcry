package services

import (
	currencyaccount "open-outcry/pkg/models/currency_account"
	orderbook "open-outcry/pkg/models/order_book"
	tradeorder "open-outcry/pkg/models/trade_order"
	"open-outcry/pkg/utils"
	"reflect"
)

func RunTestCases(assert *ServiceTestSuite, cases []MatchingServiceTestCase) {
	for _, td := range cases {
		assert.TearDownTest()
		assert.SetupTest()

		for _, step := range td.steps {
			expect := func(expectedState AppState) {
				utils.Each(expectedState.entity1, func(account currencyaccount.CurrencyAccount) {
					currencyAccount := currencyaccount.FindCurrencyAccountByAppEntityIdAndCurrencyName(assert.appEntity1, account.Currency)
					assert.Equal(account.Amount, currencyAccount.Amount)
					assert.Equal(account.AmountAvailable, currencyAccount.AmountAvailable)
					assert.Equal(account.AmountReserved, currencyAccount.AmountReserved)
				})
				if fieldExists(expectedState, "tradeCount") {
					assert.Equal(expectedState.tradeCount, GetTradeCount())
				}
				if fieldExists(expectedState, "orderBookStates") {
					utils.Each(expectedState.orderBookStates.BuySide, func(level orderbook.PriceVolume) {
						assert.Equal(level.Volume, GetAvailableLimitVolume(tradeorder.Buy, tradeorder.OrderPrice(level.Price)))
					})
					utils.Each(expectedState.orderBookStates.SellSide, func(level orderbook.PriceVolume) {
						assert.Equal(level.Volume, GetAvailableLimitVolume(tradeorder.Sell, tradeorder.OrderPrice(level.Price)))
					})
				}
			}
			expect(step.initialState)
			utils.Each(step.orders, func(order tradeorder.TradeOrder) {
				orderId, err := ProcessTradeOrder(assert.instrumentAccount1,
					"BTC_EUR",
					order.Type,
					order.Side,
					order.Price,
					order.Amount,
					order.TimeInForce,
				)
				assert.Nil(err)
				assert.NotNil(orderId)
			})
			expect(step.expectedState)
		}
	}
}

func fieldExists(s interface{}, fieldName string) bool {
	structType := reflect.TypeOf(s)
	_, found := structType.FieldByName(fieldName)
	return found
}
