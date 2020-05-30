package core

import (
	"github.com/squirrel/types"
	"github.com/squirrel/generator"
)

// Quote evaluates to himself
func quote(c *types.Cell) *types.Cell {
	return generator.Cons(QUOTE, generator.Cons(c, NIL))
}

// Car gets first cell of cons cell
func car(c *types.Cell) *types.Cell {
	if c.IsCons() {
		return c.Car
	}
	return NIL
}

// Cdr gets second cell of cons cell
func cdr(c *types.Cell) *types.Cell {
	if c.IsCons() {
		return c.Cdr
	}
	return NIL
}

// Cons create a cons from two cells
func cons(x, y *types.Cell) *types.Cell {
	return generator.Cons(x, y)
}

// -------------------------------------------------------------------------------------------------


// Backquote is like quote for macros but in combination
// with unquote
func Backquote(c *types.Cell) *types.Cell {
	return generator.Cons(BACKQUOTE, generator.Cons(c, NIL))
}

// Unquote is used in macros to enable evaluation
func Unquote(c *types.Cell) *types.Cell {
	return generator.Cons(UNQUOTE, generator.Cons(c, NIL))
}

// -------------------------------------------------------------------------------------------------

// Sym create a symbol from string
func Sym(s string) *types.Cell {
	return generator.Sym(s)
}

// Tag tags a cell with string t
func Tag(c *types.Cell, t string) *types.Cell {
	return generator.Tag(c, t)
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


// -------------------------------------------------------------------------------------------------


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

// l -->[]-->[]-->nil
// l -->[]-->[]     -->nil
// l -->[]-->[]-->[]-->nil
//
// TODO: Rename - maybe JavaScript - push
//
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