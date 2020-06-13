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

// -------------------------------------------------------------------------------------------------

func main() {

 	ui := getFlagUI()
 	fmt.Printf("\nUI set to: %v \n\n", ui)
 	
 	fileParser  := getFileName(ui, "parser", "1.0.0")
 	filePrinter := getFileName(ui, "printer", "1.0.0")
 	
 	parser  := loadParserPlugin(fileParser)
 	printer := loadPrinterPlugin(filePrinter)
 	
	fmt.Println(welcome)
	
	repl.Repl(parser, printer)
}

// -------------------------------------------------------------------------------------------------

func getFileName(ui string, pluginName string, version string) string {
	file := "../bin/"+pluginName+"_"+ui+"."+version+".so"
	return file
}

// loadParserPlugin loads the parser plugin
func loadParserPlugin(file string) Parser {

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
	
	fmt.Printf("Plugin '%v' loaded. \n", file)
	
	return parser
}

// loadPrinterPlugin loads the printer plugin
func loadPrinterPlugin(file string) Printer {


	plugIn, err := plugin.Open(file)	//(*Plugin, error)
	if err != nil {
		panic(err)
	}

	printerSym, err := plugIn.Lookup("Printer")	// func (p *Plugin) Lookup(symName string) (Symbol, error)
	if err != nil {
		panic(err)
	}
	
	fmt.Printf("Plugin '%v' loaded. \n", file)
	
	return printerSym.(Printer)
}