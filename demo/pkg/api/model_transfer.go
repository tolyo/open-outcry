package api

// TransferEntry struct for TransferEntry
type TransferEntry struct {
	Id                      string       `json:"id"`
	Type                    TransferType `json:"type"`
	Amount                  float64      `json:"amount"`
	Currency                string       `json:"currency"`
	SenderAccountId         string       `json:"senderAccountId"`
	BeneficiaryAccountId    string       `json:"beneficiaryAccountId"`
	Details                 string       `json:"details"`
	ExternalReferenceNumber string       `json:"externalReferenceNumber"`
	Status                  string       `json:"status"`
	DebitBalanceAmount      *float64     `json:"debitBalanceAmount,omitempty"`
	CreditBalanceAmount     *float64     `json:"creditBalanceAmount,omitempty"`
}
