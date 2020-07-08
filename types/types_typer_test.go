package types

import (
	"testing"
)

func TestType_(t *testing.T) {

	fn := &Cell{Type: Type{Cell: ATOM, Atom: FUNC},
				Val : "func",
				Tag : "func",
				Car : nil,
				Cdr : nil}
				
	specs := []struct {
		exp 	*Cell
		want 	string
	}{
		{ fn, "func#func"},
	}
	
	for _, spec := range specs {
	
		got := spec.exp.Type_()
	
		if got != spec.want {
			t.Errorf("Type - got: %v, want: %v", got, spec.want)
		}
	
	}
	
}