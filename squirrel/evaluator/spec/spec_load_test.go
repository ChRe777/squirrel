package spec

import (
	"testing"
)
	
func TestLoad(t *testing.T) {

	specs := []spec2 {
		{ "(load \"data/macros.cell\")", "(apply 'list '(1 2))" , "(1 2)" 	},
		{ "(load \"data/macros.cell\")", "(when (is 'a 'a) 't)", "t" 		},
	}
	
	test2(specs, t)
}
