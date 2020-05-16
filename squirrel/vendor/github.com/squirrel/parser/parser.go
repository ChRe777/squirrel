package parser

import(
	"github.com/squirrel/types"
	"github.com/squirrel/builtin"
	"github.com/squirrel/generator"
	"github.com/squirrel/scanner"
)

/*
PROCEDURE sexpr(level: INTEGER): cell;
	VAR e: cell;
*/
func sexpr(level *int) *types.Cell {

	/*	
	PROCEDURE atom(level: INTEGER): cell;
	VAR a: cell;
	*/
	atom := func(level *int) *types.Cell {
	
		/*
		PROCEDURE symbol(level: INTEGER): cell; 
		BEGIN incL(level); debug("symbol", level);
			RETURN LSG.atom(LSS.id); (* TYPE = SYM *)
		END symbol;
		*/
		symbol := func(level *int) *types.Cell {
			incLevel(level); debug("symbol", level)
			return generator.Atom(scanner.IdStr(), types.SYMBOL)
		}

		/*
		PROCEDURE string(level: INTEGER): cell; 
		BEGIN incL(level); debug("string", level);
			RETURN LSG.atom("STRING"); (* TYPE = STRING*)
		END string;
		*/
		string := func(level *int) *types.Cell {
			incLevel(level); debug("string", level)
			return generator.Atom(scanner.IdStr(), types.STRING)
		}


		/*		
		PROCEDURE number(level: INTEGER): cell; 
		BEGIN  incL(level); debug("number", level);
			RETURN LSG.atom("NUMBER");
		END number;
		*/
		number := func(level *int) *types.Cell {
			incLevel(level); debug("number", level)
			n, err := generator.NumberFromString(scanner.IdStr())
			if err != nil {
				return generator.Atom(err.Error(), types.STRING) 
			}
			return generator.Atom(n, types.NUMBER)
		}

		/*		
		BEGIN (* atom *) incL(level); debug("atom", level);
						
			CASE LSS.sym OF
				   LSS.symbol: RETURN symbol(level); |			
				   LSS.string: RETURN string(level); |
				   LSS.number : RETURN number(level);
			END;
			
		END atom;
		*/

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
	
	} // atom

	/*	
	PROCEDURE list(level: INTEGER): cell;
		VAR list, e: cell;
	BEGIN incL(level); list := LSG.list(); debug("list", level); 
		
		IF LSS.sym = LSS.lparen THEN LSS.GetSym; ELSE error(1); END;
		
		WHILE LSS.sym < LSS.rparen DO 
				e := sexpr(level); list := LSG.add(list, e); 
				LSS.GetSym;
		END;	
	
		IF LSS.sym = LSS.rparen THEN (*LSS.GetSym;*) ELSE error(2); END;
																															
		RETURN list;
	END list;
	*/
	list := func(level *int) *types.Cell {
	
		incLevel(level); list := generator.Nil(); debug("list", level)
	
		if scanner.Sym == scanner.Lparen {
			debug("list lparen", level)
			scanner.GetSym()
		} else {
			// TODO: return error tuple (nil, error)
			return generator.Error("Left paren is missing")
		}
		
		for ;scanner.Sym < scanner.Rparen; {
			e := sexpr(level); list = builtin.Add(list, e)
			scanner.GetSym()
		}
		
		if scanner.Sym == scanner.Rparen {
			debug("list rparen", level)
		} else {
			// TODO: return error tuple (nil, error)
			return generator.Error("Right paren is missing") 	// Right paren missing
		}
		
		return list
	} // end of list

	/*	
	PROCEDURE quote(level: INTEGER): cell;
	BEGIN debug("quote", level);
		LSS.GetSym; RETURN LSG.quote(sexpr(level));
	END quote;
	*/
	quote := func(level *int) *types.Cell {
		debug("quote", level); scanner.GetSym();
		cell := builtin.Quote(sexpr(level))
		return cell
	}


	/*
	BEGIN (* sexpr *) level := level + 3; debug("sexpr", level);
	
		CASE LSS.sym OF
			(* atoms  *) LSS.symbol: RETURN atom(level); | LSS.number: RETURN atom(level); | LSS.string: RETURN atom(level); | 
			(* list   *) LSS.lparen: RETURN list(level); | 
			(* quotes *) LSS.quote : RETURN quote(level); 
		END;
			
	END sexpr;
	*/
	*level += 3; debug("sexpr", level)
	
	switch scanner.Sym {
	
		case scanner.Symbol:
			return atom(level)

		case scanner.Number: 
			return atom(level)

		case scanner.String:
			return atom(level)

		case scanner.Lparen:
			return list(level)
			
		case scanner.Quote:
			return quote(level)
	}
	
	return generator.Nil()

}

/*
PROCEDURE Parse*(T: Texts.Text): cell;
	BEGIN debug("Parse",level); level := 0;
		LSS.Init(T, 0); LSS.GetSym; RETURN sexpr(level);
	END Parse;
*/
func Parse(b []byte) *types.Cell {

	debug("Parse", &level)
	
	scanner.Init(b)
	scanner.GetSym()
	
	return sexpr(&level)
}