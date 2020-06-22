package types

import (
	"fmt"
)

// String return a cell type as string
func (c *Cell) Type_() string {
	t := fmt.Sprintf("%v", c.Type)
	if c.Tag != nil {
		t = t + "#" + fmt.Sprintf("%v", c.Tag)
	}
	return t
}
