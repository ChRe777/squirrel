package spec

import (
	"testing"
)
	
func TestMac(t *testing.T) {

	specs := []spec2 {
		{ "(mac foo (x  y)  `(list ,x ,y))"		, "(foo 1 2)"					, "(1 2)"   		},
		{ "(mac bar (xs)    `(list ,@xs))" 		, "(bar '(1 2 3))"				, "(1 2 3)" 		},
		{ "(mac biz (xs ys) `(list ,@xs ,@ys))" , "(biz '(1 2 3) '(4 5 6))"		, "(1 2 3 4 5 6)" 	},
		{ "(mac zap (xs ys) `(list ,@xs ,@ys))" , "(zap '(1 2 3) '('a 'b 'c))"	, "(1 2 3 a b c)" 	},
	}
	
	test2(specs, t)
}