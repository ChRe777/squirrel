package evaluator

import (
	"fmt"
)

import (
	"github.com/squirrel/types"
	"github.com/squirrel/builtin"
	//"github.com/squirrel/generator"
)

// Eval evals expression e with environment env and returns result
func Eval(e, env *types.Cell) *types.Cell {
	return eval(e, env)
}

/*			
PROCEDURE eval*(e: cell; a: cell): cell;
	VAR c: cell;
BEGIN
	IF e IS atomCell THEN RETURN assoc(e, a) END; 
	c := car(e);
	IF c IS atomCell THEN
		IF eq(car(e), QUOTE) = T THEN RETURN quote(e);
		ELSIF eq(car(e), ATOM) = T THEN RETURN atom(eval(cadr(e), a));
		ELSIF eq(car(e), EQ) = T THEN RETURN eq(eval(cadr(e), a), eval(caddr(e), a));
		ELSIF eq(car(e), CAR) = T THEN RETURN car(eval(cadr(e), a));
		ELSIF eq(car(e), CDR) = T THEN RETURN cdr(eval(cadr(e), a));
		ELSIF eq(car(e), CONS) = T THEN RETURN cons(eval(cadr(e), a), eval(caddr(e), a));
		ELSIF eq(car(e), COND) = T THEN RETURN evcon(cdr(e), a);
		ELSIF eq(car(e), TAG) = T THEN RETURN tag(eval(cadr(e), a), eval(caddr(e), a));
		ELSIF eq(car(e), TYPE0) = T THEN RETURN type0(eval(cadr(e), a));
		ELSIF eq(car(e), REP) = T THEN RETURN rep(eval(cadr(e), a));
		ELSE RETURN eval(cons(assoc(car(e), a), cdr(e)), a); END;
	END;
	IF eq(caar(e), LABEL) = T THEN 
			RETURN eval( 
				cons(caddar(e), cdr(e)),
				cons(list(cadar(e), car(e)), a)
			);
	END;
	IF eq(caar(e), LAMBDA) = T THEN  
		RETURN eval(
				caddar(e),
				append(pair(cadar(e), evlis(cdr(e), a)), a)
			);
	END;
END eval;
*/
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
		
			// 7 core axioms - root of lisp (McCarthy)
			case c.Equal(builtin.QUOTE): return quote(e) 
			case c.Equal(builtin.ATOM ): return atom (eval(cadr(e), a))
			case c.Equal(builtin.EQ   ): return eq   (eval(cadr(e), a), eval(caddr(e), a))
			case c.Equal(builtin.CAR  ): return car  (eval(cadr(e), a))
			case c.Equal(builtin.CDR  ): return cdr  (eval(cadr(e), a))
			case c.Equal(builtin.CONS ): return cons (eval(cadr(e), a), eval(caddr(e), a))
			case c.Equal(builtin.COND ): return evcon(cdr(e), a)
						
			// 3 extra core axioms from Arc (Paul Graham)
			case c.Equal(builtin.TAG  ): return tag  (eval(cadr(e), a), eval(caddr(e), a))
			case c.Equal(builtin.TYPE0): return type0(eval(cadr(e), a))
			case c.Equal(builtin.REP  ): return rep  (eval(cadr(e), a))
			
			// Extra axioms in environment e.g. (no '()) -> t
			default: { // ToSimplify
				label := assoc(car(e), a)
				if notFound(label) {
					return builtin.Err("reference to undefined identifier: %v", car(e))
				}
				ee := cons(label, cdr(e))
				fmt.Printf("ee: %v \n",ee)
				return eval(ee, a)
			}
		}
	}
				
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
	// expression with the inner lambda expression substituted for the label
	// expression.
	
	
	if caar(e).Equal(builtin.LABEL) {
			
		ee   := cons(caddar(e), cdr(e))
		name := cadar(e)
		fnn  := caddar(e)		
		aa   := cons(list(name, fnn), a)

		res  := eval(ee, aa)
		
		fmt.Printf("\n\n label - exp:%v, env:%v \n name:%v, fnn: %v \n\n", ee, aa, name, fnn)
		return res
	}
	
	// d) Function calls 
	//           (         f          params ) -> ?
	//		e.g. ( (func (x) (car x)) '(1 2) ) -> 1
	if caar(e).Equal(builtin.FUNC) {
		ee := caddar(e)
		aa := append(pair(cadar(e), evlis(cdr(e), a)), a)
		fmt.Printf("func call ee:%v, aa:%v \n", ee, aa)
		return eval(ee,aa)
	}
		
	return builtin.Err("Something got wrong in eval")
}

func notFound(a *types.Cell) bool {
 	return a.Equal(builtin.NIL)
}

func evatom(e, a *types.Cell) *types.Cell {
	if e.IsSymbol() {
		// ToDO: not found in assoc and nil evals to nil are the same
		if e.Equal(builtin.NIL) {
			return builtin.NIL
		}
		// Todo: Hash-table
		x := assoc(e, a) // nil means also not found
		if x.Equal(builtin.NIL) {
			// TODO: Rename error message
			return builtin.Err("reference to undefined identifier: %v", e)
		}
		return x
	}
	return e
}

func evcon(c, a *types.Cell) *types.Cell {
	if eval(caar(c), a).Equal(builtin.T) { 
		return eval(cadar(c), a) 
	} else { 
		return evcon(cdr(c), a)
	}
}

func evlis(m, a *types.Cell) *types.Cell {
	if m.Equal(builtin.NIL) {
		return builtin.NIL
	} else {
		return cons(eval(car(m), a), evlis(cdr(m), a))
	}
}

/*
type cell 	*types.Cell
type coreFn func(e, a cell) cell

var coreFnsMap = map[string]coreFn { 
	// 7 core
 	builtin.ID_QUOTE: func(e, a cell) cell { return quote(e)  									},
	builtin.ID_ATOM : func(e, a cell) cell { return atom(eval(cadr(e), a)) 						},
	builtin.ID_EQ   : func(e, a cell) cell { return eq(eval(cadr(e), a), eval(caddr(e), a)) 	},
	builtin.ID_CAR  : func(e, a cell) cell { return car(eval(cadr(e), a)) 						},
	builtin.ID_CDR  : func(e, a cell) cell { return cdr(eval(cadr(e), a)) 						},
	builtin.ID_CONS : func(e, a cell) cell { return cons(eval(cadr(e), a), eval(caddr(e), a)) 	},
	builtin.ID_COND : func(e, a cell) cell { return evcon(cdr(e), a) 							},
	// 3 extra
	builtin.ID_TAG  : func(e, a cell) cell { return tag(eval(cadr(e), a), eval(caddr(e), a)) 	},
	builtin.ID_TYPE0: func(e, a cell) cell { return type0(eval(cadr(e), a)) 					},
	builtin.ID_REP  : func(e, a cell) cell { return rep(eval(cadr(e), a)) 						},

}

func eval2(e, a *types.Cell) *types.Cell {

	// a) Atom look up in environment
	if e.IsAtom() {
		if e.IsSymbol() {	
			return assoc(e, a)			// TODO: Use hash-Table for SPEEEED
		}
		return e
	} 
	
	// b) Cons take name and look in core e.g. (car '(1 2 3))
	c := car(e)	
	if c.IsAtom() {
		identifier, ok1 := c.Val.(string)
		fn, ok2 := coreFnsMap[identifier]
		if ok1 && ok2 {
			return fn(e, a)
		} else {
			fnName := assoc(car(e), a); 
			if fnName.Equal(builtin.NIL) {
				return generator.Error(fmt.Sprintf("reference to undefined identifier: %v", car(e)))
			}
			return eval(cons(fnName, cdr(e)), a)
		}
	}
	
	// c) ...
	if caar(e).Equal(builtin.LABEL) {
		return eval(cons(caddar(e), cdr(e)),
			        cons(list(cadar(e), car(e)), a))
	}
	
	// d) ...
	if caar(e).Equal(builtin.FUNC) {
		return eval(caddar(e),
					append(pair(cadar(e), evlis(cdr(e), a)), a))
	}
		
	return generator.Error("Something got wrong in eval")
}

*/








 
