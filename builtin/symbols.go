package builtin

import (
	"github.com/mysheep/squirrel/core"
)

//
// functions group 1
//
const (
	ID_NO     	= "no"
	ID_NOT    	= "not"
	ID_AND    	= "and"
	ID_PAIR   	= "pair"
	ID_ASSOC  	= "assoc"

	ID_APPEND 	= "append"
)
//
// functions group 2
//
const (
	ID_ENV		= "env"
	ID_VAR		= "var"
	ID_LET		= "let"
	ID_DEF		= "def"
	ID_FUNC 	= "func"

	ID_MAC  	= "mac"
	ID_LIST		= "list"
	ID_DO   	= "do"
	ID_COND 	= "cond"
)
//
// functions for macro
//
const (
	ID_BACKQUOTE 		= "backquote"	
	ID_UNQUOTE   		= "unquote"
	ID_UNQUOTESPLICING  = "unquote_splicing"
)

var (
	NO     		= core.Sym_(ID_NO    )
	NOT    		= core.Sym_(ID_NOT   )
	AND    		= core.Sym_(ID_AND   )
	PAIR   		= core.Sym_(ID_PAIR  )
	ASSOC  		= core.Sym_(ID_ASSOC )

	APPEND 		= core.Sym_(ID_APPEND)
)

var (
	COND		= core.Sym_(ID_COND	)
	DEF 		= core.Sym_(ID_DEF	)
	DO			= core.Sym_(ID_DO	)
	ENV 		= core.Sym_(ID_ENV	)
	FUNC 		= core.Sym_(ID_FUNC	)

	LET 		= core.Sym_(ID_LET	)
	LIST 		= core.Sym_(ID_LIST	)
	MAC			= core.Sym_(ID_MAC	)
	VAR 		= core.Sym_(ID_VAR	)
)

var (
	BACKQUOTE 			= core.Sym_(ID_BACKQUOTE)			// For Macros
	UNQUOTE   			= core.Sym_(ID_UNQUOTE) 			// For Macros, 
	UNQUOTE_SPLICING   	= core.Sym_(ID_UNQUOTESPLICING) 	// For Macros, 
)
