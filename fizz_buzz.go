package module01

import (
	"fmt"
)

// FizzBuzz will print out all of the numbers
// from 1 to N replacing any divisible by 3
// with "Fizz", and divisible by 5 with "Buzz",
// and any divisible by both with "Fizz Buzz".
//
// Note: The test for this is a little
// complicated so that you can just use the
// `fmt` package and print to standard out.
// I wouldn't normally recommend this, but did
// it here to make life easier for beginners.

// Technically correct, but test case doesn't like the format....
/*
func FizzBuzz(n int) {
	for i := 1; i <= n; i++ {
		switch {
		case i%15 == 0:
			fmt.Printf("%d:\t%s\n", i, "Fizz Buzz")
		case i%5 == 0:
			fmt.Printf("%d:\t%s\n", i, "Buzz")
		case i%3 == 0:
			fmt.Printf("%d:\t%s\n", i, "Fizz")
		default:
			fmt.Printf("%d\n", i)
		}
	}
}
*/

// Because the test case if extremely picky, we make sure
// to give proper formatting and add a Println()
func FizzBuzz(n int) {
	for i := 1; i <= n; i++ {
		switch {
		case i%15 == 0:
			fmt.Print("Fizz Buzz")
		case i%5 == 0:
			fmt.Print("Buzz")
		case i%3 == 0:
			fmt.Print("Fizz")
		default:
			fmt.Print(i)
		}
		if i != n {
			fmt.Print(", ")
		}
	}
	fmt.Println()
}
