package builtin

import (
	"os"
	"io/ioutil"
)

import (
	"github.com/squirrel/types"
	"github.com/squirrel/io/parser"
	"github.com/squirrel/core"
)

// Load loads file given in s from filesystem and return expression
func Load(s *types.Cell) *types.Cell {

	if s.IsStr() {
	
		name, err := s.AsStr()			
		
		if err != nil {
			return core.Err(err.Error())
		}
		
		bs, err := readAllBytes(name)
		
		if err != nil {
			return core.Err(err.Error())
		}

		e := parser.Parse(bs)
		return e
		
	} else {
		return core.Err("file name must be a string")
	}
}

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