package spec

import (
	"testing"
)
	
func TestSym(t *testing.T) {

	specs := []spec {
		{ "(sym \"a\")"		, "a"									},
		{ "(sym 1)"			, "Error: \"no atom and/or string\""	},
		{ "(sym '(1 2))"	, "Error: \"no atom and/or string\""	},
	}
	
	test(specs, t)
}
