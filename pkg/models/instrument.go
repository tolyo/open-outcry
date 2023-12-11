package models

import "open-outcry/pkg/db"

// InstrumentId `instrument.pub_id` db reference
type InstrumentId string

// InstrumentName Ticker-like name of the instrument. For monetary instruments, a currency pair is used.
type InstrumentName string

// InstrumentBaseCurrency The underlying currency of the FX instrument
type InstrumentBaseCurrency CurrencyName

// InstrumentQuoteCurrency The default currency for market quotes of the instrument
type InstrumentQuoteCurrency CurrencyName

type Instrument struct {
	Id            InstrumentId `db:"pub_id"`
	Name          InstrumentName
	BaseCurrency  InstrumentBaseCurrency  `db:"base_currency"`
	QuoteCurrency InstrumentQuoteCurrency `db:"quote_currency"`
	Active        bool
}

func GetInstruments() []Instrument {
	return db.QueryList[Instrument](`SELECT pub_id, name, quote_currency FROM instrument WHERE fx_instrument = FALSE`)
}

func GetFxInstruments() []Instrument {
	return db.QueryList[Instrument](`SELECT pub_id, name, base_currency, quote_currency FROM instrument WHERE fx_instrument = TRUE`)
}
