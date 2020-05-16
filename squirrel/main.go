package main

import(
	"fmt"
)

import (
	"github.com/squirrel/repl"
)

func main() {

	fmt.Println("Hello World, My name is squirrel.")
	fmt.Println("A fast, small and multi talented language.")
	fmt.Println("Just like a squirrel animal.")
	
	repl.Repl()
}