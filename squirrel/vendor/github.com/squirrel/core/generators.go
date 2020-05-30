package core

import (
	"github.com/squirrel/types"
	"github.com/squirrel/generator"
)

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

// Cons_ create a cons cell from two cells
func Cons_(x, y *types.Cell) *types.Cell {
	return generator.Cons(x, y)
}

// Quote_ create a quoted cell
func Nil_() *types.Cell {
	return generator.Nil()
}

// Quote_ create a quoted cell
func Quote_(c *types.Cell) *types.Cell {
	return generator.Cons(QUOTE, generator.Cons(c, NIL))
}

// Backquote is like quote for macros but in combination
// with unquote
func Backquote_(c *types.Cell) *types.Cell {
	return generator.Cons(BACKQUOTE, generator.Cons(c, NIL))
}

// Unquote is used in macros to enable evaluation
func Unquote_(c *types.Cell) *types.Cell {
	return generator.Cons(UNQUOTE, generator.Cons(c, NIL))
}