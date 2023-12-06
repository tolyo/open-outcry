package models

import "open-outcry/pkg/db"

type CurrencyName string

type CurrencyPrecision int

type Currency struct {
	Name      CurrencyName
	Precision CurrencyPrecision
}

// GetCurrencies return
func GetCurrencies() []Currency {
	res := db.QueryList[Currency](`SELECT * FROM currency`)
	return res
}
