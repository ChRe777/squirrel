package spec

import (
	"testing"
)
	
func TestApply(t *testing.T) {

	specs := []spec {
//		{ "(apply +  '(1 2 4))" , "7"	},	// (+ 1 2 3)
//		{ "(apply car '(a b))"	 , "a"	},	// (car a b) -> error
	}
	
	test(specs, t)
}