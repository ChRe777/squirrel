package core

// Core
//
const (

	// 7 Core Primitives
	//
	ID_QUOTE  	= "quote"
	ID_ATOM   	= "atom"
	ID_IS	  	= "is"				// was "eq"
	ID_CAR    	= "car"				// can be first		// go slices  xs[0:]
	ID_CDR	  	= "cdr"				// can be rest		// go slices  xs[:len-1]
	ID_CONS	  	= "cons"	
	ID_COND   	= "cond"	
)

// Core
//
var (	
	QUOTE 		= Sym_(ID_QUOTE)
	ATOM 		= Sym_(ID_ATOM)
	IS 			= Sym_(ID_IS)		// was EQ
	CAR 		= Sym_(ID_CAR)
	CDR 		= Sym_(ID_CDR) 
	CONS 		= Sym_(ID_CONS)
	COND 		= Sym_(ID_COND) 
)

// CarCdrs
//
const (
	ID_CAAR   	= "caar"  
	ID_CADR   	= "cadr" 
	ID_CDDR   	= "cddr"  
	ID_CADAR  	= "cadar" 
	ID_CDDDR  	= "cdddr" 
	ID_CADDR  	= "caddr" 
	ID_CADDAR 	= "caddar"
	ID_CADDDR 	= "cadddr"
)

// CarCdrs
//
var (
	CAAR   		= Sym_(ID_CAAR  )
	CADR   		= Sym_(ID_CADR  )
	CDDR   		= Sym_(ID_CDDR  )
	CADAR  		= Sym_(ID_CADAR )
	CDDDR  		= Sym_(ID_CDDDR )
	CADDR  		= Sym_(ID_CADDR )
	CADDAR 		= Sym_(ID_CADDAR)
	CADDDR 		= Sym_(ID_CADDDR)
)

// Boolean
//
const (
	ID_T	  	= "t"
//	ID_NIL	  	= "nil				// Nil in generator
)

// Boolean
//
var (		
	T 			= Sym_(ID_T) 		// TRUE
	NIL 		= Nil_() 			// NIL or FALSE
)

// Extended core
//
const (
	ID_TYPE 	= "type"
	ID_FUNC		= "func"
)

// Extended core
//
var (		
	TYPE 		= Sym_(ID_TYPE)
	FUNC 		= Sym_(ID_FUNC)
	
//	TAG 		= Sym_(ID_TAG)		// TODO: Tagging
//	TAGGED 		= Sym_(ID_TAGGED)
//	SYM 		= Sym_(ID_SYM)
//	TYPE0 		= Sym_(ID_TYPE0)							
//	REP 		= Sym_(ID_REP)	
	
)



