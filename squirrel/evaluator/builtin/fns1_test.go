package builtin

import (
	"testing"
)

import (
	"github.com/mysheep/squirrel/types"
	"github.com/mysheep/squirrel/evaluator/core"
	"github.com/mysheep/squirrel/plugins/readerwriters/lisp/parser"	

)

func TestPair(t *testing.T) {

	specs := []struct {
		xs		string
		ys		string
		want  	string
	} {
		{"(x y . z)", "(1 2 3 4)", "((x 1) (y 2) (z (3 4)))"},
		{"(x y z w)", "(1 2 . 3)", "((x 1) (y 2) ((z w) 3)))"},
		{"(x y z)"	, "(1 2 3 4)", "((x 1) (y 2) (z 3)))"	},
		{"(x y z w)", "(1 2 3)"	 , "((x 1) (y 2) (z 3)))"	},
		{"()"		, "(1 2)"	 , "()"	},
		{"(x y)"	, "()"	 	 , "()"	},
		{"()"	    , "()"	 	 , "()"	},
	}	
	
	for _, spec := range specs {
	
		xs := parser.Parse([]byte(spec.xs))
		ys := parser.Parse([]byte(spec.ys))
		
		got := Pair(xs, ys)
		want := parser.Parse([]byte(spec.want))
			
		if got.NotEqual(want) {
			t.Errorf("TestPair - got: %v, want: %v", got, spec.want)
		}
	}
}

// -------------------------------------------------------------------------------------------------

func cons(x, y *types.Cell) *types.Cell {
	return core.Cons(x, y)
}

func num(x string) *types.Cell {
	return core.Num_(x)
}

func sym(x string) *types.Cell {
	return core.Sym_(x)
}