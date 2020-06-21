package evaluator

import (
	"testing"
)

import (
	"github.com/mysheep/squirrel/types"
	"github.com/mysheep/squirrel/evaluator/core"
	"github.com/mysheep/squirrel/plugins/readerwriters/lisp/parser"

)

func TestEvalAtom(t *testing.T) {

	s := "((t t) (nil nil))"
	envBuiltin := parser.Parse([]byte(s))

	specs := []struct {
		e 	 *types.Cell
		want *types.Cell
	}{
		{ core.Sym_("t"  ), core.Sym_("t"  ) },
		{ core.Sym_("nil"), core.Sym_("nil") },
		{ core.Num_("1"  ), core.Num_("1"  ) },
		{ core.Str_("a"  ), core.Str_("a"  ) },
		{ core.Sym_("b"  ), core.Err_("reference to undefined identifier: b") },
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
		{ p("((func (x)(car x)) '(a b c))")	, core.Sym_("a") },
	}
	
	for _, spec := range specs {
		got := eval(spec.e, envBuiltin)
	
		if got.NotEqual(spec.want) {
			t.Errorf("eval func e: %v - got: %v, want: %v", spec.e, got, spec.want)
		}
	}
		
}

func TestEvalFuncEnv(t *testing.T) {

	p := func(s string) *types.Cell {
		return parser.Parse([]byte(s))
	}
	
	s := "((t t) (nil nil) (bar (func (x) x)))"
	envBuiltin := p(s)
	
	specs := []struct {
		e 	 *types.Cell
		want *types.Cell
	}{
		{ p("(foo 'a)")	, core.Err_("reference to undefined identifier: foo") },
		{ p("(bar 'a)")	, core.Sym_("a") },
	}
	
	for _, spec := range specs {
		
		got := evalFuncEnv(spec.e, envBuiltin)
	
		if got.NotEqual(spec.want) {
			t.Errorf("evalFuncEnv e: %v - got: %v, want: %v", spec.e, got, spec.want)
		}
	}
}

func TestEvalFuncOrMacCall(t *testing.T) {
	// TODO
}

func TestFuncCall(t *testing.T) {
	// TODO
}

func TestMacCall(t *testing.T) {
	// TODO
}
