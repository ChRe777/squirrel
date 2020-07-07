package spec

import (
	"testing"
)
	
func TestQuote(t *testing.T) {
	
	specs := []spec {
		{ "'(a)"		, "(a)"		},
		{ "'(a b)"		, "(a b)" 	},
		{ "(quote ())"	, "nil" 	},
		{ "(quote nil)"	, "nil" 	},
	}
	
	test(specs, t)
}
