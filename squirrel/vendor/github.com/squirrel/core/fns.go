package core

import (
	"github.com/squirrel/types"
	"github.com/squirrel/generator"
)

/*

	7 primitive core operators:

	1. Quote
	2. Atom
	3. Eq
	4. Car
	5. Cdr
	6. Cons
	7. Cond

*/


func Quote(x *types.Cell) *types.Cell {
	return cadr(x)  // (quote a) -> cdr -> (a) -> car -> a
}

func Atom(x *types.Cell) *types.Cell {
	if x.IsAtom() {
		return T
	} else {
		return NIL
	}
}

func Is(x, y *types.Cell) *types.Cell {	
	if x.Equal(y) {
	 	return T	
	}
	return NIL 	// FALSE
}

func Car(e *types.Cell) *types.Cell {
	if e.Equal(NIL) {
		return NIL
	} else {
		if e.IsCons() {
			return car_(e) 
		} else {
			return Err("Can't take car of %v", e)
		}
	}
}

func Cdr(e *types.Cell) *types.Cell {
	
	if e.Equal(NIL) {
		return NIL
	} else {
		if e.IsCons() {
			return cdr_(e)
		} else {
			return Err("Can't take cdr of %v", e)
		}
	}
	
}

func Cons(x, y *types.Cell) *types.Cell {
	return Cons_(x, y)
}

func Cond(x *types.Cell) *types.Cell {

	if x.IsCons() {
		if caar(x).Equal(T) {
			return cadar(x)
		} else {
			return Cond(cdr_(x))
		}
	} else {
		return Err("x must be a list of form ((p1 e1) (p2 e2) .. (pn en))")
	}
	
}

// END 7 primitive core operators

// -------------------------------------------------------------------------------------------------
// Helpers
// -------------------------------------------------------------------------------------------------

func car_(c *types.Cell) *types.Cell {
	if c.IsCons() {
		return c.Car
	}
	return NIL
}

func cdr_(c *types.Cell) *types.Cell {
	if c.IsCons() {
		return c.Cdr
	}
	return NIL
}

// -------------------------------------------------------------------------------------------------
// Shortcuts
// -------------------------------------------------------------------------------------------------

func caar  (e *types.Cell) *types.Cell { return car_(car_(e))             }
func cadr  (e *types.Cell) *types.Cell { return car_(cdr_(e))             }
func cddr  (e *types.Cell) *types.Cell { return cdr_(cdr_(e))             }
func cadar (e *types.Cell) *types.Cell { return car_(cdr_(car_(e)))       } 
func cdddr (e *types.Cell) *types.Cell { return cdr_(cdr_(cdr_(e)))       } 
func caddr (e *types.Cell) *types.Cell { return car_(cdr_(cdr_(e)))       }
func caddar(e *types.Cell) *types.Cell { return car_(cdr_(cdr_(car_(e)))) }
func cadddr(e *types.Cell) *types.Cell { return car_(cdr_(cdr_(cdr_(e)))) } 	


// -------------------------------------------------------------------------------------------------
//  e.g. 
//		(1 2)
//
//      last
//       v
//	xs->NIL

//      last
//       v
//	xs->   ->NIL
//
//      last
//		 v
//  xs->[1]->NIL
//
//      last
//		 v
//  xs->[1]->	->NIL
//
//			 last
//            v	
//	xs->[1]->[2]->NIL
//
func Push(xs, x, last *types.Cell) (*types.Cell, *types.Cell) {

	new_ := generator.Cons(x, NIL)
	
	if last.Equal(NIL) {
		xs = new_
	} else {
		last.Cdr = new_
	}
	
	last = new_

	return xs, last

}