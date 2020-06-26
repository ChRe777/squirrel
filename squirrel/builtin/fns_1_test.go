package builtin

import (
	"testing"
)

import (
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

func TestNo(t *testing.T) {

	specs := []struct {
		xs		string
		want  	string
	} {
		{"nil"	 	, "t"},
		{"()"  		, "t"},
		{"1"   		, "nil"},
		{"(1 2)"  	, "nil"},
	}	
	
	for _, spec := range specs {
	
		xs := parser.Parse([]byte(spec.xs))
		
		got := No(xs)
		want := parser.Parse([]byte(spec.want))
			
		if got.NotEqual(want) {
			t.Errorf("TestNo - got: %v, want: %v", got, spec.want)
		}
	}
}

// -------------------------------------------------------------------------------------------------

func TestNot(t *testing.T) {

	specs := []struct {
		x		string
		want  	string
	} {
		{"t"	 	, "nil"},
		{"nil"   	, "t"},
	}	
	
	for _, spec := range specs {
	
		x := parser.Parse([]byte(spec.x))
		
		got := Not(x)
		want := parser.Parse([]byte(spec.want))
			
		if got.NotEqual(want) {
			t.Errorf("TestNot - got: %v, want: %v", got, spec.want)
		}
	}
}

// -------------------------------------------------------------------------------------------------

func TestAnd(t *testing.T) {

	specs := []struct {
		x		string
		y		string
		want  	string
	} {
		{"t", "t"	 , "t"  },
		{"nil", "t"	 , "nil"},
		{"t", "nil"  , "nil"},
	}	
	
	for _, spec := range specs {
	
		x := parser.Parse([]byte(spec.x))
		y := parser.Parse([]byte(spec.y))
		
		got := And(x,y)
		want := parser.Parse([]byte(spec.want))
			
		if got.NotEqual(want) {
			t.Errorf("TestAnd - got: %v, want: %v", got, spec.want)
		}
	}
}