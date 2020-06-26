package functions

import (
	_ "github.com/mysheep/squirrel/core"
	_ "github.com/mysheep/squirrel/evaluator"
	_ "github.com/mysheep/squirrel/generator"
	_ "github.com/mysheep/squirrel/types"
)

// (def map (f xs) (cond ((no xs) nil) ('t (cons (f (car xs)) (map f (cdr xs))))))

//var (
//	mapFn = "(map	(func (f x)  (cond ((no x) nil) ('t (cons (f (car x)) (map f (cdr x)))))))"
//)

func Map(fn, xs *types.Cell) *types.Cell {
	if builtin.No(xs) {
		return core.NIL
	} else {
		return core.Cons(fn(core.Car(xs)),
					     Map(fn, core.Cdr(xs)))
	}
}