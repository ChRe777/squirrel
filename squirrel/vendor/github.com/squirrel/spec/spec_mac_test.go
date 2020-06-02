package spec

import (
	"testing"
)
	
func TestMac(t *testing.T) {

	specs := []spec2 {
		{ "(mac foo   (x  y) `(list ,x ,y))", "(foo 1 2)"			, "(list 1 2)"	  },
		{ "(mac apply (f xs) `(,f ,@xs))"   , "(apply 'no '(1 2 3))", "(nil nil nil)" },
	}
	
	test2(specs, t)
}