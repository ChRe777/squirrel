package scanner

import (
	"testing"
)

func TestGetSym(t *testing.T) {

	s := []byte("(123.4 \"foo\" 'biz `foo ,bar)")
	
	specs := []spec {
		{ "(", 		Lparen 		},
		{ "123.4", 	Number 		},
		{ "foo", 	String 		},
		{ "'", 		Quote  		},
		{ "biz", 	Symbol 		},
		{ "`", 		Backquote  	},
		{ "foo", 	Symbol 		},
		{ ",", 		Unquote  	},
		{ "bar", 	Symbol 		},
		{ ")", 		Rparen 		},
	}
	
	Init(s)
	
	for _, spec := range specs {
	
		GetSym(); id := asStr(Id); sym := Sym
		
		if isNotEq(spec, id, sym) {
			t.Errorf("got: Id %v Sym %v, want: Id %v, Sym %v", id, sym, spec.Id, spec.Sym)
		}
	}
	
}