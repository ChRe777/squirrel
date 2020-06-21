package builtin

import (
	"github.com/mysheep/squirrel/evaluator/core"
)

//
// fns
//
const (

	ID_NO     = "no"    
	ID_NOT    = "not"   
	ID_AND    = "and"   
	ID_PAIR   = "pair"  
	ID_ASSOC  = "assoc" 
	ID_APPEND = "append"
	
)

var (

	NO     = core.Sym_(ID_NO    )
	NOT    = core.Sym_(ID_NOT   )
	AND    = core.Sym_(ID_AND   )
	PAIR   = core.Sym_(ID_PAIR  )
	ASSOC  = core.Sym_(ID_ASSOC )
	APPEND = core.Sym_(ID_APPEND)	
	
)

//
// fns2
//
const (

	ID_ENV	= "env"
	
	ID_VAR	= "var"
	ID_LET	= "let"
	ID_DEF	= "def"
	
	ID_FUNC = "func"
	ID_MAC  = "mac"
	
	ID_LIST	= "list"
	ID_DO   = "do"
	
	ID_COND = "cond"
	
)

var (	

	ENV 	= core.Sym_(ID_ENV	)
	
	VAR 	= core.Sym_(ID_VAR	)
	LET 	= core.Sym_(ID_LET	)
	DEF 	= core.Sym_(ID_DEF	)

	FUNC 	= core.Sym_(ID_FUNC	) 
	MAC		= core.Sym_(ID_MAC	)
	
	LIST 	= core.Sym_(ID_LIST	)
	DO		= core.Sym_(ID_DO	)
	
	COND	= core.Sym_(ID_COND	)
)

//
// macro
//
const (
	
	ID_BACKQUOTE 		= "backquote"	
	ID_UNQUOTE   		= "unquote"
	ID_UNQUOTESPLICING  = "unquote_splicing"

)

var (

	BACKQUOTE 			= core.Sym_(ID_BACKQUOTE)			// For Macros
	UNQUOTE   			= core.Sym_(ID_UNQUOTE) 			// For Macros, 
	UNQUOTE_SPLICING   	= core.Sym_(ID_UNQUOTESPLICING) 	// For Macros, 

)