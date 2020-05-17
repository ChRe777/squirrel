package evaluator

import (
	"testing"
	"github.com/squirrel/types"
	"github.com/squirrel/builtin"
	"github.com/squirrel/generator"
)

/*
func quote(x *types.Cell) *types.Cell {
	return cadr(x)
}
*/
func TestQuote(t *testing.T) {
	
	e := builtin.Quote(builtin.Sym("a"))
		
	got  := quote(e)
	want := builtin.Sym("a")	
	
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
		{ builtin.Sym("a")	, builtin.T     },
		{ builtin.NIL 		, builtin.T 	},
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
		{builtin.Sym("a"), builtin.Sym("a"), builtin.T  },
		{builtin.Sym("a"), builtin.Sym("b"), builtin.NIL},
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
		{builtin.Sym("a")	, generator.Error("Can't take car of a")},
		// (car nil) -> nil
		{builtin.NIL   		, builtin.NIL},
		// (car (a b)) -> a
		{builtin.List(builtin.Sym("a"), builtin.Sym("b")), builtin.Sym("a")},
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
		{builtin.Sym("a")	, generator.Error("Can't take cdr of a")},
		
		// (car nil) -> nil
		{builtin.NIL   		, builtin.NIL},
		
		// (car (a b)) -> a
		{builtin.List(builtin.Sym("a"), builtin.Sym("b")), builtin.List(builtin.Sym("b"))},
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
		{builtin.Sym("a"), builtin.NIL, 	 builtin.List(builtin.Sym("a"))},
		{builtin.Sym("a"), builtin.Sym("b"), generator.Cons(builtin.Sym("a"), builtin.Sym("b"))},
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
	e := builtin.List(
		builtin.List(builtin.Sym("nil"), builtin.Sym("a")),
		builtin.List(builtin.Sym("t")  , builtin.Sym("b")),
	)
	want := builtin.Sym("b")

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

