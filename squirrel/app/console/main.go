package main

import(
	"os"
	"fmt"
	"flag"
	"plugin"
)

import (
	"github.com/mysheep/squirrel/ui/console/repl"
	"github.com/mysheep/squirrel/interfaces"	
)

const (
	myName  = "squirrel"
	welcome = "Hello World, my name is *"+myName+"*.       \n" +
			  "A fast, small and multi talented language.\n" +
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
 	
 	readerWriter := loadCellReaderWriterPlugin(getFileNameReaderWriter(ui, PLUGIN_IO_READER_WRITER, PLUGIN_VERSION))
 	opsBuiltin   := loadCellBuiltinPlugin(getFileNameOpsBuiltin(PLUGIN_OPS_BUILTIN, PLUGIN_VERSION))
 	storage 	 := loadIOStoragePlugin(getFileNameStorage(PLUGIN_STORAGE, PLUGIN_VERSION))
 	
 	fmt.Println()
	fmt.Println(welcome)
	
	repl.Repl(readerWriter, opsBuiltin, storage)
}

// -------------------------------------------------------------------------------------------------

func getFileNameReaderWriter(ui string, pluginName string, version string) string {
	file := PLUGIN_PATH + "io_" + ui + "_" + pluginName + "." + version + PLUGIN_SUFFIX
	return file
}

func getFileNameOpsBuiltin(pluginName string, version string) string {
	file := PLUGIN_PATH + pluginName + "." + version + PLUGIN_SUFFIX
	return file
}

func getFileNameStorage(pluginName string, version string) string {
	file := PLUGIN_PATH + pluginName + "." + version + PLUGIN_SUFFIX
	return file
}

// loadCellReaderWriterPlugin loads the reader write plugin
func loadCellReaderWriterPlugin(file string) interfaces.CellReadWriter {

	plugIn, err := plugin.Open(file)
	if err != nil {
		panic(err)
	}

	sym, err := plugIn.Lookup("ReaderWriter")
	if err != nil {
		panic(err)
	}
	
	var readerWriter interfaces.CellReadWriter
	readerWriter, ok := sym.(interfaces.CellReadWriter)
	if !ok {
		fmt.Println("unexpected type from module symbol:" +file)
		os.Exit(1)
	}
	
	fmt.Printf("Plugin '%v' loaded. \n", file)
	
	return readerWriter
}

// loadCellBuiltinPlugin loads the builtin operators plugin
func loadCellBuiltinPlugin(file string) interfaces.OpEvaluator {

	plugIn, err := plugin.Open(file)
	if err != nil {
		panic(err)
	}

	sym, err := plugIn.Lookup("Evaler")
	if err != nil {
		panic(err)
	}
	
	var opEvaluator interfaces.OpEvaluator
	opEvaluator, ok := sym.(interfaces.OpEvaluator)
	if !ok {
		fmt.Println("unexpected type from module symbol:" +file)
		os.Exit(1)
	}
	
	fmt.Printf("Plugin '%v' loaded. \n", file)
	
	return opEvaluator
}

// loadIOStoragePlugin loads the storage plugin
func loadIOStoragePlugin(file string) interfaces.OpEvaluator {

	plugIn, err := plugin.Open(file)
	if err != nil {
		panic(err)
	}

	sym, err := plugIn.Lookup("LoaderStorer")
	if err != nil {
		panic(err)
	}
	
	var storage interfaces.OpEvaluator
	storage, ok := sym.(interfaces.OpEvaluator)
	if !ok {
		fmt.Println("unexpected type from module symbol:" +file)
		os.Exit(1)
	}
	
	fmt.Printf("Plugin '%v' loaded. \n", file)
	
	return storage
}
