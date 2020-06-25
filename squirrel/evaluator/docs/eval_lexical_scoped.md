# Lexical Scoping - The Art of Interpretor

(DEFINE (EVAL EXP ENV 

	(COND
		
		((ATOM EXP) 
			(COND
				((NUMERP EXP) EXP)	-- NUMBER EVAL TO THEMSELVES
				(T (VALUE EXP ENV))	-- GET VALUE FROM ENV
			)
		)
		
		((EQ (CAR EXP) 'QUOTE)
			 (CADR EXP)
		)
		
		((EQ (CAR EXP) 'LAMBDA)		-- (LAMBDA (X) (* X L))
		     --                 
			 --  	(&PROCEDURE (X)        (NO X L)    ((a 1) (b 2))	)
			 --
			 (LIST '&PROCEDURE  (CADR EXP) (CADDR EXP) ENV             	)	!!! <-- KEY
		)
		
		((EQ (CAR EXP) 'COND)
			 (EVCOND (CDR EXP) ENV))
		)
		 
		(T (APPLY (EVAL  (CAR EXP) ENV)
				  (EVLIS (CDR EXP) ENV)
			)
		
		)
	) -- COND

) -- DEFINE

(DEFINE (APPLY FUN ARGS)
	(COND 
	
		-- 1. --
		--------------------------------------------------------------------------------------------
		((PRIMOP FUN) (PRIMOP-APPLY FUN ARGS)
		
		-- 2. --							PROC-TAG	 PARAMS     BODY         ENV
		--------------------------------------------------------------------------------------------
		
		((EQ (CAR FUN) '&PROCEDURE)		-- (&PROCEDURE   (X)        (* X 2)     ((A 1) (B 2))	)
		
			(EVAL (CADDR FUN) 			-- (* X 2)
			
				  --    PARAMS	    ARGS  ENV
				  ----------------------------------------------------------------------------------
				  --    (X)         ARGS  ((A 1) (B 2))
				  
				  (BIND (CADR FUN)  ARGS  (CADDDR FUN))	
			)
		)
	
		-- 3. --
		--------------------------------------------------------------------------------------------
		(T (ERROR))
		
	)
)

(DEFINE (BIND VARS ARGS ENV)

	-- 1. --
	(COND ( (= 
				(LENGTH VARS) 
				(LENGTH ARGS)
			)
			(CONS (CONS VARS ARGS) ENV)	--(	((x y) 1 2)  .. env ..)
		  )
		  
		  -- 2. --
		  (T (ERROR))
	)
)

A symbol table has a list of buckets, a bucket is ((x y z) 1 2 3)
a list where car is the list of variable names and the cdr is
a list of corresponding values

e.g. a symbol table 

(
	((x y z) 1 2 3)		-- a bucket
	((a b)   5 6  )     -- another bucket
	...
	...
	
)

-- interface of value
(DEFINE (VALUE NAME ENV)
	(VALUE1 NAME (LOOKUP NAME ENV))
)

(DEFINE (VALUE1 NAME SLOT)									-- ((x y) 1 2)
	(COND
		((EQ SLOT '&UNBOUND) 	(ERROR))	
		(T 						(CAR SLOT))					-- (x y)	
	)
)

-- Interface of lookup
(DEFINE (LOOKUP NAME ENV)
	(COND
		((NULL VARS) 	(LOOKUP  NAME (CDR ENV))
		(T 				(LOOKUP1 NAME (CAAR ENV) (CADR ENV) ENV))
	)
)

(DEFINE (LOOKUP1 NAME VARS VALS ENV)
	(COND
		((NULL VARS) 			(LOOKUP NAME (CDR ENV))
		((EQ NAME (CAR VARS)) 	VALS)
		(T 						(LOOKUP1 NAME (CDR VARS) (CDR VALS) ENV)))
	)
)