package evaluator

import (
	"testing"
	"fmt"
)

import (
	"github.com/squirrel/types"
	"github.com/squirrel/parser"
	"github.com/squirrel/core"
)

func TestEvalAtom(t *testing.T) {

	s := "((t t) (nil nil))"
	envBuiltin := parser.Parse([]byte(s))

	specs := []struct {
		e 	 *types.Cell
		want *types.Cell
	}{
		{ core.Sym("t"  ), core.Sym("t"  ) },
		{ core.Sym("nil"), core.Sym("nil") },
		{ core.Num("1"  ), core.Num("1"  ) },
		{ core.Str("a"  ), core.Str("a"  ) },
		{ core.Sym("b"  ), core.Err("reference to undefined identifier: b") },
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
		{ p("((func (x)(car x)) '(a b c))")	, core.Sym("a") },
	}
	
	for _, spec := range specs {
		got := eval(spec.e, envBuiltin)
	
		if got.NotEqual(spec.want) {
			t.Errorf("eval func e: %v - got: %v, want: %v", spec.e, got, spec.want)
		}
	}
		
}

func TestEvalVar(t *testing.T) {

	p := func(s string) *types.Cell {
		return parser.Parse([]byte(s))
	}
	
	s := "((t t) (nil nil))"
	envBuiltin := p(s)

	specs := []struct {
		e		*types.Cell
		want 	*types.Cell
	} {
		{ p("(var a 1)		"), p("(a 1)"    )  },
		{ p("(var b '(1 2))	"), p("(b (1 2))")  },
	}
	
	for _, spec := range specs {
	
		eval(spec.e, envBuiltin)
		
		fmt.Printf("e: %v, envBuiltin: %v  \n", spec.e, envBuiltin)
		
		got := core.Car(envBuiltin) 
		
		if got.NotEqual(spec.want) {
			t.Errorf("eval var e: %v - got: %v, want: %v", spec.e, got, spec.want)
		}
	}
}

func TestEvalDef(t *testing.T) {

	p := func(s string) *types.Cell {
		return parser.Parse([]byte(s))
	}
	
	s := "((t t) (nil nil))"
	envBuiltin := p(s)

	specs := []struct {
		e1		*types.Cell
		e2		*types.Cell
		want 	*types.Cell
	} {
		{ p("(def foo(x) (no x))")	, p("(foo '(1 2))")	, core.NIL  		},
		{ p("(def bar(x) (no x))")	, p("(bar '())"   )	, core.Sym("t")  },
	}
	
	for _, spec := range specs {
			   eval(spec.e1, envBuiltin)
		got := eval(spec.e2, envBuiltin)
		
		if got.NotEqual(spec.want) {
			t.Errorf("eval def e: %v - got: %v, want: %v", spec.e2, got, spec.want)
		}
	}
	
}

func TestEvalMac(t *testing.T) {

	p := func(s string) *types.Cell {
		return parser.Parse([]byte(s))
	}
	
	s := "((t t) (nil nil))"
	envBuiltin := p(s)

	specs := []struct {
		e1		*types.Cell
		e2		*types.Cell
		want 	*types.Cell
	} {
		{ p("(mac foo(x)   `(no ,x))"     )	, p("(foo '(1 2))")	, core.NIL  	},
		{ p("(mac bar(x y) `(list ,x ,y))")	, p("(bar 1 2)")	, p("(1 2)")  	},
	}
	
	for _, spec := range specs {
			   eval(spec.e1, envBuiltin)
		got := eval(spec.e2, envBuiltin)
		
		if got.NotEqual(spec.want) {
			t.Errorf("eval mac e: %v - got: %v, want: %v", spec.e2, got, spec.want)
		}
	}
	
}

func TestEvalLet(t *testing.T) {

	p := func(s string) *types.Cell {
		return parser.Parse([]byte(s))
	}
	
	s := "((t t) (nil nil))"
	envBuiltin := p(s)

	specs := []struct {
		e		*types.Cell
		want 	*types.Cell
	} {
		{ p("(let xs '(1 2) (no  xs))"), core.NIL  	   },
		{ p("(let ys '(1 2) (car ys))"), core.Num("1") },
		{ p("(let fn (func (x) (list x x)) (fn 1))"), p("(1 1)") },
	}
	
	for _, spec := range specs {
		got := eval(spec.e, envBuiltin)
		
		if got.NotEqual(spec.want) {
			t.Errorf("eval let e: %v - got: %v, want: %v", spec.e, got, spec.want)
		}
	}
	
}