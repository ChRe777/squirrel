package macro

const (
	
	ID_BACKQUOTE 		= "backquote"	
	ID_UNQUOTE   		= "unquote"
	ID_UNQUOTESPLICING  = "unquote_splicing"

)

var (

	BACKQUOTE 			= Sym_(ID_BACKQUOTE)			// For Macros
	UNQUOTE   			= Sym_(ID_UNQUOTE) 			// For Macros, 
	UNQUOTE_SPLICING   	= Sym_(ID_UNQUOTESPLICING) 	// For Macros, 

)