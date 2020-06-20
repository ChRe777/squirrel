package generator

import (
	"fmt"
)

import (
	"github.com/mysheep/squirrel/types"
)

// Error creates an error from error string
func Err(format string, args ...interface{}) *types.Cell {
	
	val := fmt.Sprintf("Error: \"%s\"", fmt.Sprintf(format, args...))

	return &types.Cell {
		Type: types.Type{Cell: types.ATOM, Atom: types.ERROR},
		Val : val,
		Car : NIL,
		Cdr : NIL,
	}
}
