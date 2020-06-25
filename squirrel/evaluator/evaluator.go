package evaluator

import (
	"fmt"
)

import (
	"github.com/mysheep/squirrel/builtin"
	"github.com/mysheep/squirrel/core"
	"github.com/mysheep/squirrel/interfaces"
	"github.com/mysheep/squirrel/types"
)

// -------------------------------------------------------------------------------------------------
// Type aliases
//
type Cell = types.Cell

// -------------------------------------------------------------------------------------------------

var (
	evaluators []interfaces.Evaluator
)

func SetEvaluators(evs []interfaces.Evaluator) {
	evaluators = evs
}

// -------------------------------------------------------------------------------------------------

// Eval evaluates expression 'exp' with environment list 'env' and returns result
func Eval(exp, env *Cell) *Cell {
	return eval(exp, env)
}

// -------------------------------------------------------------------------------------------------

func eval(exp, env *Cell) *Cell {

	// Lisp dialects like Arc have env data type most languages don't:
	// symbols.  We've already seen one: + is env symbol.  Symbols don't
	// evaluate to themselves the way numbers and strings do.  They return
	// whatever value they've been assigned.

	var res *Cell
	var err error

	// Lookup Value of Symbol in Environment
	if exp.IsAtom() {
		return evalAtom(exp, env)
	}

	// Plugin Operators bases on core and primitives
	for _, evaluator_ := range evaluators { // Plugin Evaluators like Storage operators like load, save
		res, err := evaluator_.Eval(exp, env)
		if err == nil {
			return res
		}
	}

	// Builtin Operators based on primitives
	res, err = builtin.Eval(exp, env, eval) // Builtin like pair, no, not, append, ...
	if err == nil {
		return res // found
	}

	// Primitive Core Operators (car, cdr, cons, ... )
	res, err = core.Eval(exp, env, eval) 	// Core like car, cadr, cons, cond, ..
	if err == nil {
		return res // found
	}

	// User-defined functions
	fun  := Eval  (Car(exp), env)			//(foo 1 2) -> (func (x y) (cons x y))
	args := EvList(Cdr(exp), env)			//(foo 1 2) -> (1 2)
	return apply(fun,
				 args,		// TODO: Macro receive there values UNEVALUATED !!!
				 env)

	/*else {
		// APPLY
		if Caar(exp).Equal(core.FUNC) { // Func or macros call with arguments
			return evalFuncOrMacCall(exp, env)
		}
		// (foo 1 2) ->
		// ((func (x y) (cons x y)) '(1 2))
		return evalApplyFunc(exp, env) // User-defined funcs created by (def ...)
	} // (foo 				    '(1 2))
	// ((func (x y) (cons x y)) '(1 2))
	*/

}


	// Exp										// Env
	// ---------------------------------------------------------------------------------------------
	// (def foo (x y) (cons x y))				// [(foo (func (x y) (cons x y))) (a 3) (b 4) ...]
	// (foo 				      1 2 )			// [(foo (func (x y) (cons x y))) (a 3) (b 4) ...]
	// ((func (x y) (cons x y))  (1 2))			// [(foo (func (x y) (cons x y))) (a 3) (b 4) ...]
	// (cons x y)								// [((x y) 1 2) (foo (func (x y) (cons x y))) ...]
	// (cons 1 2)								// [((x y) 1 2) (foo (func (x y) (cons x y))) ...]
	// (1 . 2)									// [(foo (func (x y) (cons x y))) (a 3) (b 4) ...]
		
	// Lexical Scoping - page 24 - The Art of Interpretor

	//   FUN                 ARGS
	// ( (func foo(x)(no x)) (1 2) )

	// (&PROCEDURE foo 1 2 3) -> fun = foo, args = (1 2 3)

	// (DEFINE (EVAL EXP ENV)
	//		(COND
	//			((ATOM EXP)
	//				(COND (NUMBERP ...)
	//				(T (VALUE EXP ENV)))
	//			((EQ (CAR EXP) 'QUOTE)
	//				(CADR EXP)
	//			((EQ (CAR EXP) 'LAMBDA
	//				(LIST '&PROCEDURE (CADR EXP) (CADDR EXP) ENV)
	//			(T (APPLY (EVAL (CAR EXP) ENV)
	//					  (EVLIS (CDR EXP) ENV))))
	//
	// (DEFINE APPLY (FUN ARGS)
	//		(COND
	//			((PRIMOP fun)  ->  (PRIMOP-APPLY fun args))				// if is primitive operator
	//			((EQ (CAR fun)) '&PROCEDURE)							// when is tagged as procedure
	//				(EVAL (CADDR fun)									// get 'fun from environment
	//					(BIND (CADR fun) args (CADDDR fun))				// (bind foo (1 2 3)   )
	//				)
	//		)
	//																	// vars -> (x y z)
	//	(DEFINE (BIND VARS ARGS ENV)									// env  -> ( ((x y z) (1 2 3)) ... env )
	//		(COND ( (=LENGTH VARS) (LENGTH ARGS))
	//				(CONS (CONS VARS ARGS) ENV)
	//			  )
	//			  (T (ERROR))
	//		)
	//	)


func apply(fun, args, env *Cell) *Cell {

	vars := Cadr(fun)
	kvps := Pair(vars, args)

	return eval(Caddr(fun),			// (cons x y)
				Bind(kvps, env))	// ((x 1) (y 2) ... (foo (func (x y) (cons x y)) ...)
}

//	------------------------------------------------------------------------------------------------

// evalAtom evals atom from environment
// e.g.
//		> (env) 	-> ((a 1)(b 1))
//  	> a 		-> 1
//  	> b 		-> 2
//
func evalAtom(exp, env *Cell) *Cell {

	if exp.IsSymbol() {
	
		if exp.Equal(core.NIL) { return core.NIL }	// This can go into environment
		if exp.Equal(core.T)   { return core.T   }

		x := Value(exp, env)
		if x.IsErr() {
			return Err("reference to undefined identifier: %v", exp.Val) // TODO: Rename error message
		}
		return x
	}
	return exp
}

//	------------------------------------------------------------------------------------------------

// evalApplyFunc eval func from environment
// 	e.g.
//		> (env) 	-> (foo (func (x) (no x)))
//  	> (foo nil)
//
func evalApplyFunc(exp, env *Cell) *Cell {

	fmt.Printf("evalApplyFunc - exp: %v, env :%v \n", exp, env)

	key := Car(exp)    							// env   -> (foo  (func (x) (is x (quote nil))))
	val := Value(key, env) 						// value -> (func (x) (is x (quote nil)))

	if val.IsErr() {
		return Err("reference to undefined identifier: %v", key.Val) // TODO: Rename error message
	}

	args := Cdr(exp)        					// '(1 2)
	fnExp := Cons(val, args) 					// (cons 1 2)

	return eval(fnExp, env)
}

//	------------------------------------------------------------------------------------------------

// evalFunc evals function (or macro) calls with arguments
//	e.g.
//		( (func (x)  (car  x)) '(1 2) ) -> 1
//		( (func (x) `(cdr ,x)) '(1 2) ) -> 1			// Func tagged macros
func evalFuncOrMacCall(exp, env *Cell) *Cell {

	fmt.Printf("evalFuncOrMacCall exp: %v, env: %v\n", exp, env)

	// From Arc Reference:
	// -------------------
	// Comparing ac-mac-call to ac-call shows why macros receive their arguments unevaluated.
	// ac-mac-call applies the macro function to the arguments,
	// while ac-call maps ac on the arguments before applying the function,
	// causing the arguments to be evaluated.

	funcCall := func(exp, env *Cell) *Cell {

		keys := Cadar(exp)                        			// keys -> (x y z)
		vals := EvList(Cdr(exp), env)                     	// vals -> (1 2 3), values are evaluated
		
		fmt.Printf("funcCall - keys: %v, vals: %v \n", keys, vals)
		
		
		kvps := Pair(keys, vals)                       		// kvps -> ((x 1) (y 2) (z 3))
		fn_exp := Caddar(exp)

		// Append is called Bind
		env = Bind(kvps, env) 								// SLOT in SCHEME: ( ((x y z) 1 2 3) ...)

		fmt.Printf("funcCall - fn_exp: %v, env: %v \n", fn_exp, env)

		return eval(fn_exp, env)
	}

	macroCall := func(exp, env *Cell) *Cell {

		keys := Cadar(exp)
		vals := Cdr(exp) 									// values are not evaluated in macros
		kvps := Pair(keys, vals)

		mac_exp := Caddar(exp)
		env = Bind(kvps, env)
		fn_exp := MacroExpand(mac_exp, env, eval) 			// if macro expand first and then evaluate

		fmt.Printf("macroCall - fnExp: %v, env: %v \n", fn_exp, env)

		return eval(fn_exp, env)
	}

	fmt.Printf("evalFuncOrMacCall - exp: %v, env: %v \n", exp, env)

	if Car(exp).IsTagged(builtin.ID_MAC) {
		return macroCall(exp, env)
	} else {
		return funcCall(exp, env)
	}
}




// ------------------------------------------------------------------------------------------------
// Shortcut or interfaces
//

func MacroExpand(exp, env *types.Cell, eval func(*types.Cell, *types.Cell) *types.Cell) *types.Cell {
	return builtin.MacroExpand(exp, env, eval)
}

func Bind(kvps *Cell, env *Cell) *Cell {
	return builtin.Append(kvps, env)
}

func Sym(s string) *Cell {
	return core.Sym_(s)
}

func Cons(x, y *Cell) *Cell {
	return core.Cons(x, y)
}

func Car(x *Cell) *Cell {
	return core.Car(x)
}

func Caar(x *Cell) *Cell {
	return core.Caar(x)
}

func Cadr(x *Cell) *Cell {
	return core.Cadr(x)
}

func Cadar(exp *Cell) *Cell {
	return core.Cadar(exp)
}

func Caddr(exp *Cell) *Cell {
	return core.Caddr(exp)
}

func Caddar(exp *Cell) *Cell {
	return core.Caddar(exp)
}

func Cadddr(exp *Cell) *Cell {
	return core.Cadddr(exp)
}

func Cdr(exp *Cell) *Cell {
	return core.Cdr(exp)
}

func Err(s string, a ...interface{}) *Cell {
	return core.Err_(s, a...)
}

func EvList(exp *Cell, env *Cell) *Cell {
	return builtin.List(exp, env, eval)
}

func Pair(xs, ys *Cell) *Cell {
	return builtin.Pair(xs, ys)
}

func Value(key, env *Cell) *Cell {
	return builtin.Assoc(key, env)
}


