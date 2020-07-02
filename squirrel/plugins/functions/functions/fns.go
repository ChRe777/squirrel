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
