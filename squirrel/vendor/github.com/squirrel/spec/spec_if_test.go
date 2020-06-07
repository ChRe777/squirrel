package spec

import (
	"testing"
)
	
func TestIf(t *testing.T) {

	specs := []spec2 {
		{ "(mac if (c a b) `(cond (,c ,a) ('t ,b)))"    , "(if (is 'a 'a) 't nil)"	, "t"	},
		{ "(mac if (c a b) `(cond (,c ,a) ('t ,b)))"    , "(if (is 'a 'b) 't nil)"	, "nil"	},
	}
	
	test2(specs, t)
}