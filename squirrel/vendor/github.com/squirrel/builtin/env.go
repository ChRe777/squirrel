package builtin


var (
	t		 = "(t t)"
	n		 = "(nil nil)"
	noFn     = "(no     (func (x)   (eq x '())))"
	andFn    = "(and    (func (x y) (cond (x (cond (y 't) ('t '())))('t '()))))"
	notFn    = "(not    (func (x)   (cond (x '()) ('t 't))))"
	appendFn = "(append (func (x y) (cond ((no x) y) ('t (cons (car x) (append (cdr x)  y))))))"
	pairFn   = "(pair   (func (x y) (cond ((and (no x) (no y)) '()) ((and (not (atom x)) (not (atom y))) (cons (list (car x) (car y))(pair (cdr x) (cdr y)))))) )"
	listFn   = "(list   (func (x y) (cons x (cons y '()))))"
	assocFn  = "(assoc  (func (x y) (cond ((eq (caar y) x) (cadar y)) ('t (assoc x (cdr y))))))"
    mapFn    = "(map	(func (f x) (cond ((no x) nil) ('t (cons (f (car x)) (map f (cdr x)))))))"

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
		listFn	, 
		assocFn , 
		mapFn,
	}
	
	return xs
}