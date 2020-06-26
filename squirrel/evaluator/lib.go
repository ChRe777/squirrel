package evaluator

import (
	"github.com/mysheep/squirrel/builtin"
	"github.com/mysheep/squirrel/core"
	"github.com/mysheep/squirrel/types"
)

// -------------------------------------------------------------------------------------------------
// Type aliases
//
type Cell = types.Cell

// ------------------------------------------------------------------------------------------------
// Environment functions for manipulating the symbol table
//
func EvList(exp *Cell, env *Cell) *Cell {
	return builtin.List(exp, env, eval)
}

func Bind(kvs *Cell, env *Cell) *Cell {
	return builtin.Append(kvs, env)
}

func Pair(xs, ys *Cell) *Cell {
	return builtin.Pair(xs, ys)
}

func Value(key, env *Cell) *Cell {
	return builtin.Assoc(key, env)
}

// ------------------------------------------------------------------------------------------------
// Shortcut or interfaces for core or builtin functions
//
func Sym(s string) *Cell {
	return core.Sym_(s)
}

func Cons(x, y *Cell) *Cell {
	return core.Cons(x, y)
}

func Car(x *Cell) *Cell {
	return core.Car(x)
}

func Caar(x *Cell) *Cell {
	return core.Caar(x)
}

func Cadr(x *Cell) *Cell {
	return core.Cadr(x)
}

func Cadar(exp *Cell) *Cell {
	return core.Cadar(exp)
}

func Caddr(exp *Cell) *Cell {
	return core.Caddr(exp)
}

func Caddar(exp *Cell) *Cell {
	return core.Caddar(exp)
}

func Cadddr(exp *Cell) *Cell {
	return core.Cadddr(exp)
}

func Cdr(exp *Cell) *Cell {
	return core.Cdr(exp)
}

func Err(s string, a ...interface{}) *Cell {
	return core.Err_(s, a...)
}