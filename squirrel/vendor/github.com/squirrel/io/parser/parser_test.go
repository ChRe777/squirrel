package parser

import (
//	"fmt"
	"testing"
)

import (
	"github.com/squirrel/generator"
	"github.com/squirrel/types"
)

func TestParseAtom(t *testing.T) {

	specs := []spec {
		{ "t"		, generator.Sym("t") 	},
		{ "foo"		, generator.Sym("foo") 	},
		{ "\"a\""	, generator.Str("\"a\"") 	},		// TODO: STRING are with " stored =???
		{ "-1.23"	, generator.Num("-1.23") },
	} 

	test(specs, t)
}

func TestParseCons(t *testing.T) {

	specs := []spec {
		{ "(1)"		    , generator.Cons(generator.Num("1"), generator.Nil())},
		{ "(1 2)"		, generator.Cons(generator.Num("1"), generator.Cons(generator.Num("2"), generator.Nil())) },
		{ "(a . b)"		, generator.Cons(generator.Sym("a"), generator.Sym("b")) },
		{ "(a b . c)"	, generator.Cons(generator.Sym("a"), generator.Cons(generator.Sym("b"), generator.Sym("c"))) },
		{ "(a b . c d)"	, generator.Cons(generator.Sym("a"), generator.Cons(generator.Sym("b"), generator.Sym("c"))) },

	} 

	test(specs, t)
}

func TestParseQuote(t *testing.T) {
	specs := []spec {
		{ "'a" 	, generator.Cons( generator.Sym("quote") , generator.Cons(generator.Sym("a"), generator.Nil())) },
	}

	test(specs, t)
}

func TestParseBackQuote(t *testing.T) {
	specs := []spec {
		{ "`a" 	, generator.Cons( generator.Sym("backquote") , generator.Cons(generator.Sym("a"), generator.Nil())) },
	}

	test(specs, t)
}

// ------------------ LIBRARY --------------------------------------------------

type spec struct {
	expression string
	want	   *types.Cell
}

func test(specs []spec, t *testing.T) {
	
	for _, spec := range specs {
	
		e := []byte(spec.expression); got := Parse(e)
				
		if got.NotEqual(spec.want) {
			t.Errorf("Spec of expression \"%v\" was incorrect, got: \"%v\", want: \"%v\"", spec.expression, got, spec.want)
		}
	}		
}