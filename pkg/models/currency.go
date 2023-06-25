package models

type CurrencyName string

type CurrencyPrecision int

type Currency struct {
	Name      CurrencyName
	Precision CurrencyPrecision
}
