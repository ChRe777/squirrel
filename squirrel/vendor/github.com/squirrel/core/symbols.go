package core

// Identifiers of built-in core symbols
const (

	// 7 Core Primitives
	//
	ID_QUOTE  = "quote"
	ID_ATOM   = "atom"
	ID_IS	  = "is"		// was "eq"
	ID_CAR    = "car"		// can be first		// go slices  xs[0:]
	ID_CDR	  = "cdr"		// can be rest		// go slices  xs[:len-1]
	ID_CONS	  = "cons"	
	ID_COND   = "cond"	
	
	// For Macros
	//
	ID_BACKQUOTE 		= "backquote"	
	ID_UNQUOTE   		= "unquote"
	ID_UNQUOTESPLICING  = "unquote_splicing"

	// New core axioms
	//
	ID_TAG 	  = "tag"
	ID_TAGGED = "tagged"
	ID_SYM	  = "sym"
	ID_TYPE0  = "type"
	ID_REP    = "rep"
		
	// Boolean
	//
	ID_T	  = "t"
//	ID_NIL	  = "nil	// Nil in generator
	
	// Builtin-Core
	//
	ID_ENV	= "env"
	
	ID_VAR	= "var"
	ID_LET	= "let"
	
	ID_DEF	= "def"
	
	ID_FUNC = "func"
	ID_MAC  = "mac"
	
	ID_LIST	= "list"
	ID_LOAD = "load"

)

// Core symbols of language
var (	
	QUOTE 	= Sym(ID_QUOTE)
	ATOM 	= Sym(ID_ATOM)
	IS 		= Sym(ID_IS)							// was EQ
	CAR 	= Sym(ID_CAR)
	CDR 	= Sym(ID_CDR) 
	CONS 	= Sym(ID_CONS)
	COND 	= Sym(ID_COND) 
)
	
// Macros
//
var (
	BACKQUOTE 			= Sym(ID_BACKQUOTE)			// For Macros
	UNQUOTE   			= Sym(ID_UNQUOTE) 			// For Macros, 
	UNQUOTE_SPLICING   	= Sym(ID_UNQUOTESPLICING) 	// For Macros, 
)
	
// Boolean
//
var (		
	T 		= Sym(ID_T) 							// TRUE
	NIL 	= Nil_() 								// NIL or FALSE
)

// Builtin Core
//
var (	
	VAR 	= Sym(ID_VAR)
	ENV 	= Sym(ID_ENV)
	LET 	= Sym(ID_LET)
	DEF 	= Sym(ID_DEF)

	FUNC 	= Sym(ID_FUNC) 
	MAC		= Sym(ID_MAC)
	
	LIST 	= Sym(ID_LIST)
	LOAD 	= Sym(ID_LOAD)
)


// Tagging	(Used different)
var (
	TAG 	= Sym(ID_TAG)
	TAGGED 	= Sym(ID_TAGGED)
	SYM 	= Sym(ID_SYM)
	TYPE0 	= Sym(ID_TYPE0)							
	REP 	= Sym(ID_REP)	
)
