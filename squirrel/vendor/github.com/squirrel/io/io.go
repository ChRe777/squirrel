package io

import (
	"github.com/squirrel/types"
)

type Parser interface {
	Parse(b []byte) *types.Cell 
}

type Printer interface {
	Sprint(c *types.Cell) []byte
}



