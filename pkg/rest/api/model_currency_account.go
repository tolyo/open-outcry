package api

// CurrencyAccount - Currency account available to user
type CurrencyAccount struct {
	Id string `json:"id,omitempty"`

	// ISO 4217 Currency symbol
	Currency string `json:"currency,omitempty"`

	Amount float64 `json:"amount,omitempty"`

	AmountReserved float64 `json:"amountReserved,omitempty"`

	AmountAvailable float64 `json:"amountAvailable,omitempty"`
}

// AssertCurrencyAccountRequired checks if the required fields are not zero-ed
func AssertCurrencyAccountRequired(obj CurrencyAccount) error {
	return nil
}

// AssertCurrencyAccountConstraints checks if the values respects the defined constraints
func AssertCurrencyAccountConstraints(obj CurrencyAccount) error {
	return nil
}
