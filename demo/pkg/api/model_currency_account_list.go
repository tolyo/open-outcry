package api

// CurrencyAccountList List of currency accounts available to user
type CurrencyAccountList struct {
	Data []CurrencyAccount `json:"data,omitempty"`
}
