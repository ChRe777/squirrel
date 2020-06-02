package builtin

import (
	"github.com/squirrel/types"
	"github.com/squirrel/parser"
)

var fnsMap = map[string]string{
	"t" 		: "t",
	"nil" 		: "nil",
	"no"     	: "(func (x)    (is x '()))",
	"and"    	: "(func (x y)  (cond (x (cond (y 't) ('t '())))('t '())))",
	"not"    	: "(func (x)    (cond (x '()) ('t 't)))",
	"append" 	: "(func (x y)  (cond ((no x) y) ('t (cons (car x) (append (cdr x) y)))))",
	"pair"   	: "(func (x y)  (cond ((and (no x) (no y)) '()) ((and (not (atom x)) (not (atom y))) (cons (list (car x) (car y))(pair (cdr x) (cdr y))))))",
//	"list"   	: "(func (x y)  (cons x (cons y '())))",
	"assoc"  	: "(func (x ys) (cond ((no ys) nil) ((is x (car (car ys))) (car (cdr (car ys)))) ('t (assoc x (cdr ys)))))",
    "map"		: "(func (f xs) (cond ((no xs) nil) ('t (cons (f (car xs)) (map f (cdr xs))))))",
}

// -------------------------------------------------------------------------------------------------

var (
	Fns = make(map[string](*types.Cell))
)

func init() {

	//	no     (func (x)    (is x '()))
	//	and    (func (x y)  (cond (x (cond (y 't) ('t '())))('t '())))

	for name, fnStr := range fnsMap {
		Fns[name] = parser.Parse([]byte(fnStr))
	}	
}