# ToDo

An lisp interpreter exists of following components:

	* Scanner
	* Parser
	* Evaluater
	* Generator
	* Tester
	* Spec
	* Repl
	
## ToDo

	* [ ] Implement "def" e.g. (def add (x y) (+ x y))
    * [ ] Load Squirrel from file (load "code.sqr")
	* [ ] Macro support with tagging	
	
## ToRefactor

	* [v] Symbols (sym "a") produces symbol from string
	* [ ] Full decimal number support e.g. "-123.45e-12"
	* [v] Fix >(type 12) -> sym // because nil
	* [v] TODO: Error Type - nested error
		- arc> (type (car 'a))
		- Error: "Can't take car of a"
	* [ ] FIND - error(2) - Right paren is missing - go test in evaluator
    * [ ] Check evaluator "null"-func
	* [v] Dotted pair (cons 1 2) -> (1 . 2)
	* [ ] types.go - Equal method with cmp or manual
	* [v] tests run parallel - go test -parallel 2

## ToOptimize

	* [ ] evaluator.go - eval - HASHTABLE map[string]func instead of switch -> SPEED
	* [ ] evaluator.go - environment - should be hash table
	
## Nice ToHave
	
	* [ ] channels use ->  <- chars and use f for chars
	* [ ] build in tagging into cell type "cell.Tag"
	* [ ] use facebook buck for tests and build
	* [ ] Double linked list
	  	  +-------+ <---o +-------+
  		  | cell1 | 	  | cell2 |
 		  +-------+ o---> +-------+
		- Fast insert at end
		- Fast reverse 
		- PushFront, PushEnd, Pop usw.. Like JavaScript
		

