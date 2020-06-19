package loader

import(
	"os"
	"io/ioutil"
)

import (
	"github.com/mysheep/squirrel/types"
	"github.com/mysheep/squirrel/generator"
	"github.com/mysheep/squirrel/plugins/reader_writers/lisp/parser"
)

func Load(s *types.Cell ) *types.Cell {

	if s.IsStr() == false {
		return generator.Error("file name must be a string")
	}
	
	name, err := s.AsStr()			
	if err != nil {
		return generator.Error(err.Error())
	}
	
	bs, err := readAllBytes(name)
	if err != nil {		
		return generator.Error(err.Error())
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
