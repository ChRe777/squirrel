package builtin

import (
	"errors"
	"fmt"
)

import (
	"github.com/mysheep/squirrel/core"
	"github.com/mysheep/squirrel/types"
)

// -------------------------------------------------------------------------------------------------

var (
	OperatorNotFound = errors.New("operator not found")
)

// -------------------------------------------------------------------------------------------------

func Eval(e, a *types.Cell, eval FuncType) (*types.Cell, error)  {

	if c := core.Car(e); c.IsAtom() {
		if fn, found := builtinFuncs[*c]; found {
			return fn(e, a, eval), nil
		}
	}
	
	return nil, OperatorNotFound
}

// -------------------------------------------------------------------------------------------------

type FuncType 	 = func(*types.Cell, *types.Cell) *types.Cell
type MapFuncType = func(*types.Cell, *types.Cell, FuncType) *types.Cell

var builtinFuncs = map[types.Cell] MapFuncType {
	
	// function group 1
	//
	*NO     : No_		,
	*NOT    : Not_		,
	*AND    : And_		,
	*PAIR   : Pair_		,
	*LIST   : List_		,
	*ASSOC  : Assoc_	,
	*APPEND : Append_	,
	
	// function group 2
	//
	*VAR	: Var_		,
	*DEF	: Def_ 		,
	*LET	: Let_		,
	*FUNC	: Func_		,		
	*ENV	: Env_		,
	*MAC	: Mac_		,
	*DO		: Do_		,
	
	// macro
	//
	*BACKQUOTE			: Backquote_		,
//	*UNQUOTE   			: Unquote_ 			,		// used internal
//	*UNQUOTE_SPLICING   : Unquotesplicing_	,		// used internal

}

// -------------------------------------------------------------------------------------------------

func No_(e, a *types.Cell, eval FuncType) *types.Cell {
	x := eval(core.Cadr(e), a)
	fmt.Printf("No - x:%v \n", x)
	return No(x)
}

func Not_(e, a *types.Cell, eval FuncType) *types.Cell {
	x := eval(core.Cadr(e), a)
	return Not(x)
}

func And_(e, a *types.Cell, eval FuncType) *types.Cell {
	x := eval(core.Cadr(e), a)
	y := eval(core.Caddr(e), a)
	return And(x, y)
}

func Append_(e, a *types.Cell, eval FuncType) *types.Cell {
	xs := eval(core.Cadr(e), a)
	ys := eval(core.Caddr(e), a)
	return Append (xs, ys)
}

func Pair_(e, a *types.Cell, eval FuncType) *types.Cell {
	xs := eval(core.Cadr(e), a)
	ys := eval(core.Caddr(e), a)
	return Pair (xs, ys)
}

func List_(e, a *types.Cell, eval FuncType) *types.Cell {
	xs := core.Cdr(e)
	return List(xs, a, eval)
}

func Assoc_(e, a *types.Cell, eval FuncType) *types.Cell {
	x  := eval(core.Cadr(e), a)
	ys := eval(core.Caddr(e), a)	
	return Assoc(x, ys)
}

// -------------------------------------------------------------------------------------------------

func Def_(e, a *types.Cell, eval FuncType) *types.Cell {
	return Def(e, a, eval)
}

func Var_(e, a *types.Cell, eval FuncType) *types.Cell {
	return Var(e, a, eval)
}

func Let_(e, a *types.Cell, eval FuncType) *types.Cell {
	return Let(e, a, eval)
}

func Env_(e, a *types.Cell, eval FuncType) *types.Cell {
	return Env(e, a)
}

func Func_(e, a *types.Cell, eval FuncType) *types.Cell {
	return Fun(e, a)
}

func Mac_(e, a *types.Cell, eval FuncType) *types.Cell {
	return Mac(e, a, eval)
}

func Do_(e, a *types.Cell, eval FuncType) *types.Cell {
	return Do(e, a, eval)
}

// -------------------------------------------------------------------------------------------------

func Backquote_(e, a *types.Cell, eval FuncType) *types.Cell {
	return Backquote(e, a, eval)
}

// -------------------------------------------------------------------------------------------------
