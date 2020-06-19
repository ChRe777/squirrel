package loader

import(
	"os"
	"log"
	"fmt"
	"errors"
	"plugin"
    "path/filepath"
)

import (
	"github.com/mysheep/squirrel/plugins"
	"github.com/mysheep/squirrel/interfaces"
)


// -------------------------------------------------------------------------------------------------

func Load(ui string, pluginPath string) *plugins.Plugins {
	
	loadedPlugins := &plugins.Plugins{
		ReaderWriter: nil,
		Evaluators	: []interfaces.Evaluator{ }, // Empty list
	}
	
	for _, file := range getPluginFiles(pluginPath) {
	
		sym, plugInType, err := loadPlugin(file)

		if err == nil {
		
			switch plugInType {

			//
			//	READER WRITER
			//
			case plugins.PLUGIN_TYPE_READER_WRITER:
				var readerWriter interfaces.CellReadWriter
				readerWriter, ok := sym.(interfaces.CellReadWriter)
				if ok {
					fmt.Printf("Plugin '%v' loaded!\n", file)
					loadedPlugins.ReaderWriter = readerWriter
				}
			
			//
			// EVALUATOR
			//
			case plugins.PLUGIN_TYPE_EVALUATOR:
				var evaluator interfaces.Evaluator
				evaluator, ok := sym.(interfaces.Evaluator)
				if ok {
					fmt.Printf("Plugin '%v' loaded!\n", file)
					loadedPlugins.Evaluators = append(loadedPlugins.Evaluators, evaluator)
				}
			}
		} 
	
	}
	
	if loadedPlugins.ReaderWriter == nil {
		panic("ReaderWriter plugin is a must!!")
	}

	return loadedPlugins
}

func visit(files *[]string) filepath.WalkFunc {
    return func(path string, info os.FileInfo, err error) error {
        if err != nil {
            log.Fatal(err)
        }
        
        // Avoid to walk of directories
        if info.IsDir() {
    		return nil
		}

		// Only files with .so
		if filepath.Ext(path) != plugins.PLUGIN_SUFFIX {
    		return nil
		}

        *files = append(*files, path)
        
        return nil
    }
}

func getPluginFiles(path string) []string {
	
	files := []string{}
	err := filepath.Walk(path, visit(&files))
	
	if err != nil {
		log.Fatal(err)
	}
	
	return files
}

func loadPlugin(file string) (plugin.Symbol, string, error) {

	plugIn, err := plugin.Open(file)
	if err != nil {
		return nil, "", err
	}
	
	// Try this
	
	sym, err := plugIn.Lookup(plugins.PLUGIN_TYPE_READER_WRITER)
	if err == nil {
		return sym, plugins.PLUGIN_TYPE_READER_WRITER, nil
	}
	
	// Try next
	
	sym, err = plugIn.Lookup(plugins.PLUGIN_TYPE_EVALUATOR)
	if err == nil {
		return sym, plugins.PLUGIN_TYPE_EVALUATOR, nil
	}

	return nil, "", errors.New("No known plugin type detected !!")
}


// -------------------------------------------------------------------------------------------------

/*
func getFileNameReaderWriter(ui string, pluginName string, version string) string {
	file := plugins.PLUGIN_PATH + "io_" + ui + "_" + pluginName + "." + version + PLUGIN_SUFFIX
	return file
}

func getFileNameOpsBuiltin(pluginName string, version string) string {
	file := plugins.PLUGIN_PATH + pluginName + "." + version + PLUGIN_SUFFIX
	return file
}

func getFileNameStorage(pluginName string, version string) string {
	file := plugins.PLUGIN_PATH + pluginName + "." + version + PLUGIN_SUFFIX
	return file
}


// loadCellReaderWriterPlugin loads the reader write plugin
func loadCellReaderWriterPlugin(file string) (interfaces.CellReadWriter, error) {

	plugIn, err := plugin.Open(file)
	if err != nil {
		return nil, err
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

*/