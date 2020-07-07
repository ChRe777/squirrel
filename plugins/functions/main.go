package main

import (
	//"fmt"
)

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
	
	*functions.ADD     	: AddList_	,
	*functions.SUB     	: SubList_	,
	*functions.DIV     	: DivList_	,
	*functions.MUL     	: MulList_	,

}

// -------------------------------------------------------------------------------------------------

func AddList_(exp, env *types.Cell, eval FuncType) *types.Cell {
	xs := core.Cdr(exp)		
	return functions.AddList(xs, env, eval)
}

func SubList_(exp, env *types.Cell, eval FuncType) *types.Cell {
	xs := core.Cdr(exp)		
	return functions.SubList(xs, env, eval)
}

func DivList_(exp, env *types.Cell, eval FuncType) *types.Cell {
	xs := core.Cdr(exp)		
	return functions.DivList(xs, env, eval)
}

func MulList_(exp, env *types.Cell, eval FuncType) *types.Cell {
	xs := core.Cdr(exp)		
	return functions.MulList(xs, env, eval)
}

// -------------------------------------------------------------------------------------------------


