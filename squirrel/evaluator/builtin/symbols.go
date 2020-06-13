package builtin

import (
	"github.com/mysheep/squirrel/evaluator/core"
)

const (
	ID_NO     = "no"    
	ID_NOT    = "not"   
	ID_AND    = "and"   
	ID_PAIR   = "pair"  
	ID_LIST   = "list"  			// TODO: More then two params
	ID_ASSOC  = "assoc" 
	ID_APPEND = "append"
)

var (
	NO     = core.Sym_(ID_NO    )
	NOT    = core.Sym_(ID_NOT   )
	AND    = core.Sym_(ID_AND   )
	PAIR   = core.Sym_(ID_PAIR  )
	LIST   = core.Sym_(ID_LIST  )	// TODO: More then two params
	ASSOC  = core.Sym_(ID_ASSOC )
	APPEND = core.Sym_(ID_APPEND)	
)