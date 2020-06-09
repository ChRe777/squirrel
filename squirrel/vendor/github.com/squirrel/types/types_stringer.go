package types

import (
	"fmt"
)
	
// String return a cell as string
func (c *Cell) String() string {
	return fmt.Sprintf("type: %v, tag: %v, val: %v, car: %p, cdr: %p", c.Type, c.Tag, c.Val, c.Car, c.Cdr)
}

// String return atom type as string
func (t AtomType) String() string {
	switch t {
		case SYMBOL: return "symbol"
		case STRING: return "string"
		case NUMBER: return "number"
		case FUNC  : return "func"
		default    : return ""
	}
}

// String return cell type as string
func (t CellType) String() string {
	switch t {
		case CONS: return "cons"
		case ATOM: return "atom"
		default	 : return ""
	}
}

// String return full cell type as string	
func (t Type) String() string {
	c := t.Cell.String()
	if a := t.Atom.String(); a != "" {
		c = a
	}
	return c
}
