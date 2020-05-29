package evaluator

import (
	"fmt"
)

import (
	"github.com/squirrel/types"
	"github.com/squirrel/builtin"
)

type fnCell func(e *types.Cell, a *types.Cell) *types.Cell


// arc> (mac when (test . body)
//       (list 'if test (cons 'do body)))
// *** redefining when
// #3(tagged mac #<procedure>)

//	(mac  {name} {params} {body})
//	(func {name} {params} {body}) -- tagged as 'mac

// if func is tagged as 'mac
// then macExpand 

// evfunc


// backquote
// unquote

// backquote
// e.g. 
//		`(list ,a ,b)   a = 1, b = 2
//   
//		(backquote
//			(list (unquote a) (unquote b))
//		) 
//		-> (list 1 2)
func backquote(e *types.Cell, a *types.Cell) *types.Cell {
    x := cadr(e)
    y := mapEx(expand, x, a)
    return y
}

// mapEx - maps through a element in list and expand each element
// if the element is wrapped with (unquote) the element will be evaluated
func mapEx(fn fnCell, e *types.Cell, a *types.Cell) *types.Cell {
	if no(e).Equal(builtin.T) {
		return builtin.NIL
	} else {
		x := builtin.Car(e); xs := builtin.Cdr(e)			
		return builtin.Cons(fn(x, a), mapEx(fn, xs, a))
	}
}

// expand - fill out the placeholder marked with (unquote a)
// e.g.
//		(list (unquote a) (unquote b))
func expand(e *types.Cell, a *types.Cell) *types.Cell {	
	if e.IsAtom() {
		return e
	} else {
		c := car(e)
		if c.IsAtom() {
			switch {	
				case c.Equal(builtin.UNQUOTE): return unquote(e, a) 	
				// unquote-splicing shorcut: ,@
				//`((+ 1 2) ,(+ 3 4) ,@(list 5 6))
				// ((+ 1 2) 7 5 6)
				case c.Equal(builtin.UNQUOTE_SPLICING): return unquoteSplicing(e, a)				
			}
		}
		return e
	}
}

// unquote are used in backquote to fill in the variable
// by enabling quotes of unquoted symbol
// e.g. 
//		(unquote a) 	a = 1
//		-> 1
func unquote(e *types.Cell, a *types.Cell) *types.Cell {
	x := cadr(e); y := eval(x, a)
	return y
}

// unquoteSplicing
//	e.g.
//		`((+ 1 2) ,(+ 3 4) ,@(list 5 6))
// 		((+ 1 2) 7 5 6)
func unquoteSplicing(e *types.Cell, a *types.Cell) *types.Cell {
	fmt.Printf("unquoteSplicing - e:%v \n", e)
	x := cadr(e); y := eval(x, a)
	return y
}
