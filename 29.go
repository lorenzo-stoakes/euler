package main

import (
	"fmt"
	"math"
)

// Taken from 3.go.
// See https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes.
func sieve(n int) []int {
	// Populate 2 to n. +1 size so indexes coincide with values for readabililty.
	tmp := make([]int, n+1)
	for i := 2; i <= n; i++ {
		tmp[i] = i
	}

	curr := 2
	for curr < n {
		// 0 represents a 'crossed out' value.
		for i := 2 * curr; i <= n; i += curr {
			tmp[i] = 0
		}

		// Skip to next prime number.
		for curr = curr + 1; curr < n && tmp[curr] == 0; curr++ {
		}
	}

	// Now filter our crossed out set.
	var ret []int
	for _, n := range tmp {
		if n > 0 {
			ret = append(ret, n)
		}
	}

	return ret
}

// Taken from 3.go.
// Obtain all prime factors of number in a map.
func primeFactors(n int64) map[int64]int {
	ret := make(map[int64]int)

	add := func(prime int64) {
		if _, has := ret[prime]; has {
			ret[prime]++
		} else {
			ret[prime] = 1
		}

		n /= int64(prime)
	}

	for n > 1 {
		// At least one prime factor has to be <= sqrt(N).
		maxFactor := int(math.Sqrt(float64(n)))

		isPrime := true
		for _, prime := range sieve(maxFactor) {
			for n%int64(prime) == 0 {
				isPrime = false
				add(int64(prime))
			}
		}

		if isPrime {
			add(n)
		}
	}

	return ret
}

func main() {
	// Obtain the prime factors for each number 2 - 100, then simply multiply the powers by the
	// power we want to raise to.

	var factors []map[int64]int
	for a := 2; a <= 100; a++ {
		orig := primeFactors(int64(a))

		for b := 2; b <= 100; b++ {
			set := make(map[int64]int)

			for key, val := range orig {
				set[key] = val * b
			}

			factors = append(factors, set)
		}
	}

	// Now simply store a hash of the string representation of the sets to eliminate duplicates
	// and count :-)

	seenStr := make(map[string]bool)
	for _, factor := range factors {

		str := fmt.Sprintf("%v", factor)
		seenStr[str] = true
	}

	fmt.Println(len(seenStr))
}
