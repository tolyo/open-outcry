/*
 * OPEN OUTCRY API
 *
 * # Introduction This API is documented in **OpenAPI 3.0 format**.  This API the following operations: * Retrieve a list of available instruments * Retrieve a list of executed trades  # Basics * API calls have to be secured with HTTPS. * All data has to be submitted UTF-8 encoded. * The reply is sent JSON encoded.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package api

import (
	"context"
	"errors"
	"net/http"
	"open-outcry/pkg/models"
)

// PublicAPIService is a service that implements the logic for the PublicAPIServicer
// This service should implement the business logic for every endpoint for the PublicAPI API.
// Include any external packages or services that will be required by this service.
type PublicAPIService struct {
}

// NewPublicAPIService creates a default api service
func NewPublicAPIService() PublicAPIServicer {
	return &PublicAPIService{}
}

// GetCurrencies - Currencies list
func (s *PublicAPIService) GetCurrencies(ctx context.Context) (ImplResponse, error) {
	currencies := models.GetCurrencies()
	res := make([]Currency, 0)
	for _, v := range currencies {
		cur := Currency{
			Name:      NewInterface(v.Name),
			Precision: NewInterface(v.Precision),
		}
		res = append(res, cur)
	}

	return Response(http.StatusOK, res), nil
}

// GetFxInstruments - Fx instrument list
func (s *PublicAPIService) GetFxInstruments(ctx context.Context) (ImplResponse, error) {
	// TODO - update GetFxInstruments with the required logic for this service method.
	// Add api_public_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, interface{}{}) or use other options such as http.Ok ...
	// return Response(200, interface{}{}), nil

	// TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	// return Response(404, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetFxInstruments method not implemented")
}

// GetInstruments - Instrument list
func (s *PublicAPIService) GetInstruments(ctx context.Context) (ImplResponse, error) {
	// TODO - update GetInstruments with the required logic for this service method.
	// Add api_public_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, interface{}{}) or use other options such as http.Ok ...
	// return Response(200, interface{}{}), nil

	// TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	// return Response(404, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetInstruments method not implemented")
}

// GetOrderBook - Get order book
func (s *PublicAPIService) GetOrderBook(ctx context.Context, instrumentName interface{}) (ImplResponse, error) {
	// TODO - update GetOrderBook with the required logic for this service method.
	// Add api_public_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, OrderBook{}) or use other options such as http.Ok ...
	// return Response(200, OrderBook{}), nil

	// TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	// return Response(404, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetOrderBook method not implemented")
}
