package evaluator

import (
	//"fmt"
)

import (
	"github.com/mysheep/squirrel/types"
	"github.com/mysheep/squirrel/evaluator/core"
	"github.com/mysheep/squirrel/evaluator/builtin"
)

/*

	- evalCond
	- evalList
	
	- evalVar
	- evalDef
	- evalLet
	- evalFun	
	
	- evalEnv
	
	...
	...

*/

// evalCond evals cond (= conditions)
// e.g. 
//		> (cond (
//			(nil b) 
//			( 't a))) 	-> a
//
func evalCond(c, a *types.Cell) *types.Cell {

	if c.Equal(core.NIL) {
		return core.NIL
	}
		
	y := eval(builtin.Caar(c), a)
	if y.IsErr() {
		return y;
	}

	if y.Equal(core.T) { 
		return eval(builtin.Cadar(c), a) 
	} else { 
		return evalCond(core.Cdr(c), a)
	}

}

// evalLst evals each item of a list
func evalLst(m, a *types.Cell) *types.Cell {

	if m.Equal(core.NIL) {
		return core.NIL
	}
	
	y := eval(core.Car(m), a)
	if y.IsErr() {
		return y;
	}
	
	return core.Cons(y, evalLst(core.Cdr(m), a))
}

// isMac checks if car(e) is tagged as macro
// e.g.
//		e = (func (x) (no x))
//		car(e) -> func (Tagged with mac) 
func isMac(e *types.Cell) bool {
	return core.Car(e).IsTagged(core.ID_MAC)
}

// evalDef eval 'def and creates a function in environment
// e.g.
//  	(def {name} {params}_{body})
//  	({name} (func {params}_{body})
func evalDef(e, a *types.Cell) *types.Cell {
	name := builtin.Cadr(e); params_body := builtin.Cddr(e)
	key := name; val := core.Cons(core.FUNC, params_body)		// TODO: REFACTOR
	
	core.Tag(val, core.ID_FUNC)
	a = core.Add(builtin.List_(key, val), a)
	
	return eval(key, a)
}

// evset evals expression e.g. (set a 1)
// add a key value pair on top of environment like push on a stack
// 	e.g.
//		> (env) 		-> ((t t))
// 		> (var a 1) 	-> 1
// 		> (env) 		-> ((a 1) (t t))
//
func evalVar(e, a *types.Cell) *types.Cell {
	key := builtin.Cadr(e)
	val := eval(builtin.Caddr(e), a)
	
	a = core.Add(builtin.List_(key, val), a)
	
	return eval(key, a)
}

// evalLet eval 'let, see example below
// 	e.g. 
//		> (let {key} {val} {body} )
//		> (let xs '(1 2 3) (car xs)) 	->  1
//		
func evalLet(e, a *types.Cell) *types.Cell {
	key := builtin.Cadr(e)
	val := eval(builtin.Caddr(e), a)	
	
	ee := core.Car(builtin.Cdddr(e))
	aa := core.Cons(builtin.List_(key, val), a)	
	
	return eval(ee, aa)
}

// envenv only print environment for debug purpose
// e.g.
//		env = ((a 1) (b 1))
//
//		> (env) -> ((a 1) (b 1))
//
func evalEnv(e, a *types.Cell) *types.Cell {
	return a
}


// evalFunc
// e.g.
//		(func (x) (car x))  -> func
func evalFun(e, a *types.Cell) *types.Cell {
	
	v := e; core.Tag(v, core.ID_FUNC)
	
	return v
}

// evmac eval 'mac and create a macros in environment
// 	e.g.
//	 	(mac {name} {params}_{body})
//
func evalMac(e, a *types.Cell) *types.Cell {

	name := builtin.Cadr(e); params_body := builtin.Cddr(e)	
		
	key := name; val := core.Cons(core.FUNC, params_body)	
		
	core.Tag(val, core.ID_MAC)	// A macros is a func tagged as macro (Paul Graham - Arc)

	a = core.Add(builtin.List_(key, val), a)
	
	return eval(key, a)
}

//	------------------------------------------------------------------------------------------------
/*
// evalLoad evals load function
// e.g.
//		(load "test.cell")
func evalLoad(e, a *types.Cell) *types.Cell {
	name := builtin.Cadr(e); exp := builtin.Load(name)	
	return eval(exp, a)
}
*/
//	------------------------------------------------------------------------------------------------

// evalDo evals a list of expression and returns the last expression	
//	e.g.
//		(do
//			(list 1 2)
//			(no nil)
//		)
func evalDo(e, a *types.Cell) *types.Cell {

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

