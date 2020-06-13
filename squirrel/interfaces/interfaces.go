package interfaces

import (
	"github.com/mysheep/squirrel/types"
)
	
type Parser interface {
	Parse(s []byte) *types.Cell
}

type Printer interface {
	Sprint(e *types.Cell) []byte
}

type Storer interface {
	Load(s string) (*types.Cell, error)
	Save(s string, e *types.Cell) error
}