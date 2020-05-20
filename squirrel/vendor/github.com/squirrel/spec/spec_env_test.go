package spec

import (
	"bytes"
	"testing"
)

import (
	"github.com/squirrel/types"
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

	env := createEnvironment()
	
	specs := []spec {
		{ "(no '())					   "	, "t"			},
		{ "(and  't 't)				    "	, "t"			}, 
		{ "(not  't)				    "	, "nil"			},
		{ "(append '(a b)   '(c d)  )	"	, "(a b c d)"	},
		{ "(pair   '(a b c) '(x y z))	"	, "((a x) (b y) (c z))"},
	}
		
	testWithEnv(specs, t, env)
	
}

func createEnvironment() *types.Cell {
	
	createList := func(fns []string) []byte {
		var b bytes.Buffer
		b.WriteRune('('); 
			for _, fn := range fns { b.WriteString(fn) }
		b.WriteRune(')')
		return b.Bytes()
	}

	// Some builtin fns already build inside
	// the language itself - self booting
	noFn     := "(no     (func (x)   (eq x '())))"
	andFn    := "(and    (func (x y) (cond (x (cond (y 't) ('t '())))('t '()))))"
	notFn    := "(not    (func (x)   (cond (x '()) ('t 't))))"
	appendFn := "(append (func (x y) (cond ((no x) y) ('t (cons (car x) (append (cdr x)  y))))))"
	pairFn   := "(pair   (func (x y) (cond ((and (no x) (no y)) '()) ((and (not (atom x)) (not (atom y))) (cons (list (car x) (car y))(pair (cdr x) (cdr y)))))))"
	listFn   := "(list   (func (x y) (cons x (cons y '()))))"
	assocFn  := "(assoc  (func (x y) (cond ((eq (caar y) x) (cadar y)) ('t (assoc. x (cdr y))))))"
        
	xs := []string{ 
		noFn	,
	    andFn	,     
		notFn	,    
		appendFn, 
		pairFn	,   
		listFn	, 
		assocFn , 
	}
	
	env := parser.Parse(createList(xs))
	return env
}