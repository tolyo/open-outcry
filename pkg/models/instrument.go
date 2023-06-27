package models

// `instrument.pub_id` db reference
type InstrumentId string

// Ticker-like name of the instrument. For monetary instruments, a currency pair is used.
type InstrumentName string

// The underlying currency of the FX instrument
type InstrumentBaseCurrency CurrencyName

// The default currency for market quotes of the instrument
type InstrumentQuoteCurrency CurrencyName

type Instrument struct {
	Id            InstrumentId
	Name          InstrumentName
	BaseCurrency  InstrumentBaseCurrency
	QuoteCurrency InstrumentQuoteCurrency
	Active        bool
}
