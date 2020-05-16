package parser

import(
	"testing"
	"github.com/squirrel/types"
)

type cell *types.Cell

type spec struct {
	expression string
	want	   cell
}

func test(specs []spec, t *testing.T) {
	
	for _, spec := range specs {
	
		e := []byte(spec.expression); got := Parse(e)
		
		if got.NotEqual(spec.want) {
			t.Errorf("Spec of expression \"%v\" was incorrect, got: \"%v\", want: \"%v\"", spec.expression, got, spec.want)
		}
	}		
}