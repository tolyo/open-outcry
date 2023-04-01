package models

import (
	"encoding/csv"
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
	"io"
	"open-outcry/pkg/db"
	"strconv"
	"strings"
)

type Fee struct {
	Type       string
	Currency   CurrencyName
	Min        decimal.Decimal
	Max        decimal.Decimal
	Percentage int
}

func LoadFees(fees string) {
	var feeList []Fee

	r := csv.NewReader(strings.NewReader(fees))
	record, err := r.ReadAll()
	// Stop at EOF.
	if err != io.EOF && err != nil {
		panic(err)
	}

	for i, line := range record {
		if i > 0 { // omit header line
			var rec Fee
			for j, field := range line {
				if j == 0 {
					rec.Type = string(field)
				} else if j == 1 {
					rec.Currency = CurrencyName(field)
				} else if j == 2 {
					min, err := decimal.NewFromString(field)
					if err == nil {
						rec.Min = min
					}
				} else if j == 3 {
					max, err := decimal.NewFromString(field)
					if err == nil {
						rec.Max = max
					}
				} else if j == 4 {
					percentage, err := strconv.Atoi(field)
					if err == nil {
						rec.Percentage = percentage
					}
				}
			}
			feeList = append(feeList, rec)
		}
	}

	for _, val := range feeList {
		CreateOrUpdateFee(val)
	}

}

func CreateOrUpdateFee(fee Fee) {
	_, err := db.Instance().Exec(`
		INSERT INTO fee(type, currency_name, min, max, percentage)  
		VALUES ($1, $2, $3, $4, $5) 
		ON CONFLICT (type, currency_name) 
		DO UPDATE SET min = $3, max = $4, percentage = $5;
	`, fee.Type, fee.Currency, fee.Min, fee.Max, fee.Percentage)

	if err != nil {
		log.Fatal(err)
	}
}
