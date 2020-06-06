package spec

import (
	"testing"
)
	
func TestDo(t *testing.T) {

	exp1 := "(do " +
	"	(mac foo (x y)  `(list ,x ,y))" +
	"	(mac bar (x y)  `(list ,y ,x))" +
	")"

	specs := []spec2 {
		{ exp1, "(list (foo 1 2) (bar 1 2))"	, "((1 2)(2 1))" },

	}
	
	test2(specs, t)
}
