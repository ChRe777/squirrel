package generator

import (
	"github.com/mysheep/squirrel/types"
)

func Atom(val interface{}, at types.AtomType) *types.Cell {
	return &types.Cell {
		Type: types.Type{Cell: types.ATOM, Atom: at},
		Val : val,
	}
}

func Cons(car, cdr *types.Cell) *types.Cell {	
	return &types.Cell {
		Type: types.Type{Cell: types.CONS},
		Car : car,
		Cdr : cdr,
	}
}