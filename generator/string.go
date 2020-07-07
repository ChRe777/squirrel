package generator

import (
	"github.com/mysheep/squirrel/types"
)

// Sym creates a string from string
func Str(s string) *types.Cell {
	return &types.Cell {
		Type: types.Type { Cell: types.ATOM, Atom: types.STRING },
		Val: s,
	}
}