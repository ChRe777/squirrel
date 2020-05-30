package evaluator

import (
	"testing"
	"github.com/squirrel/types"
	"github.com/squirrel/core"
	"github.com/squirrel/generator"
)

/*
func quote(x *types.Cell) *types.Cell {
	return cadr(x)
}
*/
func TestQuote(t *testing.T) {
	
	e := core.Quote(core.Sym("a"))
		
	got  := quote(e)
	want := core.Sym("a")	
	
	if got.NotEqual(want) {
		t.Errorf("Quote failed, got: %v, want: %v", got, want)
	}
}

/*
func atom(x *types.Cell) *types.Cell {
	if isAtom(x) {
		return T
	} else {
		return NIL
	}
}
*/
func TestAtom(t *testing.T) {
	
	specs := []struct {
		expr *types.Cell
		want *types.Cell
	} {
		{ core.Sym("a")	, core.T },
		{ core.NIL 		, core.T },
	}
	
	for _, spec := range specs {
	
		got := atom(spec.expr)
		
		if got != spec.want {
			t.Errorf("Failed atom(a), got:%v, want:%v", got, spec.want)
		}
	}
		
}

/*
func eq(x, y *types.Cell) *types.Cell {	
	if x.Equal(y) {
	 	return T	
	}
	return NIL 	// FALSE
}
*/
func TestEq(t *testing.T) {

	specs := []struct{
		x		*types.Cell
		y		*types.Cell
		want 	*types.Cell
	}{
		{core.Sym("a"), core.Sym("a"), core.T  },
		{core.Sym("a"), core.Sym("b"), core.NIL},
		// ... TODO MANY TESTS ...
	}

	for _, spec := range specs {
		
		got := eq(spec.x, spec.y)
		
		if got != spec.want {
			t.Errorf("%v eq %v failed, got: %v, want: %v", spec.x, spec.y, got, spec.want)
		}
	}
	
}

/*
func car(e *types.Cell) *types.Cell {
	if e == NIL {
		return NIL
	} else {
		if isCons(e) {
			return generator.Car(e) 
		} else {
			return error(fmt.Sprintf("can not take car of %v", e))
		}
	}
}
*/
func TestCar(t *testing.T) {
	specs := []struct{
		x		*types.Cell
		want 	*types.Cell
	}{
		{core.Sym("a")	, generator.Error("Can't take car of a")},
		// (car nil) -> nil
		{core.NIL   		, core.NIL},
		// (car (a b)) -> a
		{core.List(core.Sym("a"), core.Sym("b")), core.Sym("a")},
	}
	
	for _, spec := range specs {
		
		got := car(spec.x)
		
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
		{core.Sym("a")	, generator.Error("Can't take cdr of a")},
		
		// (car nil) -> nil
		{core.NIL   		, core.NIL},
		
		// (car (a b)) -> a
		{core.List(core.Sym("a"), core.Sym("b")), core.List(core.Sym("b"))},
	}
	
	for _, spec := range specs {
		
		got := cdr(spec.x)
		
		if got.NotEqual(spec.want) {
			t.Errorf("cdr %v failed, got: %v, want: %v", spec.x, got, spec.want)
		}
	}
	
}

/*
func cons (x, y *types.Cell) *types.Cell {
	if y.IsCons() {
		return generator.Cons(x, y)
	} else {
		// TODO: dotted pair (cons 'a 'b) -> (a . b) *)
		return generator.Error(fmt.Sprintf("y must be a list"))
	}
}
*/

func TestCons(t *testing.T) {
	specs := []struct{
		x		*types.Cell
		y       *types.Cell
		want 	*types.Cell
	}{
		{core.Sym("a"), core.NIL, 	 core.List(core.Sym("a"))},
		{core.Sym("a"), core.Sym("b"), generator.Cons(core.Sym("a"), core.Sym("b"))},
	}
	
	for _, spec := range specs {
		
		got := cons(spec.x, spec.y)
		
		if got.NotEqual(spec.want) {
			t.Errorf("cons %v, %v failed, got: %v, want: %v", spec.x, spec.y, got, spec.want)
		}
	}
	
}

/*	
PROCEDURE cond(x: cell): cell;
BEGIN
	IF x IS consCell THEN
		IF eq(caar(x), T) = T THEN RETURN cadar(x);				
		ELSE RETURN cond(cdr(x)) END;
	ELSE
		error(1); (* TODO: x must be a list of from ((p1 e1) (p2 e2) .. (pn en)) *)		
	END;
END cond;
*/
/*
	(
		(p1 e1) 
		(p2 e2)  
		...
		(pn en)
	)
*/
func TestCond(t *testing.T) {

/*
	(
		(nil a)
		(t	 b)
	)

*/
	e := core.List(
		core.List(core.Sym("nil"), core.Sym("a")),
		core.List(core.Sym("t")  , core.Sym("b")),
	)
	want := core.Sym("b")

	specs := []struct{
		x		*types.Cell
		want 	*types.Cell
	}{
		{e,  want},
	}
	
	for _, spec := range specs {
		
		got := cond(spec.x)
		
		if got.NotEqual(spec.want) {
			t.Errorf("cond %v failed, got: %v, want: %v", spec.x, got, spec.want)
		}
	}
	
}

