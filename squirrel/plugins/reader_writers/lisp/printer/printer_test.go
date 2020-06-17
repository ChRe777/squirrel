package printer

import (
	"fmt"
	"testing"
)

import (
	"github.com/squirrel/types"
	"github.com/squirrel/core"
)


func TestPrintCell(t *testing.T) {
		
	atomA := core.Sym("a")
		
	listA := core.Cons(atomA, core.Nil_())
		
	llist := core.Cons(listA, listA)	
	
	dotAB := core.Cons(core.Sym("a"), core.Sym("b"))
	
	dotABC := core.Cons(core.Sym("a"), core.Cons(core.Sym("b"), core.Sym("c")))
		
	specs := []struct {
		e 		*types.Cell
		want 	string
	}{
		{ atomA	, "a"    		},
		{ listA	, "(a)"   		},
		{ llist	, "((a) a)" 	},
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

/*
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
*/