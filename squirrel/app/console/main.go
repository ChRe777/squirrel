package main

import(
//	"os"
	"fmt"
	"flag"
//	"plugin"
)

import (
	"github.com/mysheep/squirrel/ui/console/repl"
	"github.com/mysheep/squirrel/plugins/loader"	
)

const (
	myName  = "squirrel"
	welcome = "Hello World, my name is *"+myName+"*.       \n" +
			  "A fast, small and multi talented language.  \n" +
			  "Just like a "+myName+" animal.                "
)

const (
	PLUGIN_PATH				= "../../bin/"
	PLUGIN_SUFFIX			= ".so"
	PLUGIN_VERSION 			= "1.0.0"
		
	PLUGIN_IO_READER_WRITER = "reader_writer"
	PLUGIN_OPS_BUILTIN 		= "ops_builtin"
	PLUGIN_STORAGE			= "io_fs_loader_storer"
	
)

// -------------------------------------------------------------------------------------------------

func getFlagUI() string {
	uiPtr := flag.String("ui", "lisp", "io type e.g. lisp or python")
 	flag.Parse()
 	return *uiPtr
}

// -------------------------------------------------------------------------------------------------

func main() {

 	ui := getFlagUI()
 	
 	fmt.Printf("\nUI set to '%v'.\n", ui)

	plugins := loader.Load(ui, PLUGIN_PATH)
 		
 	fmt.Println()
	fmt.Println(welcome)
	
	repl.Repl(plugins)
}

// -------------------------------------------------------------------------------------------------
