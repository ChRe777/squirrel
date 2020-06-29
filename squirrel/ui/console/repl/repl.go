package repl

import (
	"os"
	"fmt"
	"bufio"
)

import(
	"github.com/mysheep/squirrel/types"
	"github.com/mysheep/squirrel/evaluator"
	"github.com/mysheep/squirrel/generator"
	"github.com/mysheep/squirrel/interfaces"
	"github.com/mysheep/squirrel/plugins"
)

// -------------------------------------------------------------------------------------------------

var (
	PS1 	= ">>> "
	BYE 	= "good bye :-)"
	QUIT 	= "quit"
	QUIT_ 	= generator.Cons(generator.Sym(QUIT), generator.NIL)
	STDIN_READER = bufio.NewReader(os.Stdin)
)

// -------------------------------------------------------------------------------------------------

func Repl(plugins *plugins.Plugins) {

	// Check MUSTS !!!
	if plugins.ReaderWriter == nil {
		panic("No ReaderWriter plugin - It's a MUST plugin!!")
	}

	env := createEnvironmentList()				// TODO: Fix EMPTY Environment '()
	var parse = plugins.ReaderWriter.Read

	printHelp()
	
	evaluator.SetEvaluators(plugins.Evaluators)					
	
    for ;; {
		expr := parse(readLine(STDIN_READER)); 
		if isQuit(expr) { 
			printBye(); break 
		}
		result := evaluator.Eval(expr, env)		// TODO: return Quit = TRUE
		printResult(plugins.ReaderWriter, result)
    }
}

// -------------------------------------------------------------------------------------------------

func createEnvironmentList() *types.Cell {
	// ((nil nil)(t t))
	//
	tt := generator.Cons(generator.Sym("t"), generator.Cons(generator.Sym("t"), generator.NIL))	
	nl := generator.Cons(generator.Sym("nil"), generator.Cons(generator.Sym("nil"), generator.NIL)) 	
	//
	return generator.Cons(nl, generator.Cons(tt, generator.NIL))
}

// -------------------------------------------------------------------------------------------------

func readLine(reader *bufio.Reader) []byte {
	fmt.Print(PS1); bs, _ := reader.ReadBytes('\n')
	return bs    
}

// -------------------------------------------------------------------------------------------------

func printResult(writer interfaces.CellWriter, e *types.Cell) {
	fmt.Printf("%v\n", string(writer.Write(e)))
}

// -------------------------------------------------------------------------------------------------

func printBye() {
	fmt.Println()
	fmt.Println(BYE)
	fmt.Println()
}

// -------------------------------------------------------------------------------------------------

func printHelp() {
	fmt.Println()
	fmt.Println("to stop enter (quit) or CTRL+C")
	fmt.Println()
}

// -------------------------------------------------------------------------------------------------

func isQuit(e *types.Cell) bool {
	return e.Equal(QUIT_)
}

