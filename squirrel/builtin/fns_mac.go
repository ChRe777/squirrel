package builtin

import (
	//"fmt"
)

import (
	"github.com/mysheep/squirrel/types"
	"github.com/mysheep/squirrel/core"

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

func MacroExpand(exp, env *types.Cell, eval func(*types.Cell, *types.Cell) *types.Cell) *types.Cell {
	return Backquote(exp, env, eval)
}

func Backquote(exp, env *types.Cell, eval func(*types.Cell, *types.Cell) *types.Cell) *types.Cell {

    x := core.Cadr(exp)           // exp -> (backquote (list (unquote x) (unquote y))
    y := expandList(x, env, eval) // fill out all the "unquote" holes
       
    return y
}

// expandList - maps through a element in list and expand each element
// if the element is wrapped with (unquote) the element will be evaluated
func expandList(exp, env *types.Cell, eval func(*types.Cell, *types.Cell) *types.Cell) *types.Cell {
	
	if No(exp).Equal(core.T) {
		return core.NIL
	} 
			
	y, explode := expand(core.Car(exp), env, eval)			// ,@xs explodes list
	ys := core.Cdr(exp)
		
	if explode {
		xs := y
		return Append(xs, expandList(ys, env, eval))
	} else {
		return core.Cons(y, expandList(ys, env, eval))
	}
	
}

// expand - fill out the placeholder marked with (unquote a)
// e.g.
//		(var a 1 b 2)
//		`(list (unquote a) (unquote b))
//		(list 1 2)
//		(1 2)
func expand(exp, env *types.Cell, eval func(*types.Cell, *types.Cell) *types.Cell) (*types.Cell, bool) {
	
	if exp.IsAtom() {
		return exp, false
	} else {
		c := core.Car(exp)
			
		splice := false 
	
		if c.IsAtom() {	  // (unquote c) or (unquote_splicing c)
			switch {	
				
				// x=1, y=2 | `(,x ,y) -> (1 2)
				//
				case c.Equal(UNQUOTE): {
					return unquote(exp, env, eval), splice
				}
				
				// unquote-splicing shorcut: ,@
				// `((+ 1 2) ,(+ 3 4) ,@(list 5 6))
				//  ((+ 1 2) 7 5 6)
				//
				case c.Equal(UNQUOTE_SPLICING): {
					splice = true 
					return unquote(exp, env, eval), splice
				}
			
			}
		} 
		
		return expandList(exp, env, eval), splice
	}
}

// unquote are used in backquote to fill in the variable
// by enabling quotes of unquoted symbol
// e.g. 
//		(unquote a) 	a = 1
//		-> 1
func unquote(exp *types.Cell, env *types.Cell, eval func(*types.Cell, *types.Cell) *types.Cell ) *types.Cell {
	y := eval(core.Cadr(exp), env) 				// fills out of unquoted places, e.g. (x 1) (y 2) -> (list ,x ,y) -> (list 1 2)
	return y
}

