package spec

import (
	"testing"
)
	
func TestList(t *testing.T) {
	
	specs := []spec {
		{ "(list)"					, "nil"			},
		{ "(list 1 2 3)"			, "(1 2 3)"		},
		{ "(list 'a (list 1 'c))"	, "(a (1 c))"	},
	}
	
	test(specs, t)
}


