package types

import (
	"testing"
)

func TestEqual(t *testing.T) {

	specs := []struct {
		x	*Cell
		y   *Cell
		want bool
	} {
		//{  },
	}

	for _, spec := range specs {
		
		got := spec.x.Equal(spec.y)
		
		if got != spec.want {
			t.Errorf("TestEqual - got: %v, want: %v", got, spec.want)
		}
	}

}