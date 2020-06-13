package spec

import (
	"testing"
)
	
func TestBackquote(t *testing.T) {

	specs := []spec2 {
		{ "(var x 1      )"	, "`(,x)"				, "(1)"   			},
		{ "(var xs '(1 2))"	, "`(,@xs)"				, "(1 2)"   		},
		{ "(var xs '(1 2))"	, "`(a ,@xs b)"			, "(a 1 2 b)"   	},
		{ "(var xs '(1 2))"	, "`(a (b ,@xs c) d)"	, "(a (b 1 2 c) d)" },
		{ "(var xs '(1 2))"	, "`(,@xs ,@xs)"		, "(1 2 1 2)"   	},
	}
	
	test2(specs, t)
}