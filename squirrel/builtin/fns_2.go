package builtin

import (
 	"fmt"
)

import (
	"github.com/mysheep/squirrel/core"
	"github.com/mysheep/squirrel/types"
)

//	------------------------------------------------------------------------------------------------
//
//  List of functions:
//
//		- List
//		- Def
//		- Var
//		- Let
//		- Env
//		- Fun
//		- Mac
//		- Do
//

//	------------------------------------------------------------------------------------------------

// List evals each item of a list
func List(xs, env *types.Cell, eval func(*types.Cell, *types.Cell) *types.Cell) *types.Cell {

	if xs.Equal(core.NIL) {
		return core.NIL
	}
	
	y := eval(core.Car(xs), env)
	if y.IsErr() {
		return y
	}
	
	return core.Cons(y, List(core.Cdr(xs), env, eval))
}

//	------------------------------------------------------------------------------------------------

// evalDef eval 'def and creates a function in environment
// e.g.
//  	(def {name} {params}_{body})
//  	({name} (func {params}_{body})
func Def(exp, env *types.Cell, eval func(*types.Cell, *types.Cell) *types.Cell) *types.Cell {
	name := core.Cadr(exp); argsAndBody := core.Cddr(exp)
	key := name; val := core.Cons(core.Tag(core.Sym_(core.ID_FUNC), ID_FUNC), argsAndBody)

	env = core.Add(list__(key, val), env)
	
	return eval(key, env)
}

//	------------------------------------------------------------------------------------------------

// evset evals expression e.g. (set a 1)
// add a key value pair on top of environment like push on a stack
// 	e.g.
//		> (env) 		-> ((t t))
// 		> (var a 1) 	-> 1
// 		> (env) 		-> ((a 1) (t t))
//
func Var(exp, env *types.Cell, eval func(*types.Cell, *types.Cell) *types.Cell) *types.Cell {
	key := core.Cadr(exp)
	val := eval(core.Caddr(exp), env)
	
	env = core.Add(list__(key, val), env)
	
	return eval(key, env)
}

//	------------------------------------------------------------------------------------------------

// Let eval 'let, see example below
// 	e.g. 
//		> (let {key} {val} {body} )
//		> (let xs '(1 2 3) (car xs)) 	->  1
//		
func Let(exp, env *types.Cell, eval func(*types.Cell, *types.Cell) *types.Cell) *types.Cell {
	
	key := core.Cadr(exp)
	val := eval(core.Caddr(exp), env)
	aa := core.Cons(list__(key, val), env)
	
										// exp = (let x 1 (foo x))	
	ee := core.Car(core.Cdddr(exp))		//  ee = (foo x)

	fmt.Printf("Let - exp: %v\n", exp)
	fmt.Printf("Let - ee: %v\n", ee)
	fmt.Printf("Let - aa: %v\n", aa)
	
	return eval(ee, aa)
}

// Env prints environment for debug purpose
// e.g.
//		env = ((a 1) (b 1))
//
//		> (env) -> ((a 1) (b 1))
//
func Env(exp, env *types.Cell) *types.Cell {
	return env
}

//	------------------------------------------------------------------------------------------------

// Fun creates a function without name
//
// e.g.
//		(func (x) (car x))  -> func
//
func Fun(exp, env *types.Cell) *types.Cell {
	v := exp
	core.Tag(core.Car(exp), ID_FUNC)
	core.Tag(v, ID_FUNC)
	return v
}

//	------------------------------------------------------------------------------------------------

// Mac eval 'mac and create a macros in environment
//
// 	e.g.
//	 	(mac {name} {params}_{body})
//
func Mac(exp, env *types.Cell, eval func(*types.Cell, *types.Cell) *types.Cell) *types.Cell {

	name := core.Cadr(exp);
	
	// A macros is env func tagged as macro (Paul Graham - Arc)
	val := core.Tag(core.Cons(core.Tag(core.FUNC, ID_MAC), core.Cddr(exp)), ID_MAC)	// ToDo: ReThink - Tagging

	// Add at front without can of pointer to env
	// TODO: RETHINK
	core.Add(list__(name, val), env) 			
	
	return eval(name, env)
}

//	------------------------------------------------------------------------------------------------

// Do evals a list of expression and returns the last expression	
//
//	e.g.
//		(do
//			(list 1 2)
//			(no nil)    <-- last
//		)
//
func Do(exp, env *types.Cell, eval func(*types.Cell, *types.Cell) *types.Cell) *types.Cell {

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

	return doList(core.Cdr(exp), core.NIL, env)
}

//	------------------------------------------------------------------------------------------------

