package main

import (
	"fmt"
	"math"
)

const N = 600851475143

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
	max := int64(0)

	for prime, _ := range primeFactors(N) {
		if prime > max {
			max = prime
		}
	}

	fmt.Println(max)
}
