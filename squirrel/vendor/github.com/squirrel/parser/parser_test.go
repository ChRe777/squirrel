package parser

import (
	"testing"
	"github.com/squirrel/generator"
	"github.com/squirrel/types"
)

func TestParseAtom(t *testing.T) {

	specs := []spec {
		{ "t"		, generator.Atom("t"	, types.SYM   ) },
		{ "foo"		, generator.Atom("foo"	, types.SYMBOL) },
		{ "\"a\""	, generator.Atom("a"	, types.STRING) },
		{ "1.23"	, generator.Atom("1.23"	, types.NUMBER) },
	} 

	test(specs, t)
}

func TestParseCons(t *testing.T) {

	specs := []spec {1
		{ "(1)"			, generator.Add(generator.List(), generator.Atom(1, types.NUMBER)) },
		{ "(1 2)"		, generator.Add(generator.Add(generator.List(), generator.Atom(1, types.NUMBER)), generator.Atom(2, types.NUMBER)) },
		{ "(1 (foo))"	, generator.Add(generator.Add(generator.List(), generator.Atom(1, types.NUMBER)), generator.Add(generator.List(), generator.Atom("foo", types.SYMBOL))) },
	} 

	test(specs, t)
}