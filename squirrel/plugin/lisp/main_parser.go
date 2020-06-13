package main

import (
	"fmt"
)

import (
	"github.com/mysheep/squirrel/types"
  //"github.com/mysheep/squirrel/interfaces"
	"github.com/mysheep/squirrel/plugin/lisp/parser"
)

type parsing string

type greeting string

func (p parsing) Parse(s []byte) *types.Cell {
	return parser.Parse(s)	
}

func (p greeting) Greet() string {
	return fmt.Sprintln("Hello World")
}

// Export parser object
//
var Parser  parsing
var Greeter greeting