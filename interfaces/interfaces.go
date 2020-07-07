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

type FuncType 	 = func(*types.Cell, *types.Cell) *types.Cell
type MapFuncType = func(*types.Cell, *types.Cell, FuncType) *types.Cell

// -------------------------------------------------------------------------------------------------

type CellReader interface {
	Read(s []byte) *types.Cell
}

type CellWriter interface {
	Write(exp *types.Cell) []byte
}

type CellReadWriter interface {
	CellReader
	CellWriter
}

// -------------------------------------------------------------------------------------------------
// Evaluator plugin interface
// -------------------------------------------------------------------------------------------------

type Evaluator interface {
	//Eval(exp, env *types.Cell) (*types.Cell, error)  
	Eval(exp, env *types.Cell, eval FuncType) (*types.Cell, error)
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
	Evaluator
	Loader				// TODO: ReThink
	Storer				// TODO: ReThink
}

// -------------------------------------------------------------------------------------------------



