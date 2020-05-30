package core

// Identifiers of built-in core symbols
const (
	// 7 Primitives
	ID_QUOTE  = "quote"
	ID_ATOM   = "atom"
	ID_EQ	  = "is"		// was EQ
	ID_CAR    = "car"		// can be first
	ID_CDR	  = "cdr"		// can be rest
	ID_CONS	  = "cons"	
	ID_COND   = "cond"	
	
	// For Macros
	ID_BACKQUOTE 		= "backquote"	
	ID_UNQUOTE   		= "unquote"
	ID_UNQUOTESPLICING  = "unquote_splicing"

	// New core axioms
	ID_TAG 	  = "tag"
	ID_TAGGED = "tagged"
	ID_SYM	  = "sym"
	ID_TYPE0  = "type"
	ID_REP    = "rep"
	
	// Ext ??
	ID_LABEL  = "label"		// name
	
	// Bool
	ID_T	  = "t"
	ID_NIL	  = "nil"
	
	// 
	ID_VAR	= "var"
	ID_ENV	= "env"
	ID_LET	= "let"
	ID_DEF	= "def"
	ID_FUNC = "func"		// was LAMBDA
	ID_MAC  = "mac"

)

// Core symbols of language
var (
	
	QUOTE 		= Sym(ID_QUOTE)
	ATOM 		= Sym(ID_ATOM)
	EQ 			= Sym(ID_EQ)
	CAR 		= Sym(ID_CAR)
	CDR 		= Sym(ID_CDR) 
	CONS 		= Sym(ID_CONS)
	COND 		= Sym(ID_COND) 
	
	BACKQUOTE 			= Sym(ID_BACKQUOTE)			// For Macros
	UNQUOTE   			= Sym(ID_UNQUOTE) 			// For Macros, 
	UNQUOTE_SPLICING   	= Sym(ID_UNQUOTESPLICING) 	// For Macros, 
		
	LABEL 		= Sym(ID_LABEL)
	FUNC 		= Sym(ID_FUNC) 
	MAC			= Tag(Sym(ID_FUNC), ID_MAC)
	T 			= Sym(ID_T) 		// TRUE
	NIL 		= Sym(ID_NIL) 		// NIL or FALSE
	
	TAG 	= Sym(ID_TAG)
	TAGGED 	= Sym(ID_TAGGED)
	SYM 	= Sym(ID_SYM)
	TYPE0 	= Sym(ID_TYPE0)		// TODO: Better name - type is reserved word in GO
	REP 	= Sym(ID_REP)	
	
	VAR = Sym(ID_VAR)
	ENV = Sym(ID_ENV)
	LET = Sym(ID_LET)
	DEF = Sym(ID_DEF)
	//MAC = Sym(ID_MAC)

)


