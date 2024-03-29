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

// AdminAPIController binds http requests to an api service and writes the service results to the http response
type AdminAPIController struct {
	service      AdminAPIServicer
	errorHandler ErrorHandler
}

// AdminAPIOption for how the controller is set up.
type AdminAPIOption func(*AdminAPIController)

// WithAdminAPIErrorHandler inject ErrorHandler into controller
func WithAdminAPIErrorHandler(h ErrorHandler) AdminAPIOption {
	return func(c *AdminAPIController) {
		c.errorHandler = h
	}
}

// NewAdminAPIController creates a default api controller
func NewAdminAPIController(s AdminAPIServicer, opts ...AdminAPIOption) Router {
	controller := &AdminAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the AdminAPIController
func (c *AdminAPIController) Routes() Routes {
	return Routes{
		"CreateAdminPayment": Route{
			strings.ToUpper("Post"),
			"/apps/payments",
			c.CreateAdminPayment,
		},
		"GetAdminPaymentById": Route{
			strings.ToUpper("Post"),
			"/apps/payments/{payment_id}",
			c.GetAdminPaymentById,
		},
		"GetAppEntities": Route{
			strings.ToUpper("Get"),
			"/apps",
			c.GetAppEntities,
		},
		"GetAppEntity": Route{
			strings.ToUpper("Get"),
			"/apps/{app_entity_id}",
			c.GetAppEntity,
		},
	}
}

// CreateAdminPayment - Create admin payment
func (c *AdminAPIController) CreateAdminPayment(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.CreateAdminPayment(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// GetAdminPaymentById - Get payment
func (c *AdminAPIController) GetAdminPaymentById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	paymentIdParam := params["payment_id"]
	if paymentIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"payment_id"}, nil)
		return
	}
	result, err := c.service.GetAdminPaymentById(r.Context(), paymentIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// GetAppEntities - Get application entities
func (c *AdminAPIController) GetAppEntities(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.GetAppEntities(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// GetAppEntity - Get application entity
func (c *AdminAPIController) GetAppEntity(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	appEntityIdParam := params["app_entity_id"]
	if appEntityIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"app_entity_id"}, nil)
		return
	}
	result, err := c.service.GetAppEntity(r.Context(), appEntityIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}
