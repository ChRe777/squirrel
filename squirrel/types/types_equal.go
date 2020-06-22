package types

import (
	"reflect" // TODO: Use cmp oder create your own equal
)

// Equal check if x equals y
func (c *Cell) Equal(y *Cell) bool {

	if c.Type.Cell == ATOM && y.Type.Cell == ATOM &&
		c.Type.Atom == y.Type.Atom {
		if reflect.DeepEqual(c, y) { // TODO: Use cmp
			return true
		}
	}

	if c.Type.Cell == CONS && y.Type.Cell == CONS {
		if c.Car.Equal(y.Car) && c.Cdr.Equal(y.Cdr) {
			return true
		}
	}

	return false
}

// NotEqual check if y not equal y
func (c *Cell) NotEqual(y *Cell) bool {
	return c.Equal(y) == false
}
