package module01

// Fibonacci returns the nth fibonacci number.
//
// A Fibonacci number N is defined as:
//
//   Fibonacci(N) = Fibonacci(N-1) + Fibonacci(N-2)
//
// Where the following base cases are used:
//
//   Fibonacci(0) => 0
//   Fibonacci(1) => 1
//
//
// Examples:
//
//   Fibonacci(0) => 0
//   Fibonacci(1) => 1
//   Fibonacci(2) => 1
//   Fibonacci(3) => 2
//   Fibonacci(4) => 3
//   Fibonacci(5) => 5
//   Fibonacci(6) => 8
//   Fibonacci(7) => 13
//   Fibonacci(14) => 377
//

// It's better to use iteration, which speeds up the program
// dramatically compared to using recursion.
func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	var fn, nMinus2, nMinus1 int
	nMinus2, nMinus1 = 0, 1

	for i := 2; i <= n; i++ {
		fn = nMinus2 + nMinus1
		nMinus2 = nMinus1
		nMinus1 = fn
	}
	return fn
}

// Using recursion. The method isn't actually efficient,
// especially as n becomes higher, as repeatedly calls back
// to previous numbers in the Fibonacci sequence.
/*
func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}
*/
