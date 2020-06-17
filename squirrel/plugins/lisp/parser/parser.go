package parser

import (
	//"fmt"
)

import (
	"github.com/mysheep/squirrel/types"
	"github.com/mysheep/squirrel/generator"
	"github.com/mysheep/squirrel/plugins/lisp/scanner"
)

func sexpr(level *int) *types.Cell {

	atom := func(level *int) *types.Cell {
	
		symbol := func(level *int) *types.Cell {
			incLevel(level); debug("symbol", level)
			return generator.Atom(scanner.IdStr(), types.SYMBOL)
		}

		string := func(level *int) *types.Cell {
			incLevel(level); debug("string", level)
			return generator.Atom(scanner.IdStr(), types.STRING)
		}

		number := func(level *int) *types.Cell {
			incLevel(level); debug("number", level)
			return generator.Num(scanner.IdStr())
		}


		incLevel(level); debug("atom", level)
		
		switch {
			case scanner.Sym == scanner.Symbol:
				return symbol(level);
			
			case scanner.Sym == scanner.String:
				return string(level);
				
			case scanner.Sym == scanner.Number:
				return number(level);
		}
		
		return generator.Nil()
	
	} // end of atom

	list := func(level *int) *types.Cell {
	
		incLevel(level); list := generator.Nil(); last := generator.Nil(); debug("list", level)
	
		if scanner.Sym == scanner.Lparen {
			debug("list lparen", level)
			scanner.GetSym()
		} else {
			return generator.Error("Left paren is missing")
		}
		
		symWasDot := false;
		
		for ;scanner.Sym < scanner.Rparen; {
			
			e := sexpr(level); 
			
			// (a b . c)
			if symWasDot {
				last.Cdr = e; scanner.GetSym()
				break;				
			} else {
				list, last = Push(list, e, last)
			}
			
			scanner.GetSym()

			if scanner.Sym == scanner.Dot {
				symWasDot = true;
				scanner.GetSym()
			}
		}
		
		if scanner.Sym == scanner.Rparen {
			debug("list rparen", level)
		} else {
			return generator.Error("Right paren is missing") 	// Right paren missing
		}
		
		return list
	} // end of list

	quote := func(level *int) *types.Cell {
		debug("quote", level); scanner.GetSym();
		cell := generator.Quote_(sexpr(level))
		return cell
	}
	
	backquote := func(level *int) *types.Cell {
		debug("backquote", level); scanner.GetSym();
		cell := generator.Backquote_(sexpr(level))
		return cell
	}
	
	unquote := func(level *int) *types.Cell {
		debug("unquote", level); scanner.GetSym();
		cell := generator.Unquote_(sexpr(level))
		return cell
	}
	
	unquoteSplicing := func(level *int) *types.Cell {
		debug("unquoteSplicing", level); scanner.GetSym();
		cell := generator.UnquoteSplicing_(sexpr(level))
		return cell
	}

/*
	dot := func(level *int) *types.Cell {
		debug("dot", level); scanner.GetSym();
		cell := core.UnquoteSplicing_(sexpr(level))
		return cell
	}
*/
	
	*level += 3; debug("sexpr", level)
	
	switch scanner.Sym {
	
		case scanner.Symbol: 	return atom(level)

		case scanner.Number: 	return atom(level)

		case scanner.String: 	return atom(level)

		case scanner.Lparen: 	return list(level)
			
		case scanner.Quote: 	return quote(level)
			
		case scanner.Backquote:	return backquote(level)
			
		case scanner.Unquote:	return unquote(level)
			
		case scanner.UnquoteSplicing:	return unquoteSplicing(level)
		
		//case scanner.Dot:		return dot(level)
	}
	
	return generator.Nil()

}

func Parse(b []byte) *types.Cell {

	debug("Parse", &level)
	
	scanner.Init(b); scanner.GetSym()
	
	return sexpr(&level)
}