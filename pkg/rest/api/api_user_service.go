package api

import (
	"context"
	"net/http"

	appentity "open-outcry/pkg/models/app_entity"
	currencyaccount "open-outcry/pkg/models/currency_account"
	"open-outcry/pkg/models/instrument"
	instrumentaccount "open-outcry/pkg/models/instrument_account"
	"open-outcry/pkg/models/trade"
	tradeorder "open-outcry/pkg/models/trade_order"
	"open-outcry/pkg/services"
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

	return Response(http.StatusOK, mapTradeOrder(tradeorder.GetTradeOrder(orderId))), nil
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
	return Response(http.StatusOK, mapTradeOrders(orders)), nil
}

func (s *UserAPIService) GetCurrencyAccounts(ctx context.Context, appEntityId string) (ImplResponse, error) {
	accounts := currencyaccount.GetCurrencyAccountsByAppEntity(appentity.AppEntityId(appEntityId))
	return Response(http.StatusOK, CurrencyAccountList{Data: mapCurrencyAccounts(accounts)}), nil
}

func (s *UserAPIService) GetTradeById(ctx context.Context, instrumentAccountId string, tradeId string) (ImplResponse, error) {
	entry := trade.GetTrade(tradeId)
	if entry == nil {
		return Response(http.StatusNotFound, nil), nil
	}
	return Response(http.StatusOK, mapTrade(*entry)), nil
}

func (s *UserAPIService) GetTradeOrderById(ctx context.Context, instrumentAccountId string, tradeOrderId string) (ImplResponse, error) {
	return Response(http.StatusOK, mapTradeOrder(tradeorder.GetTradeOrder(tradeorder.TradeOrderId(tradeOrderId)))), nil
}

func (s *UserAPIService) GetTradeOrders(ctx context.Context, instrumentAccountId string) (ImplResponse, error) {
	orders := tradeorder.GetTradeOrdersByInstrumentAccount(instrumentaccount.InstrumentAccountId(instrumentAccountId))
	return Response(http.StatusOK, mapTradeOrders(orders)), nil
}

func (s *UserAPIService) GetTrades(ctx context.Context, instrumentAccountId string) (ImplResponse, error) {
	trades := trade.GetTradesByInstrumentAccount(instrumentaccount.InstrumentAccountId(instrumentAccountId))
	return Response(http.StatusOK, mapTrades(trades)), nil
}

func (s *UserAPIService) GetInstrumentAccount(ctx context.Context, instrumentAccountId string) (ImplResponse, error) {
	account := instrumentaccount.GetInstrumentAccount(instrumentaccount.InstrumentAccountId(instrumentAccountId))
	if account == nil {
		return Response(http.StatusNotFound, nil), nil
	}

	holdings := instrumentaccount.GetInstrumentAccountHoldings(instrumentaccount.InstrumentAccountId(instrumentAccountId))
	return Response(http.StatusOK, mapInstrumentAccount(*account, holdings)), nil
}
