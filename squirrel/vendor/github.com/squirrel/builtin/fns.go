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

func Pair (x, y *types.Cell) *types.Cell {
	if x.Equal(core.NIL) && y.Equal(core.NIL) {
		return core.NIL
	} else {
		if x.IsCons() && y.IsCons() {
			a := List_(car(x), car(y))
			b := Pair(cdr(x), cdr(y))
			return core.Cons(a,b)
		}
	}
	return generator.Error("x and y must be a cons") // TODO: Check
}

func No (x *types.Cell) *types.Cell { // call "no" instead of "null"
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

func And (x, y *types.Cell) *types.Cell {
	if x.Equal(core.T) && y.Equal(core.T) {
		return core.T
	} else {
		return core.NIL
	}
}

func Append (x, y *types.Cell) *types.Cell {
	if x.Equal(core.NIL) {
		return y
	} else {
		return core.Cons(car(x), Append(cdr(x), y))
	}
}

func List_ (x, y *types.Cell) *types.Cell {
	return core.Cons(x, core.Cons (y, core.NIL))
}

func Assoc (x, ys *types.Cell) *types.Cell {
	if ys.Equal(core.NIL) {
		return core.Err("Not found")
	} else {
		if x.Equal(Caar(ys)) {
			return Cadar(ys)
		} else {
			return Assoc(x, cdr(ys))	
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

func car (x *types.Cell) *types.Cell {
	return core.Car(x)
}

func cdr (x *types.Cell) *types.Cell {
	return core.Cdr(x)
}







