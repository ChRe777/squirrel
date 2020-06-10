package io

import (
	"github.com/squirrel/types"
	"github.com/squirrel/io/parser"
	"github.com/squirrel/io/parserCell"
	"github.com/squirrel/io/printer"
	"github.com/squirrel/io/printerCell"
)

type Parser interface {
	Parse(b []byte) *types.Cell 
}

type Printer interface {
	Sprint(c *types.Cell) []byte
}

type LispParser struct {}
type CellParser struct {}

type LispPrinter struct {}
type CellPrinter struct {}

func (lp LispParser) Parse(b []byte) *types.Cell {
	return parser.Parse(b)
}

func (cp CellParser) Parse(b []byte) *types.Cell {
	return parserCell.Parse(b)
}

func (lp LispPrinter) Sprint(c *types.Cell) []byte {
	return printer.SprintLisp(c)
}

func (lp CellPrinter) Sprint(c *types.Cell) []byte {
	return printerCell.SprintCell(c)
}





