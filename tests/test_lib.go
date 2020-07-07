package spec

import (
	"testing"
)

import (
	"github.com/mysheep/squirrel/types"
	"github.com/mysheep/squirrel/evaluator"
	"github.com/mysheep/squirrel/builtin"
	"github.com/mysheep/squirrel/plugins/readerwriters/lisp/parser"	
	"github.com/mysheep/squirrel/plugins/readerwriters/lisp/printer"	
)

type spec struct {
	expression  string
	want 		string
}

type spec2 struct {
	expr1  	string
	expr2  	string
	want 	string
}

func test(specs []spec, t *testing.T) {
	testWithEnv(specs, t, nil)
}

func test2(specs []spec2, t *testing.T) {

	p := func(s string) *types.Cell {
		return parser.Parse([]byte(s))
	}
	
	s := "((t t) (nil nil))"
	env := p(s)
		
	for _, spec := range specs {

		evaluator.Eval(p(spec.expr1), env); 
		gotExp := evaluator.Eval(p(spec.expr2), env); got := string(printer.Sprint(gotExp))
	
		if got != spec.want {
			t.Errorf("Spec eval %v was incorrect, got: %v, want: %v", spec.expr2, got, spec.want)
		}

	}		
}

func testWithEnv(specs []spec, t *testing.T, env *types.Cell) {

	p := func(s string) *types.Cell {
		return parser.Parse([]byte(s))
	}
	
	s := "((t t) (nil nil))"
	builtInEnv := p(s)
		
	if env == nil {
		env = builtInEnv
	} else {
		env = builtin.Append(builtInEnv, env)
	}
		
	for _, spec := range specs {
		
		bs := []byte(spec.expression); e := parser.Parse(bs)
		res := evaluator.Eval(e, env); got := string(printer.Sprint(res))
					
		if got != spec.want {
			t.Errorf("Spec eval %v was incorrect, got: %v, want: %v", spec.expression, got, spec.want)
		}
				
	}		
}