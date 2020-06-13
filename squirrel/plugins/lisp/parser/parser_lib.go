package parser

import(
	"fmt"
	"strings"
)

import (
	"github.com/mysheep/squirrel/types"
	"github.com/mysheep/squirrel/generator"
)

const (
	DEBUG = false
	IDENT = 2
)

var (
	level = 0
)

/*
PROCEDURE debug(msg: ARRAY OF CHAR; level: INTEGER);
	VAR sps: ARRAY 64 OF CHAR; i: INTEGER;
BEGIN i := 0;
	IF (DEBUG) THEN
		WHILE level > 0 DO sps [i] := " "; INC(i); DEC(level); END;
		sps[i] := 0X; (* END *);
		Out.String(sps); Out.String(msg); Out.Ln;
	END;
END debug;
*/
func debug(msg string, level *int) {
	if DEBUG {
		spaces := strings.Repeat(" ", *level)
		fmt.Printf("%s%s\n", spaces, msg)
	}
}

//
//PROCEDURE incL(VAR level: INTEGER); BEGIN level := level + 4; END incL;
//
func incLevel(level *int) {
	*level += IDENT
}


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

	new_ := generator.Cons(x, generator.NIL)
	
	if last.Equal(generator.NIL) {
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
	cdr := a.Cdr; new := generator.Cons(x, cdr); a.Cdr = new
	return new
}

func change(x, y *types.Cell) {
	// Change Val first and second to move new cell to front
	val := x.Val; x.Val = y.Val; y.Val = val
	// Change Car first and second to move new cell to front
	car := x.Car; x.Car = y.Car; y.Car = car
}

