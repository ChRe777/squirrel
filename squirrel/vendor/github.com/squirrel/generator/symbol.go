package generator

import (
	"github.com/squirrel/types"
)

// Sym creates a symbol from string
func Sym(s string) *types.Cell {
	return &types.Cell {
		Type: types.Type{Cell: types.ATOM, Atom: types.SYMBOL},
		Val : s,
	}
}

// Nil return THE ONLY one nil cell
func Nil() *types.Cell{
	return NIL
}

// ToRefactor
var (
	NIL = Atom("nil", types.SYMBOL)
)