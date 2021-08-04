package module01

// First instruction was to do so with a simple list of the first prime numbers:
// My solution
/*
var primes = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}
func GCD(a, b int) int {
	var abPrimes []int
	var answer int

	aPrimes := Factor(primes, a)
	bPrimes := Factor(primes, b)

	for _, v := range aPrimes {
		for _, v2 := range bPrimes {
			if v == v2 {
				abPrimes = append(abPrimes, v)
			}
		}
	}

	if len(abPrimes) > 1 {
		allPrimes := make(map[int]bool)
		primesList := []int{}

		for _, prime := range abPrimes {
			if _, v := allPrimes[prime]; !v {
				allPrimes[prime] = true
				primesList = append(primesList, prime)
			}
		}

		if len(primesList) == 1 {
			return primesList[0]
		}

		sum := 1
		for i := 0; i < len(primesList); i++ {
			sum *= primesList[i]
		}
		fmt.Println(sum)
		answer = sum
	}

	if len(abPrimes) == 1 {
		answer = abPrimes[0]
	}

	if abPrimes == nil {
		answer = 1
	}
	return answer
}
*/

// Using a proper alogrithm, my solution:
// Had to read about Euclidean algorithm to figure it out.
/*
func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
*/

// Done with recursion:
func GCD(a, b int) int {
	if b == 0 {
		return a
	}
	return GCD(b, a%b)
}

/*
// Same thing:
func GCD(a, b int) int {
	for b != 0 {
		return GCD(b, a%b)
	}
	return a
}
*/