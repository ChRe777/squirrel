package printer

import (
	"fmt"
	"bytes"
)

import (
	"github.com/squirrel/types"
)

const (
	NIL    = "nil"
	SPACE  = ' '
	DOT    = '.'
	LPAREN = '('
	RPAREN = ')'
)

func SprintCell(c *types.Cell) string {	
	if c == nil {
		return ""
	}
	switch c.Type.Cell { 
		case types.CONS: return sprintCons(c)
		case types.ATOM: return sprintAtom(c)
		default  : return ""
	}
}

func sprintAtom(c *types.Cell) string {
	if c == nil || c.Val == nil {
		return ""
	}
	
	if c.IsStr() {
		s, _ := c.Val.(string)
		return fmt.Sprintf("\"%s\"", s)
	}
	
	return fmt.Sprintf("%v", c.Val)
}

func sprintCons(c *types.Cell) string {

	if c.Tag != nil {
		return fmt.Sprintf("%v", c.Tag)		// mac or func
	}

	//fmt.Printf("sprintCons - c.Cdr.IsAtom(): %v, c.Val: %v\n", c.Cdr.IsAtom(), c.Val);

	//if c.Cdr.IsAtom() {
	//	return sprintDottedPair(c)
	//}
	
	return sprintList(c)
}

// Print list e.g. (1 2 3)
//
// 	[ ]-->[ ]-->[ ]-->nil
// ( 1     2     3 )
//
func sprintList(c *types.Cell) string {

	//fmt.Printf("sprintList - c.Cdr.IsAtom(): %v, c.Val: %v\n", c.Cdr.IsAtom(), c.Val);

	var buffer bytes.Buffer

	printCell := func(cc *types.Cell) *types.Cell {
		buffer.WriteString(SprintCell(cc.Car))
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
		buffer.WriteString(sprintDottedPair(cc))	
	} else {
	
//
//      +---+---+   +---+---+	+---+---+
//  l-->| a |   |-->| b	|	|-->| c |   |-->nil		(a b c)
//      +---+---+  	+---+---+	+---+---+
//

		for ;cc.Cdr != nil; {
			cc = printCell(cc)
		
//				    cc
//					|		
//					v
//      +---+---+   +---+---+
//  a-->| a |   |-->| b	| c	| 		(a b . c)
//      +---+---+  	+---+---+
//
			if cc.Cdr.IsAtom() {		
				buffer.WriteString(sprintDottedPair(cc))
				break;
			}
		
		}
	}
	
	buffer.WriteRune(RPAREN)
	
	return buffer.String()
}

// Print dotted pair e.g. (a . b)
// e.g.
//		(cons a b)   		-> "(a . b)"
//		(cons a (cons b c)) -> "(a b . c)"
// 		(cons (1) b) 		-> "((1) . b)"
func sprintDottedPair(c *types.Cell) string {
	
	var buffer bytes.Buffer
	
	buffer.WriteString(SprintCell(c.Car))
	if c.Cdr.Val != NIL {
		buffer.WriteRune(SPACE); buffer.WriteRune(DOT); buffer.WriteRune(SPACE)
		buffer.WriteString(SprintCell(c.Cdr))
	}
	
	return buffer.String()
}
