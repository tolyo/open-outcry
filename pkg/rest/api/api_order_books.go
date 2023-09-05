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
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// OrderBooksAPIController binds http requests to an api service and writes the service results to the http response
type OrderBooksAPIController struct {
	service      OrderBooksAPIServicer
	errorHandler ErrorHandler
}

// OrderBooksAPIOption for how the controller is set up.
type OrderBooksAPIOption func(*OrderBooksAPIController)

// WithOrderBooksAPIErrorHandler inject ErrorHandler into controller
func WithOrderBooksAPIErrorHandler(h ErrorHandler) OrderBooksAPIOption {
	return func(c *OrderBooksAPIController) {
		c.errorHandler = h
	}
}

// NewOrderBooksAPIController creates a default api controller
func NewOrderBooksAPIController(s OrderBooksAPIServicer, opts ...OrderBooksAPIOption) Router {
	controller := &OrderBooksAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the OrderBooksAPIController
func (c *OrderBooksAPIController) Routes() Routes {
	return Routes{
		"GetOrderBook": Route{
			strings.ToUpper("Get"),
			"/order_books/{instrument_name}",
			c.GetOrderBook,
		},
	}
}

// GetOrderBook - get order books
func (c *OrderBooksAPIController) GetOrderBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	instrumentNameParam := params["instrument_name"]
	result, err := c.service.GetOrderBook(r.Context(), instrumentNameParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}
