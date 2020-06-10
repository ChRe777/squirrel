package io

import (
	"github.com/squirrel/types"
	"github.com/squirrel/io/parser"
	"github.com/squirrel/io/printer"
)

type LispParser struct {}

type LispPrinter struct {}

func (lp LispParser) Parse(b []byte) *types.Cell { return parser.Parse(b) }

func (lp LispPrinter) Sprint(c *types.Cell) []byte { return printer.SprintLisp(c) }
