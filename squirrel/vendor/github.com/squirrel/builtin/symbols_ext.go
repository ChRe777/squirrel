package builtin

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
	NO     = Sym(ID_NO    )
	NOT    = Sym(ID_NOT   )
	AND    = Sym(ID_AND   )
	PAIR   = Sym(ID_PAIR  )
	LIST   = Sym(ID_LIST  )	// TODO: More then two params
	ASSOC  = Sym(ID_ASSOC )
	APPEND = Sym(ID_APPEND)
)