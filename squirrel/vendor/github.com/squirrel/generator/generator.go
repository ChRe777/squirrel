package generator

import (
	"fmt"
)

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



// Sym creates a symbol from string
func Sym(s string) *types.Cell {
	return &types.Cell {
		Type: types.Type{Cell: types.ATOM, Atom: types.SYMBOL},
		Val : s,
	}
}

// Error creates an error from error string
func Error(s string) *types.Cell {
	val := fmt.Sprintf("Error: \"%s\"", s)
	return &types.Cell {
		Type: types.Type{Cell: types.ATOM, Atom: types.ERROR},
		Val : val,
	}
}

// Nil return THE ONLY one nil cell
func Nil() *types.Cell{
	return NIL
}

var (
	NIL = Atom("nil", types.SYMBOL)
)
