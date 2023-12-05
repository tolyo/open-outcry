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
	"net/http"
	"open-outcry/pkg/models"
)

// CurrenciesAPIService is a service that implements the logic for the CurrenciesAPIServicer
// This service should implement the business logic for every endpoint for the CurrenciesAPI API.
// Include any external packages or services that will be required by this service.
type CurrenciesAPIService struct {
}

// NewCurrenciesAPIService creates a default api service
func NewCurrenciesAPIService() CurrenciesAPIServicer {
	return &CurrenciesAPIService{}
}

// GetCurrencies - currencies list
func (s *CurrenciesAPIService) GetCurrencies(ctx context.Context) (ImplResponse, error) {
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