package models

import (
	"github.com/cockroachdb/apd"
)

type PriceLevel struct {
	Price  apd.Decimal
	Volume apd.Decimal
}
