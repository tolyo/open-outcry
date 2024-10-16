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

// TradeOrderType the model 'TradeOrderType'
type TradeOrderType string

// List of TradeOrderType
const (
	LIMIT     TradeOrderType = "LIMIT"
	MARKET    TradeOrderType = "MARKET"
	STOPLOSS  TradeOrderType = "STOPLOSS"
	STOPLIMIT TradeOrderType = "STOPLIMIT"
)

// All allowed values of TradeOrderType enum
var AllowedTradeOrderTypeEnumValues = []TradeOrderType{
	"LIMIT",
	"MARKET",
	"STOPLOSS",
	"STOPLIMIT",
}

func (v *TradeOrderType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TradeOrderType(value)
	for _, existing := range AllowedTradeOrderTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TradeOrderType", value)
}

// NewTradeOrderTypeFromValue returns a pointer to a valid TradeOrderType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTradeOrderTypeFromValue(v string) (*TradeOrderType, error) {
	ev := TradeOrderType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TradeOrderType: valid values are %v", v, AllowedTradeOrderTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TradeOrderType) IsValid() bool {
	for _, existing := range AllowedTradeOrderTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TradeOrderType value
func (v TradeOrderType) Ptr() *TradeOrderType {
	return &v
}

type NullableTradeOrderType struct {
	value *TradeOrderType
	isSet bool
}

func (v NullableTradeOrderType) Get() *TradeOrderType {
	return v.value
}

func (v *NullableTradeOrderType) Set(val *TradeOrderType) {
	v.value = val
	v.isSet = true
}

func (v NullableTradeOrderType) IsSet() bool {
	return v.isSet
}

func (v *NullableTradeOrderType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTradeOrderType(val *TradeOrderType) *NullableTradeOrderType {
	return &NullableTradeOrderType{value: val, isSet: true}
}

func (v NullableTradeOrderType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTradeOrderType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
