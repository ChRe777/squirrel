package functions

import (
	"github.com/mysheep/squirrel/core"
)

// functions
//
const (
	ID_ADD  	= "add"
	ID_SUB		= "sub"
	ID_DIV  	= "div"
	ID_MUL  	= "mul"
)

//	------------------------------------------------------------------------------------------------

// functions
//
var (	
	ADD 		= core.Sym_(ID_ADD)
	SUB			= core.Sym_(ID_SUB)
	DIV 		= core.Sym_(ID_DIV)
	MUL 		= core.Sym_(ID_MUL)
)


