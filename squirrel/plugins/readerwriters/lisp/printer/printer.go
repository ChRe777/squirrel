package printer

import (
	"bytes"
	"fmt"
)

import (
	"github.com/mysheep/squirrel/types"
)

const (
	NIL    = "nil"
	SPACE  = ' '
	DOT    = '.'
	LPAREN = '('
	RPAREN = ')'
)

func Sprint(c *types.Cell) []byte {

	var buffer bytes.Buffer

	if c == nil {
		buffer.WriteString("")
		return buffer.Bytes()
	}
	switch c.Type.Cell {
	case types.CONS:
		return sprintCons(c)
	case types.ATOM:
		return sprintAtom(c)
	default:
		buffer.WriteString("")
		return buffer.Bytes()
	}
}

func sprintAtom(c *types.Cell) []byte {

	var buffer bytes.Buffer

	if c == nil || c.Val == nil {
		buffer.WriteString("")
		return buffer.Bytes()
	}

	if c.IsStr() {
		s, _ := c.Val.(string)
		ss := fmt.Sprintf("\"%s\"", s)
		buffer.WriteString(ss)
		return buffer.Bytes()
	}

	s := fmt.Sprintf("%v", c.Val)
	buffer.WriteString(s)
		
	if c.HasTag() {
		t := fmt.Sprintf("#%v", c.Tag)
		buffer.WriteString(t)
	}
	
	return buffer.Bytes()
}

func sprintCons(c *types.Cell) []byte {
/*
	var buffer bytes.Buffer

	if c.Tag != nil {
		s := fmt.Sprintf("%v", c.Tag) // mac or func or ...
		buffer.WriteString(s)
		return buffer.Bytes()
	}
*/
	return sprintList(c)
}

// Print list e.g. (1 2 3)
//
// 	[ ]-->[ ]-->[ ]-->nil
// ( 1     2     3 )
//
func sprintList(c *types.Cell) []byte {

	//fmt.Printf("sprintList - c.Cdr.IsAtom(): %v, c.Val: %v\n", c.Cdr.IsAtom(), c.Val);

	var buffer bytes.Buffer

	printCell := func(cc *types.Cell) *types.Cell {
		buffer.Write(Sprint(cc.Car))
		cc = cc.Cdr
		if cc.Cdr != nil {
			buffer.WriteRune(SPACE)
		}
		return cc
	}

	buffer.WriteRune(LPAREN)
	cc := c

	// (a . b)
	if cc.Cdr.IsAtom() {
		buffer.Write(sprintDottedPair(cc))
	} else {

		//
		//      +---+---+   +---+---+	+---+---+
		//  l-->| a |   |-->| b	|	|-->| c |   |-->nil		(a b c)
		//      +---+---+  	+---+---+	+---+---+
		//

		for cc.Cdr != nil {
			cc = printCell(cc)

			//				    cc
			//					|
			//					v
			//      +---+---+   +---+---+
			//  a-->| a |   |-->| b	| c	| 		(a b . c)
			//      +---+---+  	+---+---+
			//
			if cc.Cdr.IsAtom() {
				buffer.Write(sprintDottedPair(cc))
				break
			}

		}
	}

	buffer.WriteRune(RPAREN)

	return buffer.Bytes()
}

// Print dotted pair e.g. (a . b)
// e.g.
//		(cons a b)   		-> "(a . b)"
//		(cons a (cons b c)) -> "(a b . c)"
// 		(cons (1) b) 		-> "((1) . b)"
func sprintDottedPair(c *types.Cell) []byte {

	var buffer bytes.Buffer

	buffer.Write(Sprint(c.Car))
	if c.Cdr.Val != NIL {
		buffer.WriteRune(SPACE)
		buffer.WriteRune(DOT)
		buffer.WriteRune(SPACE)
		buffer.Write(Sprint(c.Cdr))
	}

	return buffer.Bytes()
}
