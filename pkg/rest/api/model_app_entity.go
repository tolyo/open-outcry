/*
 * OPEN OUTCRY API
 *
 * # Introduction This API is documented in **OpenAPI 3.0 format**.  This API the following operations: * Retrieve a list of available instruments * Retrieve a list of executed trades  # Basics * API calls have to be secured with HTTPS. * All data has to be submitted UTF-8 encoded. * The reply is sent JSON encoded.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package api

// AppEntity - Registered user
type AppEntity struct {
	Id string `json:"id,omitempty"`

	// External id
	ExternalId string `json:"external_id,omitempty"`
}

// AssertAppEntityRequired checks if the required fields are not zero-ed
func AssertAppEntityRequired(obj AppEntity) error {
	return nil
}

// AssertAppEntityConstraints checks if the values respects the defined constraints
func AssertAppEntityConstraints(obj AppEntity) error {
	return nil
}
