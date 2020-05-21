# Squirrel

Squirrel is a new language based on root of lisp and many
other ideas connected together. Tagging. Connectable objects.

## Design

	+---------------+
 	|   Evaluator   |   
 	+---------------+
 	        |
 	+---------------+
 	|    Parser     |--- [ Generator ]
 	+---------------+
 			|
 	+---------------+
 	|    Scanner    |   
 	+---------------+
 
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