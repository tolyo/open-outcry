package models

import (
	"open-outcry/pkg/db"

	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
)

type PaymentJournalId string

type PaymentType string

const (
	Deposit        PaymentType = "DEPOSIT"
	Withdrawal     PaymentType = "WITHDRAWAL"
	Transfer       PaymentType = "TRANSFER"
	InstrumentBuy  PaymentType = "INSTRUMENT_BUY"
	InstrumentSell PaymentType = "INSTRUMENT_SELL"
	Charge         PaymentType = "CHARGE"
)

type PaymentAmount decimal.Decimal

type PaymentDetails string

type PaymentExternalReferenceNumber string

// Payment represents a row from the payment table with joined pub_ids for accounts.
type Payment struct {
	Id                      string
	Type                    PaymentType
	Amount                  float64
	Currency                CurrencyName
	SenderAccountId         PaymentAccountId
	BeneficiaryAccountId    PaymentAccountId
	Details                 PaymentDetails
	ExternalReferenceNumber PaymentExternalReferenceNumber
	Status                  string
	DebitBalanceAmount      float64
	CreditBalanceAmount     float64
}

const paymentBaseQuery = `
	SELECT
	  p.pub_id,
	  p.type::text,
	  p.amount,
	  p.currency_name,
	  spa.pub_id,
	  bpa.pub_id,
	  p.details,
	  COALESCE(p.external_reference_number, ''),
	  p.status,
	  p.debit_balance_amount,
	  p.credit_balance_amount
	FROM payment p
	INNER JOIN payment_account spa ON p.sender_payment_account_id = spa.id
	INNER JOIN payment_account bpa ON p.beneficiary_payment_account_id = bpa.id
`

func GetPayment(id string) *Payment {
	var p Payment
	err := db.Instance().QueryRow(paymentBaseQuery+`WHERE p.pub_id = $1`, id).Scan(
		&p.Id, &p.Type, &p.Amount, &p.Currency,
		&p.SenderAccountId, &p.BeneficiaryAccountId,
		&p.Details, &p.ExternalReferenceNumber,
		&p.Status, &p.DebitBalanceAmount, &p.CreditBalanceAmount,
	)
	if err != nil {
		log.Error(err)
		return nil
	}
	return &p
}

func GetPaymentsByAppEntity(appEntityId AppEntityId) []Payment {
	query := paymentBaseQuery + `
		WHERE spa.app_entity_id = (SELECT id FROM app_entity WHERE pub_id = $1)
		   OR bpa.app_entity_id = (SELECT id FROM app_entity WHERE pub_id = $1)
		ORDER BY p.created_at DESC
	`
	rows, err := db.Instance().Query(query, appEntityId)
	if err != nil {
		log.Error(err)
		return nil
	}
	defer rows.Close()

	var payments []Payment
	for rows.Next() {
		var p Payment
		err := rows.Scan(
			&p.Id, &p.Type, &p.Amount, &p.Currency,
			&p.SenderAccountId, &p.BeneficiaryAccountId,
			&p.Details, &p.ExternalReferenceNumber,
			&p.Status, &p.DebitBalanceAmount, &p.CreditBalanceAmount,
		)
		if err != nil {
			log.Error(err)
			continue
		}
		payments = append(payments, p)
	}
	return payments
}

// PaymentJournal represents a double-entry journal for monetary transfers.
// Each journal groups exactly two PaymentJournalEntry records that must balance.
type PaymentJournal struct {
	Id                      string
	Type                    PaymentType
	Currency                CurrencyName
	Details                 PaymentDetails
	ExternalReferenceNumber PaymentExternalReferenceNumber
	Status                  string
	CreatedAt               string
}

// PaymentJournalEntry represents one side of a double-entry monetary transfer.
// Positive amount = debit (increase), negative amount = credit (decrease).
type PaymentJournalEntry struct {
	Id               string
	JournalId        string
	PaymentAccountId PaymentAccountId
	Currency         CurrencyName
	Amount           decimal.Decimal
	ResultingBalance decimal.Decimal
	CreatedAt        string
}
