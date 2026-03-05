package api

import (
	"encoding/json"
	"fmt"
)

// TransferType the model 'TransferType'
type TransferType string

// List of TransferType
const (
	DEPOSIT         TransferType = "DEPOSIT"
	WITHDRAWAL      TransferType = "WITHDRAWAL"
	TRANSFER        TransferType = "TRANSFER"
	INSTRUMENT_BUY  TransferType = "INSTRUMENT_BUY"
	INSTRUMENT_SELL TransferType = "INSTRUMENT_SELL"
	CHARGE          TransferType = "CHARGE"
)

// All allowed values of TransferType enum
var AllowedTransferTypeEnumValues = []TransferType{
	"DEPOSIT",
	"WITHDRAWAL",
	"TRANSFER",
	"INSTRUMENT_BUY",
	"INSTRUMENT_SELL",
	"CHARGE",
}

func (v *TransferType) UnmarshalJSON(src []byte) error {
	var value string
	if err := json.Unmarshal(src, &value); err != nil {
		return err
	}
	enumTypeValue := TransferType(value)
	for _, existing := range AllowedTransferTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}
	return fmt.Errorf("%+v is not a valid TransferType", value)
}

func NewTransferTypeFromValue(v string) (*TransferType, error) {
	ev := TransferType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TransferType: valid values are %v", v, AllowedTransferTypeEnumValues)
}

func (v TransferType) IsValid() bool {
	for _, existing := range AllowedTransferTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
