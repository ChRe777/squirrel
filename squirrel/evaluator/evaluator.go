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

var (
	evaluators []interfaces.Evaluator
)

func SetEvaluators(evs []interfaces.Evaluator) {
	evaluators = evs
}

// -------------------------------------------------------------------------------------------------

// Eval evaluates expression 'exp' with environment list 'env' and returns result
func Eval(exp, env *types.Cell) *types.Cell {
	return eval(exp, env)
}

// -------------------------------------------------------------------------------------------------

func eval(exp, env *types.Cell) *types.Cell {

	// Lisp dialects like Arc have env data type most languages don't:
	// symbols.  We've already seen one: + is env symbol.  Symbols don't
	// evaluate to themselves the way numbers and strings do.  They return
	// whatever value they've been assigned.

	if exp.IsAtom() {
		return evalAtom(exp, env)
	}

	for _, evaluator_ := range evaluators { // Plugin Evaluators like Storage operators like load, save
		res, err := evaluator_.Eval(exp, env)
		if err == nil {
			return res
		}
	}

	res, err := builtin.Eval(exp, env, eval) // Builtin like pair, no, not, append, ...
	if err == nil {
		return res
	}

	if c := core.Car(exp); c.IsAtom() {

		switch { // 7 core axioms - "The Roots of lisp" (McCarthy, Paul Graham
		case c.Equal(core.QUOTE):
			return core.Quote(exp)
		case c.Equal(core.ATOM):
			return core.Atom(eval(core.Cadr(exp), env))
		case c.Equal(core.IS):
			return core.Is(eval(core.Cadr(exp), env), eval(core.Caddr(exp), env))
		case c.Equal(core.CAR):
			return core.Car(eval(core.Cadr(exp), env))
		case c.Equal(core.CDR):
			return core.Cdr(eval(core.Cadr(exp), env))
		case c.Equal(core.CONS):
			return core.Cons(eval(core.Cadr(exp), env), eval(core.Caddr(exp), env))
		case c.Equal(core.TYPE):
			return core.Type(eval(core.Cadr(exp), env))
			//case c.Equal(core.TAG) 		: return core.Tag(...)

		default:
			return evalFuncEnv(exp, env) // Builtin and others
		}
	}

	if core.Caar(exp).Equal(core.FUNC) { // Func or macros call with arguments
		return evalFuncOrMacCall(exp, env)
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
func evalAtom(exp, env *types.Cell) *types.Cell {

	if exp.IsSymbol() {
		if exp.Equal(core.NIL) {
			return core.NIL
		}
		if exp.Equal(core.T) {
			return core.T
		}

		x := builtin.Assoc(exp, env)
		if x.IsErr() {
			return core.Err_("reference to undefined identifier: %v", exp.Val) // TODO: Rename error message
		}
		return x
	}
	return exp
}

//	------------------------------------------------------------------------------------------------

// evalFuncEnv eval func from environment
// 	e.g.
//		> (env) 	-> (foo (func (x) (no x)))
//  	> (foo nil)
//
func evalFuncEnv(exp, env *types.Cell) *types.Cell {

	key := core.Car(exp)           // env   -> (foo  (func (x) (is x (quote nil))))
	val := builtin.Assoc(key, env) // value -> (func (x) (is x (quote nil)))

	if val.IsErr() {
		return core.Err_("reference to undefined identifier: %v", key.Val) // TODO: Rename error message
	}

	fnExp := core.Cons(val, core.Cdr(exp)) // ((func (x) (is x (quote nil))) '(1 2))

	return eval(fnExp, env)
}

//	------------------------------------------------------------------------------------------------

// evalFunc evals function (or macro) calls with arguments
//	e.g.
//		( (func (x)  (car  x)) '(1 2) ) -> 1
//		( (func (x) `(cdr ,x)) '(1 2) ) -> 1	// Func tagged macros
func evalFuncOrMacCall(exp, env *types.Cell) *types.Cell {

	// From Arc Reference:
	// -------------------
	// Comparing ac-mac-call to ac-call shows why macros receive their arguments unevaluated.
	// ac-mac-call applies the macro function to the arguments,
	// while ac-call maps ac on the arguments before applying the function,
	// causing the arguments to be evaluated.

	funcCall := func(exp, env *types.Cell) *types.Cell {

		keys := core.Cadar(exp)                        // keys -> (x y z)
		vals := builtin.List(core.Cdr(exp), env, eval) // vals -> (1 2 3)
		kvps := builtin.Pair(keys, vals)               // kvps -> ((x 1) (y 2) (z 3))

		fnExp := core.Caddar(exp)
		env = builtin.Append(kvps, env) // ToDo: Check if this is ok in Multi-Threaded

		return eval(fnExp, env)
	}

	macroCall := func(exp, env *types.Cell) *types.Cell {

		keys := core.Cadar(exp)
		vals := core.Cdr(exp) // values are not evaluated in macros
		kvps := builtin.Pair(keys, vals)

		macExp := core.Caddar(exp)
		env2 := builtin.Append(kvps, env)                // replace e.g. (list {a} 1), (a 0) -> (list 0 1)
		fnExp := builtin.MacroExpand(macExp, env2, eval) // if macro expand first and then evaluate

		return eval(fnExp, env)
	}

	if core.Car(exp).IsTagged(builtin.ID_MAC) {
		return macroCall(exp, env)
	} else {
		return funcCall(exp, env)
	}
}

//	------------------------------------------------------------------------------------------------
