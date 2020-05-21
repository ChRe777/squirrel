package evaluator

import (
	"github.com/squirrel/types"
	"github.com/squirrel/builtin"
	"github.com/squirrel/generator"
)

// * pair
// * no
// * not
// * and
// * append
// * list
// * assoc

// * caar
// * cadr
// * cadar
// * caddr
// * caddar

// TODO: below should/can be implemented in LISP itself

/*
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
	IF (x IS consCell) THEN  <<--- WRONG
		IF x = EMPTY THEN RETURN T; END;
	END;
	RETURN EMPTY;
END null;
*/
func no(x *types.Cell) *types.Cell { // call "no" instead of "null"
	if x.Equal(builtin.NIL) {
		return builtin.T
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
		return builtin.Err("Not found")
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
func caar  (e *types.Cell) *types.Cell { return car(car(e))           }
func cadr  (e *types.Cell) *types.Cell { return car(cdr(e))           }
func cadar (e *types.Cell) *types.Cell { return car(cdr(car(e)))      } 
func caddr (e *types.Cell) *types.Cell { return car(cdr(cdr(e)))      }
func caddar(e *types.Cell) *types.Cell { return car(cdr(cdr(car(e)))) }


// > (set a 1) -> 1
// > a -> 1
func set(k, v *types.Cell, a *types.Cell) *types.Cell {
	// Add key-value-pair (k v) to environment
	a = cons(list(k, v), a)
	return v
}

