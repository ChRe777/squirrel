package types

type Identifier string

type CellType 	byte

type AtomType 	byte

// Type of Cells atom or cons
// if cell is atom then also
// the type of the atom symbol, string, number, bool or error
type Type struct {
	Cell	CellType
	Atom 	AtomType
}

// Cell types like Atom or Cons cells
const (
	CONS 	CellType = iota + 1
	ATOM
)

// Atom types like Strings or Numbers or Booleans
const (
	SYMBOL 	AtomType = iota + 1
	STRING  // e.g. "foo", "bar"
	NUMBER	// e.g. 123.4e-10	
	FUNC	// e.g. (func (x) (+ x 1))
	//BOOL    // e.g. true, false
	ERROR	// e.g.	error("Can't take car of a")
)

//  Cell
//  ----
//
//  [ type | val | car | cdr ]
//                  |     |
//					v     v
//
//  TODO: Double linked list
//  ------------------------
//
//  +-------+ <---o	+-------+
//  | cell1 | 		| cell2 |
// 	+-------+ o--->	+-------+
//
type Cell struct {
    Type 	Type
    Tag     interface{}
    Val 	interface{}
    Car 	*Cell
	Cdr 	*Cell
}

// IsCons checks, if cell is a cons
func (c *Cell) IsCons() bool {
	return c.Type.Cell == CONS
}

// IsAtom checks, if cell is an atom
func (c *Cell) IsAtom() bool {
	return c.Type.Cell == ATOM
}

// IsSymbol checks, if cell is a atom of type symbol
func (c *Cell) IsSymbol() bool {
	return c.IsAtom() && (c.Type.Atom == SYMBOL)
}


