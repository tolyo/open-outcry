package currency

import "open-outcry/pkg/db"

type CurrencyName string

type CurrencyPrecision int

type Currency struct {
	Name      CurrencyName
	Precision CurrencyPrecision
}

func GetCurrencies() []Currency {
	return db.QueryList[Currency](`SELECT * FROM currency`)
}
