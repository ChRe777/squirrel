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
		
	llist := &Cell{
		Type: Type{Cell: CONS},
		Car : listA,
		Cdr : listA,
	}

	specs := []struct {
		e 		*Cell
		want 	string
	}{
		{ atomA	, "a"    		},
		{ listA	, "(a)"   		},
		{ llist	, "((a) a)" 	},
	}
	
	for _, spec := range specs {
	
		got := SprintCell(spec.e)
		
		fmt.Printf("%v\n", got)
		
		if got != spec.want {
			t.Errorf("print cell got:%v, want:%v", got, spec.want)
		}
	}

}

func TestPrintCell2(t *testing.T) {
	
	atomA := &Cell{
		Type: Type{Cell: ATOM, Atom: SYMBOL},
		Val : "a",
	}
	
	atomB := &Cell{
		Type: Type{Cell: ATOM, Atom: SYMBOL},
		Val : "b",
	}
	
	atomC := &Cell{
		Type: Type{Cell: ATOM, Atom: SYMBOL},
		Val : "c",
	}
	
	dotBC := &Cell{
		Type: Type{Cell: CONS},
		Car : atomB,
		Cdr : atomC,
	}

	dotAB := &Cell{
		Type: Type{Cell: CONS},
		Car : atomA,
		Cdr : atomB,
	}
		
	dotABC := &Cell{
		Type: Type{Cell: CONS},
		Car : atomA,
		Cdr : dotBC,
	}

	specs := []struct {
		e 		*Cell
		want 	string
	}{
		{ dotAB	 , "(a . b)" 	},
		{ dotABC , "(a b . c)" 	},
	}
	
	for _, spec := range specs {
	
		got := SprintCell(spec.e)
		
		fmt.Printf("%v\n", got)
		
		if got != spec.want {
			t.Errorf("print cell got:%v, want:%v", got, spec.want)
		}
	}

}