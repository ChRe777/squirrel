package spec

import (
	"testing"
)

func TestTagging (t *testing.T) {
	
	specs := []spec {
		{ "(type (tag 'num 'a))"	, "num" },  
		{ "(rep  (tag 'num 'a))"	, "a"	},
		{ "(type 'a)"				, "sym"	},
		{ "(type '(a b))"			, "cons"},	
	}
	
	test(specs, t)
}
