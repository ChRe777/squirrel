package builtin

import (
	_ "github.com/mysheep/squirrel/core"
	_ "github.com/mysheep/squirrel/evaluator"
	_ "github.com/mysheep/squirrel/generator"
	_ "github.com/mysheep/squirrel/types"
)

// (def map (f x) (cond ((no x) nil) ('t (cons (f (car x)) (map f (cdr x))))))

var (
	mapFn = "(map	(func (f x)  (cond ((no x) nil) ('t (cons (f (car x)) (map f (cdr x)))))))"
)
