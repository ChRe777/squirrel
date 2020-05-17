package evaluator

import (
	"fmt"
)

import (
	"github.com/squirrel/types"
	"github.com/squirrel/builtin"
)

// Paul Graham add 2-3 new primitive core operators in Arc
// to make macros much simpler and other things
// A normal function is tagged with 'mac

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
func tag(t, r *types.Cell) *types.Cell {
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
