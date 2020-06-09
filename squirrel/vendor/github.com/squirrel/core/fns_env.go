package core

import (
	"github.com/squirrel/types"
	"github.com/squirrel/generator"
)

// -------------------------------------------------------------------------------------------------
// Push is need to add new cells on after another to the end of the list
// To add fast there is a pointer to the last element returned	
//
//  e.g. 
//			(1 2)
//	
//  	    last
//  	     v
//		xs->NIL
//	
//  	    last
//  	     v
//		xs->   ->NIL
//	
//  	    last
//			 v
//  	xs->[1]->NIL
//	
//  	    last
//			 v
//  	xs->[1]->	->NIL
//	
//				 last
//  	          v	
//		xs->[1]->[2]->NIL
//	
func Push(xs, x, last *types.Cell) (*types.Cell, *types.Cell) {

	new_ := generator.Cons(x, NIL)
	
	if last.Equal(NIL) {
		xs = new_
	} else {
		last.Cdr = new_
	}
	
	last = new_

	return xs, last

}

//      +---+---+   +---+---+
//  a-->| B |   |-->| A |   |-->nil
//      +---+---+   +---+---+
//
//               +---+---+
//            +->| x |   |--+
//			  |	 +---+---+  |
//            |             v
//      +---+---+   	  +---+---+
//  a-->| B |   |	      | A |   |-->nil
//      +---+---+   	  +---+---+
//
//      +---+---+   +---+---+	+---+---+
//  a-->| B |   |-->| x	|	|-->| A |   |-->nil
//      +---+---+  	+---+---+	+---+---+
//
//      +---+---+   +---+---+	+---+---+
//  a-->| x |   |-->| B	|	|-->| A |   |-->nil
//      +---+---+  	+---+---+	+---+---+

// Add is a special add that adds a new cell at the front of the environment
// but LET the Pointer to first element the SAME !!!
func Add(x *types.Cell, a *types.Cell ) *types.Cell {
	
	new := addAtSecond(x, a)
		
	change(new, a)

	// So the pointer to a are the same
	return a 
}

func addAtSecond(x *types.Cell, a *types.Cell) *types.Cell {
	// Hang in the a new cell (with points with car to x) as second
	cdr := a.Cdr; new := Cons(x, cdr); a.Cdr = new
	return new
}

func change(x, y *types.Cell) {
	// Change Val first and second to move new cell to front
	val := x.Val; x.Val = y.Val; y.Val = val
	// Change Car first and second to move new cell to front
	car := x.Car; x.Car = y.Car; y.Car = car
}