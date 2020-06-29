package evaluator

import (
	"fmt"
)

import (
	"github.com/mysheep/squirrel/builtin"
	"github.com/mysheep/squirrel/core"
	"github.com/mysheep/squirrel/interfaces"
)

const (
	REF_TO_UNDEFINED_ID = "reference to undefined identifier: %v"
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
//	
// Lexical Scoping - The Art of Interpretor (page 24)
// --------------------------------------------------
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

	fmt.Printf("eval - %65s <- exp | env-> %v \n", fmt.Sprintf("%v", exp), env)

	// Lisp dialects like Arc have env data type most languages don't:
	// symbols.  We've already seen one: + is env symbol.  Symbols don't
	// evaluate to themselves the way numbers and strings do.  They return
	// whatever value they've been assigned.

	var res *Cell
	var err error

	// Lookup Value of Symbol in Environment
	//
	if exp.IsAtom() {
		return evalAtom(exp, env)
	}
	
	// Try to evaluate functions in plugged in evaluator (like (load..) or (save ..))
	//
	for _, evaluator_ := range evaluators {
		res, err := evaluator_.Eval(exp, env, eval)
		if err == nil {
			return res
		}
	}

	// Try to eval builtin operators based on primitives (e.g. pair, no, not, append)
	//
	res, err = builtin.Eval(exp, env, eval)
	if err == nil {
		return res 
	}

	// Try to eval primitive core operators (e.g. car, cdr, cons, ... )
	//
	res, err = core.Eval(exp, env, eval) 
	if err == nil {
		return res
	}

	// User-defined functions or macros stored in evironment (e.g. ( .. (foo (func (x) (no x)))) ..)
	//
	return apply(exp, env)
}

//	------------------------------------------------------------------------------------------------

func apply(exp, env *Cell) *Cell {

	var args *Cell
	
	envFromFn := Cadddr(exp)
	
	fmt.Printf("apply - envFromFn: %v, env: %v \n", envFromFn, env)
	
	fnOrMac := eval(Car(exp), env); isMacro := isMac(fnOrMac)				
	
	if isMacro {								
		args = Cdr(exp)									// A macro receives arguments UNEVALUATED !!!
	} else {
		args = EvList(Cdr(exp), env)					// A func receives evaluated arguments
	}

	res := eval(Caddr(fnOrMac), Bind(Pair(Cadr(fnOrMac), args), envFromFn))
							
	if isMacro {										
		res = eval(res, env)									
	}
	
	return res
}

//	------------------------------------------------------------------------------------------------

func isMac(e *Cell) bool {
	return Car(e).IsTagged(builtin.ID_MAC)
}

//	------------------------------------------------------------------------------------------------
	

// evalAtom atom by return it or lookup in environment
//
//		Example "symbol table"
//		----------------------
//      
//   	symbol  | value
//      --------+-----------------------------------------------------------------------------------
//		dotted 	| (func (x) (cons x x))
//		map	   	| (func (fn xs)  (cond ((no xs) nil) ('t (cons (fn (car xs)) (map fn (cdr xs))))))
//		a		| 1
//		b		| 2
//		xs		| (1 2 3)
//	
func evalAtom(exp, env *Cell) *Cell {

	if exp.IsSymbol() {
	
		if exp.Equal(core.NIL) { return core.NIL }	// Hint: This can go into environment
		if exp.Equal(core.T)   { return core.T   }

		x := Value(exp, env)						// Lookup value of expression in env
		if x.IsErr() {
			return Err(REF_TO_UNDEFINED_ID, exp.Val) 			
		}
		return x
	}
	return exp
}

//	------------------------------------------------------------------------------------------------
