package core

import (
	"github.com/mysheep/squirrel/generator"
	"github.com/mysheep/squirrel/types"
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
	
	New core operators:
	
	8. Type
	9. Tag
	
	Other operators:
	
	10. Sym
	11. Err
	12. Num
	13. Str
	14. ...
	
*/

// -------------------------------------------------------------------------------------------------

func Quote(x *types.Cell) *types.Cell {
	return Cadr(x)
}

// -------------------------------------------------------------------------------------------------

func Atom(x *types.Cell) *types.Cell {
	if x.IsAtom() {
		return T
	} else {
		return NIL
	}
}

// -------------------------------------------------------------------------------------------------

func Is(x, y *types.Cell) *types.Cell {	
	if x.Equal(y) {
	 	return T	
	}
	return NIL 	// FALSE
}

// -------------------------------------------------------------------------------------------------

func Car(e *types.Cell) *types.Cell {
	if e.Equal(NIL) {
		return NIL
	} else {
		if e.IsCons() {
			return car_(e) 
		} else {
			return Err_("Can't take car of atom")
		}
	}
}

// -------------------------------------------------------------------------------------------------

func Cdr(e *types.Cell) *types.Cell {
	
	if e.Equal(NIL) {
		return NIL
	} else {
		if e.IsCons() {
			return cdr_(e)
		} else {
			return Err_("Can't take cdr of atom")
		}
	}
	
}

// -------------------------------------------------------------------------------------------------

func Cons(x, y *types.Cell) *types.Cell {
	return generator.Cons(x, y)
}

// -------------------------------------------------------------------------------------------------

// Cond evals cond (= conditions)
// e.g. 
//		> (cond (
//			(nil b) 
//			( 't a))) 	-> a
//
func Cond(c, a *types.Cell, eval func(*types.Cell, *types.Cell) *types.Cell) *types.Cell {

	if c.Equal(NIL) {
		return NIL
	}
		
	y := eval(Caar(c), a)
	if y.IsErr() {
		return y
	}

	if y.Equal(T) { 
		return eval(Cadar(c), a) 
	} else { 
		return Cond(Cdr(c), a, eval)
	}

}

// END 7 primitive core operators

// -------------------------------------------------------------------------------------------------

func Type(c *types.Cell) *types.Cell {
	return Sym_(c.Type_())
}

func Tag(c *types.Cell, t string) *types.Cell {
	return generator.Tag(c, t)
}

// -------------------------------------------------------------------------------------------------

func Nil_() *types.Cell {
	return generator.Nil()
}

func Num_(s string) *types.Cell {
	return generator.Num(s)
}

func Sym_(s string) *types.Cell {
	return generator.Sym(s)
}

func Str_(s string) *types.Cell {
	return generator.Str(s)
}

func Err_(s string, a ...interface{}) *types.Cell {
	return generator.Err(s, a...)  // Don't forget the ... !!
}



// -------------------------------------------------------------------------------------------------

// -------------------------------------------------------------------------------------------------
// Shortcuts for car, cdr functions
// -------------------------------------------------------------------------------------------------

func Caar  (e *types.Cell) *types.Cell { return car_(car_(e))             }
func Cadr  (e *types.Cell) *types.Cell { return car_(cdr_(e))             }
func Cddr  (e *types.Cell) *types.Cell { return cdr_(cdr_(e))             }
func Cadar (e *types.Cell) *types.Cell { return car_(cdr_(car_(e)))       } 
func Cdddr (e *types.Cell) *types.Cell { return cdr_(cdr_(cdr_(e)))       } 
func Caddr (e *types.Cell) *types.Cell { return car_(cdr_(cdr_(e)))       }
func Caddar(e *types.Cell) *types.Cell { return car_(cdr_(cdr_(car_(e)))) }
func Cadddr(e *types.Cell) *types.Cell { return car_(cdr_(cdr_(cdr_(e)))) }

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

func list2_(x, y *types.Cell) *types.Cell {
	return Cons(x, Cons(y, NIL))
}

// -------------------------------------------------------------------------------------------------

