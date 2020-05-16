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

// BEGIN  7 primitive operators

/*						

												
PROCEDURE quote(x: cell): cell;
BEGIN RETURN cadr(x);
END quote;
*/
func quote(x *types.Cell) *types.Cell {
	return cadr(x)
}

/*		
PROCEDURE atom(x: cell): cell;
BEGIN
	IF x IS atomCell THEN RETURN T;
	ELSE RETURN EMPTY END;
END atom;
*/
func atom(x *types.Cell) *types.Cell {
	if x.IsAtom() {
		return builtin.T
	} else {
		return builtin.NIL
	}
}


/*	
PROCEDURE eq(x: cell; y: cell): cell; 
BEGIN
	IF (x = EMPTY)  &  (y = EMPTY) THEN RETURN T;
	ELSIF (x IS atomCell) & (y IS atomCell) THEN
			IF x(atomCell).name = y(atomCell).name THEN RETURN T END;
	END;
	RETURN EMPTY;	
END eq;
*/
func eq(x, y *types.Cell) *types.Cell {	
	if x.Equal(y) {
	 	return builtin.T	
	}
	return builtin.NIL 	// FALSE
}

/*
PROCEDURE car(x: cell): cell;
BEGIN
	IF x IS consCell THEN
		IF x = EMPTY THEN RETURN EMPTY;
		ELSE RETURN LSG.car(x) END;
	ELSE
		error(1); (* TODO - error: can not take car of foo *)
	END;
END car;
*/
func car(e *types.Cell) *types.Cell {

	if e.Equal(builtin.NIL) {
		return builtin.NIL
	}
	
	if e.IsCons() {
		return builtin.Car(e) 
	} else {
		return generator.Error(fmt.Sprintf("Can't take car of %v",e))
	}
	
}


/*	
PROCEDURE cdr(e: cell): cell;
BEGIN
	IF e IS consCell THEN
		IF e = EMPTY THEN RETURN EMPTY;
		ELSE RETURN LSG.cdr(e) END;
	ELSE
		error(1); (* TODO - error: can not take cdr of foo *)
	END;
END cdr;
*/
func cdr(e *types.Cell) *types.Cell {
	
	if e.Equal(builtin.NIL) {
		return builtin.NIL
	}
	
	if e.IsCons() {
		return builtin.Cdr(e)
	} else {
		return generator.Error(fmt.Sprintf("Can't take cdr of %v", e))
	}
	
}


/*	
PROCEDURE cons(x:cell; y:cell): cell; 
BEGIN
	IF y IS consCell THEN
		RETURN LSG.cons(x, y);
	ELSE
		error(1); (* TODO - error: y must be a list *)
	END;
END cons;
*/
func cons (x, y *types.Cell) *types.Cell {
	return generator.Cons(x, y)
}

/*	
PROCEDURE cond(x: cell): cell;
BEGIN
	IF x IS consCell THEN
		IF eq(caar(x), T) = T THEN RETURN cadar(x);				
		ELSE RETURN cond(cdr(x)) END;
	ELSE
		error(1); (* TODO: x must be a list of from ((p1 e1) (p2 e2) .. (pn en)) *)		
	END;
END cond;


*/
func cond(x *types.Cell) *types.Cell {

	if x.IsCons() {
		if caar(x).Equal(builtin.T) {
			return cadar(x)
		} else {
			return cond(cdr(x))
		}
	} else {
		return generator.Error(fmt.Sprintf("x must be a list of form ((p1 e1) (p2 e2) .. (pn en))"))
	}
	
}

// END 7 primitive operator 

// BEGIN 2-3 NEW primitive operator 

/*	
PROCEDURE isTagged(e: cell): BOOLEAN;
	VAR x: cell;
BEGIN

	(* ('TAGGED #TYPE #REP) 
		((QUOTE TAGGED) #TYPE #REP)
	*)

	IF  (e IS consCell) THEN
		x := car(e);
		IF (x IS consCell) & (eq(caar(e), QUOTE) = T) &  (eq(cadar(e), TAGGED) = T) THEN
			RETURN TRUE;
		ELSE
			RETURN FALSE;
		END;
	ELSE 
		RETURN FALSE;
	END;
END isTagged;
*/
func isTagged(e *types.Cell) *types.Cell {

	if e.IsCons() {
		x := car(e)
		if x.IsCons() && 
		    caar(e).Equal(builtin.QUOTE) && 
		   cadar(e).Equal(builtin.TAGGED) {
			return builtin.T
		} else {
			return builtin.NIL
		}
	} else {
		return builtin.NIL
	}
	
}

/*	
PROCEDURE tag*(type: cell; rep: cell): cell;
	VAR tc, t: cell;
BEGIN
	(* ('tag TYPE REP) *)
	tc := cons(A("QUOTE"), cons(A("TAGGED"), EMPTY));
	t := cons(tc, cons(type, cons(rep, EMPTY)));	
	RETURN t;
END tag;
*/
func tag(t,r *types.Cell) *types.Cell {
	tc := cons(builtin.QUOTE, cons(builtin.TAGGED, builtin.NIL))
	return cons(tc, cons(t, cons(r, builtin.NIL)))
}

/*
PROCEDURE type*(e: cell): cell;
BEGIN
	IF e IS consCell THEN 
		RETURN A("CONS");
	ELSE 
		RETURN A("SYM");
	END;
END type;
*/
func type_(e *types.Cell) *types.Cell {
	if e.IsCons() {
		return builtin.CONS
	} else {
		// TODO: ???
		return builtin.Sym(fmt.Sprintf("%v", e.Type.Atom))
	}
}

/*
PROCEDURE type0*(e: cell): cell;
		
	PROCEDURE tagType(e: cell): cell;
	BEGIN
		RETURN cadr(e);
	END tagType;
	
BEGIN (* type *)
	IF isTagged(e) THEN
		RETURN tagType(e);
	ELSE
		RETURN type(e);			
	END;
END type0;
*/

func type0(e *types.Cell) *types.Cell {

	tagType := func (e *types.Cell) *types.Cell {
		return cadr(e)
	}
	
	if isTagged(e).Equal(builtin.T) {
		return tagType(e)
	} else {
		return type_(e)
	}
}

/*	
PROCEDURE rep*(e: cell): cell;
BEGIN
	(* ((QUOTE TAGGED) #TYPE #REP) *)
	IF isTagged(e) THEN
		RETURN caddr(e);
	ELSE
		RETURN cadr(e);
	END;
END rep;
*/
func rep(e *types.Cell) *types.Cell {
	if isTagged(e).Equal(builtin.T) {
		return caddr(e)
	} else {
		return cadr(e)
	}
}

/*	

(* TODO below should be implemented in LISP itself *)

PROCEDURE pair(x: cell; y: cell): cell;
BEGIN
	IF  (x = EMPTY) & (y = EMPTY) THEN RETURN EMPTY;
	ELSE
		IF (x IS consCell) & (y IS consCell) THEN
			RETURN cons(list(car(x),car(y)), pair(cdr(x), cdr(y)));
		END;
	END; 
END pair;
*/
func pair(x, y *types.Cell) *types.Cell {
	if x.Equal(builtin.NIL) && y.Equal(builtin.NIL) {
		return builtin.NIL
	} else {
		if x.IsCons() && y.IsCons() {
			a := list(car(x), car(y))
			b := pair(cdr(x), cdr(y))
			return cons(a,b)
		}
	}
	return generator.Error("x and y must be a cons") // TODO: Check
}
/*	
PROCEDURE null(x: cell): cell;
BEGIN
	IF (x IS consCell) THEN
		IF x = EMPTY THEN RETURN T; END;
	END;
	RETURN EMPTY;
END null;
*/
func no(x *types.Cell) *types.Cell {
	if x.IsCons() {
		if x.Equal(builtin.NIL) {
			return builtin.T
		}
	}
	return builtin.NIL
}

/*		
PROCEDURE not(x: cell): cell;
BEGIN
	IF eq(x, T) = T THEN RETURN EMPTY;
	ELSE RETURN T END;
END not;
*/
func not (x *types.Cell) *types.Cell {
	if x.Equal(builtin.T) {
		return builtin.NIL
	} else {
		return builtin.T
	}
}

/*	
PROCEDURE and(x:cell; y: cell): cell;
BEGIN
	IF eq(x,T) = T THEN
		IF eq(y,T) = T THEN RETURN T;
		ELSE RETURN EMPTY END;
	ELSE
		RETURN EMPTY;
	END;
END and;
*/
func and(x, y *types.Cell) *types.Cell {
	if x.Equal(builtin.T) && y.Equal(builtin.T) {
		return builtin.T
	} else {
		return builtin.NIL
	}
}

/*	
PROCEDURE append(x: cell; y: cell): cell;
BEGIN
	IF x = EMPTY THEN RETURN y;
	ELSE
		RETURN cons(car(x), append(cdr(x), y));
	END;
END append;
*/
func append(x, y *types.Cell) *types.Cell {
	if x.Equal(builtin.NIL) {
		return y
	} else {
		return cons(car(x), append(cdr(x), y))
	}
}

/*	
PROCEDURE list(x: cell; y: cell): cell;
BEGIN
	RETURN cons(x, cons(y, EMPTY));
END list;
*/
func list(x, y *types.Cell) *types.Cell {
	return cons(x, cons (y, builtin.NIL))
}


/*					
PROCEDURE assoc(x: cell; y: cell): cell;
BEGIN
	IF y = EMPTY THEN RETURN EMPTY;
	ELSE
		IF eq(caar(y), x) = T THEN RETURN cadar(y);
		ELSE RETURN assoc(x, cdr(y)); END;
	END;
END assoc;
*/

func assoc(x, y *types.Cell) *types.Cell {
	if y.Equal(builtin.NIL) {
		return builtin.NIL
	} else {
		if eq(caar(y), x).Equal(builtin.T) {
			return cadar(y)
		} else {
			return assoc(x, cdr(y))	
		}
	}
}

/*		
PROCEDURE caar  (e: cell): cell; BEGIN RETURN car(car(e)) END caar;
PROCEDURE cadr  (e: cell): cell; BEGIN RETURN car(cdr(e)) END cadr;
PROCEDURE cadar (e: cell): cell; BEGIN RETURN car(cdr(car(e))) END cadar;
PROCEDURE caddr (e: cell): cell; BEGIN RETURN car(cdr(cdr(e))) END caddr;
PROCEDURE caddar(e: cell): cell; BEGIN RETURN car(cdr(cdr(car(e)))) END caddar;
*/

func caar  (e *types.Cell) *types.Cell { return car(car(e)) }
func cadr  (e *types.Cell) *types.Cell { return car(cdr(e)) }
func cadar (e *types.Cell) *types.Cell { return car(cdr(car(e))) }
func caddr (e *types.Cell) *types.Cell { return car(cdr(cdr(e))) }
func caddar(e *types.Cell) *types.Cell { return car(cdr(cdr(car(e)))) }

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

	// TODO: Use hash-Table for SPEEEED
	
	if e.IsAtom() { return assoc(e, a) } // Take it from environment
	
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
			
			// Function call of functions stored in environment
			default: {
				label := assoc(car(e), a)
				if label.Equal(builtin.NIL) {
					return generator.Error(fmt.Sprintf("reference to undefined identifier: %v", car(e)))
				}
				return eval(cons(label, cdr(e)), a)
			}
			// (label 'foo)
			// TODO: Error: "reference to undefined identifier: _label"
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





 
