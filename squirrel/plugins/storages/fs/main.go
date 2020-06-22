package main

import (
	"errors"
)

import (
	"github.com/mysheep/squirrel/core"
	"github.com/mysheep/squirrel/evaluator"
	"github.com/mysheep/squirrel/generator" // TODO: Use core
	"github.com/mysheep/squirrel/plugins/storages/fs/loader"
	"github.com/mysheep/squirrel/plugins/storages/fs/storer"
	"github.com/mysheep/squirrel/types"
)

// -------------------------------------------------------------------------------------------------

type any string // Could be any type

var Evaluator any // Name ist important to detected plugin

// -------------------------------------------------------------------------------------------------

func (p any) Eval(e, a *types.Cell) (*types.Cell, error) {

	if c := core.Car(e); c.IsAtom() {
		if op, found := builtOps[*c]; found {
			return op(e, a), nil
		}
	}

	return nil, errors.New("Operator not found")
}

// -------------------------------------------------------------------------------------------------

func (p any) Load(location string) (*types.Cell, error) {

	s := generator.Str(location)

	res := loader.Load(s)

	if res.IsErr() {
		msg, _ := res.AsStr()
		return nil, errors.New(msg)
	}

	return res, nil
}

func (p any) Store(location string, c *types.Cell) error {

	s := generator.Str(location)
	res := storer.Store(s, c)

	if res.IsErr() {
		msg, _ := res.AsStr()
		errors.New(msg)
	}

	return nil
}

// -------------------------------------------------------------------------------------------------

type OpFunc func(e, a *types.Cell) *types.Cell

var builtOps = map[types.Cell]OpFunc{
	*loader.LOAD:  load_,
	*storer.STORE: store_,
}

// -------------------------------------------------------------------------------------------------

func load_(e, a *types.Cell) *types.Cell {

	loc := evaluator.Eval(cadr_(e), a)
	exp := loader.Load(loc)

	return evaluator.Eval(exp, a)

}

func store_(e, a *types.Cell) *types.Cell {

	loc := evaluator.Eval(cadr_(e), a)
	exp := evaluator.Eval(caddr_(e), a)

	return storer.Store(loc, exp)
}

// -------------------------------------------------------------------------------------------------

func cadr_(e *types.Cell) *types.Cell {
	return core.Car(core.Cdr(e))
}

func caddr_(e *types.Cell) *types.Cell {
	return core.Car(core.Cdr(core.Cdr(e)))
}
