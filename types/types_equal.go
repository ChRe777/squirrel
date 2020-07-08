package types

import (
	"github.com/shopspring/decimal"
)

import (
	"reflect" // TODO: Use cmp oder create your own equal
)

// Equal check if x equals y
func (c *Cell) Equal(y *Cell) bool {

	if c.Type.Cell == ATOM && y.Type.Cell == ATOM &&
	   c.Type.Atom == y.Type.Atom {
	   
	    if c.IsNumber() && y.IsNumber() {
	    	c_, _ := c.Val.(decimal.Decimal)
	    	y_, _ := y.Val.(decimal.Decimal)
	    	return c_.Equal(y_)
	    }
	    
	    if c.IsStr() && y.IsStr() {
	    	c_, _ := c.Val.(string)
	    	y_, _ := y.Val.(string)
	    	return c_ == y_
	    }
	    
	    if c.IsSymbol() && y.IsSymbol() {
	    	c_, _ := c.Val.(string)
	    	y_, _ := y.Val.(string)
	    	return c_ == y_
	    }
		
		// All other types (error, func, ...) use 
		//
		if reflect.DeepEqual(c, y) {
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
