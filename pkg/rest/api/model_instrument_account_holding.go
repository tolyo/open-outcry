package api

type InstrumentAccountHolding struct {
	Name string `json:"name"`

	Amount float64 `json:"amount"`

	AmountReserved float64 `json:"amountReserved"`

	AmountAvailable float64 `json:"amountAvailable"`

	Value float64 `json:"value"`

	// ISO 4217 Currency symbol
	Currency string `json:"currency"`
}

// AssertInstrumentAccountHoldingRequired checks if the required fields are not zero-ed
func AssertInstrumentAccountHoldingRequired(obj InstrumentAccountHolding) error {
	elements := map[string]interface{}{
		"name":            obj.Name,
		"amount":          obj.Amount,
		"amountReserved":  obj.AmountReserved,
		"amountAvailable": obj.AmountAvailable,
		"value":           obj.Value,
		"currency":        obj.Currency,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertInstrumentAccountHoldingConstraints checks if the values respects the defined constraints
func AssertInstrumentAccountHoldingConstraints(obj InstrumentAccountHolding) error {
	return nil
}
