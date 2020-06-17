package builtin

/*
var (
	t		 = "(t t)"
	n		 = "(nil nil)"
	noFn     = "(no     (func (x)    (is x '())))"
	andFn    = "(and    (func (x y)  (cond (x (cond (y 't) ('t '())))('t '()))))"
	notFn    = "(not    (func (x)    (cond (x '()) ('t 't))))"
	appendFn = "(append (func (x y)  (cond ((no x) y) ('t (cons (car x) (append (cdr x)  y))))))"
	pairFn   = "(pair   (func (x y)  (cond ((and (no x) (no y)) '()) ((and (not (atom x)) (not (atom y))) (cons (list (car x) (car y))(pair (cdr x) (cdr y)))))) )"
	
	// TODO: Error: Not Found instead of nil - nil can be a valid value
	//
	assocFn  = "(assoc  (func (x ys) (cond ((no ys) nil) ((is x (car (car ys))) (car (cdr (car ys)))) ('t (assoc x (cdr ys))))))"    
    mapFn    = "(map	(func (f x)  (cond ((no x) nil) ('t (cons (f (car x)) (map f (cdr x)))))))"

)

func Env() []string {

	xs := []string{ 
		t		,
		n		,
		noFn	,
	    andFn	,     
		notFn	,    
		appendFn, 
		pairFn	,   
//		listFn	, 
		assocFn , 
		mapFn,
	}
	
	return xs
}
*/