package evaluator

import (
	//"fmt"
)

import (
	"github.com/mysheep/squirrel/types"
	"github.com/mysheep/squirrel/interfaces"
	"github.com/mysheep/squirrel/evaluator/core"
	"github.com/mysheep/squirrel/evaluator/builtin"
)

var (
	evaluators 	[]interfaces.Evaluator
)

// -------------------------------------------------------------------------------------------------

func SetEvaluators(evs []interfaces.Evaluator) {
	evaluators = evs
}

// -------------------------------------------------------------------------------------------------

// Eval evals expression e with environment env and returns result
func Eval(e, env *types.Cell) *types.Cell {
	return eval(e, env)
}

// -------------------------------------------------------------------------------------------------

func eval(e, a *types.Cell) *types.Cell {

 	// Lisp dialects like Arc have a data type most languages don't:
	// symbols.  We've already seen one: + is a symbol.  Symbols don't
	// evaluate to themselves the way numbers and strings do.  They return
	// whatever value they've been assigned.

	// a) Atom e.g. "foo" -> "foo"
	if e.IsAtom() {
	
		// LISP boolean values
		//
		if e.Equal(core.NIL) {
			return core.NIL
		}
		if e.Equal(core.T) {
			return core.T
		}
	
		return evalAtom(e, a)
	} 
	
	
	// Plugin Evaluators like Storage operators like load, save
	// 
	for _, evaluator_ := range evaluators {
		res, err := evaluator_.Eval(e, a)
		if err == nil {
			return res
		}
	}
	
	// Core line car, cdr, cons, cond, ..
	//
	//res, err := core.Eval(e, a, eval)
	//if err == nil {
	//	return res
	//}
	
	// Builtin like pair, no, not, append, ...
	//
	res, err := builtin.Eval(e, a, eval)
	if err == nil {
		return res
	}
	
	// Functions from environment like ( (foo (func (x)(no x))) (bar (func (y)(not y))) ... )
	//
	//res, err := envFunc.Eval(e, a, eval)
	//if err == nil {
	//	return res
	//}
								
	// b.1) Functions e.g. (car '(1 2)) -> 1	
	if c := core.Car(e); c.IsAtom() {
	
		switch {	
		
			// 7 core axioms - "The Roots of lisp" (McCarthy, Paul Graham			//
			case c.Equal(core.QUOTE) 		: return core.Quote(e) 
			case c.Equal(core.ATOM ) 		: return core.Atom(eval(core.Cadr(e), a))
			case c.Equal(core.IS   ) 		: return core.Is  (eval(core.Cadr(e), a), eval(core.Caddr(e), a))
			case c.Equal(core.CAR  ) 		: return core.Car (eval(core.Cadr(e), a))
			case c.Equal(core.CDR  ) 		: return core.Cdr (eval(core.Cadr(e), a))
			case c.Equal(core.CONS ) 		: return core.Cons(eval(core.Cadr(e), a), eval(core.Caddr(e), a))			
			// New extra core
			case c.Equal(core.TYPE) 		: return core.Type(eval(core.Cadr(e), a))			
			//case c.Equal(core.TAG) 		: return core.Tag(...)
									
			// Extra axioms in environment e.g. (no '()) -> t
			default: return evalFuncEnv(e, a)									// Builtin and others
		}
	} 

	// c) Function or macros call with parameter values		// e.g. (call {fn} {param-values})
	if core.Caar(e).Equal(core.FUNC) {
		return evalFuncOrMacCall(e, a)
	}
	
	return core.Err_("Wrong expression")
}

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
			return core.Err_("reference to undefined identifier: %v", e.Val) // TODO: Rename error message
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

	key := core.Car(e)						// a     = (foo  (func (x) (is x (quote nil))) )
	value := builtin.Assoc(key, a)			// value = (func (x) (is x (quote nil)))
	
	//fmt.Printf("evalFuncEnv e: %v, a: %v \n", e, a)

	if value.IsErr(){
		return core.Err_("reference to undefined identifier: %v", key.Val) // TODO: Rename error message
	}
	
	// Function call with parameter values
	ee := core.Cons(value, core.Cdr(e))			// ((func (x) (is x (quote nil))) '(1 2))
	
	//fmt.Printf("evalFuncEnv ee: %v, a: %v \n", ee, a)
		
	return eval(ee, a)
}

//	------------------------------------------------------------------------------------------------

// evalFunc evals function (or macro) calls with arguments
//	e.g.
//		( (func (x)  (car  x)) '(1 2) ) -> 1
//		( (func (x) `(cdr ,x)) '(1 2) ) -> 1	// Func tagged macros
func evalFuncOrMacCall(e, a *types.Cell) *types.Cell {

	//fmt.Printf("evalFuncOrMacCall e: %v, a: %v \n", e, a)

	// Comparing ac-mac-call to ac-call shows why macros receive their arguments unevaluated. 
	// ac-mac-call applies the macro function to the arguments, 
	// while ac-call maps ac on the arguments before applying the function, 
	// causing the arguments to be evaluated.
		
	if isMac(e) {
		return macCall(e, a)
	}

	return funcCall(e, a)
}

func funcCall(e, a *types.Cell) *types.Cell {
			
	keys := core.Cadar(e); vals := builtin.List(core.Cdr(e), a, eval)	
	
	//fmt.Printf("funcCall keys: %v, vals: %v \n", keys, vals)
	
	//
	// keys: (x y . z), vals: (1 2 3 4)
	//
	
	kvs := builtin.Pair(keys, vals)
	
	//
	// kvs: ((x 1) (y 2) (z (3 4))
	//
	
	ee  := core.Caddar(e); aa := builtin.Append(kvs, a)		
	
	res := eval(ee, aa)			// will call func or expand backquotes and unquotes
	
	return res
}

func macCall(e, a *types.Cell) *types.Cell {
		
	// Comparing ac-mac-call to ac-call shows why macros receive their arguments unevaluated. 
	// ac-mac-call applies the macro function to the arguments, 
	// while ac-call maps ac on the arguments before applying the function, 
	// causing the arguments to be evaluated.

	keys := core.Cadar(e); vals := core.Cdr(e)	
	
	//fmt.Printf("macCall keys: %v, vals: %v \n", keys, vals)
	
	kvs := builtin.Pair(keys, vals)		

	ee  := core.Caddar(e); aa := builtin.Append(kvs, a)	
	//
	// xs = (quote (1 2 3))	
	// (unquote_splicing xs)
	// ... quote (1 2 3) ...
	// (list quote (1 2 3))
	//
	ff  := builtin.MacEx(ee, aa, eval)		// if macro expand first and then evaluate
	
	//fmt.Printf("macCall ff: %v \n", ff)
	
	// macCall keys: (xs), vals: ((quote (1 2 3))) 
	// macCall ff: (list quote (1 2 3)) 
	
	res := eval(ff, a)
	
	return res
}


//	------------------------------------------------------------------------------------------------

// isMac checks if car(e) is tagged as macro
// e.g.
//		e = (func (x) (no x))
//		car(e) -> func (Tagged with mac) 
func isMac(e *types.Cell) bool {
	return core.Car(e).IsTagged(builtin.ID_MAC)
}










