package evaluator

import (
	"github.com/squirrel/types"
	"github.com/squirrel/builtin"
	"github.com/squirrel/generator"
)

// * pair
// * no
// * not
// * and
// * append
// * list
// * assoc

// * caar
// * cadr
// * cadar
// * caddr
// * caddar

// TODO: below should/can be implemented in LISP itself
// TODO: Put it into autoload environment

func pair(x, y *types.Cell) *types.Cell {
	if x.Equal(builtin.NIL) && y.Equal(builtin.NIL) {
		return builtin.NIL
	} else {
		if x.IsCons() && y.IsCons() {
			a := list(car(x), car(y))
			b := pair(cdr(x), cdr(y))
			return cons(a,b)
		}
	}
	return generator.Error("x and y must be a cons") // TODO: Check
}

func no(x *types.Cell) *types.Cell { // call "no" instead of "null"
	if x.Equal(builtin.NIL) {
		return builtin.T
	}
	return builtin.NIL
}

func not (x *types.Cell) *types.Cell {
	if x.Equal(builtin.T) {
		return builtin.NIL
	} else {
		return builtin.T
	}
}

func and(x, y *types.Cell) *types.Cell {
	if x.Equal(builtin.T) && y.Equal(builtin.T) {
		return builtin.T
	} else {
		return builtin.NIL
	}
}

func append(x, y *types.Cell) *types.Cell {
	if x.Equal(builtin.NIL) {
		return y
	} else {
		return cons(car(x), append(cdr(x), y))
	}
}

func list(x, y *types.Cell) *types.Cell {
	return cons(x, cons (y, builtin.NIL))
}

func assoc(x, y *types.Cell) *types.Cell {
	if y.Equal(builtin.NIL) {
		return builtin.Err("Not found")
	} else {
		if eq(caar(y), x).Equal(builtin.T) {
			return cadar(y)
		} else {
			return assoc(x, cdr(y))	
		}
	}
}

func caar  (e *types.Cell) *types.Cell { return car(car(e))           }
func cadr  (e *types.Cell) *types.Cell { return car(cdr(e))           }
func cddr  (e *types.Cell) *types.Cell { return cdr(cdr(e))           }
func cadar (e *types.Cell) *types.Cell { return car(cdr(car(e)))      } 
func cdddr (e *types.Cell) *types.Cell { return cdr(cdr(cdr(e)))      } 
func caddr (e *types.Cell) *types.Cell { return car(cdr(cdr(e)))      }
func caddar(e *types.Cell) *types.Cell { return car(cdr(cdr(car(e)))) }
func cadddr(e *types.Cell) *types.Cell { return car(cdr(cdr(cdr(e)))) } 	

// > (set a 1) -> 1
// > a -> 1
func set(k, v *types.Cell, a *types.Cell) *types.Cell {
	// Add key-value-pair (k v) to environment
	a = cons(list(k, v), a)
	return v
}

// TODO: CHECK to USE append instead e.g. ((a 1)) append ((b 2) (c 3))

// addEnv add a new cell at the front of the environment
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

