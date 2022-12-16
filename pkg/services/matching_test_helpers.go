package services

import (
	"open-outcry/pkg/db"
	"open-outcry/pkg/models"
)

func GetSellBookOrderCount() int {
	return db.QueryVal[int](
		`
		SELECT COUNT(*) 
		FROM trade_order t 
		    INNER JOIN book_order b 
			ON t.id = b.trade_order_id 
		WHERE t.side = 'SELL'
	`)
}
func GetBuyBookOrderCount() int {
	return db.QueryVal[int](
		`
		SELECT COUNT(*) 
		FROM tradeOrder t 
		    INNER JOIN book_order b 
			ON t.id = b.trade_order_id 
		WHERE t.side = 'BUY'
	`)
}

func GetTradeCount() int {
	return db.QueryVal[int]("SELECT COUNT(*) FROM trade")
}

//
//  func get_trade_prices() {
//    DB.query_list("SELECT (price) FROM trade ORDER BY created_at ASC")
//    |> Enum.map(&Decimal.to_float(&1))
//  }
//
//  @spec get_crossing_limit_orders(number(), TradeOrder.Side.t(), Decimal.t()) [any]
//  func get_crossing_limit_orders(instrument_id, side, price) {
//    [instrument_id, side |> Atom.to_string(), price, 100_000]
//    |> DB.query_list("SELECT get_crossing_limit_orders($1, $2, $3, $4)")
//  }

func GetAvailableLimitVolume(side models.OrderSide, price models.OrderPrice) float64 {
	return db.QueryVal[float64]("SELECT get_available_limit_volume(1, $1::order_side, $2)", side, price)
}

//
//  @spec get_payment_account(TradingAccount.id(), PaymentAccount.currency()) ::
//          :none | PaymentAccount.t()
//  func get_payment_account(tradingAccount, currency) {
//    TradingAccount.get(tradingAccount).appEntityId
//    |> FindPaymentAccountByAppEntityIdAndCurrencyName(currency)
//  }
//}
