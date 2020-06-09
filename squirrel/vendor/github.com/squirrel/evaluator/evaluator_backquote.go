package evaluator

import (
	//"fmt"
)

import (
	"github.com/squirrel/types"
	"github.com/squirrel/core"
	"github.com/squirrel/builtin"
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
func evalBackquote(e *types.Cell, a *types.Cell) *types.Cell {
   
    x := builtin.Cadr(e)
    y := macExpand(x, a)	// fill out all the "unquote" holes
    
    //fmt.Printf("evalBackquote - y: %v \n", y)	
    
    return y
}

// mapEx - maps through a element in list and expand each element
// if the element is wrapped with (unquote) the element will be evaluated
func macExpand(e *types.Cell, a *types.Cell) *types.Cell {
	if builtin.No(e).Equal(core.T) {
		return core.NIL
	} else {
			
		y, explode := expand(core.Car(e), a)
		ys := core.Cdr(e)
	
		//fmt.Printf("macExpand - y: %v, explode: %v \n", y, explode)
	
		if explode {
			xs := y
			return builtin.Append(xs, macExpand(ys, a))
		} else {
			return core.Cons(y, macExpand(ys, a))
		}
	}
}

// expand - fill out the placeholder marked with (unquote a)
// e.g.
//		(var a 1 b 2)
//		`(list (unquote a) (unquote b))
//		(list 1 2)
//		(1 2)
func expand(e *types.Cell, a *types.Cell) (*types.Cell, bool) {	

	if e.IsAtom() {
		return e, false
	} else {
		c := core.Car(e)
		
	//	fmt.Printf("expand - c: %v, e: %v \n", c, e)	// (0 (unquote_splicing xs))
	
		splice := false 
	
		if c.IsAtom() {	  // (unquote c) or (unquote_splicing c)
			switch {	
				
				// x=1, y=2 | `(,x ,y) -> (1 2)
				//
				case c.Equal(core.UNQUOTE): {
					return unquote(e, a), splice
				}
				
				// unquote-splicing shorcut: ,@
				// `((+ 1 2) ,(+ 3 4) ,@(list 5 6))
				//  ((+ 1 2) 7 5 6)
				//
				case c.Equal(core.UNQUOTE_SPLICING): {
					splice = true 
					return unquote(e, a), splice
				}
			
			}
		} 
		
		return macExpand(e, a), splice
	}
}

// unquote are used in backquote to fill in the variable
// by enabling quotes of unquoted symbol
// e.g. 
//		(unquote a) 	a = 1
//		-> 1
func unquote(e *types.Cell, a *types.Cell) *types.Cell {
	
	x := builtin.Cadr(e)
	
	y := eval(x, a)	// TODO: Should fill in, but not EVAL ..???
	
	//fmt.Printf("unquote - x: %v, y: %v \n", x, y)
	
	return y
}

// unquoteSplicing
//	e.g.
//	   `((+ 1 2) ,(+ 3 4) ,@(list 5 6))
// 		((+ 1 2) 7 5 6)
/*
func unquoteSplicing(e *types.Cell, a *types.Cell) *types.Cell {
	
	x := builtin.Cadr(e)
	y := eval(x, a)	
	
	return y
}
*/
