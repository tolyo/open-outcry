/*
 * OPEN OUTCRY API
 *
 * # Introduction This API is documented in **OpenAPI 3.0 format**.  This API the following operations: * Retrieve a list of available instruments * Retrieve a list of executed trades  # Basics * API calls have to be secured with HTTPS. * All data has to be submitted UTF-8 encoded. * The reply is sent JSON encoded.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package api

type OrderBook struct {
	Sell *interface{} `json:"sell,omitempty"`

	Buy *interface{} `json:"buy,omitempty"`
}

// AssertOrderBookRequired checks if the required fields are not zero-ed
func AssertOrderBookRequired(obj OrderBook) error {
	return nil
}

// AssertOrderBookConstraints checks if the values respects the defined constraints
func AssertOrderBookConstraints(obj OrderBook) error {
	return nil
}
