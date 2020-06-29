package functions

import (
	"github.com/mysheep/squirrel/builtin"
	"github.com/mysheep/squirrel/core"
	"github.com/mysheep/squirrel/types"
)

// (def map (f xs) (cond ((no xs) nil) ('t (cons (f (car xs)) (map f (cdr xs))))))

//var (
//	mapFn = "(map	(func (f x)  (cond ((no x) nil) ('t (cons (f (car x)) (map f (cdr x)))))))"
//)

func Map(fn, xs, env *types.Cell, eval func(*types.Cell, *types.Cell) *types.Cell) *types.Cell {
	if builtin.No(xs).Equal(core.T) {
		return core.NIL
	} else {
		return core.Cons(eval(core.Cons(fn, core.Car(xs)), env),
					     Map(fn, core.Cdr(xs), env, eval))
	}
}