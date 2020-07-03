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
	
	*functions.ADD     	: Add_		,
	*functions.SUB     	: Sub_		,

}

// -------------------------------------------------------------------------------------------------

func Add_(exp, env *types.Cell, eval FuncType) *types.Cell {
	x := eval(core.Cadr (exp), env)
	y := eval(core.Caddr(exp), env)
	return functions.Add(x, y)
}

func Sub_(exp, env *types.Cell, eval FuncType) *types.Cell {
	x := eval(core.Cadr (exp), env)
	y := eval(core.Caddr(exp), env)
	return functions.Sub(x, y)
}

