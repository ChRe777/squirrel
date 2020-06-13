package spec

import (
	"testing"
)

func TestCond(t *testing.T) {
	
	specs := []spec {
		{ "(cond ((eq 'a 'b) 'first) ('t 'second))"		  , "second"	},
		{ "(cond ((eq 'a 'b) 'first) ((atom 'a) 'second))", "second"	},
		{ "(cond ('t 'first) ('t 'second))"				  , "first" 	},
	}
	
	test(specs, t)
}