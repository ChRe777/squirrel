package types

import (
	"errors"
	"strings"
)

type Identifier string

type CellType byte

type AtomType byte

// Type of Cells atom or cons
// if cell is atom then also
// the type of the atom symbol, string, number, bool or error
type Type struct {
	Cell CellType
	Atom AtomType
}

// Cell types like Atom or Cons cells
const (
	CONS CellType = iota + 1
	ATOM
)

// Atom types like Strings or Numbers or Booleans
const (
	SYMBOL AtomType = iota + 1
	STRING          // e.g. "foo", "bar"
	NUMBER          // e.g. 123.4e-10
	FUNC            // e.g. (func (x) (+ x 1))
	//BOOL    	// e.g. true, false (currently we have 't nil)
	ERROR // e.g.	error("Can't take car of a")
)

//  Cell
//  ----
//
//  +------+-----+-----+-----+
//  | type | val | car | cdr |
//  +------+-----+--|--+--|--+
//					v     v
//
//  Double linked list
//  ------------------------
//  We will need another pointer
//
//    car   cdr   cgr
//	+-----+-----+-----+
//  |  o  |  o  |  o  |
//  +-----+-----+-----+
//
//			   cgr
//  +-------+ <---o	+-------+
//  | cell1 | 		| cell2 |
// 	+-------+ o--->	+-------+
//			   car
//
type Cell struct {
	Type Type
	Tag  interface{}
	//  Level	interface{}		// each function is in a level -> level 2 can access level 1
	//  Sec		interface{}		// security - more information later
	Val interface{}
	Car *Cell
	Cdr *Cell
	//	Cgr		*Cell
}

//  Levels
//  ------
//
//  +--------------+
//  | Level 2  o   |		// Level 2 functions can access functions of level 1	(VIEW)
//  +----------|---+
//  | Level 1  v o |		// Level 1 functions can access functions of level 0	(PRESENTER)
//  +------------|-+
//  | Level 0    v |		// Level 0 functions can access functions of ....		(MODEL)
//  +--------------+

// IsCons checks, if cell is a cons
func (c *Cell) IsCons() bool {
	return c.Type.Cell == CONS
}

// IsAtom checks, if cell is an atom
func (c *Cell) IsAtom() bool {
	return c.Type.Cell == ATOM
}

// IsSymbol checks, if cell is an atom of type symbol
func (c *Cell) IsSymbol() bool {
	return c.IsAtom() && (c.Type.Atom == SYMBOL)
}

// IsErr checks, if cell is an atom of type error
func (c *Cell) IsErr() bool {
	return c.IsAtom() && (c.Type.Atom == ERROR)
}

// IsStr checks, if cell is an atom of type string
func (c *Cell) IsStr() bool {
	return c.IsAtom() && (c.Type.Atom == STRING)
}

// IsTagged checks, if cell is tagged with string t
func (c *Cell) IsTagged(t string) bool {
	s, ok := c.Tag.(string)
	if ok {
		return s == t
	}
	return false 
}

// HasTag checks, if cell has a tag
func (c *Cell) HasTag() bool {
	return c.Tag != nil
}

// AsStr checks, if cell is an atom of type string and return string value
func (c *Cell) AsStr() (string, error) {
	if c.IsAtom() && c.IsStr() {
		s, ok := c.Val.(string)
		if ok {
			return strings.Trim(s, "\""), nil
		}
	}
	return "", errors.New("no atom and/or string")
}
