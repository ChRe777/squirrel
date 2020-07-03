package functions

import (
	"github.com/shopspring/decimal"
)

import (
//	"github.com/mysheep/squirrel/builtin"
	"github.com/mysheep/squirrel/generator"
	"github.com/mysheep/squirrel/types"
)

func Add(x, y *types.Cell) *types.Cell {
	if x.IsNumber() && y.IsNumber() {
		x_, _ := x.Val.(decimal.Decimal)
		y_, _ := y.Val.(decimal.Decimal)
		z_ := x_.Add(y_)
		return generator.Num_(z_)
	} else {
		return generator.Err("both cells must be of type number")
	}
}

func Sub(x, y *types.Cell) *types.Cell {
	if x.IsNumber() && y.IsNumber() {
		x_, _ := x.Val.(decimal.Decimal)
		y_, _ := y.Val.(decimal.Decimal)
		z_ := x_.Sub(y_)
		return generator.Num_(z_)
	} else {
		return generator.Err("both cells must be of type number")
	}
}