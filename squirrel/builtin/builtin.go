package builtin

import (
	"errors"
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

type FuncType 	 = func(*types.Cell, *types.Cell) *types.Cell
type MapFuncType = func(*types.Cell, *types.Cell, FuncType) *types.Cell

// -------------------------------------------------------------------------------------------------

func Eval(exp, env *types.Cell, eval FuncType) (*types.Cell, error)  {

	if c := core.Car(exp); c.IsAtom() {
		if fn, found := builtinFuncs[*c]; found {
			return fn(exp, env, eval), nil
		}
	}
	
	return nil, OperatorNotFound
}

// -------------------------------------------------------------------------------------------------

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

func No_(exp, env *types.Cell, eval FuncType) *types.Cell {
	x := eval(core.Cadr(exp), env)
	return No(x)
}

func Not_(exp, a *types.Cell, eval FuncType) *types.Cell {
	x := eval(core.Cadr(exp), a)
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

func Def_(exp, env *types.Cell, eval FuncType) *types.Cell {
	return Def(exp, env, eval)
}

func Var_(exp, env *types.Cell, eval FuncType) *types.Cell {
	return Var(exp, env, eval)
}

func Let_(exp, env *types.Cell, eval FuncType) *types.Cell {
	return Let(exp, env, eval)
}

func Env_(exp, env *types.Cell, eval FuncType) *types.Cell {
	return Env(exp, env)
}

func Func_(exp, env *types.Cell, eval FuncType) *types.Cell {
	return Fun(exp, env, eval)
}

func Mac_(exp, env *types.Cell, eval FuncType) *types.Cell {
	return Mac(exp, env, eval)
}

func Do_(exp, env *types.Cell, eval FuncType) *types.Cell {
	return Do(exp, env, eval)
}

// -------------------------------------------------------------------------------------------------

func Backquote_(exp, env *types.Cell, eval FuncType) *types.Cell {
	return Backquote(exp, env, eval)
}

// -------------------------------------------------------------------------------------------------
