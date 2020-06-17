package core

import (
	"github.com/mysheep/squirrel/types"
	"github.com/mysheep/squirrel/generator"
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

// -------------------------------------------------------------------------------------------------

func Quote(x *types.Cell) *types.Cell {
	return cadr(x)  // (quote a) -> cdr -> (a) -> car -> a
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
			return Err("Can't take car of atom")
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
			return Err("Can't take cdr of atom")
		}
	}
	
}

// -------------------------------------------------------------------------------------------------

func Cons(x, y *types.Cell) *types.Cell {
	return Cons_(x, y)
}

// -------------------------------------------------------------------------------------------------

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

func Cons_(x, y *types.Cell) *types.Cell {
	return generator.Cons(x, y)
}

func Nil_() *types.Cell {
	return generator.Nil()
}

func Sym_(s string) *types.Cell {
	return generator.Sym(s)
}

func Err(s string, a ...interface{}) *types.Cell {
	return generator.Error(s, a)
}

func Type(c *types.Cell) *types.Cell {
	return Sym_(c.Type_())
}

func Tag(c *types.Cell, t string) *types.Cell {
	return generator.Tag(c, t)
}

// -------------------------------------------------------------------------------------------------