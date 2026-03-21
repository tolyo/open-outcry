package api

import (
	"context"
	"net/http"
	currency "open-outcry/pkg/models/currency"
	instrument "open-outcry/pkg/models/instrument"
	"open-outcry/pkg/services"
)

type PublicAPIService struct{}

func NewPublicAPIService() PublicAPIServicer {
	return &PublicAPIService{}
}

func (s *PublicAPIService) GetCurrencies(ctx context.Context) (ImplResponse, error) {
	currencies := currency.GetCurrencies()
	var result []Currency
	for _, c := range currencies {
		result = append(result, Currency{Name: string(c.Name), Precision: int32(c.Precision)})
	}
	return Response(http.StatusOK, CurrencyList{Data: result}), nil
}

func (s *PublicAPIService) GetFxInstruments(ctx context.Context) (ImplResponse, error) {
	instruments := instrument.GetFxInstruments()
	var result []FxInstrument
	for _, i := range instruments {
		result = append(result, FxInstrument{Id: string(i.Id), Name: string(i.Name), BaseCurrency: string(i.BaseCurrency), QuoteCurrency: string(i.QuoteCurrency)})
	}
	return Response(http.StatusOK, result), nil
}

func (s *PublicAPIService) GetInstruments(ctx context.Context) (ImplResponse, error) {
	instruments := instrument.GetInstruments()
	var result []Instrument
	for _, i := range instruments {
		result = append(result, Instrument{Id: string(i.Id), Name: string(i.Name), QuoteCurrency: string(i.QuoteCurrency)})
	}
	return Response(http.StatusOK, result), nil
}

func (s *PublicAPIService) GetOrderBook(ctx context.Context, instrumentName string) (ImplResponse, error) {
	ob := services.GetOrderBook(instrument.InstrumentName(instrumentName))

	var sellSide []PriceVolume
	for _, pv := range ob.SellSide {
		sellSide = append(sellSide, PriceVolume{Price: float32(pv.Price), Volume: float32(pv.Volume)})
	}

	var buySide []PriceVolume
	for _, pv := range ob.BuySide {
		buySide = append(buySide, PriceVolume{Price: float32(pv.Price), Volume: float32(pv.Volume)})
	}

	var buySideInterface interface{} = buySide
	return Response(http.StatusOK, OrderBook{Sell: sellSide, Buy: &buySideInterface}), nil
}
