package spec

import (
	"testing"
)
	
func TestEq(t *testing.T) {
	
	specs := []spec {
		{ "(eq 'a 'a)"	, "t"	},
		{ "(eq 'a 'b)"	, "nil"	},
		{ "(eq '() '())", "t" 	},
	}
	
	test(specs, t)
}
