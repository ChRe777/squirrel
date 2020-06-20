package plugins

import (
	"github.com/mysheep/squirrel/interfaces"
)

const (
	PLUGIN_PATH				= "../../bin/"
	PLUGIN_SUFFIX			= ".so"
	PLUGIN_VERSION 			= "1.0.0"
		
	PLUGIN_IO_READER_WRITER = "reader_writer"
	PLUGIN_OPS_BUILTIN 		= "ops_builtin"
	PLUGIN_STORAGE			= "io_fs_loader_storer"	
)

const (
	PLUGIN_TYPE_READER_WRITER 	= "ReaderWriter"		// TODO: ReThink Naming
	PLUGIN_TYPE_EVALUATOR		= "Evaluator"			// TODO: ReThink Naming
)

var (
	ALL_PLUGIN_TYPES = []string {
		PLUGIN_TYPE_READER_WRITER,
		PLUGIN_TYPE_EVALUATOR,
	}
)

type Plugins struct {
	ReaderWriter	  interfaces.CellReadWriter
	Evaluators 		[]interfaces.Evaluator
}