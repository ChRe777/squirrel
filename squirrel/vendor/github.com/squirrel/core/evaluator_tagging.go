package core

import (
	//"fmt"
)

import (
	//"github.com/squirrel/types"
)


// Paul Graham add 2-3 new primitive core operators in Arc
// to make macros much simpler and other things
// A normal function is tagged with 'mac

/*
func isTagged(e *types.Cell) *types.Cell {

	if e.IsCons() {
		x := car(e)
		if x.IsCons() && 
		    caar(e).Equal(core.QUOTE) && 
		   cadar(e).Equal(core.TAGGED) {
			return core.T
		} else {
			return core.NIL
		}
	} else {
		return core.NIL
	}
	
}

func tag(t, r *types.Cell) *types.Cell {
	tc := cons(core.QUOTE, cons(core.TAGGED, core.NIL))
	return cons(tc, cons(t, cons(r, core.NIL)))
}

func type_(e *types.Cell) *types.Cell {
	if e.IsCons() {
		return core.CONS
	} else {
		// TODO: ???
		return core.Sym(fmt.Sprintf("%v", e.Type.Atom))
	}
}

func type0(e *types.Cell) *types.Cell {

	tagType := func (e *types.Cell) *types.Cell {
		return cadr(e)
	}
	
	if isTagged(e).Equal(core.T) {
		return tagType(e)
	} else {
		return type_(e)
	}
}

func rep(e *types.Cell) *types.Cell {
	if isTagged(e).Equal(core.T) {
		return caddr(e)
	} else {
		return cadr(e)
	}
}
*/