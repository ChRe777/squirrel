package types

import (
	"fmt"
	"testing"
)

func TestPrintCell(t *testing.T) {

	symNil := &Cell{
		Type: Type{Cell: ATOM, Atom: SYMBOL},
		Val : "nil",
	}
	
	atomA := &Cell{
		Type: Type{Cell: ATOM, Atom: SYMBOL},
		Val : "a",
	}
	
	listA := &Cell{
		Type: Type{Cell: CONS},
		Car : atomA,
		Cdr : symNil,
	}

	dotP := &Cell{
		Type: Type{Cell: CONS},
		Car : atomA,
		Cdr : atomA,
	}
	
	llist := &Cell{
		Type: Type{Cell: CONS},
		Car : listA,
		Cdr : listA,
	}

	specs := []struct {
		e 		*Cell
		want 	string
	}{
		{ atomA,    "a"    },
		{ listA,   "(a)"   },
		{  dotP, "(a . a)" },
		{ llist, "((a) a)" },
	}
	
	for _, spec := range specs {
	
		got := SprintCell(spec.e)
		
		fmt.Printf("%v\n", got)
		
		if got != spec.want {
			t.Errorf("print cell got:%v, want:%v", got, spec.want)
		}
	}

}