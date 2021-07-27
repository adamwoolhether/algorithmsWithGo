package module01

import "fmt"

// BaseToDec takes in a number and the base it is currently
// in and returns the decimal equivalent as an integer.
//
// Eg:
//
//   BaseToDec("E", 16) => 14
//   BaseToDec("1110", 2) => 14
//
/*
func BaseToDec(value string, base int) int {
	var charset = "0123456789ABCDEF"
	multiplier := 1
	res := 0

	for i := len(value) - 1; i >= 0; i-- {
		index := -1
		for k, char := range charset {
			if char == rune(value[i]) {
				index = k
				break
			}
		}
		if index < 0 {
			panic("uh oh")
		}
		res += index*multiplier
		multiplier *= base
	}

	return res
}
*/

// This one is much better:
func BaseToDec(value string, base int) int {
	res := 0
	multiplier := 1

	for i := len(value) - 1; i >= 0; i-- {
		var val int
		fmt.Sscanf(string(value[i]), "%X", &val)
		res += multiplier * val
		multiplier *= base
	}
	return res
}
