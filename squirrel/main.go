package main

import(
	"fmt"
)

import (
	"github.com/squirrel/repl"
)

const (
	welcomeMessage = "Hello World, My name is squirrel.\n" +
					 "A fast, small and multi talented language.\n" +
					 "Just like a squirrel animal."
)

func main() {

	fmt.Println(welcomeMessage)
	
	repl.Repl()
}