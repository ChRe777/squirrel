package builtin

import(
//	"fmt"
)

import (
	"github.com/mysheep/squirrel/types"
//	"github.com/mysheep/squirrel/evaluator/"
	"github.com/mysheep/squirrel/evaluator/core"		// builtin layer based on core layer
)
/*

	Funcs:

		Pair
		No
		Not
		And
		Append
		//List
		Assoc
		
*/

// -------------------------------------------------------------------------------------------------

func Pair (xs, ys *types.Cell) *types.Cell {
	
	if xs.Equal(core.NIL) || 
	   ys.Equal(core.NIL) {
		return core.NIL
	} 

	if xs.IsCons() && ys.IsCons() {		// (x y z) (1 2 3)
	
		x := core.Car(xs)
		y := core.Car(ys)
	
		ws := core.Cdr(xs)
		zs := core.Cdr(ys)
		
		a := List_(x, y)
		b := Pair(ws, zs)
		
		return core.Cons(a, b)
	
	} else {							// (x y . z) (1 2 3 4)
		return core.Cons(List_(xs, ys), core.NIL)
	}
	
}

// -------------------------------------------------------------------------------------------------

func No (x *types.Cell) *types.Cell { // call "no" instead of "null"
	if x.Equal(core.NIL) {
		return core.T
	}
	return core.NIL
}

// -------------------------------------------------------------------------------------------------

func Not (x *types.Cell) *types.Cell {
	if x.Equal(core.T) {
		return core.NIL
	} else {
		return core.T
	}
}

// -------------------------------------------------------------------------------------------------

func And (x, y *types.Cell) *types.Cell {
	if x.Equal(core.T) && y.Equal(core.T) {
		return core.T
	} else {
		return core.NIL
	}
}

// -------------------------------------------------------------------------------------------------

func Append (x, y *types.Cell) *types.Cell {
	if x.Equal(core.NIL) {
		return y
	} else {
		return core.Cons(core.Car(x), Append(core.Cdr(x), y))
	}
}

// -------------------------------------------------------------------------------------------------

func List_ (x, y *types.Cell) *types.Cell {
	return core.Cons(x, core.Cons (y, core.NIL))
}

// -------------------------------------------------------------------------------------------------

/*
func List (xs, a *types.Cell) *types.Cell {

	if xs.Equal(core.NIL) {
		return core.NIL
	} else {
		y  := evaluator.Eval(car(xs), a)
		ys := cdr(xs)
		return core.Cons(y, List(ys, a))
	}
}
*/

// -------------------------------------------------------------------------------------------------

func Assoc (x, ys *types.Cell) *types.Cell {
	if ys.Equal(core.NIL) {
		return core.Err_("Not found")
	} else {
		if x.Equal(core.Caar(ys)) {
			return core.Cadar(ys)
		} else {
			return Assoc(x, core.Cdr(ys))	
		}
	}
}

