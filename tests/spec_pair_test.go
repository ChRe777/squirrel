package spec

import (
	"testing"
)
	
func TestPair(t *testing.T) {
	
	specs := []spec {
		{ "(pair '(x y) 	'(1 2))"		, "((x 1) (y 2))"				},
		{ "(pair '(x y . z) '(1 2 3 4))"	, "((x 1) (y 2) (z (3 4)))"		},
		{ "(pair '(x y z w) '(1 2 3))"		, "((x 1) (y 2) (z 3))"			},
	}
	
	test(specs, t)
}
