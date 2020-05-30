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
		return evatom(e, a)
	} 
	
	// b) Functions e.g. (car '(1 2)) -> 1	
	c := core.Car(e)
	if c.IsAtom() {
		switch {	
		
			// 7 core axioms - "The Roots of lisp" (McCarthy, Paul Graham)
			//
			case c.Equal(core.QUOTE): return core.Quote(e) 
			case c.Equal(core.ATOM ): return core.Atom(eval(builtin.Cadr(e), a))
			case c.Equal(core.IS   ): return core.Is  (eval(builtin.Cadr(e), a), eval(builtin.Caddr(e), a))
			case c.Equal(core.CAR  ): return core.Car (eval(builtin.Cadr(e), a))
			case c.Equal(core.CDR  ): return core.Cdr (eval(builtin.Cadr(e), a))
			case c.Equal(core.CONS ): return core.Cons(eval(builtin.Cadr(e), a), eval(builtin.Caddr(e), a))
			case c.Equal(core.COND ): return evcon(core.Cdr(e), a)
									

			// 3 extra core axioms from Arc (Paul Graham)
			//
//			case c.Equal(core.TAG  ): return core.Tag  (eval(builtin.cadr(e), a), eval(builtin.caddr(e), a))
//			case c.Equal(core.TYPE0): return core.Type0(eval(builtin.cadr(e), a))
//			case c.Equal(core.REP  ): return core.Rep  (eval(builtin.cadr(e), a))		

			// TEST - REFACTOR
			case c.Equal(core.VAR): return evvar(e, a)		
			case c.Equal(core.ENV): return evenv(e, a)
			case c.Equal(core.LET): return evlet(e, a)
			case c.Equal(core.DEF): return evdef(e, a)
			case c.Equal(core.MAC): return evmac(e, a)
			
			// For macros
			case c.Equal(core.BACKQUOTE): return Backquote(e, a) 	// TODO: evBackQuote

			
			// 7 extension functions from "The Roots of Lisp" (McCarthy, Paul Graham)
			//
			case c.Equal(builtin.NO    ): return builtin.No    (eval(builtin.Cadr(e), a))
			case c.Equal(builtin.NOT   ): return builtin.Not   (eval(builtin.Cadr(e), a))
			case c.Equal(builtin.AND   ): return builtin.And   (eval(builtin.Cadr(e), a), eval(builtin.Caddr(e), a))
			case c.Equal(builtin.PAIR  ): return builtin.Pair  (eval(builtin.Cadr(e), a), eval(builtin.Caddr(e), a))
			case c.Equal(builtin.LIST  ): return builtin.List_ (eval(builtin.Cadr(e), a), eval(builtin.Caddr(e), a))
			case c.Equal(builtin.ASSOC ): return builtin.Assoc (eval(builtin.Cadr(e), a), eval(builtin.Caddr(e), a))
			case c.Equal(builtin.APPEND): return builtin.Append(eval(builtin.Cadr(e), a), eval(builtin.Caddr(e), a))

			
			// Extra axioms in environment e.g. (no '()) -> t
			default: return evfunc(e, a)
		}
	}
	
	// c) Labels calls
	if builtin.Caar(e).Equal(core.LABEL) {
		return evlabel(e, a)
	} 
	
	// d) Function calls 
	if builtin.Caar(e).Equal(core.FUNC) {
		return evfuncCall(e, a)
	}
		
	return core.Err("Wrong expression")
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
func evlabel(e, a *types.Cell) *types.Cell {
	label := builtin.Cadar(e); fn := builtin.Caddar(e)
		
	ee := core.Cons(builtin.Caddar(e),core.Cdr(e))		
	aa := core.Cons(builtin.List_(label, fn), a)  // ( (no (func (x) (eq x nil)) (a 1) (b 2) )
				
	return eval(ee, aa)
}


// evfunc evals function calls with arguments
//	e.g.
//		( (func (x) (car x)) '(1 2) ) -> 1
func evfuncCall(e, a *types.Cell) *types.Cell {
	k := builtin.Cadar(e); v := evlis(core.Cdr(e), a)	
		
	ee := builtin.Caddar(e)
	aa := builtin.Append(builtin.Pair(k, v), a)		
			
	r := eval(ee, aa)		// will expand backquotes and unquotes
	if isMac(e) {
		return eval(r, aa)
	}
	
	return r
}


// isMac checks if caar(e) is tagged as macro
// e.g.
//		e = ((func (x) (no x))
//		caar(e) -> func 
func isMac(e *types.Cell) bool {
	return builtin.Caar(e).IsTagged(core.ID_MAC)
}

// evdef eval 'def and creates a function in environment
// e.g.
//  	(def {name} {params} {body})
//  	(var {name} (func {params} {body}) )
func evdef(e, a *types.Cell) *types.Cell {
	name := builtin.Cadr(e); params_body := builtin.Cddr(e)
	k := name; v := core.Cons(core.FUNC, params_body)		// TODO: REFACTOR
	core.Tag(v, core.ID_FUNC)
	a = addEnv(builtin.List_(k, v), a)
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

// evlet eval 'let, see example below
// 	e.g. 
//		(let xs '(1 2 3) (car xs)) ->  1
//		(let {key} {val} {body} )
func evlet(e, a *types.Cell) *types.Cell {
	k := builtin.Cadr(e);  v := eval(builtin.Caddr(e), a)	
	ee := core.Car(builtin.Cdddr(e)); aa := core.Cons(builtin.List_(k, v), a)	
	
	fmt.Printf("evlet e: %v ee:%v, aa: %v \n", e, ee, aa)
	
	return eval(ee, aa)
}

// envenv only print environment for debug purpose
func evenv(e, a *types.Cell) *types.Cell {
	fmt.Printf("evenv - a: %v ap:%p \n\n", a, a)
	return core.NIL
}

// evset evals expression e.g. (set a 1)
// add a key value pair on top of environment
// like push on a stack
// 	env := (
//		(t t)
// 	)
//
// 	> (var a 1) ->
//
// 	env = (
//		(a 1)
//		(t t)
//	)
func evvar(e, a *types.Cell) *types.Cell {
	k := builtin.Cadr(e); v := eval(builtin.Caddr(e), a)
	fmt.Printf("evalcar k:%v v:%v e:%v \n",k,v, e)
	a = addEnv(builtin.List_(k, v), a)
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
		x := builtin.Assoc(e, a) // ToDo: Hash-table // nil means also not found !!!	
		if x.IsErr() {
			return core.Err("reference to undefined identifier: %v", e) // TODO: Rename error message
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
	name := builtin.Assoc(core.Car(e), a)
	if name.IsErr() {
		return core.Err("reference to undefined identifier: %v", core.Car(e)) // TODO: Rename error message
	}
	ee := core.Cons(name, core.Cdr(e))
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
	if eval(builtin.Caar(c), a).Equal(core.T) { 
		return eval(builtin.Cadar(c), a) 
	} else { 
		return evcon(core.Cdr(c), a)
	}
}

// evlis evals each item of a list
func evlis(m, a *types.Cell) *types.Cell {
	if m.Equal(core.NIL) {
		return core.NIL
	} else {
		return core.Cons(eval(core.Car(m), a), evlis(core.Cdr(m), a))
	}
}


// ---------------------------------
// Just ALIAS for better readability
// ---------------------------------

/*
func car(x *types.Cell) *types.Cell {
	return core.Car(x)
}

func cdr(x *types.Cell) *types.Cell {
	return core.Cdr(x)
}

func cons(x, y *types.Cell) *types.Cell {
	return core.Cons(x,y)
}

func eq(x, y *types.Cell) *types.Cell {
	return core.Eq(x,y)
}
*/


// ---------------------------------
// TODO: CHECK FOR SPEED
// ---------------------------------


// > (set a 1) -> 1
// > a -> 1
func set(k, v *types.Cell, a *types.Cell) *types.Cell {
	// Add key-value-pair (k v) to environment
	a = core.Cons(builtin.List_(k, v), a)
	return v
}

// addEnv is a special add that adds a new cell at the front of the environment
// but LET the Pointer to first element the SAME !!!
func addEnv(kv *types.Cell, a *types.Cell ) *types.Cell {
	// Hang in new as second
	cdr := a.Cdr; new := core.Cons(kv, cdr); a.Cdr = new
	// Change Val first and second to move new second to front
	val := new.Val; new.Val = a.Val; a.Val = val
	// Change Car first and second to move new seocen to front
	car := new.Car; new.Car = a.Car; a.Car = car
	// So the pointer to a stays the same // Side effects // ToReThink: ?
	return a
}



 
