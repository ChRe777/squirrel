package types

import (
	"fmt"
	"bytes"
)

const (
	NIL    = "nil"
	SPACE  = ' '
	DOT    = '.'
	LPAREN = '('
	RPAREN = ')'
)

func SprintCell(c *Cell) string {	
	if c == nil {
		return ""
	}
	switch c.Type.Cell { 
		case CONS: return sprintCons(c)
		case ATOM: return sprintAtom(c)
		default  : return ""
	}
}

func sprintAtom(c *Cell) string {
	if c == nil || c.Val == nil {
		return ""
	}
	return fmt.Sprintf("%v", c.Val)
}

func sprintCons(c *Cell) string {

	// Print dotted pair e.g. (a . b)
	sprintDottedPair := func(c *Cell) string {
		
		var buffer bytes.Buffer
		buffer.WriteRune(LPAREN)
		
		buffer.WriteString(SprintCell(c.Car))
		if c.Cdr.Val != NIL {
			buffer.WriteRune(SPACE)
			buffer.WriteRune(DOT)
			buffer.WriteRune(SPACE)
			buffer.WriteString(SprintCell(c.Cdr))
		}
		
		buffer.WriteRune(RPAREN)
		return buffer.String()
	}
	
	// Print list e.g. (1 2 3)
	//
	// 	[ ]-->[ ]-->[ ]-->nil
	// ( 1     2     3 )
	//
	sprintList := func(c *Cell) string {
	
		var buffer bytes.Buffer
	
		printCell := func(cc *Cell) *Cell {
			buffer.WriteString(SprintCell(cc.Car))
			cc = cc.Cdr
			if cc.Cdr != nil {
				buffer.WriteRune(SPACE)
			}
			return cc
		}
	
		buffer.WriteRune(LPAREN)
		cc := c
		for ;cc.Cdr != nil; {
			cc = printCell(cc)
		}
		buffer.WriteRune(RPAREN)
		
		return buffer.String()
	}


	
	// Dotted Pair
	// -------------------------
	// (cons a b)   -> (a . b)
	// (cons (1) b) -> ((1) .b)
	
	if c.Cdr.IsAtom() {
		return sprintDottedPair(c)
	}
	
	//  List e.g. (1 2 3)
	//
	//	CONS	ATOM  
	//  [o|o]-->[nil] Symbol
	//   |
	//   v
	//   1 ATOM
	
	return sprintList(c)
}

