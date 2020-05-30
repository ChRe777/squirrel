package core

import (
	"fmt"
)

import (
	"github.com/squirrel/types"
)

// evmac eval 'mac and create a macros in environment
// 	e.g.
//	 	(mac {name} {params} {body})
//  	(var {name} (mac {params} {body}) )
func evmac(e, a *types.Cell) *types.Cell {
	name := cadr(e); params_body := cddr(e)
	k := name; v := cons(core.MAC, params_body)
	core.Tag(v, core.ID_MAC)
	a = addEnv(list(k, v), a)
	return eval(k, a)
}

// backquote
// e.g. 
//		`(list ,a ,b)   a = 1, b = 2
//   
//		(backquote
//			(list (unquote a) (unquote b))
//		) 
//		-> (list 1 2)
func Backquote(e *types.Cell, a *types.Cell) *types.Cell {
    x := cadr(e)
    y := mapEx(expand, x, a)
    return y
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


//	--------------------------------
//	Helpers for backquote and macros
//  --------------------------------

type fnCell func(e *types.Cell, a *types.Cell) *types.Cell

// mapEx - maps through a element in list and expand each element
// if the element is wrapped with (unquote) the element will be evaluated
func mapEx(fn fnCell, e *types.Cell, a *types.Cell) *types.Cell {
	if no(e).Equal(core.T) {
		return core.NIL
	} else {
		x := core.Car(e); xs := core.Cdr(e)			
		return core.Cons(fn(x, a), mapEx(fn, xs, a))
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
				case c.Equal(core.UNQUOTE): return unquote(e, a) 	
				// unquote-splicing shorcut: ,@
				//`((+ 1 2) ,(+ 3 4) ,@(list 5 6))
				// ((+ 1 2) 7 5 6)
				case c.Equal(core.UNQUOTE_SPLICING): return unquoteSplicing(e, a)				
			}
		}
		return e
	}
}
