package models

type PaymentId string

type PaymentAmount decimal

type PaymentCurrency PaymentAccountCurrency

type PaymentDetails string

type PaymentExternaReferenceNumber string

type Payment struct {
	PubId                   string
	Number                  string
	Type                    string
	Amount                  string
	Currency                string
	SenderAccountId         string
	BeneficiaryAccountId    string
	Details                 string
	ExternalReferenceNumber string
	FeeSender               string
	FeeBeneficiary          string
	Status                  string
	TotalAmount             string
	DebitBalanceAmount      string
	CreditBalanceAmount     string
}
