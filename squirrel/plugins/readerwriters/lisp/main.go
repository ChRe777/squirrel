package main

import (
	"github.com/mysheep/squirrel/types"
	"github.com/mysheep/squirrel/plugins/readerwriters/lisp/printer"
	"github.com/mysheep/squirrel/plugins/readerwriters/lisp/parser"
)

// -------------------------------------------------------------------------------------------------

type any string 		// could any type

var ReaderWriter any	// Name important to detect

// -------------------------------------------------------------------------------------------------

func (p any) Read(bs []byte) *types.Cell  {
	return parser.Parse(bs)
}

func (p any) Write(e *types.Cell) []byte {
	return printer.Sprint(e)	
}

// -------------------------------------------------------------------------------------------------
