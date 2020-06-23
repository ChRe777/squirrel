package main

import (
	"errors"
)

import (
	"github.com/mysheep/squirrel/core"
	"github.com/mysheep/squirrel/types"
)

// -------------------------------------------------------------------------------------------------

type any string // could any type

var Evaluator any // Name is used to detect plugin type

// -------------------------------------------------------------------------------------------------

func (p any) Eval(e, a *types.Cell) (*types.Cell, error) {

	if c := core.Car(e); c.IsAtom() {
		if op, found := builtFuncs[*c]; found {
			return op(e, a), nil
		}
	}

	return nil, errors.New("Operator not found")
}

// -------------------------------------------------------------------------------------------------

var builtFuncs = map[types.Cell] func(e, a *types.Cell) *types.Cell {
	
	*MAP     		: Map_		,
	// ...

}

// TODO:

//    mapFn    = "(map	(func (fn xs)  (cond ((no xs) nil) ('t (cons (fn (car xs)) (map fn (cdr xs)))))))"

// -------------------------------------------------------------------------------------------------

func Map_ (e, a *types.Cell) *types.Cell {
	fn := evaluator.Eval(builtin.Cadr (e), a)
	xs := evaluator.Eval(builtin.Caddr(e), a)
	return Map(fn, xs)
}

