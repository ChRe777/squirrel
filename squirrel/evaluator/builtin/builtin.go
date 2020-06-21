package builtin

import (
//	"fmt"
	"errors"
)

import (
	"github.com/mysheep/squirrel/types"
	"github.com/mysheep/squirrel/evaluator/core"
)

// -------------------------------------------------------------------------------------------------

func Eval(e, a *types.Cell, eval func(*types.Cell, *types.Cell) *types.Cell) (*types.Cell, error)  {

	if c := core.Car(e); c.IsAtom() {
		if fn, found := builtinFuncs[*c]; found {
			return fn(e, a, eval), nil
		}
	}
	return nil, OPERATOR_NOT_FOUND
}

// -------------------------------------------------------------------------------------------------

var builtinFuncs = map[types.Cell] func(*types.Cell, *types.Cell, func(*types.Cell, *types.Cell) *types.Cell) *types.Cell {
	
	//
	// fns1
	//
	*NO     : No_		,
	*NOT    : Not_		,
	*AND    : And_		,
	*PAIR   : Pair_		,
	*LIST   : List_		,
	*ASSOC  : Assoc_	,
	*APPEND : Append_	,
	
	//
	// fns2
	//
	*COND	: Cond_		,
	*VAR	: Var_		,
	*DEF	: Def_ 		,
	*LET	: Let_		,
	*FUNC	: Func_		,		
	*ENV	: Env_		,
	*MAC	: Mac_		,
	*DO		: Do_		,
	
	//
	// macro
	//
	*BACKQUOTE			: Backquote_		,
//	*UNQUOTE   			: Unquote_ 			,		// used internal
//	*UNQUOTE_SPLICING   : Unquotesplicing_	,		// used internal

}

// -------------------------------------------------------------------------------------------------

var (
	OPERATOR_NOT_FOUND = errors.New("Operator not found")
)

// -------------------------------------------------------------------------------------------------

func No_(e, a *types.Cell, eval func(*types.Cell, *types.Cell) *types.Cell) *types.Cell {
	x := eval(core.Cadr(e), a)
	return No(x)
}

func Not_(e, a *types.Cell, eval func(*types.Cell, *types.Cell) *types.Cell) *types.Cell {
	x := eval(core.Cadr(e), a)
	return Not(x)
}

func And_(e, a *types.Cell, eval func(*types.Cell, *types.Cell) *types.Cell) *types.Cell {
	x := eval(core.Cadr(e), a)
	y := eval(core.Caddr(e), a)
	return And(x, y)
}

func Append_(e, a *types.Cell, eval func(*types.Cell, *types.Cell) *types.Cell) *types.Cell {
	xs := eval(core.Cadr(e), a)
	ys := eval(core.Caddr(e), a)
	return Append (xs, ys)
}

func Pair_(e, a *types.Cell, eval func(*types.Cell, *types.Cell) *types.Cell) *types.Cell {
	xs := eval(core.Cadr(e), a)
	ys := eval(core.Caddr(e), a)
	return Pair (xs, ys)
}

func List_(e, a *types.Cell, eval func(*types.Cell, *types.Cell) *types.Cell) *types.Cell {
	xs := core.Cdr(e)
	return List(xs, a, eval)
}

func Assoc_(e, a *types.Cell, eval func(*types.Cell, *types.Cell) *types.Cell) *types.Cell {
	x  := eval(core.Cadr(e), a)
	ys := eval(core.Caddr(e), a)	
	return Assoc(x, ys)
}

// -------------------------------------------------------------------------------------------------

func Cond_(e, a *types.Cell, eval func(*types.Cell, *types.Cell) *types.Cell) *types.Cell {
	cs := core.Cdr(e)
	return Cond(cs, a, eval)
}

func Def_(e, a *types.Cell, eval func(*types.Cell, *types.Cell) *types.Cell) *types.Cell {
	return Def(e, a, eval)
}

func Var_(e, a *types.Cell, eval func(*types.Cell, *types.Cell) *types.Cell) *types.Cell {
	return Var(e, a, eval)
}

func Let_(e, a *types.Cell, eval func(*types.Cell, *types.Cell) *types.Cell) *types.Cell {
	return Let(e, a, eval)
}

func Env_(e, a *types.Cell, eval func(*types.Cell, *types.Cell) *types.Cell) *types.Cell {
	return Env(e, a)
}

func Func_(e, a *types.Cell, eval func(*types.Cell, *types.Cell) *types.Cell) *types.Cell {
	return Fun(e, a)
}

func Mac_(e, a *types.Cell, eval func(*types.Cell, *types.Cell) *types.Cell) *types.Cell {
	return Mac(e, a, eval)
}

func Do_(e, a *types.Cell, eval func(*types.Cell, *types.Cell) *types.Cell) *types.Cell {
	return Do(e, a, eval)
}

// -------------------------------------------------------------------------------------------------

func Backquote_(e, a *types.Cell, eval func(*types.Cell, *types.Cell) *types.Cell) *types.Cell {
	return Backquote(e, a, eval)
}



