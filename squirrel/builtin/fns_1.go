package builtin

import (
	"github.com/mysheep/squirrel/core" // builtin layer based on core layer
	"github.com/mysheep/squirrel/types"
)

/*

	Functions are:

	- Pair
	- No
	- Not
	- And
	- Append
	- Assoc
		
*/

// -------------------------------------------------------------------------------------------------

func Pair (xs, ys *types.Cell) *types.Cell {
	
	if xs.Equal(core.NIL) || 
	   ys.Equal(core.NIL) {
		return core.NIL
	} 

	if xs.IsCons() && ys.IsCons() {		// (x y z) (1 2 3)
		ws := list__(core.Car(xs), core.Car(ys))
		zs := Pair  (core.Cdr(xs), core.Cdr(ys))
		return core.Cons(ws, zs)
	} else {							// (x y . z) (1 2 3 4)
		return core.Cons(list__(xs, ys), core.NIL)
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

// -------------------------------------------------------------------------------------------------
// Helper functions
//

func list__(x, y *types.Cell) *types.Cell {
	return core.Cons(x, core.Cons (y, core.NIL))
}

// -------------------------------------------------------------------------------------------------

