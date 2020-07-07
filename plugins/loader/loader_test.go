package loader 

import (
	"fmt"
	"testing"
)

import (
	"github.com/mysheep/squirrel/plugins"
)	

func TestGetPluginFiles(t *testing.T) {

	files := getPluginFiles(plugins.PLUGIN_PATH)
	for _, file := range files {
		fmt.Printf("file: %v \n", file)
	}
		
}

func TestLoadPluginFile(t *testing.T) {

	specs := []struct {
		file 	string
		want 	string
	}{
		{ plugins.PLUGIN_PATH+"io_reader_writer_lisp.1.0.0.so"	, plugins.PLUGIN_TYPE_READER_WRITER	},
		{ plugins.PLUGIN_PATH+"ops_builtin.1.0.0.so"			, plugins.PLUGIN_TYPE_EVALUATOR		},
	}
	
	for _, spec := range specs {
	
		_, got, err := tryLoadKnownPlugin(spec.file)
		if got != spec.want {
			t.Errorf("Load plugin got type: %v, want: %v, file: '%v' \n", got, spec.want, spec.file)
		
			if err != nil {
				t.Errorf("Load plugin got err: %v \n", err)
			}
		}
	}
			
}

