package main

import(
	"os"
	"fmt"
	"flag"
	"plugin"
)

import (
	"github.com/mysheep/squirrel/ui/repl"
	"github.com/mysheep/squirrel/types"		
)

// -------------------------------------------------------------------------------------------------

type Parser interface {
	Parse(s []byte) *types.Cell
}

// -------------------------------------------------------------------------------------------------

type Printer interface {
	Sprint(e *types.Cell) []byte
}

// -------------------------------------------------------------------------------------------------

type Greeter interface {
	Greet() string
}

// -------------------------------------------------------------------------------------------------

const (
	welcome = "Hello World, my name is *squirrel*.       \n" +
			  "A fast, small and multi talented language.\n" +
			  "Just like a squirrel animal.                "
)

func getFlagUI() string {
	uiPtr := flag.String("ui", "lisp", "io type e.g. lisp or python")
 	flag.Parse()
 	return *uiPtr
}

func main() {



 	ui := getFlagUI()
 	
 	fmt.Printf("\nUI set to: %v \n\n", ui)
 	
 	loadGreeterPlugin(ui)
 	
 	parser  := loadParserPlugin(ui)
 	printer := loadPrinterPlugin(ui)
 	
	fmt.Println(welcome)
	
	repl.Repl(parser, printer)
	
}


// loadParserPlugin loads the parser plugin
func loadGreeterPlugin(ui string)  {

	file := "../bin/parser_"+ui+".1.0.0.so"

	plugIn, err := plugin.Open(file)
	if err != nil {
		panic(err)
	}

	parserSym, err := plugIn.Lookup("Greeter")
	if err != nil {
		panic(err)
	}
	
	var greeter Greeter
	greeter, ok := parserSym.(Greeter)
	if !ok {
		fmt.Println("unexpected type from module symbol:" +file)
		os.Exit(1)
	}
	
	fmt.Println(greeter.Greet())
}

// -------------------------------------------------------------------------------------------------

// loadParserPlugin loads the parser plugin
func loadParserPlugin(ui string) Parser {

	file := "../bin/parser_"+ui+".1.0.0.so"

	plugIn, err := plugin.Open(file)
	if err != nil {
		panic(err)
	}

	parserSym, err := plugIn.Lookup("Parser")
	if err != nil {
		panic(err)
	}
	
	var parser Parser
	parser, ok := parserSym.(Parser)
	if !ok {
		fmt.Println("unexpected type from module symbol:" +file)
		os.Exit(1)
	}
	
	return parser
}

// loadPrinterPlugin loads the printer plugin
func loadPrinterPlugin(ui string) Printer {

	plugIn, err := plugin.Open("../bin/printer_"+ui+".1.0.0.so")	//(*Plugin, error)
	if err != nil {
		panic(err)
	}

	printerSym, err := plugIn.Lookup("Printer")	// func (p *Plugin) Lookup(symName string) (Symbol, error)
	if err != nil {
		panic(err)
	}
	
	return printerSym.(Printer)
}