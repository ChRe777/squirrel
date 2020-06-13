package repl

import (
	"os"
	"fmt"
	"bufio"
	"bytes"
)

import(
	"github.com/mysheep/squirrel/types"
	"github.com/mysheep/squirrel/evaluator"
	"github.com/mysheep/squirrel/interfaces"
)

var (
	PS1 = ">>> "
	BYE = "good bye :-)"
//	QUIT = "quit"
//	QUIT_ = core.Cons(core.Sym(QUIT), core.NIL)
)

func readLine(reader *bufio.Reader) []byte {
	fmt.Print(PS1); bs, _ := reader.ReadBytes('\n')
	return bs    
}

func printRes(printer interfaces.Printer, e *types.Cell) {
	fmt.Printf("%v\n", string(printer.Sprint(e)))
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

//func isQuit(e *types.Cell) bool {
//	return e.Equal(QUIT_)
//}

func createList (fns []string) []byte {
	var b bytes.Buffer
	b.WriteRune('('); for _, fn := range fns { b.WriteString(fn) }; b.WriteRune(')')
	return b.Bytes()
}

func Repl(parser interfaces.Parser, printer interfaces.Printer) {

	reader := bufio.NewReader(os.Stdin)
	
	// env := parser.Parse(createList(builtin.Env()));
		
	printHelp()
	
    for ;; {
		exp := parser.Parse(readLine(reader)); 
		//if isQuit(exp) { 
		//	printBye(); break 
		//}
		res := evaluator.Eval(exp, nil)		// TODO: return Quit = TRUE
		printRes(printer, res)
    }
}
