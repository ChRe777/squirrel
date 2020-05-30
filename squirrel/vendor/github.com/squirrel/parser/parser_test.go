package parser

import (
	"fmt"
	"testing"
)

import (
	"github.com/squirrel/generator"
	"github.com/squirrel/types"
)

func TestParseAtom(t *testing.T) {

	specs := []spec {
		{ "t"		, generator.Atom("t"	, types.SYMBOL) },
		{ "foo"		, generator.Atom("foo"	, types.SYMBOL) },
		{ "\"a\""	, generator.Atom("\"a\"", types.STRING) },
		{ "1.23"	, generator.Atom("1.23"	, types.NUMBER) },
	} 

	test(specs, t)
}

func TestParseCons(t *testing.T) {

	specs := []spec {
		{   "(1)"	, generator.Cons(generator.Atom( 1 , types.NUMBER), generator.Nil()) },
		{  "(1 2)"	, generator.Cons(generator.Atom( 1 , types.NUMBER), generator.Cons(generator.Atom(2, types.NUMBER), generator.Nil())) },
		{ "(a . b)"	, generator.Cons(generator.Atom("a", types.SYMBOL), generator.Atom("b", types.SYMBOL)) },
	} 

	test(specs, t)
}

func TestParseBackQuote(t *testing.T) {
	specs := []spec {
		{ "`a" 	, generator.Cons( generator.Atom("backquote", types.SYMBOL) , generator.Cons(generator.Atom("a", types.SYMBOL), generator.Nil())) },
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
		
		fmt.Printf("test - got :%v", got)
		
		if got.NotEqual(spec.want) {
			t.Errorf("Spec of expression \"%v\" was incorrect, got: \"%v\", want: \"%v\"", spec.expression, got, spec.want)
		}
	}		
}