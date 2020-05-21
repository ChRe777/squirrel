package builtin

import (
	"github.com/squirrel/types"
	"github.com/squirrel/generator"
)

/*
PROCEDURE quote*(c: cell): cell;
BEGIN
	RETURN cons(atom("QUOTE"), cons(c, EMPTY));
END quote;
*/
func Quote(c *types.Cell) *types.Cell {
	return generator.Cons(QUOTE, generator.Cons(c, NIL))
}


// l -->[]-->[]-->nil
// l -->[]-->[]     -->nil
// l -->[]-->[]-->[]-->nil
// TODO: Rename - maybe JavaScript - push
func Add(l, c *types.Cell)  *types.Cell {
	li := l
	if l.IsCons() && l.NotEqual(NIL) { 
	
		// TODO: Speed Up With Pointer on LAST element
	
		for ;l.Cdr.NotEqual(NIL); {
			l = l.Cdr
		}
	
		l.Cdr = generator.Cons(c, NIL)
	
	} else {
		li = generator.Cons(c, NIL)
	}
	return li
}

// (list 1 2 3 4)
// (cons 1 (cons 2 (cons 3 (cons 4 ()))))
//
// TODO: Make this more effective
// because Add adds cells at the end
// List create a list of a list of cells
func List(xs ...*types.Cell) *types.Cell {
	l := NIL
	for _, x := range xs {
		l = Add(l, x)
	}
	return l
}


// Car gets first cell of cons cell
func Car(c *types.Cell) *types.Cell {
	if c.IsCons() {
		return c.Car
	}
	return NIL
}

// Cdr gets second cell of cons cell
func Cdr(c *types.Cell) *types.Cell {
	if c.IsCons() {
		return c.Cdr
	}
	return NIL
}

// Cons create a cons from two cells
func Cons(x, y *types.Cell) *types.Cell {
	return generator.Cons(x, y)
}

// Sym create a symbol from string
func Sym(s string) *types.Cell {
	return generator.Sym(s)
}

// Num create a number from string
func Num(s string) *types.Cell {
	return generator.Num(s)
}

// Str create a string from string
func Str(s string) *types.Cell {
	return generator.Str(s)
}

// Error create an error from string
func Err(s string, a ...interface{}) *types.Cell {
	return generator.Error(s, a...)
}