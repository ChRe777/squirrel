package main

import (
	"fmt"
)

import (
	"github.com/mysheep/squirrel/types"
  //"github.com/mysheep/squirrel/interfaces"
	"github.com/mysheep/squirrel/plugin/lisp/printer"
)

type printing string
type greeting string

func (p printing) Sprint(e *types.Cell) []byte {
	return printer.Sprint(e)	
}

func (p greeting) Greet() string {
	return fmt.Sprintln("Hello, I am the 'lisp' printer")
}

// -------------------------------------------------------------------------------------------------

// Exported objects
//
var Printer printing
var Greeter greeting
