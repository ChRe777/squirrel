package builtin

import (
	//"fmt"
)

import (
	"github.com/mysheep/squirrel/types"
	"github.com/mysheep/squirrel/evaluator/core"
)

/*

	- Cond
	- List
	
	- Var
	- Def
	- Let
	- Fun	
	
	- Env
	- Mac
	- Do
	
*/


// Cond evals cond (= conditions)
// e.g. 
//		> (cond (
//			(nil b) 
//			( 't a))) 	-> a
//
func Cond(c, a *types.Cell, eval func(*types.Cell, *types.Cell) *types.Cell) *types.Cell {

	if c.Equal(core.NIL) {
		return core.NIL
	}
		
	y := eval(core.Caar(c), a)
	if y.IsErr() {
		return y;
	}

	if y.Equal(core.T) { 
		return eval(core.Cadar(c), a) 
	} else { 
		return Cond(core.Cdr(c), a, eval)
	}

}

// List evals each item of a list
func List(xs, a *types.Cell, eval func(*types.Cell, *types.Cell) *types.Cell) *types.Cell {

	if xs.Equal(core.NIL) {
		return core.NIL
	}
	
	y := eval(core.Car(xs), a)
	if y.IsErr() {
		return y;
	}
	
	return core.Cons(y, List(core.Cdr(xs), a, eval))
}



// evalDef eval 'def and creates a function in environment
// e.g.
//  	(def {name} {params}_{body})
//  	({name} (func {params}_{body})
func Def(e, a *types.Cell, eval func(*types.Cell, *types.Cell) *types.Cell) *types.Cell {
	name := core.Cadr(e); params_body := core.Cddr(e)
	key := name; val := core.Cons(FUNC, params_body)		// TODO: REFACTOR
	
	core.Tag(val, ID_FUNC)
	a = core.Add(list2_(key, val), a)
	
	return eval(key, a)
}

// evset evals expression e.g. (set a 1)
// add a key value pair on top of environment like push on a stack
// 	e.g.
//		> (env) 		-> ((t t))
// 		> (var a 1) 	-> 1
// 		> (env) 		-> ((a 1) (t t))
//
func Var(e, a *types.Cell, eval func(*types.Cell, *types.Cell) *types.Cell) *types.Cell {
	key := core.Cadr(e)
	val := eval(core.Caddr(e), a)
	
	a = core.Add(list2_(key, val), a)
	
	return eval(key, a)
}

// Let eval 'let, see example below
// 	e.g. 
//		> (let {key} {val} {body} )
//		> (let xs '(1 2 3) (car xs)) 	->  1
//		
func Let(e, a *types.Cell, eval func(*types.Cell, *types.Cell) *types.Cell) *types.Cell {
	key := core.Cadr(e)
	val := eval(core.Caddr(e), a)	
	
	ee := core.Car(core.Cdddr(e))
	aa := core.Cons(list2_(key, val), a)	
	
	return eval(ee, aa)
}

// Env prints environment for debug purpose
// e.g.
//		env = ((a 1) (b 1))
//
//		> (env) -> ((a 1) (b 1))
//
func Env(e, a *types.Cell) *types.Cell {
	return a
}


// Fun create a function
// e.g.
//		(func (x) (car x))  -> func
func Fun(e, a *types.Cell) *types.Cell {
	
	v := e; core.Tag(v, ID_FUNC)
	
	return v
}

// Mac eval 'mac and create a macros in environment
// 	e.g.
//	 	(mac {name} {params}_{body})
//
func Mac(e, a *types.Cell, eval func(*types.Cell, *types.Cell) *types.Cell) *types.Cell {

	name := core.Cadr(e); params_body := core.Cddr(e)	
	
	key := name; val := core.Cons(FUNC, params_body)	
		
	core.Tag(val, ID_MAC)	// A macros is a func tagged as macro (Paul Graham - Arc)

	aa := core.Add(list2_(key, val), a)
	
	return eval(key, aa)
}

//	------------------------------------------------------------------------------------------------

// Do evals a list of expression and returns the last expression	
//	e.g.
//		(do
//			(list 1 2)
//			(no nil)
//		)
func Do(e, a *types.Cell, eval func(*types.Cell, *types.Cell) *types.Cell) *types.Cell {

	var doList func(e, last, a *types.Cell) *types.Cell

	doList = func(e, last, a *types.Cell) *types.Cell {
		if e.Equal(core.NIL) {
			return last
		} else {
			x := core.Car(e); xs := core.Cdr(e)
			l := eval(x, a)
			return doList(xs, l, a)
		}	
	}

	return doList(core.Cdr(e), core.NIL, a)
}


