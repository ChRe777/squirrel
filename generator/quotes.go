package generator

import (
	"github.com/mysheep/squirrel/types"
)

var (
	// 7 Core Primitives
	//
	ID_QUOTE  = "quote"

	// For Macros
	//
	ID_BACKQUOTE 		= "backquote"
	ID_UNQUOTE   		= "unquote"
	ID_UNQUOTESPLICING  = "unquote_splicing"
)

// Core symbols of language
var (
	QUOTE 				= Sym(ID_QUOTE)
	BACKQUOTE 			= Sym(ID_BACKQUOTE)			// For Macros
	UNQUOTE   			= Sym(ID_UNQUOTE) 			// For Macros,
	UNQUOTE_SPLICING   	= Sym(ID_UNQUOTESPLICING) 	// For Macros,
)

// Quote_ create a quoted cell
func Quote_(c *types.Cell) *types.Cell {
	return Cons(QUOTE, Cons(c, NIL))
}

// Backquote is like quote for macros but in combination
// with unquote
func Backquote_(c *types.Cell) *types.Cell {
	return Cons(BACKQUOTE, Cons(c, NIL))
}

// Unquote is used in macros to enable evaluation
func Unquote_(c *types.Cell) *types.Cell {
	return Cons(UNQUOTE, Cons(c, NIL))
}

// UnquoteSplicing is used in macros to enable evaluation
func UnquoteSplicing_(c *types.Cell) *types.Cell {
	return Cons(UNQUOTE_SPLICING, Cons(c, NIL))
}
