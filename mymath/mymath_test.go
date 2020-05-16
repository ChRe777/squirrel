package mymath

import (
	"testing"
)

// Write test FIRST !!!

func TestSum(t *testing.T) {
	sum := Sum(1, 2)
	if sum != 3 {
		t.Errorf("got: %d, want: %d", sum, 3)
	}
}

func TestAdd(t *testing.T) {
	add := Add(1,2)
	if add != 3 {
		t.Errorf("got: %d, want: %d", add, 3)
	}
}

func TestMin(t *testing.T) {
	min := Min(1,2)
	if min != 1 {
		t.Errorf("got: %d, want: %d", min, 1)
	}
}
