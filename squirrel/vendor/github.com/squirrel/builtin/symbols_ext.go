package builtin

import (
	"github.com/squirrel/core"
)

const (
	ID_NO     = "no"    
	ID_NOT    = "not"   
	ID_AND    = "and"   
	ID_PAIR   = "pair"  
	ID_LIST   = "list"  	// TODO: More then two params
	ID_ASSOC  = "assoc" 
	ID_APPEND = "append"
)

var (
	NO     = core.Sym(ID_NO    )
	NOT    = core.Sym(ID_NOT   )
	AND    = core.Sym(ID_AND   )
	PAIR   = core.Sym(ID_PAIR  )
	LIST   = core.Sym(ID_LIST  )	// TODO: More then two params
	ASSOC  = core.Sym(ID_ASSOC )
	APPEND = core.Sym(ID_APPEND)
)