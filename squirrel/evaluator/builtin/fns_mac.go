package builtin

import (
	//"fmt"
)

import (
	"github.com/mysheep/squirrel/types"
	"github.com/mysheep/squirrel/evaluator/core"

)

// backquote
// e.g. 
//		`(list ,a ,b)   a = 1, b = 2
//   
//		(backquote
//			(list (unquote a) (unquote b))
//		) 
//		-> (list 1 2)
//
//		(var xs '(1 2) 
//			 x  'a
//		)
//	   `(list ,@xs
//			(list ,x ,@xs)
//		)
//
// 		(list 1 2 (list a 1 2))
//		(1 2 (a 1 2))
//

func MacEx(e, a *types.Cell, eval func(*types.Cell, *types.Cell) *types.Cell) *types.Cell {
	return Backquote(e, a, eval)
}

func Backquote(e, a *types.Cell, eval func(*types.Cell, *types.Cell) *types.Cell) *types.Cell {
   
   	// e.g. (backquote (list ,x, y))
   
    x := core.Cadr(e)
    y := expandList(x, a, eval)	// fill out all the "unquote" holes
       
    return y
}

// expandList - maps through a element in list and expand each element
// if the element is wrapped with (unquote) the element will be evaluated
func expandList(e, a *types.Cell, eval func(*types.Cell, *types.Cell) *types.Cell) *types.Cell {
	
	if No(e).Equal(core.T) {
		return core.NIL
	} 
			
	y, explode := expand(core.Car(e), a, eval)
	ys := core.Cdr(e)
		
	if explode {
		xs := y
		return Append(xs, expandList(ys, a, eval))
	} else {
		return core.Cons(y, expandList(ys, a, eval))
	}
	
}

// expand - fill out the placeholder marked with (unquote a)
// e.g.
//		(var a 1 b 2)
//		`(list (unquote a) (unquote b))
//		(list 1 2)
//		(1 2)
func expand(e, a *types.Cell, eval func(*types.Cell, *types.Cell) *types.Cell) (*types.Cell, bool) {	
	
	if e.IsAtom() {
		return e, false
	} else {
		c := core.Car(e)
			
		splice := false 
	
		if c.IsAtom() {	  // (unquote c) or (unquote_splicing c)
			switch {	
				
				// x=1, y=2 | `(,x ,y) -> (1 2)
				//
				case c.Equal(UNQUOTE): {
					return unquote(e, a, eval), splice
				}
				
				// unquote-splicing shorcut: ,@
				// `((+ 1 2) ,(+ 3 4) ,@(list 5 6))
				//  ((+ 1 2) 7 5 6)
				//
				case c.Equal(UNQUOTE_SPLICING): {
					splice = true 
					return unquote(e, a, eval), splice
				}
			
			}
		} 
		
		return expandList(e, a, eval), splice
	}
}

// unquote are used in backquote to fill in the variable
// by enabling quotes of unquoted symbol
// e.g. 
//		(unquote a) 	a = 1
//		-> 1
func unquote(e *types.Cell, a *types.Cell, eval func(*types.Cell, *types.Cell) *types.Cell ) *types.Cell {
	
	x := core.Cadr(e)
	y := eval(x, a)	// TODO: Should fill in, but not EVAL ..???
		
	return y
}

