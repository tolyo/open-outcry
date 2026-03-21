package api

import (
	"time"

	appentity "open-outcry/pkg/models/app_entity"
	"open-outcry/pkg/models/currency"
	currencyaccount "open-outcry/pkg/models/currency_account"
	"open-outcry/pkg/models/instrument"
	instrumentaccount "open-outcry/pkg/models/instrument_account"
	orderbook "open-outcry/pkg/models/order_book"
	"open-outcry/pkg/models/trade"
	tradeorder "open-outcry/pkg/models/trade_order"
	"open-outcry/pkg/models/transfer"
)

func mapAppEntity(entity appentity.AppEntity) AppEntity {
	return AppEntity{
		Id:         string(entity.Id),
		ExternalId: string(entity.ExternalId),
	}
}

func mapAppEntities(entities []appentity.AppEntity) []AppEntity {
	result := make([]AppEntity, 0, len(entities))
	for _, entity := range entities {
		result = append(result, mapAppEntity(entity))
	}
	return result
}

func mapCurrency(entry currency.Currency) Currency {
	return Currency{
		Name:      string(entry.Name),
		Precision: int32(entry.Precision),
	}
}

func mapCurrencies(entries []currency.Currency) []Currency {
	result := make([]Currency, 0, len(entries))
	for _, entry := range entries {
		result = append(result, mapCurrency(entry))
	}
	return result
}

func mapCurrencyAccount(account currencyaccount.CurrencyAccount) CurrencyAccount {
	return CurrencyAccount{
		Id:              string(account.Id),
		Currency:        string(account.Currency),
		Amount:          account.Amount,
		AmountReserved:  account.AmountReserved,
		AmountAvailable: account.AmountAvailable,
	}
}

func mapCurrencyAccounts(accounts []currencyaccount.CurrencyAccount) []CurrencyAccount {
	result := make([]CurrencyAccount, 0, len(accounts))
	for _, account := range accounts {
		result = append(result, mapCurrencyAccount(account))
	}
	return result
}

func mapInstrument(entry instrument.Instrument) Instrument {
	return Instrument{
		Id:            string(entry.Id),
		Name:          string(entry.Name),
		QuoteCurrency: string(entry.QuoteCurrency),
		Enabled:       entry.Active,
	}
}

func mapInstruments(entries []instrument.Instrument) []Instrument {
	result := make([]Instrument, 0, len(entries))
	for _, entry := range entries {
		result = append(result, mapInstrument(entry))
	}
	return result
}

func mapFxInstrument(entry instrument.Instrument) FxInstrument {
	return FxInstrument{
		Id:            string(entry.Id),
		Name:          string(entry.Name),
		BaseCurrency:  string(entry.BaseCurrency),
		QuoteCurrency: string(entry.QuoteCurrency),
		Enabled:       entry.Active,
	}
}

func mapFxInstruments(entries []instrument.Instrument) []FxInstrument {
	result := make([]FxInstrument, 0, len(entries))
	for _, entry := range entries {
		result = append(result, mapFxInstrument(entry))
	}
	return result
}

func mapPriceVolume(entry orderbook.PriceVolume) PriceVolume {
	return PriceVolume{
		Price:  float32(entry.Price),
		Volume: float32(entry.Volume),
	}
}

func mapPriceVolumes(entries []orderbook.PriceVolume) []PriceVolume {
	result := make([]PriceVolume, 0, len(entries))
	for _, entry := range entries {
		result = append(result, mapPriceVolume(entry))
	}
	return result
}

func mapOrderBook(book orderbook.OrderBook) OrderBook {
	buy := mapPriceVolumes(book.BuySide)
	var buyValue interface{} = buy
	return OrderBook{
		Sell: mapPriceVolumes(book.SellSide),
		Buy:  &buyValue,
	}
}

func mapInstrumentAccountHolding(holding instrumentaccount.InstrumentAccountHolding) InstrumentAccountHolding {
	return InstrumentAccountHolding{
		Name:            string(holding.Name),
		Amount:          holding.AmountAvailable + holding.AmountReserved,
		AmountReserved:  holding.AmountReserved,
		AmountAvailable: holding.AmountAvailable,
		Value:           holding.Value,
		Currency:        string(holding.Currency),
	}
}

func mapInstrumentAccountHoldings(holdings []instrumentaccount.InstrumentAccountHolding) []InstrumentAccountHolding {
	result := make([]InstrumentAccountHolding, 0, len(holdings))
	for _, holding := range holdings {
		result = append(result, mapInstrumentAccountHolding(holding))
	}
	return result
}

func mapInstrumentAccount(account instrumentaccount.InstrumentAccount, holdings []instrumentaccount.InstrumentAccountHolding) InstrumentAccount {
	return InstrumentAccount{
		Id:          string(account.Id),
		Instruments: mapInstrumentAccountHoldings(holdings),
	}
}

func mapTrade(entry trade.Trade) Trade {
	return Trade{Id: entry.Id}
}

func mapTrades(entries []trade.Trade) []Trade {
	result := make([]Trade, 0, len(entries))
	for _, entry := range entries {
		result = append(result, mapTrade(entry))
	}
	return result
}

func mapTradeOrder(order tradeorder.TradeOrder) TradeOrder {
	created, _ := time.Parse(time.RFC3339Nano, order.Created)
	return TradeOrder{
		Id:          string(order.Id),
		Instrument:  string(order.InstrumentName),
		Side:        TradeOrderSide(order.Side),
		Type:        TradeOrderType(order.Type),
		TimeInForce: TradeOrderTimeInForce(order.TimeInForce),
		Status:      TradeOrderStatus(order.Status),
		Price:       float64(order.Price),
		Amount:      order.Amount,
		OpenAmount:  order.OpenAmount,
		Created:     created,
	}
}

func mapTradeOrders(orders []tradeorder.TradeOrder) []TradeOrder {
	result := make([]TradeOrder, 0, len(orders))
	for _, order := range orders {
		result = append(result, mapTradeOrder(order))
	}
	return result
}

func mapTransferEntry(entry transfer.TransferEntry) TransferEntry {
	return TransferEntry{
		Id:                      entry.Id,
		Type:                    TransferType(entry.Type),
		Amount:                  entry.Amount,
		Currency:                string(entry.Currency),
		SenderAccountId:         string(entry.SenderAccountId),
		BeneficiaryAccountId:    string(entry.BeneficiaryAccountId),
		Details:                 string(entry.Details),
		ExternalReferenceNumber: string(entry.ExternalReferenceNumber),
		Status:                  entry.Status,
		DebitBalanceAmount:      entry.DebitBalanceAmount,
		CreditBalanceAmount:     entry.CreditBalanceAmount,
	}
}
