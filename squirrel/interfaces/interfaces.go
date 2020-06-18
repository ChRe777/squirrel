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

// -------------------------------------------------------------------------------------------------
// Cell reader and writer plugin interface (as user interface for the language)
// -------------------------------------------------------------------------------------------------

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
// Evaluator plugin interface
// -------------------------------------------------------------------------------------------------

type OpEvaluator interface {
	EvalOp(e, a *types.Cell) (*types.Cell, error)  
}

// -------------------------------------------------------------------------------------------------
// Storage plugin interface
// -------------------------------------------------------------------------------------------------

type Loader interface {
	Load(location string) (*types.Cell, error)
}

type Storer interface {
	Store(location string, e *types.Cell) error
}

type Storage interface {
	Loader
	Storer
}

// -------------------------------------------------------------------------------------------------



