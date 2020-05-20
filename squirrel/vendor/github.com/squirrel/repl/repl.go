package repl

import (
	"os"
	"fmt"
	"bufio"
	"bytes"
)

import(
	"github.com/squirrel/types"
	"github.com/squirrel/parser"
	"github.com/squirrel/builtin"
//	"github.com/squirrel/generator"
	"github.com/squirrel/evaluator"
)

var (
	PS1 = "> "
	BYE = "good bye :-)"
	QUIT = "quit"
	QUIT_ = builtin.List(builtin.Sym(QUIT))
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

func createEnv() *types.Cell {

	createList := func(fns []string) []byte {
		var b bytes.Buffer
		b.WriteRune('(')
		for _, fn := range fns { 
			b.WriteString(fn)
		}
		b.WriteRune(')')
		return b.Bytes()
	}
	
	t		 := "(t t)"
	n		 := "(nil nil)"
//	noFn     := "(no     (func (x)   (eq x '())))"
//	andFn    := "(and    (func (x y) (cond (x (cond (y 't) ('t '())))('t '()))))"
//	notFn    := "(not    (func (x)   (cond (x '()) ('t 't))))"
//	appendFn := "(append (func (x y) (cond ((no x) y) ('t (cons (car x) (append (cdr x)  y))))))"
//	pairFn   := "(pair   (func (x y) (cond ((and (no x) (no y)) '()) ((and (not (atom x)) (not (atom y))) (cons (list (car x) (car y))(pair (cdr x) (cdr y)))))) )"
//	listFn   := "(list   (func (x y) (cons x (cons y '()))))"
//	assocFn  := "(assoc  (func (x y) (cond ((eq (caar y) x) (cadar y)) ('t (assoc x (cdr y))))))"

	xs := []string{ 
		t		,
		n		,
//		noFn	,
//	    andFn	,     
//		notFn	,    
//		appendFn, 
//		pairFn	,   
//		listFn	, 
//		assocFn , 
	}
	
	env := parser.Parse(createList(xs))
	return env
}
	

func Repl() {

	reader := bufio.NewReader(os.Stdin)
	
	// TODO: PreLoaded
	env := createEnv()
	
	printHelp()
	
    for ;; {
		exp := parser.Parse(readLine(reader)); if isQuit(exp) { printBye(); break }
		res := evaluator.Eval(exp, env)
		printRes(res)
    }
}
