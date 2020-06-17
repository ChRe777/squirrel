# Squirrel

Squirrel is a new language based on root of lisp and many
other ideas connected together. Tagging. Connectable objects.

## Design


	"Uncle"-Bob Clean Architecture
	==============================

	- Evaluator   	is the 'core' business
	
	- Parser		is a plugin to the business
	- UI   			is a plugin to the business
	- Store 		is a plugin to the business

			  [plugin]
			  +--------------+
 			  |      UI      |  (Repl) 
 			  +--------------+
 			
 	+---------------+	+---------------+
 	|    Reader     |   |    Printer    |  
 	+---------------+	+---------------+		
 	
 				
	 ====================================
	||								    ||		[plugin]
	||		  +---------------+			||		+---------------+
 	||		  |   Evaluator   | 		||		|     Store     |  
 	||		  +---------------+			||		+---------------+
 	||        							||
	 ====================================
	
	
 						[plugin]
 	+---------------+	+---------------+
 	|    Scanner    |<--|    Parser     |  
 	+---------------+	+-------+-------+
								|
						+-------v-------+
						|    Generator  |	
						+---------------+	
						
						
	Flow
	====
	
	Reader -> Parser -> Evaluator -> Printer
	
	Reader reads a string from ui (=stdin)
	Parser parses the string a generates cells
	Evaluator evaluates the cells according to the rules
	Printer prints the cells to the ui (=stdout)
	Store stores the current state of evaluator
	
						
											
## Write tests first

First you should write the tests which are the specifications
of the language. It is a design of the language on paper.
Never get deceived to start coding to soon. Make first your
plan of journey and then start the journey. The plan is your
target otherwise you will not reach anything.

### car

	> (car '(1 2 3))
	> 1

### cdr

	> (cdr '(1 2 3))
	> (2 3)
	
### atom

	> (atom 1)
	> 'true
	
### 


## Dependecy Graph Tool "depth"

see https://github.com/KyleBanks/depth

### Usage

	> depth .
	> depth github.com/KyleBanks/depth/cmd/depth

## VENDOR

evaluator/evaluator_lib.go:8:2: cannot find package "github.com/squirrel/types" in any of:
	/Users/christophreif/DevelopSource/Go/src/github.com/mysheep/squirrel/vendor/github.com/squirrel/types (vendor tree)
	/usr/local/go/src/github.com/squirrel/types (from $GOROOT)
	/Users/christophreif/DevelopSource/Go/src/github.com/squirrel/types (from $GOPATH)