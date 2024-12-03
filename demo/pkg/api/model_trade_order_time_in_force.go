/*
OPEN OUTCRY API

# Introduction This API is documented in **OpenAPI 3.0 format**.  This API the following operations: * Retrieve a list of available instruments * Retrieve a list of executed trades  # Basics * API calls have to be secured with HTTPS. * All data has to be submitted UTF-8 encoded. * The reply is sent JSON encoded.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// TradeOrderTimeInForce the model 'TradeOrderTimeInForce'
type TradeOrderTimeInForce string

// List of TradeOrderTimeInForce
const (
	GTC TradeOrderTimeInForce = "GTC"
	IOC TradeOrderTimeInForce = "IOC"
	FOK TradeOrderTimeInForce = "FOK"
	GTD TradeOrderTimeInForce = "GTD"
	GTT TradeOrderTimeInForce = "GTT"
)

// All allowed values of TradeOrderTimeInForce enum
var AllowedTradeOrderTimeInForceEnumValues = []TradeOrderTimeInForce{
	"GTC",
	"IOC",
	"FOK",
	"GTD",
	"GTT",
}

func (v *TradeOrderTimeInForce) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TradeOrderTimeInForce(value)
	for _, existing := range AllowedTradeOrderTimeInForceEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TradeOrderTimeInForce", value)
}

// NewTradeOrderTimeInForceFromValue returns a pointer to a valid TradeOrderTimeInForce
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTradeOrderTimeInForceFromValue(v string) (*TradeOrderTimeInForce, error) {
	ev := TradeOrderTimeInForce(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TradeOrderTimeInForce: valid values are %v", v, AllowedTradeOrderTimeInForceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TradeOrderTimeInForce) IsValid() bool {
	for _, existing := range AllowedTradeOrderTimeInForceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TradeOrderTimeInForce value
func (v TradeOrderTimeInForce) Ptr() *TradeOrderTimeInForce {
	return &v
}

type NullableTradeOrderTimeInForce struct {
	value *TradeOrderTimeInForce
	isSet bool
}

func (v NullableTradeOrderTimeInForce) Get() *TradeOrderTimeInForce {
	return v.value
}

func (v *NullableTradeOrderTimeInForce) Set(val *TradeOrderTimeInForce) {
	v.value = val
	v.isSet = true
}

func (v NullableTradeOrderTimeInForce) IsSet() bool {
	return v.isSet
}

func (v *NullableTradeOrderTimeInForce) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTradeOrderTimeInForce(val *TradeOrderTimeInForce) *NullableTradeOrderTimeInForce {
	return &NullableTradeOrderTimeInForce{value: val, isSet: true}
}

func (v NullableTradeOrderTimeInForce) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTradeOrderTimeInForce) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}