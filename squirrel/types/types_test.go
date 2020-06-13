package types

import (
	"fmt"
	"testing"
)

func TestTypes(t *testing.T) {
	
	atomStr := &Cell{
		Type: Type{Cell: ATOM, Atom: STRING},
		Val : "A",
	}
	
	atomNum := &Cell{
		Type: Type{Cell: ATOM, Atom: NUMBER},
		Val : 1.23e-10,
	}
	
	fmt.Printf("atomA: %v \n", atomStr)
	fmt.Printf("atomB: %v \n", atomNum)
}

func TestEqualAtom(t *testing.T) {

	x := &Cell{
		Type: Type{Cell: ATOM, Atom: STRING},
		Val : "a",
	}
	
	y := &Cell{
		Type: Type{Cell: ATOM, Atom: STRING},
		Val : "a",
	}
	
	got  := x.Equal(y)
	want := true
	
	if got != want {
		t.Errorf("%v should be equal %v", x, y)
	}
	
}

func TestEqualCons(t *testing.T) {

	a := &Cell{
		Type: Type{Cell: ATOM, Atom: STRING},
		Val : "a",
	}
	
	b := &Cell{
		Type: Type{Cell: ATOM, Atom: STRING},
		Val : "b",
	}

	x := &Cell{
		Type: Type{Cell: CONS},
		Val : "a",
		Car : a,
		Cdr : b,
	}
	
	y := &Cell{
		Type: Type{Cell: CONS},
		Val : "b",
		Car : a,
		Cdr : b,
	}
	
	got  := x.Equal(y)
	want := true
	
	if got != want {
		t.Errorf("%v should be equal %v", x, y)
	}
	
}

func TestString(t *testing.T) {
	
	//  e.g. (1 a)
	// 
	//		CONS	CONS    ATOM
	//  --->[o|o]-->[o|o]-->[nil]
	//   	 |       |
	//   	 v       v
	//  	 1       a
	//  	ATOM    ATOM
	
	eol := &Cell{
		Type: Type{Cell: ATOM, Atom: SYMBOL},
		Val : nil,
		Car : nil,
		Cdr : nil,
	}
	
	one := &Cell{
		Type: Type{Cell: ATOM, Atom: NUMBER},
		Val : 1,
		Car : nil,
		Cdr : nil,
	}
	
	aaa := &Cell{
		Type: Type{Cell: ATOM, Atom: SYMBOL},
		Val : "a",
		Car : nil,
		Cdr : nil,
	}
		
	cons := &Cell{
		Type: Type{Cell: CONS},
		Val : nil,
		Car : aaa,
		Cdr : eol,
	}
	
	list := &Cell{
		Type: Type{Cell: CONS},
		Val : nil,
		Car : one,
		Cdr : cons,
	}
	
	got  := fmt.Sprintf("%v", list)
	want := "(1 a)"
	
	if got != want {
		t.Errorf("String (1 a) - got: %v, want: %v", got, want)
	}
	
}

func TestIsSymbol(t *testing.T) {

	e := &Cell{
		Type: Type{Cell: ATOM, Atom: SYMBOL},
		Val : "a",
		Car : nil,
		Cdr : nil,
	}
	
	got := e.IsSymbol()
	want := true
	
	if got != want {
		t.Errorf("IsSymbol failed - got: %v, want: %v", got, want)
	}
}