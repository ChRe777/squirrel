package loader

import(
	"os"
	"io/ioutil"
)

import (
	"github.com/mysheep/squirrel/types"
	"github.com/mysheep/squirrel/evaluator/core"
	"github.com/mysheep/squirrel/plugins/readerwriters/lisp/parser"
)

func Load(s *types.Cell ) *types.Cell {

	if s.IsStr() == false {
		return core.Err_("file name must be a string")
	}
	
	name, err := s.AsStr()			
	if err != nil {
		return core.Err_(err.Error())
	}
	
	bs, err := readAllBytes(name)
	if err != nil {		
		return core.Err_(err.Error())
	}

	e := parser.Parse(bs)
	
	return e	
	
}

// readAllBytes all bytes from file by name
func readAllBytes(name string) ([]byte, error) {
	
	file, err := os.Open(name)
    if err != nil {
    	return nil, err
    }
    defer file.Close()

  	bs, err := ioutil.ReadAll(file)
  	if err != nil {
  		return nil, err
  	}
  		
	return bs, nil
}
