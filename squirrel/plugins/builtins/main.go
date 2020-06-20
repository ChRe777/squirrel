package main

import (
	"errors"
)

import (
	"github.com/mysheep/squirrel/types"
	"github.com/mysheep/squirrel/evaluator/core"
//	"github.com/mysheep/squirrel/evaluator/builtin"
)

// -------------------------------------------------------------------------------------------------

type any string 	// could any type

var Evaluator any	// Name is used to detect plugin type

// -------------------------------------------------------------------------------------------------

func (p any) Eval(e, a *types.Cell) (*types.Cell, error)  {

	if c := core.Car(e); c.IsAtom() {
		if op, found := builtOps[*c]; found {
			return op(e, a), nil
		}
	}

	return nil, errors.New("Operator not found")
}

// -------------------------------------------------------------------------------------------------

type OpFunc func (e, a *types.Cell) *types.Cell

var builtOps = map[types.Cell] OpFunc {
	/*
	*builtin.NO     : No_		,
	*builtin.NOT    : Not_		,
	*builtin.AND    : And_		,
	*builtin.PAIR   : Pair_		,
	*builtin.LIST   : List_		,
	*builtin.ASSOC  : Assoc_	,
	*builtin.APPEND : Append_	,
	*/
	//
	// no more caarrrrsss or cdrrrrssss :-)
	//
	/*
	*builtin.CAAR  	: Caar_		,  
	*builtin.CADR  	: Cadr_		,  
	*builtin.CDDR  	: Cddr_		,  
	*builtin.CADAR 	: Cadar_	, 
	*builtin.CDDDR 	: Cdddr_	, 
	*builtin.CADDR 	: Caddr_	, 
	*builtin.CADDAR	: Caddar_	,
	*builtin.CADDDR	: Cadddr_	, 
	*/
	
}

// TODO: 

//    mapFn    = "(map	(func (f x)  (cond ((no x) nil) ('t (cons (f (car x)) (map f (cdr x)))))))"



// -------------------------------------------------------------------------------------------------
/*
func Pair_ (e, a *types.Cell) *types.Cell {
	x := evaluator.Eval(builtin.Cadr (e), a)
	y := evaluator.Eval(builtin.Caddr(e), a)
	return builtin.Pair(x, y)
}

func No_(e, a *types.Cell) *types.Cell {
	x := evaluator.Eval(builtin.Cadr (e), a)
	return builtin.No(x)
}

func Not_(e, a *types.Cell) *types.Cell {
	x := evaluator.Eval(builtin.Cadr (e), a)
	return builtin.Not(x)
}

func And_ (e, a *types.Cell) *types.Cell {
	x := evaluator.Eval(builtin.Cadr (e), a)
	y := evaluator.Eval(builtin.Caddr(e), a)
	return builtin.And(x, y)
}

func Append_ (e, a *types.Cell) *types.Cell {
	x := evaluator.Eval(builtin.Cadr (e), a)
	y := evaluator.Eval(builtin.Caddr(e), a)
	return builtin.Append(x, y)
}

func Assoc_ (e, a *types.Cell) *types.Cell {
	x := evaluator.Eval(builtin.Cadr (e), a)
	y := evaluator.Eval(builtin.Caddr(e), a)
	return builtin.Assoc(x, y)
}

func List_(e, a *types.Cell) *types.Cell {
	ys := core.Cdr(e)
	return builtin.List(ys, a)
}
*/

// -------------------------------------------------------------------------------------------------
/*
func Caar_(e, a *types.Cell) *types.Cell {
	x := evaluator.Eval(builtin.Cadr (e), a)
	return builtin.Caar(x)
}

func Cddr_(e, a *types.Cell) *types.Cell {
	x := evaluator.Eval(builtin.Cadr (e), a)
	return builtin.Cddr(x)
}

func Cadr_(e, a *types.Cell) *types.Cell {
	x := evaluator.Eval(builtin.Cadr (e), a)
	return builtin.Cadr(x)
}

func Cadar_(e, a *types.Cell) *types.Cell {
	x := evaluator.Eval(builtin.Cadr (e), a)
	return builtin.Cadar(x)
}

func Cdddr_(e, a *types.Cell) *types.Cell {
	x := evaluator.Eval(builtin.Cadr (e), a)
	return builtin.Cdddr(x)
}

func Caddr_(e, a *types.Cell) *types.Cell {
	x := evaluator.Eval(builtin.Cadr (e), a)
	return builtin.Caddr(x)
}

func Caddar_(e, a *types.Cell) *types.Cell {
	x := evaluator.Eval(builtin.Cadr (e), a)
	return builtin.Caddar(x)
}

func Cadddr_(e, a *types.Cell) *types.Cell {
	x := evaluator.Eval(builtin.Cadr (e), a)
	return builtin.Cadddr(x)
}

// -------------------------------------------------------------------------------------------------
*/
