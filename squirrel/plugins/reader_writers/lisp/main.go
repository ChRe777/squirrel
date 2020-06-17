package main

import (
	"github.com/mysheep/squirrel/types"
	"github.com/mysheep/squirrel/plugins/reader_writers/lisp/printer"
	"github.com/mysheep/squirrel/plugins/reader_writers/lisp/parser"
)

type any string 	// could any type

func (p any) Read(bs []byte) *types.Cell  {
	return parser.Parse(bs)
}

func (p any) Write(e *types.Cell) []byte {
	return printer.Sprint(e)	
}

var ReaderWriter any