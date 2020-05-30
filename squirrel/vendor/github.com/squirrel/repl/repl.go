package repl

import (
	"os"
	"fmt"
	"bufio"
)

import(
	"github.com/squirrel/types"
	"github.com/squirrel/core"
	"github.com/squirrel/parser"
	"github.com/squirrel/evaluator"
)

var (
	PS1 = "> "
	BYE = "good bye :-)"
	QUIT = "quit"
	QUIT_ = core.List(core.Sym(QUIT))
)

func readLine(reader *bufio.Reader) []byte {
	fmt.Print(PS1); bs, _ := reader.ReadBytes('\n')
	return bs    
}

func printRes(e *types.Cell) {
	fmt.Printf("%v\n", types.SprintCell(e))
}

func printBye() {
	fmt.Println()
	fmt.Println(BYE)
	fmt.Println()
}

func printHelp() {
	fmt.Println()
	fmt.Println("to stop enter (quit) or CTRL+C")
	fmt.Println()
}

func isQuit(e *types.Cell) bool {
	return e.Equal(QUIT_)
}


func Repl() {

	reader := bufio.NewReader(os.Stdin)
		
	printHelp()
	
    for ;; {
		exp := parser.Parse(readLine(reader)); if isQuit(exp) { printBye(); break }
		res := evaluator.Eval(exp, env)
		printRes(res)
    }
}
