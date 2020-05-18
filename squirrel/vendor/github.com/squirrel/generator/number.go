package generator

import (
	"github.com/shopspring/decimal"
)

import (
	"github.com/squirrel/types"
)

func Num(s string) *types.Cell {
	
	d, err := decimal.NewFromString(s)
	
	if err == nil {
		return num(d)
	} else {
		return Error(err.Error())
	}
}

func num(d decimal.Decimal) *types.Cell {
	return &types.Cell {
		Type: types.Type{Cell: types.ATOM, Atom: types.NUMBER},
		Val : d,
	}
}