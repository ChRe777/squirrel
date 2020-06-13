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
	
	// Extra core
	//
	ID_TYPE 	= "type"
	ID_PRINTLN 	= "println"
		
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
	ID_DO   = "do"

)

// Core symbols of language
var (	
	QUOTE 	= Sym_(ID_QUOTE)
	ATOM 	= Sym_(ID_ATOM)
	IS 		= Sym_(ID_IS)							// was EQ
	CAR 	= Sym_(ID_CAR)
	CDR 	= Sym_(ID_CDR) 
	CONS 	= Sym_(ID_CONS)
	COND 	= Sym_(ID_COND) 
)
	
// Macros
//
var (
	BACKQUOTE 			= Sym_(ID_BACKQUOTE)			// For Macros
	UNQUOTE   			= Sym_(ID_UNQUOTE) 			// For Macros, 
	UNQUOTE_SPLICING   	= Sym_(ID_UNQUOTESPLICING) 	// For Macros, 
)
	
	
// Extended core
//
var (		
	TYPE 		= Sym_(ID_TYPE)
	PRINTLN 	= Sym_(ID_PRINTLN)
)

// Boolean
//
var (		
	T 		= Sym_(ID_T) 							// TRUE
	NIL 	= Nil_() 								// NIL or FALSE
)

// Builtin Core
//
var (	
	VAR 	= Sym_(ID_VAR)
	ENV 	= Sym_(ID_ENV)
	LET 	= Sym_(ID_LET)
	DEF 	= Sym_(ID_DEF)

	FUNC 	= Sym_(ID_FUNC) 
	MAC		= Sym_(ID_MAC)
	
	LIST 	= Sym_(ID_LIST)
	
	LOAD 	= Sym_(ID_LOAD)
	DO		= Sym_(ID_DO)
)


// Tagging	(Used different)
var (
	TAG 	= Sym_(ID_TAG)
	TAGGED 	= Sym_(ID_TAGGED)
	SYM 	= Sym_(ID_SYM)
	TYPE0 	= Sym_(ID_TYPE0)							
	REP 	= Sym_(ID_REP)	
)


