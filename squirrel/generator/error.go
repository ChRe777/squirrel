package generator

import (
	"fmt"
)

import (
	"github.com/mysheep/squirrel/types"
)

// Error creates an error from error string
func Err(format string, args ...interface{}) *types.Cell {
	
	msg := format
	if len(args) > 0 {
	
		fmt.Printf("generator.Err - args: %v, len(args): %v\n", args, len(args))
		msg = fmt.Sprintf(format, args...)
	}
	
	val := fmt.Sprintf("Error: \"%s\"", msg)

	return &types.Cell {
		Type: types.Type{Cell: types.ATOM, Atom: types.ERROR},
		Val : val,
		Car : NIL,
		Cdr : NIL,
	}
}
