package builtin

import (
	//"github.com/squirrel/generator"
)

// Identifiers of built-in core symbols
const (
	// 7 Primitives
	ID_QUOTE  = "quote"
	ID_ATOM   = "atom"
	ID_EQ	  = "eq"
	ID_CAR    = "car"
	ID_CDR	  = "cdr"
	ID_CONS	  = "cons"
	ID_COND   = "cond"
	// Ext
	ID_LABEL  = "label"
	ID_FUNC   = "func"		// was LAMBDA
	ID_T	  = "t"
	ID_NIL	  = "nil"
	// New Operator from Arc
	ID_TAG 	  = "tag"
	ID_TAGGED = "tagged"
	ID_SYM	  = "sym"
	ID_TYPE0  = "type"
	ID_REP    = "rep"
)

// Core symbols of language
var (
	
	QUOTE 	= Sym(ID_QUOTE)
	ATOM 	= Sym(ID_ATOM)
	EQ 		= Sym(ID_EQ)
	CAR 	= Sym(ID_CAR)
	CDR 	= Sym(ID_CDR) 
	CONS 	= Sym(ID_CONS)
	COND 	= Sym(ID_COND) 
	
	LABEL 	= Sym(ID_LABEL)
	FUNC 	= Sym(ID_FUNC) 
	
	T 		= Sym(ID_T) 		// TRUE
	NIL 	= Sym(ID_NIL) 		// NIL or FALSE
	
	TAG 	= Sym(ID_TAG)
	TAGGED 	= Sym(ID_TAGGED)
	SYM 	= Sym(ID_SYM)
	TYPE0 	= Sym(ID_TYPE0)		// TODO: Better name
	REP 	= Sym(ID_REP)	
)



/*		
BEGIN

	QUOTE := A("QUOTE"); 
	ATOM := A("ATOM");
	EQ := A("EQ"); 
	CAR := A("CAR");
	CDR := A("CDR"); 
	CONS := A("CONS");
	COND := A("COND"); 
	
	TAGGED := A("TAGGED");
	SYM := A("SYM");
	TYPE0 := A("TYPE");
	TAG := A("TAG");
	REP := A("REP");
	
	LABEL := A("LABEL");
	LAMBDA := A("LAMBDA"); 
	T := A("T"); (*TRUE*)
	
	EMPTY := LSG.EMPTY; (*EMPTY LIST, FALSE*)
	ENV := LSG.cons(NIL,NIL); (*ENVIROMENT*)	
	
END LSE.
*/

func init() {
	
}


