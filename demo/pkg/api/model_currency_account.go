package api

// CurrencyAccount Currency account available to user
type CurrencyAccount struct {
	Id              string  `json:"id,omitempty"`
	Currency        string  `json:"currency,omitempty"`
	Amount          float64 `json:"amount,omitempty"`
	AmountReserved  float64 `json:"amountReserved,omitempty"`
	AmountAvailable float64 `json:"amountAvailable,omitempty"`
}
