package generator

import (
	"github.com/shopspring/decimal"
)

import (
	"github.com/mysheep/squirrel/types"
)

// Num creates am atom of type number from string
func Num(s string) *types.Cell {

	d, err := decimal.NewFromString(s)
	
	if err == nil {
		return Num_(d)
	} else {
		return Err(err.Error())
	}
}

// num creates am atom of type number from decimal
func Num_(d decimal.Decimal) *types.Cell {
	return &types.Cell {
		Type: types.Type{ Cell: types.ATOM, Atom: types.NUMBER },
		Val : d,
	}
}