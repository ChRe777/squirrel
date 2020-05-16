package spec

import (
	"testing"
)
	
func TestCar(t *testing.T) {

	specs := []spec {
		{ "(car '1)"	, "Error: \"Can't take car of 1\""},
		{ "(car nil)"	, "nil"},
		{ "(car '())"	, "nil"},
		{ "(car '(a b))", "a"},
	}
	
	test(specs, t)
}