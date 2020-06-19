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

func Repl(readerWriter 	interfaces.CellReadWriter, opsBuiltin 	interfaces.OpEvaluator, storage interfaces.OpEvaluator ) {

	env := createEnvironmentList()	// TODO: Fix EMPTY Enviroment '()
	var parse = readerWriter.Read

	printHelp()
	
	evaluator.SetOpEvaluator(opsBuiltin)	// TODO: ReThink
	evaluator.SetStorage(storage)			// TODO: ReThink - Optional Injection
	
    for ;; {
		expr := parse(readLine(STDIN_READER)); 
		if isQuit(expr) { 
			printBye(); break 
		}
		result := evaluator.Eval(expr, env)		// TODO: return Quit = TRUE
		printResult(readerWriter, result)
    }
}

// -------------------------------------------------------------------------------------------------

func createEnvironmentList() *types.Cell {
	// (a 1)
	a1 := generator.Cons(generator.Sym("a"), generator.Cons(generator.Num("1"), generator.NIL)) 	
	// ((a 1))
	return generator.Cons(a1, generator.NIL)
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

