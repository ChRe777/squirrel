package core

import (
	"github.com/squirrel/types"
	"github.com/squirrel/generator"
)

// -------------------------------------------------------------------------------------------------
// Push adds a new cell at the end of the list
// To add fast there is a pointer to the last element	
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

// Add is a special add that adds a new cell at the front of the environment
// but LET the Pointer to first element the SAME !!!
func Add(kv *types.Cell, a *types.Cell ) *types.Cell {
	// Hang in new as second
	cdr := a.Cdr; new := Cons(kv, cdr); a.Cdr = new
	// Change Val first and second to move new second to front
	val := new.Val; new.Val = a.Val; a.Val = val
	// Change Car first and second to move new seocen to front
	car := new.Car; new.Car = a.Car; a.Car = car
	// So the pointer to a stays the same
	return a
}