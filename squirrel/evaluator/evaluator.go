package evaluator

import (
	"fmt"
)

import (
	"github.com/mysheep/squirrel/builtin"
	"github.com/mysheep/squirrel/core"
	"github.com/mysheep/squirrel/interfaces"
)

// -------------------------------------------------------------------------------------------------

var (
	evaluators []interfaces.Evaluator
)

func SetEvaluators(evs []interfaces.Evaluator) {
	evaluators = evs
}

// -------------------------------------------------------------------------------------------------

// Eval interface evaluates expression (exp) with a symbol table 
// called environment (env) and returns evaluated result
func Eval(exp, env *Cell) *Cell {
	return eval(exp, env)
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

	// try to eval builtin operators based on primitives (e.g. pair, no, not, append)
	res, err = builtin.Eval(exp, env, eval)
	if err == nil {
		return res // found
	}

	// try to eval primitive core operators (e.g. car, cdr, cons, ... )
	res, err = core.Eval(exp, env, eval) 
	if err == nil {
		return res // found
	}

	// User-defined functions or macros stored in evironment (e.g. ( .. (foo (func (x) (no x)))) ..)
	return apply(exp, env)
}

func apply(exp, env *Cell) *Cell {

	isMac := func(e *Cell) bool {
		return Car(e).IsTagged(builtin.ID_MAC)
	}

	// User-defined functions					// for example (foo  1 2) evaluates
	fun := eval(Car(exp), env)					// to (func (x y) (cons x y))
	
	var args *Cell	
	if isMac(fun) {								
		args = Cdr(exp)							// A macros receives their values UNEVALUATED !!!
		fmt.Printf("apply - args: %v \n", args)
	} else {
		args = EvList(Cdr(exp), env)			// A func received evaluated vars e.g. (a b) -> (1 2)
	}

	vars := Cadr(fun)

	exp = Caddr(fun)
	env = Bind(Pair(vars, args), env)
	
	res := eval(exp, env)						// (cons x y) or (backquote (cons (unquote x)(unquote y)))
												// ((x 1) (y 2) ... (foo (func (x y) (cons x y)) ...)
	if isMac(fun) {								// First eval only expand - fill in the args
		res = eval(res, env)					// Now evaluated the expanded expression
	}

	return res
}

//	------------------------------------------------------------------------------------------------

// evalAtom atom by return it or lookup in environment
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







