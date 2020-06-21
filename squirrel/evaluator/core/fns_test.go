package core

import (
	"testing"
)

import(	
	"github.com/mysheep/squirrel/types"
	"github.com/mysheep/squirrel/generator"
)

/*
func quote(x *types.Cell) *types.Cell {
	return cadr(x)
}
*/
func TestQuote(t *testing.T) {
	
	e := Quote(Sym_("a"))
		
	got  := Quote(e)
	want := Sym_("a")	
	
	if got.NotEqual(want) {
		t.Errorf("Quote failed, got: %v, want: %v", got, want)
	}
}


func TestAtom(t *testing.T) {
	
	specs := []struct {
		expr *types.Cell
		want *types.Cell
	} {
		{ Sym_("a")	, T },
		{ NIL 		, T },
	}
	
	for _, spec := range specs {
	
		got := Atom(spec.expr)
		
		if got != spec.want {
			t.Errorf("Failed atom(a), got:%v, want:%v", got, spec.want)
		}
	}
		
}

func TestIs(t *testing.T) {

	specs := []struct{
		x		*types.Cell
		y		*types.Cell
		want 	*types.Cell
	}{
		{Sym_("a"), Sym_("a"), T  },
		{Sym_("a"), Sym_("b"), NIL},
		// ... TODO MANY TESTS ...
	}

	for _, spec := range specs {
		
		got := Is(spec.x, spec.y)
		
		if got != spec.want {
			t.Errorf("%v eq %v failed, got: %v, want: %v", spec.x, spec.y, got, spec.want)
		}
	}
	
}

func TestCar(t *testing.T) {
	specs := []struct{
		x		*types.Cell
		want 	*types.Cell
	}{
		{Sym_("a")	, Err_("Can't take car of a")},
		// (car nil) -> nil
		{NIL   		, NIL},
		// (car (a b)) -> a
		{list2_(Sym_("a"), Sym_("b")), Sym_("a")},
	}
	
	for _, spec := range specs {
		
		got := Car(spec.x)
		
		if got.NotEqual(spec.want) {
			t.Errorf("car %v failed, got: %v, want: %v", spec.x, got, spec.want)
		}
	}
}

func TestCdr(t *testing.T) {
	specs := []struct{
		x		*types.Cell
		want 	*types.Cell
	}{
		// (car 'a)  -> error
		{Sym_("a")	, Err_("Can't take cdr of a")},
		
		// (car nil) -> nil
		{NIL   		, NIL},
		
		// (car (a b)) -> a
		{list2_(Sym_("a"), Sym_("b")), list2_(Sym_("b"), NIL)},
	}
	
	for _, spec := range specs {
		
		got := Cdr(spec.x)
		
		if got.NotEqual(spec.want) {
			t.Errorf("cdr %v failed, got: %v, want: %v", spec.x, got, spec.want)
		}
	}
	
}

func TestCons(t *testing.T) {
	specs := []struct{
		x		*types.Cell
		y       *types.Cell
		want 	*types.Cell
	}{
		{Sym_("a"), NIL, 	 list2_(Sym_("a"), NIL)},
		{Sym_("a"), Sym_("b"), generator.Cons(Sym_("a"), Sym_("b"))},
	}
	
	for _, spec := range specs {
		
		got := Cons(spec.x, spec.y)
		
		if got.NotEqual(spec.want) {
			t.Errorf("cons %v, %v failed, got: %v, want: %v", spec.x, spec.y, got, spec.want)
		}
	}
	
}

func TestCond(t *testing.T) {

/*
	(
		(nil a)
		(t	 b)
	)

*/
	e := list2_(list2_(Sym_("nil"), Sym_("a")), list2_(Sym_("t")  , Sym_("b")))
		
	want := Sym_("b")

	specs := []struct{
		x		*types.Cell
		want 	*types.Cell
	}{
		{e,  want},
	}
	
	for _, spec := range specs {
		
		got := Cond(spec.x)
		
		if got.NotEqual(spec.want) {
			t.Errorf("cond %v failed, got: %v, want: %v", spec.x, got, spec.want)
		}
	}
	
}

