package storer

import (
	"io/ioutil"
)

import (
	"github.com/mysheep/squirrel/types"
	"github.com/mysheep/squirrel/generator"			// TODO: Use core
	"github.com/mysheep/squirrel/plugins/lisp/printer"
)

// -------------------------------------------------------------------------------------------------
	
func Store(loc *types.Cell, s *types.Cell ) *types.Cell {

	if loc.IsStr() == false {
		return generator.Error("file name must be a string")
	}
	
	name, err := loc.AsStr()			
	if err != nil {
		return generator.Error(err.Error())
	}
	
	bs := printer.Sprint(s)
	
	err = writeAllBytes(name, bs)
	if err != nil {		
		return generator.Error(err.Error())
	}
	return generator.NIL
}

// -------------------------------------------------------------------------------------------------
	
func writeAllBytes(name string, bs[]byte) error {
	err := ioutil.WriteFile(name, bs, 0644)
	return err
}

// -------------------------------------------------------------------------------------------------
	