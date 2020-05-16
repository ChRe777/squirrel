package generator

import (
	"github.com/shopspring/decimal"
)

func NumberFromString(s string) (decimal.Decimal, error) {
	n, err := decimal.NewFromString(s)
	return n, err
}