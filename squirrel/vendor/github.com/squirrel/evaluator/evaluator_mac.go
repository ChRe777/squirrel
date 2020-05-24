package evaluator

import (
	//"fmt"
)

import (
	"github.com/squirrel/types"
	"github.com/squirrel/builtin"
)

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
  //	fmt.Printf("backquote - e: %v \n", e)
    x := cadr(e)
  //  fmt.Printf("backquote - x: %v \n", x)
    y := mapEx(x, a)
    return y
}

func mapEx(e *types.Cell, a *types.Cell) *types.Cell {
	if no(e).Equal(builtin.T) {
		return builtin.NIL
	} else {
		x := builtin.Car(e); xs := builtin.Cdr(e)			
		return builtin.Cons(expand(x, a), mapEx(xs, a))
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
			}
		}
		return e
	}
}

// unquote
// e.g. 
//		(unquote a) 	a = 1
//		-> 1

func unquote(e *types.Cell, a *types.Cell) *types.Cell {
	x := cadr(e); y := eval(x, a)
	return y
}
