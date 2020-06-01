package evaluator

import (
	"fmt"
)

import (
	"github.com/squirrel/types"
	"github.com/squirrel/core"
	"github.com/squirrel/builtin"
)

// evmac eval 'mac and create a macros in environment
// 	e.g.
//	 	(mac {name} {params}_{body})
//
func evalMac(e, a *types.Cell) *types.Cell {
	name := builtin.Cadr(e); params_body := builtin.Cddr(e)	
		
	key := name; val := core.Cons(core.FUNC, params_body)	// A macros is a func tagged as macro
	core.Tag(val, core.ID_MAC)
	
	a = addEnv(builtin.List_(key, val), a)
	
	return eval(key, a)
}

// backquote
// e.g. 
//		`(list ,a ,b)   a = 1, b = 2
//   
//		(backquote
//			(list (unquote a) (unquote b))
//		) 
//		-> (list 1 2)
func evalBackquote(e *types.Cell, a *types.Cell) *types.Cell {
    x := builtin.Cadr(e)
    y := mapEx(expand, x, a)
    return y
}

// unquote are used in backquote to fill in the variable
// by enabling quotes of unquoted symbol
// e.g. 
//		(unquote a) 	a = 1
//		-> 1
func unquote(e *types.Cell, a *types.Cell) *types.Cell {
	x := builtin.Cadr(e); y := eval(x, a)
	return y
}

// unquoteSplicing
//	e.g.
//		`((+ 1 2) ,(+ 3 4) ,@(list 5 6))
// 		((+ 1 2) 7 5 6)
func unquoteSplicing(e *types.Cell, a *types.Cell) *types.Cell {
	fmt.Printf("unquoteSplicing - e:%v \n", e)
	x := builtin.Cadr(e); y := eval(x, a)
	return y
}

//	--------------------------------
//	Helpers for backquote and macros
//  --------------------------------

type fnCell func(e *types.Cell, a *types.Cell) *types.Cell

// mapEx - maps through a element in list and expand each element
// if the element is wrapped with (unquote) the element will be evaluated
func mapEx(fn fnCell, e *types.Cell, a *types.Cell) *types.Cell {
	if builtin.No(e).Equal(core.T) {
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
		c := core.Car(e)
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

// addEnv is a special add that adds a new cell at the front of the environment
// but LET the Pointer to first element the SAME !!!
func addEnv(kv *types.Cell, a *types.Cell ) *types.Cell {
	// Hang in new as second
	cdr := a.Cdr; new := core.Cons(kv, cdr); a.Cdr = new
	// Change Val first and second to move new second to front
	val := new.Val; new.Val = a.Val; a.Val = val
	// Change Car first and second to move new seocen to front
	car := new.Car; new.Car = a.Car; a.Car = car
	// So the pointer to a stays the same // Side effects // ToReThink: ?
	return a
}
