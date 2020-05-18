package generator

import(
	"fmt"
)

import (
	"github.com/squirrel/types"
)

// Error creates an error from error string
func Error(s string) *types.Cell {
	val := fmt.Sprintf("Error: \"%s\"", s)
	return &types.Cell {
		Type: types.Type{Cell: types.ATOM, Atom: types.ERROR},
		Val : val,
	}
}