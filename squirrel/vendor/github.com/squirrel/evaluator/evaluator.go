package evaluator

import (
	"fmt"
)

import (
	"github.com/squirrel/types"
	"github.com/squirrel/core"
	"github.com/squirrel/builtin"
)

// Eval evals expression e with environment env and returns result
func Eval(e, env *types.Cell) *types.Cell {
	return eval(e, env)
}

func eval(e, a *types.Cell) *types.Cell {

 	// Lisp dialects like Arc have a data type most languages don't:
	// symbols.  We've already seen one: + is a symbol.  Symbols don't
	// evaluate to themselves the way numbers and strings do.  They return
	// whatever value they've been assigned.

	// a) Atom e.g. "foo" -> "foo"
	if e.IsAtom() {
		return evalAtom(e, a)
	} 
	
	// b) Functions e.g. (car '(1 2)) -> 1	
	if c := core.Car(e); c.IsAtom() {
		
		switch {	
		
			// 7 core axioms - "The Roots of lisp" (McCarthy, Paul Graham)
			//
			case c.Equal(core.QUOTE) 		: return core.Quote(e) 
			case c.Equal(core.ATOM ) 		: return core.Atom(eval(builtin.Cadr(e), a))
			case c.Equal(core.IS   ) 		: return core.Is  (eval(builtin.Cadr(e), a), eval(builtin.Caddr(e), a))
			case c.Equal(core.CAR  ) 		: return core.Car (eval(builtin.Cadr(e), a))
			case c.Equal(core.CDR  ) 		: return core.Cdr (eval(builtin.Cadr(e), a))
			case c.Equal(core.CONS ) 		: return core.Cons(eval(builtin.Cadr(e), a), eval(builtin.Caddr(e), a))			
			case c.Equal(core.COND ) 		: return evalCond(core.Cdr(e), a)
			
			
			
			
			// Macros			
			case c.Equal(core.BACKQUOTE) 	: return evalBackquote(e, a) 		// Used for macros in combination with unquote
			// Extra commands
			//
			case c.Equal(core.VAR ) 		: return evalVar(e, a)				// Tests
			case c.Equal(core.LET ) 		: return evalLet(e, a)				// Tests	
			case c.Equal(core.DEF ) 		: return evalDef(e, a)				// Tests
			case c.Equal(core.MAC ) 		: return evalMac(e, a)				// Tests
			case c.Equal(core.FUNC)			: return evalFunc(e, a)				// Tests
			
			case c.Equal(core.ENV ) 		: return evalEnv(e, a)				// Tests
			case c.Equal(core.LIST) 		: return evalList(core.Cdr(e), a)
	
			
			// 3 extra core axioms from Arc (Paul Graham)
			//
			//case c.Equal(core.TAG  ): return core.Tag  (eval(builtin.cadr(e), a), eval(builtin.caddr(e), a))
			//case c.Equal(core.TYPE0): return core.Type0(eval(builtin.cadr(e), a))
			//case c.Equal(core.REP  ): return core.Rep  (eval(builtin.cadr(e), a))		

			// Extra axioms in environment e.g. (no '()) -> t
			default: return evalFuncEnv(e, a)									// Builtin and others
		}
	}
/*	
	// c) Labels calls
	if builtin.Caar(e).Equal(core.LABEL) {
		return evlabel(e, a)
	}
*/ 

	// d) Function call with parameter values		// e.g. (call {fn} {values})
	if builtin.Caar(e).Equal(core.FUNC) {
		return evalFuncCall(e, a)
	}
			
	return core.Err("Wrong expression")
}

//	------------------------------------------------------------------------------------------------

			
// c) Labels calls 
//		e.g. 
//			( (label cadr (func (x) (car (cdr x))) ) (cadr '(1 2 3)) ) -> 2
//			
//			( (func (x) (car (cdr x))) )
//			(
//				(x '(1 2 3))
//			)         
//
// A "label" expression is evaluated by pushing a list of the function name
// and the function itself, onto the environment, and then calling eval on an
// expression with the inner lambda expression substituted for the label expression.
/*
func evlabel(e, a *types.Cell) *types.Cell {
	label := builtin.Cadar(e); fn := builtin.Caddar(e)
	ee := core.Cons(builtin.Caddar(e), core.Cdr(e))		
	aa := core.Cons(builtin.List_(label, fn), a)  // ( (no (func (x) (eq x nil)) (a 1) (b 2) )
				
	return eval(ee, aa)
}
*/

//	------------------------------------------------------------------------------------------------

// evalAtom evals atom from environment
// e.g. 
//		> (env) 	-> ((a 1)(b 1))
//  	> a 		-> 1
//  	> b 		-> 2
//
func evalAtom(e, a *types.Cell) *types.Cell {

	if e.IsSymbol() {	
		x := builtin.Assoc(e, a) // ToDo: Hash-table // nil means also not found !!!	
		if x.IsErr() {
			return core.Err("reference to undefined identifier: %v", e) // TODO: Rename error message
		}
		return x
	}
	return e

}

//	------------------------------------------------------------------------------------------------

// evalFuncEnv eval func from environment
// 	e.g.
//		> (env) 	-> (foo (func (x) (no x)))
//  	> (foo nil)
//
func evalFuncEnv(e, a *types.Cell) *types.Cell {

	key := core.Car(e)							// a = (foo  (func (x) (is x (quote nil))) )
	
	// 1. First look in builtin hash table
	keyStr := fmt.Sprintf("%v", key)			// key   = 'foo, keyStr = "foo"
	value, found := builtin.Fns[keyStr]			// value = (func (x) (is x (quote nil)))
	
	// 2. Look in environment association list
	if found == false {
		value = builtin.Assoc(key, a)			// value = (func (x) (is x (quote nil)))
	}
	
	if value.IsErr() && found == false {
		return core.Err("reference to undefined identifier: %v", key) // TODO: Rename error message
	}
	
	// Function call with parameter values
	ee := core.Cons(value, core.Cdr(e))			// ((func (x) (is x (quote nil))) '(1 2))
		
	return eval(ee, a)
}

//	------------------------------------------------------------------------------------------------

// evalFunc evals function (or macro) calls with arguments
//	e.g.
//		( (func (x)  (car  x)) '(1 2) ) -> 1
//		( (func (x) `(cdr ,x)) '(1 2) ) -> 1	// Func tagged macros
func evalFuncCall(e, a *types.Cell) *types.Cell {

	key := builtin.Cadar(e); val := evalList(core.Cdr(e), a)			
	ee  := builtin.Caddar(e); aa := builtin.Append(builtin.Pair(key, val), a)		
			
	res := eval(ee, aa)			// will call func or expand backquotes and unquotes
		
	if isMac(e) {	
		return eval(res, aa)	// and then if macros call func
	}
	
	return res
}

//	------------------------------------------------------------------------------------------------
















