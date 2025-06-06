package functions

import (
	"github.com/shopspring/decimal"
)

import (
	"github.com/mysheep/squirrel/core"
	"github.com/mysheep/squirrel/generator"
	"github.com/mysheep/squirrel/types"
)

const (
	MUST_BE_NUMBER_TYPES = "x,y must be of type number"
)

// AddList adds all items in list
func AddList(xs, env *types.Cell, eval func(*types.Cell, *types.Cell) *types.Cell) *types.Cell {

	if xs.Equal(core.NIL) {
		zero_ := decimal.NewFromInt(0)
		return generator.Num_(zero_)
	}
	
	x := eval(core.Car(xs), env)
	if x.IsErr() {
		return x
	}
	
	return add(x, AddList(core.Cdr(xs), env, eval))
}

// SubList subtracts all others items in list from first items
func SubList(xs, env *types.Cell, eval func(*types.Cell, *types.Cell) *types.Cell) *types.Cell {

	if xs.Equal(core.NIL) {
		zero_ := decimal.NewFromInt(0)
		return generator.Num_(zero_)
	}
	
	x := eval(core.Car(xs), env)
	if x.IsErr() {
		return x
	}
	
	return sub(x, AddList(core.Cdr(xs), env, eval))
}

// DivList divides all items 
func DivList(xs, env *types.Cell, eval func(*types.Cell, *types.Cell) *types.Cell) *types.Cell {

	if xs.Equal(core.NIL) {
		zero_ := decimal.NewFromInt(1)
		return generator.Num_(zero_)
	}
	
	x := eval(core.Car(xs), env)
	if x.IsErr() {
		return x
	}
	
	return div(x, MulList(core.Cdr(xs), env, eval))
}

// MulList multiplies all items 
func MulList(xs, env *types.Cell, eval func(*types.Cell, *types.Cell) *types.Cell) *types.Cell {

	if xs.Equal(core.NIL) {
		zero_ := decimal.NewFromInt(1)
		return generator.Num_(zero_)
	}
	
	x := eval(core.Car(xs), env)
	if x.IsErr() {
		return x
	}
	
	return mul(x, MulList(core.Cdr(xs), env, eval))
}

// -------------------------------------------------------------------------------------------------

func add(x, y *types.Cell) *types.Cell {
	if x.IsNumber() && y.IsNumber() {
		x_, _ := x.Val.(decimal.Decimal)
		y_, _ := y.Val.(decimal.Decimal)
		z_ := x_.Add(y_)
		return generator.Num_(z_)
	} else {
		return generator.Err(MUST_BE_NUMBER_TYPES)
	}
}

func sub(x, y *types.Cell) *types.Cell {
	if x.IsNumber() && y.IsNumber() {
		x_, _ := x.Val.(decimal.Decimal)
		y_, _ := y.Val.(decimal.Decimal)
		z_ := x_.Sub(y_)
		return generator.Num_(z_)
	} else {
		return generator.Err(MUST_BE_NUMBER_TYPES)
	}
}

func div(x, y *types.Cell) *types.Cell {
	if x.IsNumber() && y.IsNumber() {
		x_, _ := x.Val.(decimal.Decimal)
		y_, _ := y.Val.(decimal.Decimal)
		z_ := x_.Div(y_)
		return generator.Num_(z_)
	} else {
		return generator.Err("x,y must be of type number")
	}
}

func mul(x, y *types.Cell) *types.Cell {
	if x.IsNumber() && y.IsNumber() {
		x_, _ := x.Val.(decimal.Decimal)
		y_, _ := y.Val.(decimal.Decimal)
		z_ := x_.Mul(y_)
		return generator.Num_(z_)
	} else {
		return generator.Err("x,y must be of type number")
	}
}