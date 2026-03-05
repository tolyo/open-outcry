package api

// CurrencyAccountList - List of currency accounts available to user
type CurrencyAccountList struct {
	Data []CurrencyAccount `json:"data,omitempty"`
}

// AssertCurrencyAccountListRequired checks if the required fields are not zero-ed
func AssertCurrencyAccountListRequired(obj CurrencyAccountList) error {
	for _, el := range obj.Data {
		if err := AssertCurrencyAccountRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertCurrencyAccountListConstraints checks if the values respects the defined constraints
func AssertCurrencyAccountListConstraints(obj CurrencyAccountList) error {
	return nil
}
