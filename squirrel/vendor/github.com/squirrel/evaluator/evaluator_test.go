package evaluator

import (
	"testing"
)

import (
	"github.com/squirrel/types"
	"github.com/squirrel/parser"
	"github.com/squirrel/builtin"
)

func TestEvalAtom(t *testing.T) {

	s := "((t t) (nil nil))"
	envBuiltin := parser.Parse([]byte(s))

	specs := []struct {
		e 	 *types.Cell
		want *types.Cell
	}{
		{ builtin.Sym("t")		, builtin.Sym("t") 	 },
		{ builtin.Sym("nil")	, builtin.Sym("nil") },
		{ builtin.Num("1")		, builtin.Num("1")   },
		{ builtin.Str("a")		, builtin.Str("a")   },
		{ builtin.Sym("b")		, builtin.Err("reference to undefined identifier: b") },
	}

	for _, spec := range specs {
		got := eval(spec.e, envBuiltin)
	
		if got.NotEqual(spec.want) {
			t.Errorf("eval atom e: %v - got: %v, want: %v", spec.e, got, spec.want)
		}
	}
}

func TestEvalFunc(t *testing.T) {

	p := func(s string) *types.Cell {
		return parser.Parse([]byte(s))
	}

	s := "((t t) (nil nil))"
	envBuiltin := p(s)
	
	specs := []struct {
		e 	 *types.Cell
		want *types.Cell
	}{
		{ p("((func (x)(car x)) '(a b c))")	, builtin.Sym("a") },
	}
	
	for _, spec := range specs {
		got := eval(spec.e, envBuiltin)
	
		if got.NotEqual(spec.want) {
			t.Errorf("eval func e: %v - got: %v, want: %v", spec.e, got, spec.want)
		}
	}
		
}
