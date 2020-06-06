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
		{ "(type (func (x) x))"	, "cons#func"	},	
	}
	
	test(specs, t)

}

func TestType2(t *testing.T) {

	specs := []spec2 {
		{ "(mac foo (x) `(no ,x))"	, "(type foo)"	, "cons#mac" },
	}
	
	test2(specs, t)
}
