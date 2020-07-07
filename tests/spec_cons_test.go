package spec

import (
	"testing"
)
	
func TestCons(t *testing.T) {
	
	specs := []spec {
		{ "(cons ' a ' (b c)))"						, "(a b c)"		},
		{ "(cons 'a (cons 'b (cons 'c  '()))) )"	, "(a b c)"		},
		{ "(car  (cons 'a '(b c) ))"				, "a"			},
		{ "(cdr  (cons 'a '(b c) ))"				, "(b c)"   	},
		{ "(cons 'a 'b)"							, "(a . b)" 	},
		{ "(cons 'a (cons 'b 'c))"					, "(a b . c)" 	},
		{ "(cons 'a (cons 'b (cons 'c 'd)))"		, "(a b c . d)" },		
	}
	
	test(specs, t)
}