/*
 * OPEN OUTCRY API
 *
 * # Introduction This API is documented in **OpenAPI 3.0 format**.  This API the following operations: * Retrieve a list of available instruments * Retrieve a list of executed trades  # Basics * API calls have to be secured with HTTPS. * All data has to be submitted UTF-8 encoded. * The reply is sent JSON encoded.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package api

type CreateTradeRequest struct {

	// Ticker-like name of the instrument. For monetary instruments, a currency pair is used.
	Instrument *interface{} `json:"instrument"`

	Side TradeOrderSide `json:"side"`

	Type TradeOrderType `json:"type"`

	TimeInForce TradeOrderTimeInForce `json:"timeInForce"`

	Amount *interface{} `json:"amount"`

	Price *interface{} `json:"price,omitempty"`
}

// AssertCreateTradeRequestRequired checks if the required fields are not zero-ed
func AssertCreateTradeRequestRequired(obj CreateTradeRequest) error {
	elements := map[string]interface{}{
		"instrument":  obj.Instrument,
		"side":        obj.Side,
		"type":        obj.Type,
		"timeInForce": obj.TimeInForce,
		"amount":      obj.Amount,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertCreateTradeRequestConstraints checks if the values respects the defined constraints
func AssertCreateTradeRequestConstraints(obj CreateTradeRequest) error {
	return nil
}