package loader

import(
	"os"
	"log"
	"fmt"
	"errors"
	"strings"
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
	
		sym, plugInType, err := tryLoadKnownPlugin(file)

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

func tryLoadKnownPlugin(file string) (plugin.Symbol, string, error) {

	plugIn, err := plugin.Open(file)
	if err != nil {
		return nil, "", err
	}
		
	for _, knownType := range plugins.ALL_PLUGIN_TYPES {
		sym, err := plugIn.Lookup(knownType)
		if err == nil {
			return sym, knownType, nil
		}
	}
	
	str := strings.Join(plugins.ALL_PLUGIN_TYPES, ", ")
	
	return nil, "", errors.New(fmt.Sprintf("No known plugin type (%v) detected !!", str))
}


// -------------------------------------------------------------------------------------------------
