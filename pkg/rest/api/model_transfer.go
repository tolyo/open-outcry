package api

type TransferEntry struct {
	Id string `json:"id"`

	Type TransferType `json:"type"`

	Amount float64 `json:"amount"`

	// ISO 4217 Currency symbol
	Currency string `json:"currency"`

	SenderAccountId string `json:"senderAccountId"`

	BeneficiaryAccountId string `json:"beneficiaryAccountId"`

	Details string `json:"details"`

	ExternalReferenceNumber string `json:"externalReferenceNumber"`

	Status string `json:"status"`

	DebitBalanceAmount float64 `json:"debitBalanceAmount,omitempty"`

	CreditBalanceAmount float64 `json:"creditBalanceAmount,omitempty"`
}

// AssertTransferEntryRequired checks if the required fields are not zero-ed
func AssertTransferEntryRequired(obj TransferEntry) error {
	elements := map[string]interface{}{
		"id":                      obj.Id,
		"type":                    obj.Type,
		"amount":                  obj.Amount,
		"currency":                obj.Currency,
		"senderAccountId":         obj.SenderAccountId,
		"beneficiaryAccountId":    obj.BeneficiaryAccountId,
		"details":                 obj.Details,
		"externalReferenceNumber": obj.ExternalReferenceNumber,
		"status":                  obj.Status,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertTransferEntryConstraints checks if the values respects the defined constraints
func AssertTransferEntryConstraints(obj TransferEntry) error {
	return nil
}
