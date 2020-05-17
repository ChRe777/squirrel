# ToDo

An lisp interpreter exists of following components:

	* Scanner
	* Parser
	* Evaluater
	* Generator
	* Tester
	* Spec
	* Repl
	
## ToRefactor

	* [ ] Symbols (sym "a") produces symbol from string
	* [ ] Full decimal number support e.g. "-123.45e-12"
	* [ ] Fix >(type 12) -> sym // because nil
	* [ ] TODO: Error Type - nested error
		- arc> (type (car 'a))
		- Error: "Can't take car of a"
	* [ ] FIND - error(2) - Right paren is missing - go test in evaluator
    * [ ] Check evaluator "null"-func
	* [v] Dotted pair (cons 1 2) -> (1 . 2)
    * [ ] EMPTY := cons(NIL,NIL); -> NIL = NICHTS
	* [ ] types.go - Equal method
	* [v] tests run parallel - go test -parallel 2
	* [ ] evaluator.go - eval - map[string]func instead of switch -> SPEED
	* [ ] evaluator.go - environment - should be hash table
	
## Nice ToHave
	
	* [ ] build in tagging into cell type "cell.Tag"
	* [ ] use facebook buck for tests and build
	* [ ] Double linked list
	  	  +-------+ <---o +-------+
  		  | cell1 | 	  | cell2 |
 		  +-------+ o---> +-------+
		- Fast insert at end
		- Fast reverse 
		- PushFront, PushEnd, Pop usw.. Like JavaScript
		

