package services

import (
	"open-outcry/pkg/models"
	"open-outcry/pkg/utils"
	"reflect"
)

var testcases = []MatchingServiceTestCase{
	{steps: []TestStep{
		// Expectations:
		// - reserve balance should increase
		{
			initialState: AppState{
				entity1: []models.PaymentAccount{
					{
						Amount:          1000,
						AmountAvailable: 1000,
						AmountReserved:  0,
						Currency:        "EUR",
					},
				},
				entity2: nil,
			},
			orders: []models.TradeOrder{
				{Side: models.Buy, Type: models.Limit, Price: 10, Amount: 10, TimeInForce: models.GTC},
			},
			expectedState: AppState{
				entity1: []models.PaymentAccount{
					{
						Amount:          1000,
						AmountAvailable: 900,
						AmountReserved:  100,
						Currency:        "EUR",
					},
				},
				tradeCount: 0,
			},
		},

		{
			initialState: AppState{
				entity1: []models.PaymentAccount{
					{
						Amount:          1000,
						AmountAvailable: 900,
						AmountReserved:  100,
						Currency:        "EUR",
					},
				},
				entity2: nil,
			},
			orders: []models.TradeOrder{
				{Side: models.Buy, Type: models.Limit, Price: 10, Amount: 10, TimeInForce: models.GTC},
			},
			expectedState: AppState{
				entity1: []models.PaymentAccount{
					{
						Amount:          1000,
						AmountAvailable: 800,
						AmountReserved:  200,
						Currency:        "EUR",
					},
				},
				tradeCount: 0,
			},
		},
	}},

	{steps: []TestStep{
		// Expectations:
		// - reserve balance should increase
		{
			initialState: AppState{
				entity1: []models.PaymentAccount{
					{
						Amount:          1000,
						AmountAvailable: 1000,
						AmountReserved:  0,
						Currency:        "BTC",
					},
				},
				entity2: nil,
			},
			orders: []models.TradeOrder{
				{Side: models.Sell, Type: models.Limit, Price: 10, Amount: 10, TimeInForce: models.GTC},
			},
			expectedState: AppState{
				entity1: []models.PaymentAccount{
					{
						Amount:          1000,
						AmountAvailable: 990,
						AmountReserved:  10,
						Currency:        "BTC",
					},
				},
				tradeCount: 0,
			},
		},

		{
			initialState: AppState{
				entity1: []models.PaymentAccount{
					{
						Amount:          1000,
						AmountAvailable: 990,
						AmountReserved:  10,
						Currency:        "BTC",
					},
				},
				entity2: nil,
			},
			orders: []models.TradeOrder{
				{Side: models.Sell, Type: models.Limit, Price: 10, Amount: 10, TimeInForce: models.GTC},
			},
			expectedState: AppState{
				entity1: []models.PaymentAccount{
					{
						Amount:          1000,
						AmountAvailable: 980,
						AmountReserved:  20,
						Currency:        "BTC",
					},
				},
				tradeCount: 0,
			},
		},
	}},
}

func (assert *ServiceTestSuite) TestProcessLimitOrderReservedBalance() {
	RunTestCases(assert, testcases)
}

func RunTestCases(assert *ServiceTestSuite, cases []MatchingServiceTestCase) {
	for _, td := range cases {
		assert.TearDownTest()
		assert.SetupTest()

		for _, step := range td.steps {

			// given:
			expect := func(expectedState AppState) {
				utils.Each(expectedState.entity1, func(account models.PaymentAccount) {
					var paymentAccount = models.FindPaymentAccountByAppEntityIdAndCurrencyName(assert.appEntity1,
						account.Currency)
					assert.Equal(account.Amount, paymentAccount.Amount)
					assert.Equal(account.AmountAvailable, paymentAccount.AmountAvailable)
					assert.Equal(account.AmountReserved, paymentAccount.AmountReserved)
				})
				if fieldExists(expectedState, "tradeCount") {
					assert.Equal(expectedState.tradeCount, GetTradeCount())
				}

				if fieldExists(expectedState, "orderBookStates") {
					utils.Each(expectedState.orderBookStates.BuySide, func(level models.PriceVolume) {
						assert.Equal(level.Volume, GetAvailableLimitVolume(models.Buy, models.OrderPrice(level.Price)))
					})
					utils.Each(expectedState.orderBookStates.SellSide, func(level models.PriceVolume) {
						assert.Equal(level.Volume, GetAvailableLimitVolume(models.Sell, models.OrderPrice(level.Price)))
					})
				}
			}
			// then:
			expect(step.initialState)
			utils.Each(step.orders, func(order models.TradeOrder) {
				orderId, err := ProcessTradeOrder(assert.tradingAccount1,
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
