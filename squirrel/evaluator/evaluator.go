package evaluator

import (
//"fmt"
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

	if exp.IsAtom() {
		return evalAtom(exp, env)
	}

	if Caar(exp).Equal(core.FUNC) { 					// Func or macros call with arguments
		return evalFuncOrMacCall(exp, env)
	}
	
	for _, evaluator_ := range evaluators { 			// Plugin Evaluators like Storage operators like load, save
		res, err := evaluator_.Eval(exp, env)
		if err == nil {
			return res
		}
	}

	res, err := builtin.Eval(exp, env, eval) 			// Builtin like pair, no, not, append, ...
	if err == nil {
		return res
	}

	res, err = core.Eval(exp, env, eval) 				// Core like car, cadr, cons, cond, ..
	if err == nil {
		return res
	} else {
		return evalApplyFunc(exp, env)
	}
/*
	if c := Car(exp); c.IsAtom() {

		switch { 										// 7 core axioms - "The Roots of lisp" (McCarthy, Paul Graham
		case c.Equal(core.QUOTE):
			return core.Quote(exp)
		case c.Equal(core.ATOM):
			return core.Atom(eval(core.Cadr(exp), env))
		case c.Equal(core.IS):
			return core.Is(eval(core.Cadr(exp), env), eval(core.Caddr(exp), env))
		case c.Equal(core.CAR):
			return Car(eval(core.Cadr(exp), env))
		case c.Equal(core.CDR):
			return core.Cdr(eval(Cadr()(exp), env))
		case c.Equal(core.CONS):
			return core.Cons(eval(core.Cadr(exp), env), eval(core.Caddr(exp), env))
		case c.Equal(core.TYPE):
			return core.Type(eval(core.Cadr(exp), env))
		//case c.Equal(core.TAG) : return core.Tag(...)
		default:
			return evalApplyFunc(exp, env)
		}

		// Lexical Scoping - page 24 - The Art of Interpretor

		//   FUN                 ARGS
		// ( (func foo(x)(no x)) (1 2) )

		// (&PROCEDURE foo 1 2 3) -> fun = foo, args = (1 2 3)
		//
		// (DEFINE APPLY (FUN ARGS)
		//		(COND 	( (PRIMOP fun)  ->  (PRIMOP-APPLY fun args))		// if is primitive operator
		//				( (EQ (CAR fun)) '&PROCEDURE)						// when is tagged as procedure
		//				  -> (EVAL (CADDR fun)								// get 'fun from environment
		//						   (BIND (CADR fun) args (CADDDR fun))		// (bind foo (1 2 3)   )
		//					 )
		//		)
		//																	// vars -> (x y z)
		//	(DEFINE (BIND VARS ARGS ENV)									// env  -> ( ((x y z) (1 2 3)) ... env )
		//		(COND ( (=LENGTH VARS) (LENGTH ARGS))
		//				(CONS (CONS VARS ARGS) ENV)
		//			  )
		//			  (T (ERROR))
		//		)
		//	)

	}

	// APPLY
	// e.g.
	//	(  FUNC 				   ARGS  )
	//	( (func (x y) (cons x y)) '(1 2) )
	//
	if Caar(exp).Equal(core.FUNC) { // Func or macros call with arguments
		return evalFuncOrMacCall(exp, env)
	}
*/

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
		if exp.Equal(core.T)   { return core.T }

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

	key := Car(exp)    							// env   -> (foo  (func (x) (is x (quote nil))))
	val := Value(key, env) 						// value -> (func (x) (is x (quote nil)))

	if val.IsErr() {
		return Err("reference to undefined identifier: %v", key.Val) // TODO: Rename error message
	}

	args := Cdr(exp)        					// '(1 2)
	fnExp := Cons(val, args) 					// ((func (x y) (cons x y)) '(1 2))

	return eval(fnExp, env)
}



//	------------------------------------------------------------------------------------------------

// evalFunc evals function (or macro) calls with arguments
//	e.g.
//		( (func (x)  (car  x)) '(1 2) ) -> 1
//		( (func (x) `(cdr ,x)) '(1 2) ) -> 1			// Func tagged macros
func evalFuncOrMacCall(exp, env *Cell) *Cell {

	// From Arc Reference:
	// -------------------
	// Comparing ac-mac-call to ac-call shows why macros receive their arguments unevaluated.
	// ac-mac-call applies the macro function to the arguments,
	// while ac-call maps ac on the arguments before applying the function,
	// causing the arguments to be evaluated.

	funcCall := func(exp, env *Cell) *Cell {

		keys := Cadar(exp)                        			// keys -> (x y z)
		vals := List(exp, env)                     			// vals -> (1 2 3)
		kvps := Pair(keys, vals)                       		// kvps -> ((x 1) (y 2) (z 3))
		fnExp := Caddar(exp)

		env = Append(kvps, env) // ToDo: Check if this is ok in Multi-Threaded

		return eval(fnExp, env)
	}

	macroCall := func(exp, env *Cell) *Cell {

		keys := Cadar(exp)
		vals := Cdr(exp) 									// values are not evaluated in macros
		kvps := Pair(keys, vals)

		macExp := Caddar(exp)
		env2 := Append(kvps, env)                        	// replace e.g. (list {a} 1), (a 0) -> (list 0 1)
		fnExp := builtin.MacroExpand(macExp, env2, eval) 	// if macro expand first and then evaluate

		return eval(fnExp, env)
	}

	if Car(exp).IsTagged(builtin.ID_MAC) {
		return macroCall(exp, env)
	} else {
		return funcCall(exp, env)
	}
}




// ------------------------------------------------------------------------------------------------
// Shortcut or interfaces
//

func Append(kvps *types.Cell, env *Cell) *types.Cell {
	return builtin.Append(kvps, env)
}

func Car(x *Cell) *Cell {
	return core.Car(x)
}

func Caar(x *Cell) *Cell {
	return core.Caar(x)
}

func Cadr() func(e *Cell) *Cell {
	return core.Cadr
}

func Cadar(exp *Cell) *types.Cell {
	return core.Cadar(exp)
}

func Caddar(exp *Cell) *types.Cell {
	return core.Caddar(exp)
}

func Cdr(exp *Cell) *Cell {
	return core.Cdr(exp)
}

func Cons(x, y *Cell) *Cell {
	return core.Cons(x, y)
}

func Err(s string, a ...interface{}) *Cell {
	return core.Err_(s, a...)
}

func List(exp *Cell, env *Cell) *Cell {
	return builtin.List(Cdr(exp), env, eval)
}

func Pair(xs, ys *Cell) *Cell {
	return builtin.Pair(xs, ys)
}

func Value(key, env *Cell) *Cell {
	return builtin.Assoc(key, env)
}


