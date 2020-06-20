package builtin


var (
    mapFn    = "(map	(func (f x)  (cond ((no x) nil) ('t (cons (f (car x)) (map f (cdr x)))))))"

)

// (def map (f x) (cond ((no x) nil) ('t (cons (f (car x)) (map f (cdr x))))))



import (
	"github.com/mysheep/squirrel/types"
	"github.com/mysheep/squirrel/evaluator"
	"github.com/mysheep/squirrel/evaluator/core"
	"github.com/mysheep/squirrel/generator"
)
/*

	Funcs:

		Pair
		No
		Not
		And
		Append
		List
		Assoc
		
		Caar  
		Cadr  
		Cddr  
		Cadar 
		Cdddr 
		Caddr 
		Caddar
		Cadddr

*/

func Pair (xs, ys *types.Cell) *types.Cell {
	
	if xs.Equal(core.NIL) || 
	   ys.Equal(core.NIL) {
		return core.NIL
	} 

	if xs.IsCons() && ys.IsCons() {		// (x y z) (1 2 3)
	
		x := car(xs)
		y := car(ys)
	
		ws := cdr(xs)
		zs := cdr(ys)
		
		a := List_(x, y)
		b := Pair(ws, zs)
		
		return core.Cons(a, b)
	
	} else {							// (x y . z) (1 2 3 4)
		return core.Cons(List_(xs, ys), core.NIL)
	}
	
}

func List_ (x, y *types.Cell) *types.Cell {
	return core.Cons(x, core.Cons (y, core.NIL))
}

func No (x *types.Cell) *types.Cell { // call "no" instead of "null"
	if x.Equal(core.NIL) {
		return core.T
	}
	return core.NIL
}

func Not (x *types.Cell) *types.Cell {
	if x.Equal(core.T) {
		return core.NIL
	} else {
		return core.T
	}
}

func And (x, y *types.Cell) *types.Cell {
	if x.Equal(core.T) && y.Equal(core.T) {
		return core.T
	} else {
		return core.NIL
	}
}

func Append (x, y *types.Cell) *types.Cell {
	if x.Equal(core.NIL) {
		return y
	} else {
		return core.Cons(car(x), Append(cdr(x), y))
	}
}

func List (xs, a *types.Cell) *types.Cell {

	if xs.Equal(core.NIL) {
		return core.NIL
	} else {
		y  := evaluator.Eval(car(xs), a)
		ys := cdr(xs)
		return core.Cons(y, List(ys, a))
	}
}

func Assoc (x, ys *types.Cell) *types.Cell {
	if ys.Equal(core.NIL) {
		return core.Err_("Not found")
	} else {
		if x.Equal(Caar(ys)) {
			return Cadar(ys)
		} else {
			return Assoc(x, cdr(ys))	
		}
	}
}


//    mapFn    = "(map	(func (f x)  (cond ((no x) nil) ('t (cons (f (car x)) (map f (cdr x)))))))"



func Caar  (e *types.Cell) *types.Cell { return car(car(e))           }
func Cadr  (e *types.Cell) *types.Cell { return car(cdr(e))           }
func Cddr  (e *types.Cell) *types.Cell { return cdr(cdr(e))           }
func Cadar (e *types.Cell) *types.Cell { return car(cdr(car(e)))      } 
func Cdddr (e *types.Cell) *types.Cell { return cdr(cdr(cdr(e)))      } 
func Caddr (e *types.Cell) *types.Cell { return car(cdr(cdr(e)))      }
func Caddar(e *types.Cell) *types.Cell { return car(cdr(cdr(car(e)))) }
func Cadddr(e *types.Cell) *types.Cell { return car(cdr(cdr(cdr(e)))) } 	


// -------------------------------------------------------------------------------------------------
// Just ALIAS for better readability
// -------------------------------------------------------------------------------------------------

func car (x *types.Cell) *types.Cell {
	return core.Car(x)
}

func cdr (x *types.Cell) *types.Cell {
	return core.Cdr(x)
}

func Sym_(s string) *types.Cell {
	return generator.Sym(s)
}

func list_ (x, y *types.Cell) *types.Cell {
	return core.Cons(x, core.Cons(y, core.NIL))
}






