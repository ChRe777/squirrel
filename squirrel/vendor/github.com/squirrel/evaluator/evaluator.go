package evaluator

import (
	"fmt"
)

import (
	"github.com/squirrel/types"
	"github.com/squirrel/builtin"
	"github.com/squirrel/generator"
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

	// > 1 -> 1
	// > "2" -> "2"
	// > a -> error: undefined identifier
	// > t -> t
	// > nil -> ''
	if e.IsAtom() {
		if e.IsSymbol() {	
			return assoc(e, a)			// TODO: Use hash-Table for SPEEEED
		}
		return e
	} 
	
	c := car(e)	
	if c.IsAtom() {
		switch {	
			//
			// TODO: map[string]func instead of switch -> SPEED
			//
			case c.Equal(builtin.QUOTE): return quote(e) 
			case c.Equal(builtin.ATOM ): return atom(eval(cadr(e), a))
			case c.Equal(builtin.EQ   ): return eq(eval(cadr(e), a), eval(caddr(e), a))
			case c.Equal(builtin.CAR  ): return car(eval(cadr(e), a))
			case c.Equal(builtin.CDR  ): return cdr(eval(cadr(e), a))
			case c.Equal(builtin.CONS ): return cons(eval(cadr(e), a), eval(caddr(e), a))
			case c.Equal(builtin.COND ): return evcon(cdr(e), a)
						
			// Extra core from Arc
			case c.Equal(builtin.TAG)  : return tag(eval(cadr(e), a), eval(caddr(e), a))
			case c.Equal(builtin.TYPE0): return type0(eval(cadr(e), a))
			case c.Equal(builtin.REP)  : return rep(eval(cadr(e), a))
			
			// Function calls by environment
			default: {
				label := assoc(car(e), a)
				if label.Equal(builtin.NIL) {
					return generator.Error(fmt.Sprintf("reference to undefined identifier: %v", car(e)))
				}
				return eval(cons(label, cdr(e)), a)
			}
		}
	}
	/*
		IF eq(caar(e), LABEL) = T THEN 
				RETURN eval( 
					cons(caddar(e), cdr(e)),
					cons(list(cadar(e), car(e)), a)
				);
		END;
	*/
	if caar(e).Equal(builtin.LABEL) {
		return eval(cons(caddar(e), cdr(e)),
			        cons(list(cadar(e), car(e)), a))
	}
	/*
		IF eq(caar(e), LAMBDA) = T THEN  
			RETURN eval(
					caddar(e),
					append(pair(cadar(e), evlis(cdr(e), a)), a)
				);
		END;
	*/
	
	if caar(e).Equal(builtin.FUNC) {
		return eval(caddar(e),
					append(pair(cadar(e), evlis(cdr(e), a)), a))
	}
		
	return generator.Error("Something got wrong in eval")
}


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


/*		
PROCEDURE evcon(c: cell; a: cell): cell;
BEGIN	
	IF eq(eval(caar(c), a), T) = T THEN RETURN eval(cadar(c), a);
	ELSE RETURN evcon(cdr(c), a) END;
END evcon;
*/
func evcon(c, a *types.Cell) *types.Cell {
	
	if eval(caar(c), a).Equal(builtin.T) { 
		return eval(cadar(c), a) 
	} else { 
		return evcon(cdr(c), a)
	}
}

/*	
PROCEDURE evlis(m: cell; a: cell): cell;
BEGIN
	 IF m = EMPTY THEN RETURN EMPTY;
	 ELSE
	 	 RETURN cons(eval(car(m), a), evlis(cdr(m), a))
	 END;
END evlis;
*/
func evlis(m, a *types.Cell) *types.Cell {
	if m.Equal(builtin.NIL) {
		return builtin.NIL
	} else {
		return cons(eval(car(m), a), evlis(cdr(m), a))
	}
}





 
