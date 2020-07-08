package functions

import (
	"testing"
)

import (
	"github.com/shopspring/decimal"
)

import (
	"github.com/mysheep/squirrel/generator"
	"github.com/mysheep/squirrel/evaluator"
	"github.com/mysheep/squirrel/types"
)


func TestAddList(t *testing.T) {

	num := func(i int64) *types.Cell {
		return generator.Num_(decimal.NewFromInt(i))
	}
		
	sym := func(s string) *types.Cell {
		return generator.Sym(s)
	}
	
	err := func(s string) *types.Cell {
		return generator.Err(s)
	}
	
	str := func(s string) *types.Cell {
		return generator.Str(s)
	}
	
	env := generator.Nil()
	
	xs1 := generator.Cons(num(1)  , generator.Cons(num(2), generator.Cons(num(3), generator.Nil())))
	xs2 := generator.Cons(sym("a"), generator.Cons(num(2), generator.Cons(num(3), generator.Nil())))
	xs3 := generator.Cons(str("1"), generator.Cons(num(2), generator.Cons(num(3), generator.Nil())))
		
	specs := []struct {
		xs		*types.Cell
		env		*types.Cell
		eval    func(*types.Cell, *types.Cell) *types.Cell
		want	*types.Cell
	} {
		{ xs1, env, evaluator.Eval, num(6) },
		{ xs2, env, evaluator.Eval, err("reference to undefined identifier: a") },
		{ xs3, env, evaluator.Eval, err(MUST_BE_NUMBER_TYPES) },
	}

	for _, spec := range specs {
		
		got := AddList(spec.xs, spec.env, spec.eval)

		if got.NotEqual(spec.want) {
			t.Errorf("TestAdd - got: %v, want: %v", got, spec.want)
		}
	
	}
}


func TestSubList(t *testing.T) {

	num := func(i int64) *types.Cell {
		return generator.Num_(decimal.NewFromInt(i))
	}
		
	sym := func(s string) *types.Cell {
		return generator.Sym(s)
	}
	
	err := func(s string) *types.Cell {
		return generator.Err(s)
	}
	
	str := func(s string) *types.Cell {
		return generator.Str(s)
	}
	
	env := generator.Nil()
	
	xs1 := generator.Cons(num(5)  , generator.Cons(num(2), generator.Cons(num(3), generator.Nil())))
	xs2 := generator.Cons(sym("a"), generator.Cons(num(2), generator.Cons(num(3), generator.Nil())))
	xs3 := generator.Cons(str("1"), generator.Cons(num(2), generator.Cons(num(3), generator.Nil())))
		
	specs := []struct {
		xs		*types.Cell
		env		*types.Cell
		eval    func(*types.Cell, *types.Cell) *types.Cell
		want	*types.Cell
	} {
		{ xs1, env, evaluator.Eval, num(0) },
		{ xs2, env, evaluator.Eval, err("reference to undefined identifier: a") },
		{ xs3, env, evaluator.Eval, err(MUST_BE_NUMBER_TYPES) },
	}

	for _, spec := range specs {
		
		got := SubList(spec.xs, spec.env, spec.eval)

		if got.NotEqual(spec.want) {
			t.Errorf("TestSubList - got: %v, want: %v", got, spec.want)
		}
	
	}
}

func TestAdd(t *testing.T) {

	num := func(i int64) *types.Cell {
		return generator.Num_(decimal.NewFromInt(i))
	}
	
	numS := func(s string) *types.Cell {
		return generator.Num(s)
	}
	
	sym := func(s string) *types.Cell {
		return generator.Sym(s)
	}
	
	err := func(s string) *types.Cell {
		return generator.Err(s)
	}
	
	str := func(s string) *types.Cell {
		return generator.Str(s)
	}
	

	specs := []struct {
		x		*types.Cell
		y 		*types.Cell
		want	*types.Cell
	} {
		{ num(1)  			, num(2)			, num(3) 						},
		{ sym("1")			, num(2)			, err(MUST_BE_NUMBER_TYPES) 	},
		{ sym("1")			, str("2")			, err(MUST_BE_NUMBER_TYPES) 	},
		{ numS("1.1")  		, numS("2.2")		, numS("3.3") 					},
		{ numS("-1.1")  	, numS("2.2")		, numS("1.1") 					},
		{ numS("-1.1e7")  	, numS("2.2e7")		, numS("1.1e7") 				},
	}

	for _, spec := range specs {
	
		got := add(spec.x, spec.y)

		if got.NotEqual(spec.want) {
			t.Errorf("TestAddList - got: %v, want: %v", got, spec.want)
		}
	
	}
}

func TestSub(t *testing.T) {

	num := func(i int64) *types.Cell {
		return generator.Num_(decimal.NewFromInt(i))
	}
	
	numS := func(s string) *types.Cell {
		return generator.Num(s)
	}
	
	sym := func(s string) *types.Cell {
		return generator.Sym(s)
	}
	
	err := func(s string) *types.Cell {
		return generator.Err(s)
	}
	
	str := func(s string) *types.Cell {
		return generator.Str(s)
	}
	
	specs := []struct {
		x		*types.Cell
		y 		*types.Cell
		want	*types.Cell
	} {
		{ num(1)  			, num(1)			, num(0) 						},
		{ num(2)  			, num(1)			, num(1) 						},
		{ sym("1")			, num(2)			, err(MUST_BE_NUMBER_TYPES) 	},
		{ sym("1")			, str("2")			, err(MUST_BE_NUMBER_TYPES) 	},
		{ numS("3.3")  		, numS("2.2")		, numS("1.1") 					},
		{ numS("-3.3")  	, numS("-2.2")		, numS("-1.1") 					},
		{ numS("-1.1e7")  	, numS("2.2e7")		, numS("-3.3e7") 				},
	}

	for _, spec := range specs {
	
		got := sub(spec.x, spec.y)

		if got.NotEqual(spec.want) {
			t.Errorf("TestAdd - got: %v, want: %v", got, spec.want)
		}
	
	}
}