package spec

import (
	"testing"
)
	
func TestAtom(t *testing.T) {

	specs := []spec {
		{ "(atom 'a)"			, "t"	},
		{ "(atom '(a b))"		, "nil"	},
		{ "(atom (atom 'a))"	, "t"	},
		{ "(atom '(atom 'a))"	, "nil"	},
	}
	
	test(specs, t)
}
