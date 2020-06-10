package main

import(
	"fmt"
	"flag"
)

import (
	"github.com/squirrel/io"
	"github.com/squirrel/repl"
)

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

func getParser(ui string) io.Parser {
	switch ui {
		case "lisp":
			return &io.LispParser{}
		case "cell":
			return &io.CellParser{}
		default:
			return &io.LispParser{}	
	}
}

func getPrinter(ui string) io.Printer {
	switch ui {
		case "lisp":
			return &io.LispPrinter{}
		case "cell":
			return &io.CellPrinter{}
		default:
			return &io.LispPrinter{}	
	}
}

func main() {

 	ui := getFlagUI()
 	fmt.Printf("\nUI set to: %v \n\n", ui)
 	
 	parser := getParser(ui)
 	printer := getPrinter(ui)
 
	fmt.Println(welcome)
	repl.Repl(parser, printer)
}