/*
 * OPEN OUTCRY API
 *
 * # Introduction This API is documented in **OpenAPI 3.0 format**.  This API the following operations: * Retrieve a list of available instruments * Retrieve a list of executed trades  # Basics * API calls have to be secured with HTTPS. * All data has to be submitted UTF-8 encoded. * The reply is sent JSON encoded.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package api

// CurrencyList - List of currencies supported by app
type CurrencyList struct {
	Data []Currency `json:"data,omitempty"`
}

// AssertCurrencyListRequired checks if the required fields are not zero-ed
func AssertCurrencyListRequired(obj CurrencyList) error {
	for _, el := range obj.Data {
		if err := AssertCurrencyRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertCurrencyListConstraints checks if the values respects the defined constraints
func AssertCurrencyListConstraints(obj CurrencyList) error {
	return nil
}
