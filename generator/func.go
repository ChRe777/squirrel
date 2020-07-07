package generator

import (
	"github.com/mysheep/squirrel/types"
)

// Fun creates a symbol of type func from string
func Fun(s string) *types.Cell {
	return &types.Cell {
		Type: types.Type { Cell: types.ATOM, Atom: types.FUNC },
		Val: s,
	}
}