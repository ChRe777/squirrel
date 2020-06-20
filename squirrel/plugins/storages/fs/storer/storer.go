package storer

import (
	"io/ioutil"
)

import (
	"github.com/mysheep/squirrel/types"
	"github.com/mysheep/squirrel/evaluator/core"
	"github.com/mysheep/squirrel/plugins/readerwriters/lisp/printer"
)

// -------------------------------------------------------------------------------------------------
	
func Store(loc *types.Cell, s *types.Cell ) *types.Cell {

	if loc.IsStr() == false {
		return core.Err_("file name must be a string")
	}
	
	name, err := loc.AsStr()			
	if err != nil {
		return core.Err_(err.Error())
	}
	
	bs := printer.Sprint(s)
	
	err = writeAllBytes(name, bs)
	if err != nil {		
		return core.Err_(err.Error())
	}
	
	return loc
}

// -------------------------------------------------------------------------------------------------
	
func writeAllBytes(name string, bs[]byte) error {
	err := ioutil.WriteFile(name, bs, 0644)
	return err
}

// -------------------------------------------------------------------------------------------------
	