package spec

import (
	"testing"
)

func TestClosure(t *testing.T) {

	specs := []spec2{
		{"(def closure (x) (func (y) (cons x y)))"	, "(closure 1)"		, "(func#func (y) (cons x y) ((x 1)))"	},
		{"(def closure (x) (func (y) (cons x y)))"	, "((closure 1) 2)"	, "(1 . 2)"								},
	}

	test2(specs, t)
}
