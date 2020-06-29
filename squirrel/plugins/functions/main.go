package main

import (
	"errors"
)

import (
	"github.com/mysheep/squirrel/core"
	"github.com/mysheep/squirrel/types"
	"github.com/mysheep/squirrel/plugins/functions/functions"
)

// -------------------------------------------------------------------------------------------------

var (
	OperatorNotFound = errors.New("operator not found")
)

// -------------------------------------------------------------------------------------------------

type FuncType 	 = func(*types.Cell, *types.Cell) *types.Cell
type MapFuncType = func(*types.Cell, *types.Cell, FuncType) *types.Cell

// -------------------------------------------------------------------------------------------------

type any string 		// Could be any type

var Evaluator any 		// Name is used to detect plugin type

// -------------------------------------------------------------------------------------------------

func (p any) Eval(exp, env *types.Cell, eval FuncType) (*types.Cell, error)  {

	if c := core.Car(exp); c.IsAtom() {
		if fn, found := functionMap[*c]; found {
			return fn(exp, env, eval), nil
		}
	}
	
	return nil, OperatorNotFound
}

// -------------------------------------------------------------------------------------------------

var functionMap = map[types.Cell] MapFuncType {
	
	*functions.MAP     		: Map_		,
	// ...

}

// TODO:

//    mapFn    = "(map	(func (fn xs)  (cond ((no xs) nil) ('t (cons (fn (car xs)) (map fn (cdr xs)))))))"

// -------------------------------------------------------------------------------------------------

func Map_(exp, env *types.Cell, eval FuncType) *types.Cell {
	fn := eval(core.Cadr (exp), env)
	xs := eval(core.Caddr(exp), env)
	return functions.Map(fn, xs, env, eval)
}

