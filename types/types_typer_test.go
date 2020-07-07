package types

import (
	"testing"
)

func TestType_(t *testing.T) {

	spec := []struct {
		exp 	*Cell
		want 	string
	}{
		{
			&Cell{
				Type: Type{Cell: ATOM, Atom: SYMBOL},
				Val : "func",
				Tag : "func",
				Car : nil,
				Cdr : nil,
			},
			"func#func"
		
		}
	}
	
	if got != spec.want {
		t.Errorf("Type - got: %v, want: %v", got, spec.want)
	}
	
}