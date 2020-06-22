package main

import (
	"github.com/mysheep/squirrel/plugins/readerwriters/lisp/parser"
	"github.com/mysheep/squirrel/plugins/readerwriters/lisp/printer"
	"github.com/mysheep/squirrel/types"
)

// -------------------------------------------------------------------------------------------------

type any string // could any type

var ReaderWriter any // Name important to detect plugin symbol

// -------------------------------------------------------------------------------------------------

func (p any) Read(bs []byte) *types.Cell {
	return parser.Parse(bs)
}

func (p any) Write(e *types.Cell) []byte {
	return printer.Sprint(e)
}

// -------------------------------------------------------------------------------------------------
