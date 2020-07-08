package spec

import (
	"testing"
)
	
func TestVar(t *testing.T) {

	specs := []spec2 {
		{ "(var a 1)"    			, "a"	, "1"		},
		{ "(var xs '(1 2 3))"   	, "xs"	, "(1 2 3)"	},
		{ "(var a \"a\")"   		, "a"	, "\"a\""	},
		
		// TODO: FIX BUG !!!
		{ "(var fn (func (x) x))"   , "fn"	, "\"a\""	},
	}
	
	test2(specs, t)
}
