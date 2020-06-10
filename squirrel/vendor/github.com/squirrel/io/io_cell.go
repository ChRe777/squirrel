package io

import (
	"github.com/squirrel/types"
	"github.com/squirrel/io/parserCell"
	"github.com/squirrel/io/printerCell"
)

type CellParser struct {}

type CellPrinter struct {}

func (cp CellParser) Parse(b []byte) *types.Cell { return parserCell.Parse(b) }

func (lp CellPrinter) Sprint(c *types.Cell) []byte { return printerCell.SprintCell(c) }

