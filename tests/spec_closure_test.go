package spec

import (
	"testing"
)

func TestClosure(t *testing.T) {

	specs := []spec2{
		{"(def closure (x) (func (y) (cons x y)))"				, "((closure 1) 2)"		, "(1 . 2)"	},
		{"(def closure (x) (func (y) (func (z) (list x y z))))"	, "(((closure 1) 2) 3)"	, "(1 2 3)"	},
	}

	test2(specs, t)
}
