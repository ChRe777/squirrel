package generator

import (
	"github.com/mysheep/squirrel/types"
)

const (
	ID_NIL = "nil"
)

// Sym creates a symbol from string
func Sym(s string) *types.Cell {
	return &types.Cell {
		Type: types.Type {
			Cell: types.ATOM, 
			Atom: types.SYMBOL,
		},
		Val: s,
	}
}

// Nil return THE-ONLY-ONE nil symbol
func Nil() *types.Cell{
	return NIL
}

// Tag tags a cell with string t
func Tag(c *types.Cell, t string) *types.Cell {
	c.Tag = t
	return c
}

// The only won nil atom in system
var (
	NIL = Atom(ID_NIL, types.SYMBOL)
)
