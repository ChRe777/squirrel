package builtin

import (
	"github.com/squirrel/types"
	"github.com/squirrel/generator"
)

/*
PROCEDURE quote*(c: cell): cell;
BEGIN
	RETURN cons(atom("QUOTE"), cons(c, EMPTY));
END quote;
*/
func Quote(c *types.Cell) *types.Cell {
	return generator.Cons(QUOTE, generator.Cons(c, NIL))
}

/*
PROCEDURE add*(l, c: cell): cell;
	VAR li: cell;
BEGIN 
	li := l;
	IF l # EMPTY THEN
		WHILE l(consCell).cdr # EMPTY DO l := l(consCell).cdr END;
		l(consCell).cdr := cons(c, EMPTY);
	ELSE
		li := cons(c, EMPTY);
	END;
	RETURN li;
END add;
*/
func Add(l, c *types.Cell)  *types.Cell {
	li := l
	if l.IsCons() && l.NotEqual(NIL) { 
	
		// TODO: Speed Up With Pointer on LAST element
	
		for ;l.Cdr.NotEqual(NIL); {
			l = l.Cdr
		}
	
		l.Cdr = generator.Cons(c, NIL)
	
	} else {
		li = generator.Cons(c, NIL)
	}
	return li
}

/*
PROCEDURE list*(): cell;
BEGIN RETURN EMPTY;
END list;
*/
func List(xs ...*types.Cell) *types.Cell {
	l := NIL
	for _, x := range xs {
		l = Add(l, x)
	}
	return l
}

/*
PROCEDURE cdr*(c: cell): cell;
BEGIN
	IF c IS consCell THEN RETURN c(consCell).cdr END;
END cdr;
*/
func Cdr(c *types.Cell) *types.Cell {
	if c.IsCons() {
		return c.Cdr
	}
	return NIL
}

/*
PROCEDURE car*(c: cell): cell;
BEGIN
	IF c IS consCell THEN RETURN c(consCell).car END;
END car;
*/
func Car (c *types.Cell) *types.Cell {
	if c.IsCons() {
		return c.Car
	}
	return NIL
}

// Sym create a symbol from string
func Sym(s string) *types.Cell {
	return generator.Sym(s)
}

// Num create a number from string
func Num(s string) *types.Cell {
	return generator.Num(s)
}

// Str create a string from string
func Str(s string) *types.Cell {
	return generator.Str(s)
}