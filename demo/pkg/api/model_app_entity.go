/*
OPEN OUTCRY API

# Introduction This API is documented in **OpenAPI 3.0 format**.  This API the following operations: * Retrieve a list of available instruments * Retrieve a list of executed trades  # Basics * API calls have to be secured with HTTPS. * All data has to be submitted UTF-8 encoded. * The reply is sent JSON encoded.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the AppEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppEntity{}

// AppEntity Registered user
type AppEntity struct {
	Id *string `json:"id,omitempty"`
	// External id
	ExternalId *string `json:"external_id,omitempty"`
}

// NewAppEntity instantiates a new AppEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppEntity() *AppEntity {
	this := AppEntity{}
	return &this
}

// NewAppEntityWithDefaults instantiates a new AppEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppEntityWithDefaults() *AppEntity {
	this := AppEntity{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AppEntity) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppEntity) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AppEntity) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AppEntity) SetId(v string) {
	o.Id = &v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *AppEntity) GetExternalId() string {
	if o == nil || IsNil(o.ExternalId) {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppEntity) GetExternalIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalId) {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *AppEntity) HasExternalId() bool {
	if o != nil && !IsNil(o.ExternalId) {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *AppEntity) SetExternalId(v string) {
	o.ExternalId = &v
}

func (o AppEntity) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppEntity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.ExternalId) {
		toSerialize["external_id"] = o.ExternalId
	}
	return toSerialize, nil
}

type NullableAppEntity struct {
	value *AppEntity
	isSet bool
}

func (v NullableAppEntity) Get() *AppEntity {
	return v.value
}

func (v *NullableAppEntity) Set(val *AppEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableAppEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableAppEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppEntity(val *AppEntity) *NullableAppEntity {
	return &NullableAppEntity{value: val, isSet: true}
}

func (v NullableAppEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
