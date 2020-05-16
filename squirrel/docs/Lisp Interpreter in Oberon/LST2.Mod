MODULE LST;

	IMPORT
		LSE, LSG, LSO, LSP, Texts, Oberon, Out;
		
	TYPE
		cell = LSG.cell;
		atomCell = LSG.atomCell;
		consCell = LSG.consCell;
		string = ARRAY OF CHAR;
		
	VAR
		env: cell;
		W: Texts.Writer;
		R: Texts.Reader;
		T: Texts.Text;
		
	PROCEDURE ReadEvalPrint*;
	
		PROCEDURE evalPrint(sexpr: cell);
		BEGIN
			LSO.print(sexpr); LSO.print(LSE.eval(sexpr, env));
		END evalPrint;
		
	BEGIN (* ReadEvalPrint *)
		evalPrint(LSP.Parse(T));
	END ReadEvalPrint;


	PROCEDURE tests*;
	
		PROCEDURE read(texpr: ARRAY OF CHAR): cell;
		BEGIN
			Texts.WriteString(W, texpr); NEW(T); Texts.Open(T, "Test"); Texts.Append(T, W.buf);
			RETURN LSP.Parse(T);
		END read;
	
		PROCEDURE test(texpr, tres: ARRAY OF CHAR);
		BEGIN test2(texpr, tres, "()")												
		END test;
		
		PROCEDURE test2(texp, tres, tenv: ARRAY OF CHAR);
			VAR exp, res, evl, env: cell;				
		BEGIN Out.Ln; 	
			
			exp := read(texp);  res := read(tres); env := read(tenv); 		   
			evl := LSE.eval(exp, env); 
			LSO.print(exp); LSO.print(evl); LSO.print(res);
													
		END test2;

		PROCEDURE testQuote;
		BEGIN
			test("'A", "A");
			test("'(A)", "(A)");
			test("'(A B))", "(A B)");
			test("(QUOTE ())", "()"); 
		END testQuote;
		
		PROCEDURE testAtom;
		BEGIN
			test("(ATOM 'A)", "T");
			test("(ATOM '(A B))", "()");
			test("(ATOM (ATOM 'A))", "T");
			test("(ATOM '(ATOM 'A))", "()");
		END testAtom;
		
		PROCEDURE testEq;
		BEGIN
			test("(EQ 'A 'A)", "T");
			test("(EQ 'A 'B)", "()");
			test("(EQ '() '())", "T");
		END testEq;

		PROCEDURE testCar;
		BEGIN
			test("(CAR '(A B))", "A");
		END testCar;
		
		PROCEDURE testCdr;
		BEGIN
			test("(CDR '(A B C))", "(B C)");
		END testCdr;
		
		PROCEDURE testCons;
		BEGIN
			test("(CONS ' A '(B C)))", "(A B C)");
			test("(CONS 'A (CONS 'B (CONS 'C  '()))) )", "(A B C)");
			test("(CAR (CONS 'A '(B C) ))", "A");
			test("(CDR (CONS 'A '(B C) ))", "(B C)");
		END testCons;
		
		PROCEDURE testCond;
		BEGIN
			test("(COND ((EQ 'A 'B) 'FIRST) ('T 'SECOND))", "SECOND");
			test("(COND ((EQ 'A 'B) 'FIRST)  ((ATOM 'A) 'SECOND)  )", "SECOND");
			test("(COND ('T 'FIRST)  ('T 'SECOND))", "FIRST");
		END testCond;
		
		PROCEDURE testLambda;
		BEGIN
			test("((LAMBDA (X)    (CONS X '(B)))   'A )", "(A B)" );  
			test("((LAMBDA (X Y) (CONS X (CDR Y)))  'Z '(A B C))", "(Z B C)");
		END testLambda;
		
		PROCEDURE testTagging;
		BEGIN
			test("(TYPE (TAG 'NUM 'A))", "NUM" );  
			test("(REP  (TAG 'NUM 'A))", "A");
			test("(TYPE  'A)", "SYM");
			test("(TYPE  '(A B))", "CONS");
		END testTagging;
		
		PROCEDURE testAssign;
		BEGIN
			(* test("(ASSIGN A B)", "A"); *)
		END testAssign;
		
		PROCEDURE testEnv;
		
		VAR 
			null: string;
			and: string;
			not: string;
			append: string;
			
		BEGIN
		
			null   = "(NULL   (LAMBDA (X)   (EQ X '())))";
			and    = "(AND    (LAMBDA (X Y) (COND (X (COND (Y 'T) ('T '())))('T '()))))";
			not    = "(NOT    (LAMBDA (X)   (COND (X '()) ('T 'T))))";
			append = "(APPEND (LAMBDA (X Y) (COND ((NULL X) Y) ('T (CONS (CAR X) (APPEND (CDR X) Y))))))";
			pair   = "(PAIR   (LAMBDA (X Y) (COND ((AND (NULL X) (NULL Y)) '()) ((AND (NOT (ATOM X)) (NOT (ATOM Y))) (CONS (LIST (CAR X) (CAR Y))(PAIR (CDR X) (CDR Y)))))))";
		
			test2("(NULL '())",  "T",  "(" + null + ")");
			
			test2("(AND 'T 'T)", "T",  "(" + and  + ")"); 
			
			test2("(NOT 'T)",    "()", "(" + not  + ")");
			
			(* Uses NULL *)
			test2("(APPEND '(A B) '(C D))", "(A B C D)", "(" + null + append + ")");
			
			(* Uses NOT, NULL, AND, LIST *)		
			test2("(PAIR '(A B C) '(X Y Z))", "((A X)(B Y)(C Z))", "(" + not + null + and + pair + ")");
			
		END testEnv;
		
		(*
			SOME FUNCTIONS FOR EVAL 
		
			(LIST e1 e2 ... en) = (CONS e1 (CONS e2 ... (CONS en '()))
			
			CAAR, CADR, CADAR, CADDR, CADDAR, ....
		
			1.
			
			(NULL
				(LAMBDA (X)
					(EQ X '())
				)
			)
			
			2.
			
			(AND
				(LAMBDA (X Y)
					(COND 
						(X (COND (Y 'T) ('T '())))
						('T '())
					)
				)
			)
		
			3.
		
			(NOT 
				(LAMBDA(X) 
					(COND (X '()) 
								('T 'T)
					) 
				) 
			)
			
			4.
			
			Uses NULL
			
			(APPEND 
				(LAMBDA (X Y)
					(COND ((NULL X) Y)
								('T (CONS (CAR X) (APPEND (CDR X)  Y)))
					)
				)
			)
			
			5.
			
			Uses NOT, NULL, AND, LIST
			
			(PAIR 
				(LAMBDA (X Y)
					(COND ((AND (NULL X) (NULL Y)) '())
								 ((AND (NOT (ATOM X)) (NOT (ATOM Y)))
										(CONS (LIST (CAR X) (CAR Y))
											(PAIR (CDR X) (CDR Y)) 
								 )
					)			
				)		
			)
			
			6.
			
			Uses CAAR, CADAR
			
			(ASSOC 
				(LAMBDA (X Y)
					(COND 
						( (EQ (CAAR Y) X) (CADAR Y) )
						( 'T (ASSOC X (CDR Y)) )
					)
				)
			)
			
			
			7.
			
			.....
			
		*)
		
		

		PROCEDURE testString;
		BEGIN
			(* test("(CAR "foo" "bar")"); *) 
		END testString;
		
		PROCEDURE testNum;
		BEGIN
			(* TODO *)
		END testNum;
		
		(*
		
			Ideas:
			
			mac, macex - code template, template expand
			null -> empty
		
		*)		

	BEGIN (* tests *)
	
		testQuote;
		testAtom;
		
		testEq;
		
		testCar;
		testCdr;
		
		testCons;
		
		testCond;
		testLambda;	
		
		testTagging;
		testEnv;
		
	END tests;

BEGIN

	Texts.OpenWriter(W);

END LST.

** TOOLS **

System.Free LST LSE LSP LSO LSG LSS ~   |   Builder.Compile *
System.Free LST ~
LST.tests