package spec

import (
	"testing"
)
	
func TestLet(t *testing.T) {

	specs := []spec {
		{ "(let x 1 x)"					, "1"	},
		{ "(let x 1 (cons x x))"		, "(1 . 1)"	},
		{ "(let x 1 (list x 2))"		, "(1 2)"	},
	}
	
	test(specs, t)
}

func TestLet2(t *testing.T) {

	specs := []spec2 {
		{ "(mac foo (x y) `(list ,x ,y)"    , "(let x 1 (foo x x))"	, "(1 1)"	},
	}
	
	test2(specs, t)
}
