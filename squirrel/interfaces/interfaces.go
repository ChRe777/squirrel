package interfaces

import (
	"github.com/mysheep/squirrel/types"
)
	
// -------------------------------------------------------------------------------------------------
// A parser and writer is just an I/O device
//
// 	see http://localhost:6060/pkg/io/#ReadWriter
//
//	type ReadWriter interface {
//		Reader
//		Writer
//	}

type CellReader interface {
	Read(s []byte) *types.Cell
}

type CellWriter interface {
	Write(e *types.Cell) []byte
}

type CellReadWriter interface {
	CellReader
	CellWriter
}

// -------------------------------------------------------------------------------------------------

type OpEvaluator interface {
	EvalOp(e, a *types.Cell) (*types.Cell, error)  
}

// -------------------------------------------------------------------------------------------------

type Storer interface {
	Load(s []byte) (*types.Cell, error)
	Save(s string, e *types.Cell) error
}


