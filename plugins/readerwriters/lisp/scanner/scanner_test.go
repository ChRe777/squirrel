package scanner

import (
	"testing"
)

func TestGetSym(t *testing.T) {

	s := []byte("(+123.4 \"foo bar\" 'biz . `foo ,bar ,@foo)")
	
	specs := []spec {
		{ "(", 		Lparen 			},
		{ "+123.4", Number 			},
		{ "foo bar", String 		},
		{ "'", 		Quote  			},
		{ "biz", 	Symbol 			},
		{ ".", 		Dot 			},
		{ "`", 		Backquote  		},
		{ "foo", 	Symbol 			},
		{ ",", 		Unquote  		},
		{ "bar", 	Symbol 			},
		{ ",@", 	UnquoteSplicing },
		{ "foo", 	Symbol 			},
		{ ")", 		Rparen 			},
	}
	
	Init(s)
	
	for _, spec := range specs {
	
		GetSym(); id := asStr(Id); sym := Sym
		
		if isNotEq(spec, id, sym) {
			t.Errorf("got: Id %v Sym %v, want: Id %v, Sym %v", id, sym, spec.Id, spec.Sym)
		}
	}
	
}

func TestGetSym2(t *testing.T) {

	s := []byte("(-123.4)")
	
	specs := []spec {
		{ "(", 		Lparen 	},
		{ "-123.4", Number 	},
		{ ")", 		Rparen 	},

	}
	
	Init(s)
	
	for _, spec := range specs {
	
		GetSym(); id := asStr(Id); sym := Sym
		
		if isNotEq(spec, id, sym) {
			t.Errorf("got: Id %v Sym %v, want: Id %v, Sym %v", id, sym, spec.Id, spec.Sym)
		}
	}
	
}