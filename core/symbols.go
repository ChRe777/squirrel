package core

// Core functions
//
const (
	ID_QUOTE  	= "quote"
	ID_ATOM   	= "atom"
	ID_IS	  	= "is"				// was "eq"
	ID_CAR    	= "car"				// can be first	 like go slices xs[0:]
	ID_CDR	  	= "cdr"				// can be rest   like go slices xs[:len-1]
	ID_CONS	  	= "cons"
    ID_COND   	= "cond"			
)

// Boolean
//
const (
	ID_T	  	= "t"				// true
)

// Extended core
//
const (
	ID_TYPE 	= "type"
	ID_FUNC		= "func"
	ID_SYM		= "sym"
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

//	------------------------------------------------------------------------------------------------

// Core functions
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
var (		
	T 			= Sym_(ID_T) 		// TRUE
	NIL 		= Nil_() 			// NIL or FALSE
)

// Extended core
//
var (		
	TYPE 		= Sym_(ID_TYPE)
	FUNC 		= Sym_(ID_FUNC)
	SYM 		= Sym_(ID_SYM)		// Create symbol from string
//	TAG 		= Sym_(ID_TAG)
//	TAGGED 		= Sym_(ID_TAGGED)
//	TYPE0 		= Sym_(ID_TYPE0)							
//	REP 		= Sym_(ID_REP)
)



