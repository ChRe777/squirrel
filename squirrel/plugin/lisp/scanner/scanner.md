# Scanner

Scanner is scanning the textfile for symbols

## EBNF of Lisp

1 sexpr = atom | list | quoted-sexpr  /* | backquoted-sexpt | backquoted-slicing */	
	1.1 atom = symbol | string | number |
	1.2 list = "("{sexpr}")"
	1.3 quoted-sexpr = "'" sexpr.
	1.4 backquoted-sexpr = "," sexpr

## Idea

Can I use "go" scanner?
see https://golang.org/pkg/text/scanner/

