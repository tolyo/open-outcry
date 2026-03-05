package api

// InstrumentAccountHolding struct for InstrumentAccountHolding
type InstrumentAccountHolding struct {
	Name            string  `json:"name"`
	Amount          float64 `json:"amount"`
	AmountReserved  float64 `json:"amountReserved"`
	AmountAvailable float64 `json:"amountAvailable"`
	Value           float64 `json:"value"`
	Currency        string  `json:"currency"`
}
