package plugins

import (
	"github.com/mysheep/squirrel/interfaces"
)

const (
	PLUGIN_PATH				= "../../bin/"				// TODO: ReThink
	PLUGIN_SUFFIX			= ".so"
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