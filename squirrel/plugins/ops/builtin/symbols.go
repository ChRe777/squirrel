package builtin

import (
	"github.com/mysheep/squirrel/evaluator/core"
)

const (

	ID_NO     = "no"    
	ID_NOT    = "not"   
	ID_AND    = "and"   
	ID_PAIR   = "pair"  
	ID_LIST   = "list"
	ID_ASSOC  = "assoc" 
	ID_APPEND = "append"
	
	ID_CAAR   = "caar"  
	ID_CADR   = "cadr" 
	ID_CDDR   = "cddr"  
	ID_CADAR  = "cadar" 
	ID_CDDDR  = "cdddr" 
	ID_CADDR  = "caddr" 
	ID_CADDAR = "caddar"
	ID_CADDDR = "cadddr"

)

var (

	NO     = core.Sym_(ID_NO    )
	NOT    = core.Sym_(ID_NOT   )
	AND    = core.Sym_(ID_AND   )
	PAIR   = core.Sym_(ID_PAIR  )
	LIST   = core.Sym_(ID_LIST  )
	ASSOC  = core.Sym_(ID_ASSOC )
	APPEND = core.Sym_(ID_APPEND)	
	
	CAAR   = core.Sym_(ID_CAAR  )
	CADR   = core.Sym_(ID_CADR  )
	CDDR   = core.Sym_(ID_CDDR  )
	CADAR  = core.Sym_(ID_CADAR )
	CDDDR  = core.Sym_(ID_CDDDR )
	CADDR  = core.Sym_(ID_CADDR )
	CADDAR = core.Sym_(ID_CADDAR)
	CADDDR = core.Sym_(ID_CADDDR)

)