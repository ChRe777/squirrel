package spec

import (
	"bytes"
	"testing"
)

import (
	"github.com/squirrel/builtin"
	"github.com/squirrel/parser"
)

func createList (fns []string) []byte {
	var b bytes.Buffer
	b.WriteRune('('); for _, fn := range fns { b.WriteString(fn) }; b.WriteRune(')')
	return b.Bytes()
}

func TestEnvironment(t *testing.T) {

	env := parser.Parse(createList(builtin.Env()));
	
	specs := []spec {
		{ "(no '())					   	  "	, "t"					},
		{ "(and 't 't)				      "	, "t"					}, 
		{ "(not 't)				      	  "	, "nil"					},
		{ "(append '(a b)   '(c d)  )	  "	, "(a b c d)"			},
		{ "(list 'a 'b)  				  "	, "(a b)"	    		},
		{ "(pair '(a b c) '(x y z))	      " , "((a x) (b y) (c z))"	},
		{ "(assoc 'b '((a 1) (b 2) (c 3)))"	, "2"					},

	}
		
	testWithEnv(specs, t, env)
	
}