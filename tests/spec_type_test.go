package spec

import (
	"testing"
)

func TestType (t *testing.T) {
	
	specs := []spec {
		{ "(type  1)"			, "number" 		},  
		{ "(type 'a)"			, "symbol"		},
		{ "(type \"a\")"		, "string"		},
		{ "(type '(a b))"		, "cons"		},
		{ "(type (func (x) x))"	, "func#func"	},	
	}
	
	test(specs, t)
}

func TestMacType(t *testing.T) {

	specs := []spec2 {
		{ "(mac foo (x) `(no ,x))"	, "(type foo)"	, "func#mac" },
	}
	
	test2(specs, t)
}

func TestFuncType(t *testing.T) {

	specs := []spec2 {
		{ "(def foo (x) (cons x x))"	, "(type foo)"	, "func#func" },
	}
	
	test2(specs, t)
}

