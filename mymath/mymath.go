// mymath is my math package
// with some incredible functions
package mymath

// Sum sums two integer together
//
func Sum (a, b int) int {
	return a + b
}

// Add adds to two integers together
//
func Add (x, y int) int {
	return x + y
}

// Min take two integer and return smaller one
//
func Min (x, y int) int {
	z := y
	if x < y {
		z = x
	}
	return z
}