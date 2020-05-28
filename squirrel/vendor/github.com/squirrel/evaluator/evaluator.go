package evaluator

import (
	"fmt"
)

import (
	"github.com/squirrel/types"
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
		return evatom(e, a)
	} 
	
	// b) Functions e.g. (car '(1 2)) -> 1	
	c := car(e)
	if  c.IsAtom() {
		switch {	
		
			// 7 core axioms - "The Roots of lisp" (McCarthy, Paul Graham)
			//
			case c.Equal(builtin.QUOTE): return quote(e) 
			case c.Equal(builtin.ATOM ): return atom (eval(cadr(e), a))
			case c.Equal(builtin.EQ   ): return eq   (eval(cadr(e), a), eval(caddr(e), a))
			case c.Equal(builtin.CAR  ): return car  (eval(cadr(e), a))
			case c.Equal(builtin.CDR  ): return cdr  (eval(cadr(e), a))
			case c.Equal(builtin.CONS ): return cons (eval(cadr(e), a), eval(caddr(e), a))
			case c.Equal(builtin.COND ): return evcon(cdr(e), a)
									
			// 7 extension functions from "The Roots of Lisp" (McCarthy, Paul Graham)
			//
			case c.Equal(builtin.NO    ): return no    (eval(cadr(e), a))
			case c.Equal(builtin.NOT   ): return not   (eval(cadr(e), a))
			case c.Equal(builtin.AND   ): return and   (eval(cadr(e), a), eval(caddr(e), a))
			case c.Equal(builtin.PAIR  ): return pair  (eval(cadr(e), a), eval(caddr(e), a))
			case c.Equal(builtin.LIST  ): return list  (eval(cadr(e), a), eval(caddr(e), a))
			case c.Equal(builtin.ASSOC ): return assoc (eval(cadr(e), a), eval(caddr(e), a))
			case c.Equal(builtin.APPEND): return append(eval(cadr(e), a), eval(caddr(e), a))

			// 3 extra core axioms from Arc (Paul Graham)
			//
			case c.Equal(builtin.TAG  ): return tag  (eval(cadr(e), a), eval(caddr(e), a))
			case c.Equal(builtin.TYPE0): return type0(eval(cadr(e), a))
			case c.Equal(builtin.REP  ): return rep  (eval(cadr(e), a))
			
			// TEST - REFACTOR
			case c.Equal(builtin.Sym("var")): return evset(e, a)		
			case c.Equal(builtin.Sym("env")): return evenv(e, a)
			case c.Equal(builtin.Sym("let")): return evlet(e, a)
			case c.Equal(builtin.Sym("def")): return evdef(e, a)
			case c.Equal(builtin.Sym("mac")): return evmac(e, a)		
				
			// TEST
			// e.g. (load "code.sqr")
			// case c.Equal(builtin.Sym("load")): return evload(e, a)
			
			// Test - MACRO SUPPORT
			case c.Equal(builtin.BACKQUOTE): return backquote(e, a) 
			
			// Extra axioms in environment e.g. (no '()) -> t
			default: return evfunc(e, a)
		}
	}
	
	// (var mapn (func (f ys) (cond (((no ys) nil)('t  (cons 	(f      (first ys)) (mapn f (rest  ys))))))))
				
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
	if caar(e).Equal(builtin.LABEL) {
		
		label := cadar(e); fn := caddar(e)
		
		ee := cons(caddar(e), cdr(e))		
		aa := cons(list(label, fn), a)  // ( (no (func (x) (eq x nil)) (a 1) (b 2) )
				
		return eval(ee, aa)
	} 
	
	// d) Function calls 
	//           (         f          params ) -> ?
	//		e.g. ( (func (x) (car x)) '(1 2) ) -> 1
	if caar(e).Equal(builtin.FUNC) {
	
		k := cadar(e); v := evlis(cdr(e), a)	
		
		ee := caddar(e)
		aa := append(pair(k, v), a)		
				
		r := eval(ee, aa)		// will expand backquotes and unquotes
		if isMac(e) {
			return eval(r, aa)
		}
		
		return r
	}
		
	return builtin.Err("Wrong expression")
}

// isMac checks if caar(e) is tagged as macro
// e.g.
//		e = ((func (x) (no x))
//		caar(e) -> func 
func isMac(e *types.Cell) bool {
	return caar(e).IsTagged(builtin.ID_MAC)
}

//  (def {name} {params} {body})
//  (var {name} (func {params} {body}) )
func evdef(e, a *types.Cell) *types.Cell {
	name := cadr(e); params_body := cddr(e)
	k := name; v := cons(builtin.FUNC, params_body)
	builtin.Tag(v, builtin.ID_FUNC)
	a = addEnv(list(k, v), a)
	return eval(k, a)
}

//  (def {name} {params} {body})
//  (var {name} (func {params} {body}) )
func evmac(e, a *types.Cell) *types.Cell {
	name := cadr(e); params_body := cddr(e)
	k := name; v := cons(builtin.MAC, params_body)
	builtin.Tag(v, builtin.ID_MAC)
	a = addEnv(list(k, v), a)
	return eval(k, a)
}

// Named functions
// ---------------
// (func foo (x) (car x))
// (func {name} {params} {body})
//
// Unnamed lambda functions
// ------------------------
// (func (x) (car x))
// (func {params} {body})

// evlet eval let see example below
// 	e.g. 
//		(let xs '(1 2 3) (car xs)) ->  1
//		(let {key} {val} {body} )
func evlet(e, a *types.Cell) *types.Cell {
	k  := cadr(e) ;  v := eval(caddr(e), a)	
	ee := cdddr(e); aa := cons(list(k, v), a)	
	return eval(ee, aa)
}

// envenv only print environment for debug purpose
func evenv(e, a *types.Cell) *types.Cell {
	fmt.Printf("evenv - a: %v ap:%p \n\n", a, a)
	return builtin.NIL
}

// evset evals expression e.g. (set a 1)
// add a key value pair on top of environment
// like push on a stack
// 	env := (
//		(t t)
// 	)
//
// 	> (set a 1) ->
//
// 	env = (
//		(a 1)
//		(t t)
//	)
func evset(e, a *types.Cell) *types.Cell {
	k := cadr(e); v := eval(caddr(e), a)
	a = addEnv(list(k, v), a)
	return eval(k, a)
}


// evatom evals atom from environment
// e.g. env := (
//		(a	1)
//		(b  1)
//	)
//
//  > a -> 1
//  > b -> 2
//
func evatom(e, a *types.Cell) *types.Cell {
	if e.IsSymbol() {	
		x := assoc(e, a) // ToDo: Hash-table // nil means also not found !!!	
		if x.IsErr() {
			return builtin.Err("reference to undefined identifier: %v", e) // TODO: Rename error message
		}
		return x
	}
	return e
}

// evfunc eval func from environment
//	env = (foo (func (x) (no x)))
//  > (foo nil)
// 	ee = ((func (x) (no x)) nil)
func evfunc(e, a *types.Cell) *types.Cell {
	name := assoc(car(e), a)
	if name.IsErr() {
		return builtin.Err("reference to undefined identifier: %v", car(e)) // TODO: Rename error message
	}
	ee := cons(name, cdr(e))
	//fmt.Printf("evfunc - ee: %v", ee)
	return eval(ee, a)
}

// evcon evals cond (= conditions)
// e.g. (cond 
//			( 
//				(nil b) 
//				( 't a)
//			) 
//		) -> a
func evcon(c, a *types.Cell) *types.Cell {
	if eval(caar(c), a).Equal(builtin.T) { 
		return eval(cadar(c), a) 
	} else { 
		return evcon(cdr(c), a)
	}
}

// evlis evals each item of a list
func evlis(m, a *types.Cell) *types.Cell {
	if m.Equal(builtin.NIL) {
		return builtin.NIL
	} else {
		return cons(eval(car(m), a), evlis(cdr(m), a))
	}
}






 
