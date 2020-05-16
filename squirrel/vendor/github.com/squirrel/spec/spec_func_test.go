package spec

import (
	"testing"
)
	
func TestFunc(t *testing.T) {
	
	specs := []spec {
		{ "((func (x)   (cons x '(b)))   'a )"			, "(a b)"  },  
		{ "((func (x y) (cons x (cdr y)))  'z '(a b c))"	, "(z b c)"},
	}
	
	test(specs, t)
}
