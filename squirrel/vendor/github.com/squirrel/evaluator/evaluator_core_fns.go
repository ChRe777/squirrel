package evaluator

import (
	"github.com/squirrel/types"
	"github.com/squirrel/builtin"
)

// 7 primitive core operators

/*																		
PROCEDURE quote(x: cell): cell;
BEGIN RETURN cadr(x);
END quote;
*/
func quote(x *types.Cell) *types.Cell {
	return cadr(x)  // (quote a) -> cdr -> (a) -> car -> a
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
	} else {
		if e.IsCons() {
			return builtin.Car(e) 
		} else {
			return builtin.Err("Can't take car of %v", e)
		}
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
	} else {
		if e.IsCons() {
			return builtin.Cdr(e)
		} else {
			return builtin.Err("Can't take cdr of %v", e)
		}
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
func cons(x, y *types.Cell) *types.Cell {
	return builtin.Cons(x, y)
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
		return builtin.Err("x must be a list of form ((p1 e1) (p2 e2) .. (pn en))")
	}
	
}

// END 7 primitive core operators
