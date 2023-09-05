/*
 * OPEN OUTCRY API
 *
 * # Introduction This API is documented in **OpenAPI 3.0 format**.  This API the following operations: * Retrieve a list of available instruments * Retrieve a list of executed trades  # Basics * API calls have to be secured with HTTPS. * All data has to be submitted UTF-8 encoded. * The reply is sent JSON encoded.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package api

type TradeOrder struct {
	Id string `json:"id,omitempty"`

	// Ticker-like name of the instrument. For monetary instruments, a currency pair is used.
	Instrument string `json:"instrument,omitempty"`

	Side TradeOrderSide `json:"side,omitempty"`

	TimeInForce TradeOrderTimeInForce `json:"timeInForce,omitempty"`

	Status TradeOrderStatus `json:"status,omitempty"`
}

// AssertTradeOrderRequired checks if the required fields are not zero-ed
func AssertTradeOrderRequired(obj TradeOrder) error {
	return nil
}

// AssertTradeOrderConstraints checks if the values respects the defined constraints
func AssertTradeOrderConstraints(obj TradeOrder) error {
	return nil
}
