package spec

import (
	"testing"
)
	
func TestList(t *testing.T) {
	
	specs := []spec {
		{ "(list '+ 1 2)"		, "(+ 1 2)"		},
	}
	
	test(specs, t)
}


