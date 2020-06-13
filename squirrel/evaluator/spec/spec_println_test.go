package spec

import (
	"testing"
)
	
func TestPrintln(t *testing.T) {
	
	specs := []spec {
//		{ "(println 'a)"			, "a"		},
//		{ "(println '(a b))"		, "(a b)" 	},
//		{ "(println \"foo\" 1 't)"	, "\"foo\"" 	},
	}
	
	test(specs, t)
}
