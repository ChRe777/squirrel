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

	s := "((t t))"
	envBuiltin := parser.Parse([]byte(s))

	specs := []struct {
		e 	 *types.Cell
		want *types.Cell
	}{
		{ builtin.Sym("t")	,  builtin.Sym("t") },
		{ builtin.Num("1")	,  builtin.Num("1") },
//		{ builtin.Str("a")	,  builtin.Str("a") },
	}

	for _, spec := range specs {
		got := eval(spec.e, envBuiltin)
	
		if got.NotEqual(spec.want) {
			t.Errorf("eval atom e: %v - got: %v, want: %v", spec.e, got, spec.want)
		}
	}
}
