package builtin

import (
	//"testing"
)

import (
//	"github.com/mysheep/squirrel/types"
//	"github.com/mysheep/squirrel/evaluator"
//	"github.com/mysheep/squirrel/core"
//	"github.com/mysheep/squirrel/plugins/readerwriters/lisp/parser"	

)
/*
func TestBackquote(t *testing.T) {

	env := "((x 1) (xs (1 2)))"
	a := parser.Parse([]byte(env))

	specs := []struct {
		exp		string
		want  	string
	} {
		{"`(,x)"				, "(1)"   			},
		{"`(,@xs)"				, "(1 2)"   		},	
		{"`(a ,@xs b)"			, "(a 1 2 b)"   	},	
		{"`(a (b ,@xs c) d)"	, "(a (b 1 2 c) d)"	},
		{"`(,@xs ,@xs)"			, "(1 2 1 2)"   	},	
	}
	
	for _, spec := range specs {
	
		e    := parser.Parse([]byte(spec.exp))
		want := parser.Parse([]byte(spec.want))
	
		got := Backquote(e, a, eval)
 
 		if got.NotEqual(want) {
 			t.Errorf("TestBackquote - got: %v, want: %v", got, want)
 		}
	
	}
}


func TestExpandList(t *testing.T) {

	specs := []struct {
		xs		string
		ys		string
		want  	string
	} {
	
	}

}

func TestExpand(t *testing.T) {

	specs := []struct {
		xs		string
		ys		string
		want  	string
	} {
	
	}

}

func TestUnquote(t *testing.T) {

	specs := []struct {
		xs		string
		ys		string
		want  	string
	} {
	
	}
}
*/