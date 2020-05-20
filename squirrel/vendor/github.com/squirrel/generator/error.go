package generator

import (
	"fmt"
)

import (
	"github.com/squirrel/types"
)

// Error creates an error from error string
func Error(s string, a ...interface{}) *types.Cell {
	m := fmt.Sprintf(s, a...)
	val := fmt.Sprintf("Error: \"%s\"", m)
	return &types.Cell {
		Type: types.Type{Cell: types.ATOM, Atom: types.ERROR},
		Val : val,
	}
}