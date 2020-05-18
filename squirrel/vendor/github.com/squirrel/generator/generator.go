package generator

import (
	"github.com/squirrel/types"
)

/*
PROCEDURE atom*(name: identifier): cell;
	VAR ac: atomCell;
BEGIN
	NEW(ac); ac.type := ATOM; ac.name := name;
	RETURN ac;
END atom;
*/
func Atom(val interface{}, at types.AtomType) *types.Cell {
	return &types.Cell {
		Type: types.Type{Cell: types.ATOM, Atom: at},
		Val : val,
	}
}

/*
PROCEDURE cons*(car: cell; cdr: cell): cell;
	VAR cc: consCell;
BEGIN
	NEW(cc); cc.type := CONS; cc.car := car; cc.cdr := cdr;
	RETURN cc;
END cons;
*/
func Cons(car, cdr *types.Cell) *types.Cell {	
	return &types.Cell {
		Type: types.Type{Cell: types.CONS},
		Car : car,
		Cdr : cdr,
	}
}