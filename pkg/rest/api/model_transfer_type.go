package api

import (
	"fmt"
)

type TransferType string

const (
	DEPOSIT         TransferType = "DEPOSIT"
	WITHDRAWAL      TransferType = "WITHDRAWAL"
	TRANSFER        TransferType = "TRANSFER"
	INSTRUMENT_BUY  TransferType = "INSTRUMENT_BUY"
	INSTRUMENT_SELL TransferType = "INSTRUMENT_SELL"
	CHARGE          TransferType = "CHARGE"
)

var AllowedTransferTypeEnumValues = []TransferType{
	"DEPOSIT",
	"WITHDRAWAL",
	"TRANSFER",
	"INSTRUMENT_BUY",
	"INSTRUMENT_SELL",
	"CHARGE",
}

var validTransferTypeEnumValues = map[TransferType]struct{}{
	"DEPOSIT":         {},
	"WITHDRAWAL":      {},
	"TRANSFER":        {},
	"INSTRUMENT_BUY":  {},
	"INSTRUMENT_SELL": {},
	"CHARGE":          {},
}

func (v TransferType) IsValid() bool {
	_, ok := validTransferTypeEnumValues[v]
	return ok
}

func NewTransferTypeFromValue(v string) (TransferType, error) {
	ev := TransferType(v)
	if ev.IsValid() {
		return ev, nil
	}
	return "", fmt.Errorf("invalid value '%v' for TransferType: valid values are %v", v, AllowedTransferTypeEnumValues)
}

func AssertTransferTypeRequired(obj TransferType) error {
	return nil
}

func AssertTransferTypeConstraints(obj TransferType) error {
	return nil
}
