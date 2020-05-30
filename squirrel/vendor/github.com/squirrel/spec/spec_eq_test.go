package spec

import (
	"testing"
)
	
func TestIs(t *testing.T) {
	
	specs := []spec {
		{ "(is 'a 'a)"	, "t"	},
		{ "(is 'a 'b)"	, "nil"	},
		{ "(is '() '())", "t" 	},
	}
	
	test(specs, t)
}
