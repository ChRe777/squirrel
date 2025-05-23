package spec

import (
	"testing"
)

func TestMac(t *testing.T) {

	specs := []spec2{
		{"(mac foo (x  y)    `(list ,x ,y))"	, "(foo 1 2)"					, "(1 2)"			},
		{"(mac bar (xs)      `(list ,@xs))"		, "(bar (1 2 3))"				, "(1 2 3)"			},
		{"(mac biz (xs ys)   `(list ,@xs ,@ys))", "(biz (1 2 3) (4 5 6))"		, "(1 2 3 4 5 6)"	},
		{"(mac zap (xs ys)   `(list ,@xs ,@ys))", "(zap (1 2 3) ('a 'b 'c))"	, "(1 2 3 a b c)"	},
		{"(mac apply (fn xs) `(,fn ,@xs))"		, "(apply list (1 2 3))"		, "(1 2 3)"			},
	}

	test2(specs, t)
}

func TestMac2(t *testing.T) {

	specs := []spec2{
		{"(mac when (c e) `(cond (,c ,e)('t nil)) )", "(when (is 'a 'a) 't)"	, "t"},
		{"(mac when (c e) `(cond (,c ,e)('t nil)) )", "(when (is 'a 'b) nil)"	, "nil"},
	}

	test2(specs, t)
}

func TestMac3(t *testing.T) {

	specs := []spec2{
		{"(mac foo (x y) `(list ,x ,y))", "(let x 1 (foo x x))", "(1 1)"},
	}

	test2(specs, t)
}
