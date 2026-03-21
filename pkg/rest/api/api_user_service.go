package api

import (
	"context"
	"net/http"
	appentity "open-outcry/pkg/models/app_entity"
	currencyaccount "open-outcry/pkg/models/currency_account"
	instrument "open-outcry/pkg/models/instrument"
	instrumentaccount "open-outcry/pkg/models/instrument_account"
	trade "open-outcry/pkg/models/trade"
	tradeorder "open-outcry/pkg/models/trade_order"
	"open-outcry/pkg/services"
	"time"
)

type UserAPIService struct{}

func NewUserAPIService() UserAPIServicer {
	return &UserAPIService{}
}

func (s *UserAPIService) CreateTrade(ctx context.Context, instrumentAccountId string, req CreateTradeRequest) (ImplResponse, error) {
	orderId, err := services.ProcessTradeOrder(
		instrumentaccount.InstrumentAccountId(instrumentAccountId),
		instrument.InstrumentName(req.Instrument),
		tradeorder.OrderType(req.Type),
		tradeorder.OrderSide(req.Side),
		tradeorder.OrderPrice(req.Price),
		req.Amount,
		tradeorder.OrderTimeInForce(req.TimeInForce),
	)
	if err != nil {
		return Response(http.StatusBadRequest, nil), err
	}

	order := tradeorder.GetTradeOrder(orderId)
	return Response(http.StatusOK, mapTradeOrder(order)), nil
}

func (s *UserAPIService) DeleteTradeOrderById(ctx context.Context, instrumentAccountId string, tradeOrderId string) (ImplResponse, error) {
	err := services.CancelTradeOrder(tradeorder.TradeOrderId(tradeOrderId))
	if err != nil {
		return Response(http.StatusNotFound, nil), err
	}
	return Response(http.StatusNoContent, nil), nil
}

func (s *UserAPIService) GetBookOrders(ctx context.Context, instrumentAccountId string) (ImplResponse, error) {
	orders := tradeorder.GetBookOrdersByInstrumentAccount(instrumentaccount.InstrumentAccountId(instrumentAccountId))
	var result []TradeOrder
	for _, o := range orders {
		result = append(result, mapTradeOrder(o))
	}
	return Response(http.StatusOK, result), nil
}

func (s *UserAPIService) GetCurrencyAccounts(ctx context.Context, appEntityId string) (ImplResponse, error) {
	accounts := currencyaccount.GetCurrencyAccountsByAppEntity(appentity.AppEntityId(appEntityId))
	var result []CurrencyAccount
	for _, a := range accounts {
		result = append(result, CurrencyAccount{Id: string(a.Id), Currency: string(a.Currency), Amount: a.Amount, AmountReserved: a.AmountReserved, AmountAvailable: a.AmountAvailable})
	}
	return Response(http.StatusOK, CurrencyAccountList{Data: result}), nil
}

func (s *UserAPIService) GetTradeById(ctx context.Context, instrumentAccountId string, tradeId string) (ImplResponse, error) {
	entry := trade.GetTrade(tradeId)
	if entry == nil {
		return Response(http.StatusNotFound, nil), nil
	}
	return Response(http.StatusOK, Trade{Id: entry.Id}), nil
}

func (s *UserAPIService) GetTradeOrderById(ctx context.Context, instrumentAccountId string, tradeOrderId string) (ImplResponse, error) {
	order := tradeorder.GetTradeOrder(tradeorder.TradeOrderId(tradeOrderId))
	return Response(http.StatusOK, mapTradeOrder(order)), nil
}

func (s *UserAPIService) GetTradeOrders(ctx context.Context, instrumentAccountId string) (ImplResponse, error) {
	orders := tradeorder.GetTradeOrdersByInstrumentAccount(instrumentaccount.InstrumentAccountId(instrumentAccountId))
	var result []TradeOrder
	for _, o := range orders {
		result = append(result, mapTradeOrder(o))
	}
	return Response(http.StatusOK, result), nil
}

func (s *UserAPIService) GetTrades(ctx context.Context, instrumentAccountId string) (ImplResponse, error) {
	trades := trade.GetTradesByInstrumentAccount(instrumentaccount.InstrumentAccountId(instrumentAccountId))
	var result []Trade
	for _, t := range trades {
		result = append(result, Trade{Id: t.Id})
	}
	return Response(http.StatusOK, result), nil
}

func (s *UserAPIService) GetInstrumentAccount(ctx context.Context, instrumentAccountId string) (ImplResponse, error) {
	account := instrumentaccount.GetInstrumentAccount(instrumentaccount.InstrumentAccountId(instrumentAccountId))
	if account == nil {
		return Response(http.StatusNotFound, nil), nil
	}

	instruments := instrumentaccount.GetInstrumentAccountHoldings(instrumentaccount.InstrumentAccountId(instrumentAccountId))
	var apiInstruments []InstrumentAccountHolding
	for _, inst := range instruments {
		apiInstruments = append(apiInstruments, InstrumentAccountHolding{Name: string(inst.Name), Amount: inst.AmountAvailable + inst.AmountReserved, AmountReserved: inst.AmountReserved, AmountAvailable: inst.AmountAvailable, Value: inst.Value, Currency: string(inst.Currency)})
	}

	return Response(http.StatusOK, InstrumentAccount{Id: string(account.Id), Instruments: apiInstruments}), nil
}

func mapTradeOrder(o tradeorder.TradeOrder) TradeOrder {
	created, _ := time.Parse(time.RFC3339Nano, o.Created)
	return TradeOrder{Id: string(o.Id), Instrument: string(o.InstrumentName), Side: TradeOrderSide(o.Side), Type: TradeOrderType(o.Type), TimeInForce: TradeOrderTimeInForce(o.TimeInForce), Status: TradeOrderStatus(o.Status), Price: float64(o.Price), Amount: o.Amount, OpenAmount: o.OpenAmount, Created: created}
}
