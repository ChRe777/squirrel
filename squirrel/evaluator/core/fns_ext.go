package core

import (
	"github.com/mysheep/squirrel/types"
)


// -------------------------------------------------------------------------------------------------
// Shortcuts for car, cdr functions
// -------------------------------------------------------------------------------------------------

func caar  (e *types.Cell) *types.Cell { return car_(car_(e))             }
func cadr  (e *types.Cell) *types.Cell { return car_(cdr_(e))             }
func cddr  (e *types.Cell) *types.Cell { return cdr_(cdr_(e))             }
func cadar (e *types.Cell) *types.Cell { return car_(cdr_(car_(e)))       } 
func cdddr (e *types.Cell) *types.Cell { return cdr_(cdr_(cdr_(e)))       } 
func caddr (e *types.Cell) *types.Cell { return car_(cdr_(cdr_(e)))       }
func caddar(e *types.Cell) *types.Cell { return car_(cdr_(cdr_(car_(e)))) }
func cadddr(e *types.Cell) *types.Cell { return car_(cdr_(cdr_(cdr_(e)))) } 	

// -------------------------------------------------------------------------------------------------
// Helpers
// -------------------------------------------------------------------------------------------------

func car_(c *types.Cell) *types.Cell {
	if c.IsCons() {
		return c.Car
	}
	return NIL
}

func cdr_(c *types.Cell) *types.Cell {
	if c.IsCons() {
		return c.Cdr
	}
	return NIL
}

// -------------------------------------------------------------------------------------------------

