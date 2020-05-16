package spec

import (
	"bytes"
	"testing"
)

import (
	"github.com/squirrel/parser"
)

/*
PROCEDURE testEnv;

VAR
	nullEx: ARRAY 256 OF CHAR;
	andEx: ARRAY 256 OF CHAR;
	notEx: ARRAY 256 OF CHAR;
	appendEx: ARRAY 256 OF CHAR;
	pairEx: ARRAY 256 OF CHAR;
	listEx: ARRAY 256 OF CHAR;
	envEx: ARRAY 2048 OF CHAR;

BEGIN

	nullEx := "(NULL (LAMBDA (X)   (EQ X '())))";
	andEx  := "(AND  (LAMBDA (X Y) (COND (X (COND (Y 'T) ('T '())))('T '()))))";
	nullEx := "(NULL (LAMBDA (X)   (EQ X '())))";
	notEx  := "(NOT  (LAMBDA (X)   (COND (X '()) ('T 'T))))";
	appendEx := "(APPEND (LAMBDA (X Y) (COND ((NULL X) Y) ('T (CONS (CAR X) (APPEND (CDR X)  Y))))))";
	pairEx := "(PAIR (LAMBDA (X Y) (COND ((AND (NULL X) (NULL Y)) '()) ((AND (NOT (ATOM X)) (NOT (ATOM Y))) (CONS (LIST (CAR X) (CAR Y))(PAIR (CDR X) (CDR Y)))))) )";
	listEx := "(LIST (LAMBDA(X Y)  (CONS X (CONS Y '()))))";
	
	Strings.Append(envEx, "(");
	Strings.Append(envEx, nullEx);
	Strings.Append(envEx, andEx);
	Strings.Append(envEx, notEx);
	Strings.Append(envEx, appendEx);
	Strings.Append(envEx, pairEx);
	Strings.Append(envEx, listEx);
	Strings.Append(envEx, ")");

	test2("(NULL '())", "T", envEx);
	test2("(AND 'T 'T)", "T", envEx); 
	test2("(NOT 'T)","()", envEx);
	test2("(APPEND '(A B) '(C D))", "(A B C D)", envEx);
	(*Uses NOT, NULL, AND, LIST*)
	test2("(PAIR '(A B C) '(X Y Z))", "((A X)(B Y)(C Z))", envEx);
	
END testEnv;
*/

func TestEnvironment(t *testing.T) {

	nullEx   := "(null   (func (x)   (eq x '())))"
	andEx    := "(and    (func (x y) (cond (x (cond (y 't) ('t '())))('t '()))))"
	notEx    := "(not    (func (x)   (cond (x '()) ('t 't))))"
	appendEx := "(append (func (x y) (cond ((null x) y) ('t (cons (car x) (append (cdr x)  y))))))"
	pairEx   := "(pair   (func (x y) (cond ((and (null x) (null y)) '()) ((and (not (atom x)) (not (atom y))) (cons (list (car x) (car y))(pair (cdr x) (cdr y)))))) )"
	listEx   := "(list   (func (x y) (cons x (cons y '()))))"

	fns := []string{ 
		nullEx	,
	    andEx	,     
		notEx	,    
		appendEx, 
		pairEx	,   
		listEx	,   
	}
	
	var b bytes.Buffer
	b.WriteRune('(')
	for _, fn := range fns {
		b.WriteString(fn)
	}
	b.WriteRune(')')
	
	env := parser.Parse(b.Bytes())
	
	// End create environment
	
	specs := []spec {
		{ "(null '())				"	, "t"			},
		{ "(and 't 't)				"	, "t"			}, 
		{ "(not 't)				 	"	, "nil"			},
		{ "(append '(a b) '(c d))	"	, "(a b c d)"	},
		//
		// uses not, null, and, list
		//
		{ "(pair '(a b c) '(x y z))	"	, "((a x) (b y) (c z))"},
	}
		
	testWithEnv(specs, t, env)
	
}