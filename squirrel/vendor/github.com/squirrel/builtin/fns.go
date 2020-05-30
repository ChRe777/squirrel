package builtin

import (
	"github.com/squirrel/types"
	"github.com/squirrel/core"
	"github.com/squirrel/generator"
)
/*

	Funcs:

		Pair
		No
		Not
		And
		Append
		List
		Assoc
		
		Caar  
		Cadr  
		Cddr  
		Cadar 
		Cdddr 
		Caddr 
		Caddar
		Cadddr

*/

func Pair(x, y *types.Cell) *types.Cell {
	if x.Equal(core.NIL) && y.Equal(core.NIL) {
		return core.NIL
	} else {
		if x.IsCons() && y.IsCons() {
			a := list(car(x), car(y))
			b := pair(cdr(x), cdr(y))
			return cons(a,b)
		}
	}
	return generator.Error("x and y must be a cons") // TODO: Check
}

func No(x *types.Cell) *types.Cell { // call "no" instead of "null"
	if x.Equal(core.NIL) {
		return core.T
	}
	return core.NIL
}

func Not (x *types.Cell) *types.Cell {
	if x.Equal(core.T) {
		return core.NIL
	} else {
		return core.T
	}
}

func And(x, y *types.Cell) *types.Cell {
	if x.Equal(core.T) && y.Equal(core.T) {
		return core.T
	} else {
		return core.NIL
	}
}

func Append(x, y *types.Cell) *types.Cell {
	if x.Equal(core.NIL) {
		return y
	} else {
		return cons(car(x), append(cdr(x), y))
	}
}

func List_(x, y *types.Cell) *types.Cell {
	return cons(x, cons (y, core.NIL))
}

func Assoc(x, y *types.Cell) *types.Cell {
	if y.Equal(core.NIL) {
		return core.Err("Not found")
	} else {
		if eq(caar(y), x).Equal(core.T) {
			return cadar(y)
		} else {
			return assoc(x, cdr(y))	
		}
	}
}

func Caar  (e *types.Cell) *types.Cell { return car(car(e))           }
func Cadr  (e *types.Cell) *types.Cell { return car(cdr(e))           }
func Cddr  (e *types.Cell) *types.Cell { return cdr(cdr(e))           }
func Cadar (e *types.Cell) *types.Cell { return car(cdr(car(e)))      } 
func Cdddr (e *types.Cell) *types.Cell { return cdr(cdr(cdr(e)))      } 
func Caddr (e *types.Cell) *types.Cell { return car(cdr(cdr(e)))      }
func Caddar(e *types.Cell) *types.Cell { return car(cdr(cdr(car(e)))) }
func Cadddr(e *types.Cell) *types.Cell { return car(cdr(cdr(cdr(e)))) } 	


// ---------------------------------
// Just ALIAS for better readability
// ---------------------------------

func car(x *types.Cell) *types.Cell {
	return core.Car(x)
}

func cdr(x *types.Cell) *types.Cell {
	return core.Cdr(x)
}

// ---------------------------------
// 
// ---------------------------------


// > (set a 1) -> 1
// > a -> 1
func set(k, v *types.Cell, a *types.Cell) *types.Cell {
	// Add key-value-pair (k v) to environment
	a = cons(list(k, v), a)
	return v
}

// addEnv is a special add that adds a new cell at the front of the environment
// but LET the Pointer to first element the SAME !!!
func addEnv(kv *types.Cell, a *types.Cell ) *types.Cell {
	// Hang in new as second
	cdr := a.Cdr; new := cons(kv, cdr); a.Cdr = new
	// Change Val first and second to move new second to front
	val := new.Val; new.Val = a.Val; a.Val = val
	// Change Car first and second to move new seocen to front
	car := new.Car; new.Car = a.Car; a.Car = car
	// So the pointer to a stays the same // Side effects // ToReThink: ?
	return a
}





