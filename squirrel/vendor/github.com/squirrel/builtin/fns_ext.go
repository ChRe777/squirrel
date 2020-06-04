package builtin

import (
	"fmt"
	"os"
	"log"
	"io/ioutil"
)

import (
	"github.com/squirrel/types"
	"github.com/squirrel/parser"
	"github.com/squirrel/core"
)

func Load (s *types.Cell) *types.Cell {

	if s.IsStr() {
	
		name := fmt.Sprintf("%v", s)
		
		file, err := os.Open(name)
    	if err != nil {
    		return core.Err(err.Error())
    	}
    	defer file.Close()

  		bs, err := ioutil.ReadAll(file)
  		
  		if err != nil {
  			log.Fatal(err)
  			return core.Err(err.Error())
  		}
	
		e := parser.Parse(bs)
		return e
		
	} else {
		return core.Err("s must be a string")
	}
}