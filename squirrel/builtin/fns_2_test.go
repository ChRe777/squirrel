package builtin

import (
//	"testing"
//	"fmt"
)

import (
//	"github.com/mysheep/squirrel/types"
//	"github.com/mysheep/squirrel/core"
//	"github.com/mysheep/squirrel/plugins/readerwriters/lisp/parser"

)

/*
	- var
	- def
	- mac
	- let
	
	... TODO
	
*/

/*
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
	
		evaluator.Eval(spec.e, envBuiltin)
		
		fmt.Printf("e: %v, envBuiltin: %v  \n", spec.e, envBuiltin)
		
		got := core.Car(envBuiltin) 
		
		if got.NotEqual(spec.want) {
			t.Errorf("eval var e: %v - got: %v, want: %v", spec.e, got, spec.want)
		}
	}
}
/*
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
		{ p("(def foo(x) (no x))")	, p("(foo '(1 2))")	, core.NIL  	  },
		{ p("(def bar(x) (no x))")	, p("(bar '())"   )	, core.Sym_("t")  },
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
		{ p("(let ys '(1 2) (car ys))"), core.Num_("1") },
		{ p("(let fn (func (x) (list x x)) (fn 1))"), p("(1 1)") },
	}
	
	for _, spec := range specs {
		got := eval(spec.e, envBuiltin)
		
		if got.NotEqual(spec.want) {
			t.Errorf("eval let e: %v - got: %v, want: %v", spec.e, got, spec.want)
		}
	}
	
}
*/