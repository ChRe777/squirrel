package evaluator

import (
	"testing"
)

import (
	"github.com/squirrel/types"
	"github.com/squirrel/builtin"
)

func TestEvalAtom(t *testing.T) {

	specs := []struct {
		e 	 *types.Cell
		want *types.Cell
	}{
		{ builtin.Sym("t"),  builtin.Sym("t") },
		// todo ...
	}

	for _, spec := range specs {
		got := eval(spec.e, builtin.NIL)
	
		if got.NotEqual(spec.want) {
			t.Errorf("eval atom e: %v - got: %v, want: %v", spec.e, got, spec.want)
		}
	}
}
