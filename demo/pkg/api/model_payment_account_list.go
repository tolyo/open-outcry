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

// checks if the PaymentAccountList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentAccountList{}

// PaymentAccountList List of payment accounts available to user
type PaymentAccountList struct {
	Data []PaymentAccount `json:"data,omitempty"`
}

// NewPaymentAccountList instantiates a new PaymentAccountList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentAccountList() *PaymentAccountList {
	this := PaymentAccountList{}
	return &this
}

// NewPaymentAccountListWithDefaults instantiates a new PaymentAccountList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentAccountListWithDefaults() *PaymentAccountList {
	this := PaymentAccountList{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PaymentAccountList) GetData() []PaymentAccount {
	if o == nil || IsNil(o.Data) {
		var ret []PaymentAccount
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentAccountList) GetDataOk() ([]PaymentAccount, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PaymentAccountList) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []PaymentAccount and assigns it to the Data field.
func (o *PaymentAccountList) SetData(v []PaymentAccount) {
	o.Data = v
}

func (o PaymentAccountList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentAccountList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePaymentAccountList struct {
	value *PaymentAccountList
	isSet bool
}

func (v NullablePaymentAccountList) Get() *PaymentAccountList {
	return v.value
}

func (v *NullablePaymentAccountList) Set(val *PaymentAccountList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentAccountList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentAccountList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentAccountList(val *PaymentAccountList) *NullablePaymentAccountList {
	return &NullablePaymentAccountList{value: val, isSet: true}
}

func (v NullablePaymentAccountList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentAccountList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
