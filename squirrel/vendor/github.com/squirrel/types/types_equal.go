package types

import (
	"reflect"	// TODO: Use cmp oder create your own
)

// Equal check if x equals y
func (x *Cell) Equal(y *Cell) bool {
	
	if x.Type.Cell == ATOM && y.Type.Cell == ATOM &&
	   x.Type.Atom == y.Type.Atom {
	   	// TODO: Use cmp
		if reflect.DeepEqual(x.Val, y.Val) {
			return true
		}
	}

	if x.Type.Cell == CONS && y.Type.Cell == CONS {
		if x.Car.Equal(y.Car) && x.Cdr.Equal(y.Cdr) {
			return true
		}
	} 

	return false
}

// NotEqual check if y not equal y
func (x *Cell) NotEqual(y *Cell) bool {	
	return x.Equal(y) == false
}