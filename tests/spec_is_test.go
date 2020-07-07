package spec

import (
	"testing"
)
	
func TestIs(t *testing.T) {
	
	specs := []spec {
		{ "(is 'a 'a)"			, "t"	},
		{ "(is 'a 'b)"			, "nil"	},
		{ "(is '() '())"		, "t" 	},
		{ "(is '(a b) '(a b))"	, "t" 	},	// nil in Arc - (iso '(a) '(a) t)
	}
	
	test(specs, t)
}
