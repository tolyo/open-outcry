package api

import (
	"context"
	"net/http"

	"open-outcry/pkg/models/currency"
	"open-outcry/pkg/models/instrument"
	"open-outcry/pkg/services"
)

type PublicAPIService struct{}

func NewPublicAPIService() PublicAPIServicer {
	return &PublicAPIService{}
}

func (s *PublicAPIService) GetCurrencies(ctx context.Context) (ImplResponse, error) {
	return Response(http.StatusOK, CurrencyList{Data: mapCurrencies(currency.GetCurrencies())}), nil
}

func (s *PublicAPIService) GetFxInstruments(ctx context.Context) (ImplResponse, error) {
	return Response(http.StatusOK, mapFxInstruments(instrument.GetFxInstruments())), nil
}

func (s *PublicAPIService) GetInstruments(ctx context.Context) (ImplResponse, error) {
	return Response(http.StatusOK, mapInstruments(instrument.GetInstruments())), nil
}

func (s *PublicAPIService) GetOrderBook(ctx context.Context, instrumentName string) (ImplResponse, error) {
	return Response(http.StatusOK, mapOrderBook(services.GetOrderBook(instrument.InstrumentName(instrumentName)))), nil
}
