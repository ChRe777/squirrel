package core

import (
	"errors"
)

import (
	"github.com/mysheep/squirrel/types"
)

// -------------------------------------------------------------------------------------------------

var (
	OperatorNotFound = errors.New("operator not found")
)

// -------------------------------------------------------------------------------------------------

func Eval(e, a *types.Cell, eval func(*types.Cell, *types.Cell) *types.Cell) (*types.Cell, error)  {

	if c := Car(e); c.IsAtom() {
		if fn, found := coreFuncMap[*c]; found {
			return fn(e, a, eval), nil
		}
	}
	
	return nil, OperatorNotFound
}

// -------------------------------------------------------------------------------------------------

type FuncType 	 = func(*types.Cell, *types.Cell) *types.Cell
type MapFuncType = func(*types.Cell, *types.Cell, FuncType) *types.Cell

var coreFuncMap = map[types.Cell] MapFuncType {

	*QUOTE	: Quote__ ,
	*ATOM	: Atom__  ,
	*IS		: Is__    ,
	*CAR	: Car__   ,
	*CDR	: Cdr__   ,
	*CONS	: Cons__  ,
    *COND   : Cond__  ,
	
	*CAAR   : Caar__  , 
	*CADR   : Cadr__  ,
	*CDDR   : Cddr__  ,
	*CADAR  : Cadar__ ,
	*CDDDR  : Cdddr__ ,
	*CADDR  : Caddr__ ,
	*CADDAR : Caddar__,

	*TYPE	: Type__  ,	
	*SYM 	: Sym__	 ,
	
//	FUNC 		= Sym_(ID_FUNC)
//	TAG 		= Sym_(ID_TAG)
//	TAGGED 		= Sym_(ID_TAGGED)

//	TYPE0 		= Sym_(ID_TYPE0)							
//	REP 		= Sym_(ID_REP)

}

// -------------------------------------------------------------------------------------------------

func Quote__(exp, env *types.Cell, eval FuncType) *types.Cell {
	return Quote(exp)
}

func Atom__(exp, env *types.Cell, eval FuncType) *types.Cell {
	x := eval(Cadr(exp), env)
	return Atom(x)
}

func Is__(exp, env *types.Cell, eval FuncType) *types.Cell {
	x := eval(Cadr(exp), env)
	y := eval(Caddr(exp), env)	
	return Is(x, y)
}

func Car__(exp, env *types.Cell, eval FuncType) *types.Cell {
	x := eval(Cadr(exp), env)
	return Car(x)
}

func Cdr__(exp, env *types.Cell, eval FuncType) *types.Cell {
	x := eval(Cadr(exp), env)
	return Cdr(x)
}

func Cons__(exp, env *types.Cell, eval FuncType) *types.Cell {
	x := eval(Cadr(exp), env)
	y := eval(Caddr(exp), env)
	return Cons(x, y)
}

func Cond__(exp, env *types.Cell, eval FuncType) *types.Cell {
	cs := Cdr(exp)
	return Cond(cs, env, eval)
}

func Type__(exp, env *types.Cell, eval FuncType) *types.Cell {
	x := eval(Cadr(exp), env)
	return Type(x)
}

func Sym__(exp, env *types.Cell, eval FuncType) *types.Cell {
	x := eval(Cadr(exp), env)
	
	s, err := x.AsStr()	
	if err == nil {
		return Sym_(s)
	} else {
		return Err_(err.Error())
	}
	
}

// -------------------------------------------------------------------------------------------------

func Caar__  (exp, env *types.Cell, eval FuncType) *types.Cell { x := eval(Cadr(exp), env); return Caar(x)   }
func Cadr__  (exp, env *types.Cell, eval FuncType) *types.Cell { x := eval(Cadr(exp), env); return Cadr(x)   }
func Cddr__  (exp, env *types.Cell, eval FuncType) *types.Cell { x := eval(Cadr(exp), env); return Cddr(x)   }
func Cadar__ (exp, env *types.Cell, eval FuncType) *types.Cell { x := eval(Cadr(exp), env); return Cadar(x)  } 
func Cdddr__ (exp, env *types.Cell, eval FuncType) *types.Cell { x := eval(Cadr(exp), env); return Cdddr(x)  } 
func Caddr__ (exp, env *types.Cell, eval FuncType) *types.Cell { x := eval(Cadr(exp), env); return Caddr(x)  }
func Caddar__(exp, env *types.Cell, eval FuncType) *types.Cell { x := eval(Cadr(exp), env); return Caddar(x) }

// -------------------------------------------------------------------------------------------------

