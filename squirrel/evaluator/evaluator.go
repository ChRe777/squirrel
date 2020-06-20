package evaluator

import (
	"fmt"
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
	
	
	// Builtin operators, like and, not, append, ...
	// Storage operators like load, save
	// ... 
	
	for _, evaluator_ := range evaluators {
		res, err := evaluator_.Eval(e, a)
		if err == nil {
			return res
		}
	}
								
	// b.1) Functions e.g. (car '(1 2)) -> 1	
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
			case c.Equal(core.TYPE) 		: return core.Type(eval(builtin.Cadr(e), a))
			
			case c.Equal(core.COND ) 		: return evalCond(core.Cdr(e), a)
			case c.Equal(core.LIST) 		: return evalLst(core.Cdr(e), a)

			// Extra core
			//
			case c.Equal(core.BACKQUOTE) 	: return evalBackquote(e, a)	
			case c.Equal(core.DO)   		: return evalDo(e, a)
			case c.Equal(core.VAR ) 		: return evalVar(e, a)				
			case c.Equal(core.LET ) 		: return evalLet(e, a)					
			case c.Equal(core.DEF ) 		: return evalDef(e, a)				
			case c.Equal(core.MAC ) 		: return evalMac(e, a)				
			case c.Equal(core.FUNC)			: return evalFun(e, a)				
			case c.Equal(core.ENV ) 		: return evalEnv(e, a)	
						
			// Extra axioms in environment e.g. (no '()) -> t
			default: return evalFuncEnv(e, a)									// Builtin and others
		}
	} 

	// c) Function or macros call with parameter values		// e.g. (call {fn} {param-values})
	if builtin.Caar(e).Equal(core.FUNC) {
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

	key := core.Car(e)						// a = (foo  (func (x) (is x (quote nil))) )
	value := builtin.Assoc(key, a)			// value = (func (x) (is x (quote nil)))
	
	if value.IsErr(){
		return core.Err_("reference to undefined identifier: %v", key.Val) // TODO: Rename error message
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
func evalFuncOrMacCall(e, a *types.Cell) *types.Cell {

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
	
	fmt.Println("funcCall")
		
	keys := builtin.Cadar(e); vals := evalLst(core.Cdr(e), a)	
	
	//
	// keys: (x y . z), vals: (1 2 3 4)
	//
	fmt.Printf("funcCall - keys: %v, vals: %v", keys, vals)
	//
	//		
	//
	
	ee  := builtin.Caddar(e); aa := builtin.Append(builtin.Pair(keys, vals), a)		
	
	res := eval(ee, aa)			// will call func or expand backquotes and unquotes
	
	return res
}

func macCall(e, a *types.Cell) *types.Cell {
	
	fmt.Println("macCall")
	
	// Comparing ac-mac-call to ac-call shows why macros receive their arguments unevaluated. 
	// ac-mac-call applies the macro function to the arguments, 
	// while ac-call maps ac on the arguments before applying the function, 
	// causing the arguments to be evaluated.

	keys := builtin.Cadar(e); vals := core.Cdr(e)			

	fmt.Printf("macCall - keys: %v, vals: %v", keys, vals)
	
	// (when (is 'a 'a) 'a)
	// macCall
	// macCall - keys: (test . body), vals: ((is (quote a) (quote a))

	ee  := builtin.Caddar(e); aa := builtin.Append(builtin.Pair(keys, vals), a)		
	
	ff  := macex(ee, aa)		// if macro expand first and then evaluate
	
	res := eval(ff, a)
	
	return res
}


//	------------------------------------------------------------------------------------------------


 












