package spec

import (
	"testing"
)
	
func TestNumber(t *testing.T) {
	
	specs := []spec {
		{ "(is (add 3 4) (sub 10 3))"	, "t"	},
	}
	
	test(specs, t)
}
