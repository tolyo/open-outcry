package services

var volumeCases = []MatchingServiceTestCase{
	// Test for available volume on the sell side. Available volume should increase
	// if the order is on the sell side and order limit price is below or equal the query limit price.
	{steps: []TestStep{
		{
			expectedState: AppState{orderBookStates: OrderBook{
				SellSide: []PriceVolume{
					{Price: 10.00, Volume: 0.0},
				},
				BuySide: []PriceVolume{
					{Price: 10.00, Volume: 0.0},
				},
			}},
		},
		{
			orders: []TradeOrder{
				{Side: Sell, Type: Limit, Price: 10, Amount: 100, TimeInForce: GTC},
			},
			expectedState: AppState{orderBookStates: OrderBook{
				SellSide: []PriceVolume{
					{Price: 10.00, Volume: 100.0},
					{Price: 11.00, Volume: 100.0},
					{Price: 9.00, Volume: 0.0},
				},
				BuySide: []PriceVolume{
					{Price: 10.00, Volume: 0.0},
					{Price: 11.00, Volume: 0.0},
					{Price: 9.00, Volume: 0.0},
				},
			}},
		},

		{
			orders: []TradeOrder{
				{Side: Sell, Type: Limit, Price: 10, Amount: 100, TimeInForce: GTC},
			},
			expectedState: AppState{orderBookStates: OrderBook{
				SellSide: []PriceVolume{
					{Price: 10.00, Volume: 200.0},
					{Price: 11.00, Volume: 200.0},
					{Price: 9.00, Volume: 0.0},
				},
				BuySide: []PriceVolume{
					{Price: 10.00, Volume: 0.0},
					{Price: 11.00, Volume: 0.0},
					{Price: 9.00, Volume: 0.0},
				},
			}},
		},

		{
			orders: []TradeOrder{
				{Side: Sell, Type: Limit, Price: 9, Amount: 100, TimeInForce: GTC},
			},
			expectedState: AppState{orderBookStates: OrderBook{
				SellSide: []PriceVolume{
					{Price: 10.00, Volume: 300.0},
					{Price: 11.00, Volume: 300.0},
					{Price: 9.00, Volume: 100.0},
					{Price: 8.00, Volume: 0.0},
				},
				BuySide: []PriceVolume{
					{Price: 10.00, Volume: 0.0},
					{Price: 11.00, Volume: 0.0},
					{Price: 9.00, Volume: 0.0},
					{Price: 8.00, Volume: 0.0},
				},
			}},
		},
	}},

	{steps: []TestStep{
		{
			orders: []TradeOrder{
				{Side: Buy, Type: Limit, Price: 10, Amount: 10, TimeInForce: GTC},
			},
			expectedState: AppState{orderBookStates: OrderBook{
				SellSide: []PriceVolume{
					{Price: 10.00, Volume: 0.0},
					{Price: 11.00, Volume: 0.0},
					{Price: 9.00, Volume: 0.0},
				},
				BuySide: []PriceVolume{
					{Price: 10.00, Volume: 10.0},
					{Price: 11.00, Volume: 0.0},
					{Price: 9.00, Volume: 10.0},
				},
			}},
		},

		{
			orders: []TradeOrder{
				{Side: Buy, Type: Limit, Price: 10, Amount: 10, TimeInForce: GTC},
			},
			expectedState: AppState{orderBookStates: OrderBook{
				SellSide: []PriceVolume{
					{Price: 10.00, Volume: 0.0},
					{Price: 11.00, Volume: 0.0},
					{Price: 9.00, Volume: 0.0},
				},
				BuySide: []PriceVolume{
					{Price: 10.00, Volume: 20.0},
					{Price: 11.00, Volume: 0.0},
					{Price: 9.00, Volume: 20.0},
				},
			}},
		},

		{
			orders: []TradeOrder{
				{Side: Buy, Type: Limit, Price: 9, Amount: 10, TimeInForce: GTC},
			},
			expectedState: AppState{orderBookStates: OrderBook{
				SellSide: []PriceVolume{
					{Price: 10.00, Volume: 0.0},
					{Price: 11.00, Volume: 0.0},
					{Price: 9.00, Volume: 0.0},
				},
				BuySide: []PriceVolume{
					{Price: 10.00, Volume: 20.0},
					{Price: 10.001, Volume: 0.0},
					{Price: 11.00, Volume: 0.0},
					{Price: 9.00, Volume: 30.0},
					{Price: 9.99, Volume: 20.0},
				},
			}},
		},
	}},
}

func (assert *ServiceTestSuite) TestGetAvailableLimitVolumeEmpty() {
	RunTestCases(assert, volumeCases)
}
