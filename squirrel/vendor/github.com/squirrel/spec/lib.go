package spec

import (
	"fmt"
	"testing"
)

import (
	"github.com/squirrel/types"
	"github.com/squirrel/generator"
	"github.com/squirrel/evaluator"
	"github.com/squirrel/parser"	
)

type spec struct {
	expression  string
	want 		string
}

func test(specs []spec, t *testing.T) {
	testWithEnv(specs, t, nil)
}

func testWithEnv(specs []spec, t *testing.T, env *types.Cell) {

	if env == nil {
		env = generator.Nil()
	}
	
	for i, spec := range specs {
		
		name := fmt.Sprintf("test%v", i);
		t.Run(name, func(t *testing.T) {
		
			bs := []byte(spec.expression); e := parser.Parse(bs)
			res := evaluator.Eval(e, env); got := fmt.Sprintf("%v", res)
		
			if got != spec.want {
				t.Errorf("Spec eval %v was incorrect, got: %v, want: %v", spec.expression, got, spec.want)
			}
		})
	}		
}