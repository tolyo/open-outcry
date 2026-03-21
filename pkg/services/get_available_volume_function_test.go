package services

import (
	order "open-outcry/pkg/models/order_book"
	"open-outcry/pkg/models/trade_order"
)

var volumeCases = []MatchingServiceTestCase{
	{steps: []TestStep{
		{
			expectedState: AppState{orderBookStates: order.OrderBook{
				SellSide: []order.PriceVolume{{Price: 10.00, Volume: 0.0}},
				BuySide:  []order.PriceVolume{{Price: 10.00, Volume: 0.0}},
			}},
		},
		{
			orders: []tradeorder.TradeOrder{{Side: tradeorder.Sell, Type: tradeorder.Limit, Price: 10, Amount: 100, TimeInForce: tradeorder.GTC}},
			expectedState: AppState{orderBookStates: order.OrderBook{
				SellSide: []order.PriceVolume{{Price: 10.00, Volume: 100.0}, {Price: 11.00, Volume: 100.0}, {Price: 9.00, Volume: 0.0}},
				BuySide:  []order.PriceVolume{{Price: 10.00, Volume: 0.0}, {Price: 11.00, Volume: 0.0}, {Price: 9.00, Volume: 0.0}},
			}},
		},
		{
			orders: []tradeorder.TradeOrder{{Side: tradeorder.Sell, Type: tradeorder.Limit, Price: 10, Amount: 100, TimeInForce: tradeorder.GTC}},
			expectedState: AppState{orderBookStates: order.OrderBook{
				SellSide: []order.PriceVolume{{Price: 10.00, Volume: 200.0}, {Price: 11.00, Volume: 200.0}, {Price: 9.00, Volume: 0.0}},
				BuySide:  []order.PriceVolume{{Price: 10.00, Volume: 0.0}, {Price: 11.00, Volume: 0.0}, {Price: 9.00, Volume: 0.0}},
			}},
		},
		{
			orders: []tradeorder.TradeOrder{{Side: tradeorder.Sell, Type: tradeorder.Limit, Price: 9, Amount: 100, TimeInForce: tradeorder.GTC}},
			expectedState: AppState{orderBookStates: order.OrderBook{
				SellSide: []order.PriceVolume{{Price: 10.00, Volume: 300.0}, {Price: 11.00, Volume: 300.0}, {Price: 9.00, Volume: 100.0}, {Price: 8.00, Volume: 0.0}},
				BuySide:  []order.PriceVolume{{Price: 10.00, Volume: 0.0}, {Price: 11.00, Volume: 0.0}, {Price: 9.00, Volume: 0.0}, {Price: 8.00, Volume: 0.0}},
			}},
		},
	}},
	{steps: []TestStep{
		{
			orders: []tradeorder.TradeOrder{{Side: tradeorder.Buy, Type: tradeorder.Limit, Price: 10, Amount: 10, TimeInForce: tradeorder.GTC}},
			expectedState: AppState{orderBookStates: order.OrderBook{
				SellSide: []order.PriceVolume{{Price: 10.00, Volume: 0.0}, {Price: 11.00, Volume: 0.0}, {Price: 9.00, Volume: 0.0}},
				BuySide:  []order.PriceVolume{{Price: 10.00, Volume: 10.0}, {Price: 11.00, Volume: 0.0}, {Price: 9.00, Volume: 10.0}},
			}},
		},
		{
			orders: []tradeorder.TradeOrder{{Side: tradeorder.Buy, Type: tradeorder.Limit, Price: 10, Amount: 10, TimeInForce: tradeorder.GTC}},
			expectedState: AppState{orderBookStates: order.OrderBook{
				SellSide: []order.PriceVolume{{Price: 10.00, Volume: 0.0}, {Price: 11.00, Volume: 0.0}, {Price: 9.00, Volume: 0.0}},
				BuySide:  []order.PriceVolume{{Price: 10.00, Volume: 20.0}, {Price: 11.00, Volume: 0.0}, {Price: 9.00, Volume: 20.0}},
			}},
		},
		{
			orders: []tradeorder.TradeOrder{{Side: tradeorder.Buy, Type: tradeorder.Limit, Price: 9, Amount: 10, TimeInForce: tradeorder.GTC}},
			expectedState: AppState{orderBookStates: order.OrderBook{
				SellSide: []order.PriceVolume{{Price: 10.00, Volume: 0.0}, {Price: 11.00, Volume: 0.0}, {Price: 9.00, Volume: 0.0}},
				BuySide:  []order.PriceVolume{{Price: 10.00, Volume: 20.0}, {Price: 10.001, Volume: 0.0}, {Price: 11.00, Volume: 0.0}, {Price: 9.00, Volume: 30.0}, {Price: 9.99, Volume: 20.0}},
			}},
		},
	}},
}

func (assert *ServiceTestSuite) TestGetAvailableLimitVolumeEmpty() {
	RunTestCases(assert, volumeCases)
}
