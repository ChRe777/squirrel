# Bugs


	> (when (is 'a 'a 'b))  -> endless loop
	
	>>> (mac if (c a b) `(cond (,c ,a) ('t ,b)))
mac
>>> (var name "foo")
"foo"
>>> (if (is name "foo") (list "name is" name) (list "name is not" name))
evalFuncCall - res :%v 
 (cond (t ("name is" "foo")) ((quote t) ("name is not" "foo")))
Error: "reference to undefined identifier: name is"