package models

// `instrument.pub_id` db reference
type InstrumentId string

// Ticker-like name of the instrument. For monetary instruments, a currency pair is used.
type InstrumentName string

// The underlying currency of the FX instrument
type InstumentBaseCurrency CurrencyName

// The default currency for market quotes of the instrument
type InstrumentQuoteCurrency CurrencyName

type Instrument struct {
	Id            InstrumentId
	Name          InstrumentName
	BaseCurrency  InstrumentBaseCurrency
	QuoteCurrency InstrumentQuoteCurrency
	Active        boolean
}

func CreateInstrument(name InstrumentName, quote_currency CurrencyName) InstrumentId {
	db.QueryVal(`
      INSERT INTO instrument(
        name,
        quote_currency
      ) VALUES (
        $1,
        $2
      )
      RETURNING pub_id;
    `, name, quote_currency)
}

func CreateFxInstrument(name InstrumentName, base_currency CurrencyName, quote_currency CurrencyName) InstrumentId {
	db.QueryVal(`
      INSERT INTO instrument(
        name,
        base_currency,
        quote_currency,
        fx_instrument
      ) VALUES (
        $1,
        $2,
        $3,
        TRUE
      )
      RETURNING pub_id;
    `, name, base_currency, quote_currency)
}
