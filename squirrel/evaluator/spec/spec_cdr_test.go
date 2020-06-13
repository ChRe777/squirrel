package spec

import (
	"testing"
)
	
func TestCdr(t *testing.T) {
	
	specs := []spec {
		{ "(cdr '())"		, "nil"},
		{ "(cdr '(a b c))"	, "(b c)"},
	}
	
	test(specs, t)
}