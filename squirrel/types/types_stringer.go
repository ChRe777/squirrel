package types

// String return a cell as string
func (c *Cell) String() string {
	return string(Sprint(c))
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
