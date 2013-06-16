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

// Taken from 7.go.
// Get n primes.
func primes(n int) []int {
	// The nth prime is roughly n log n.
	estimate := n * int(math.Log(float64(n)))

	var ret []int
	// Double estimate until we have the primes we need.
	for len(ret) < n {
		ret = sieve(estimate)
		estimate *= 2
	}

	return ret[:n]
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

func countDivisors(n int64) int {
	ret := 1

	for _, count := range primeFactors(n) {
		ret *= count + 1
	}

	return ret
}

func main() {
	// The number of divisors of a number is equal to the product of the powers of each
	// prime divisor + 1, i.e. \sigma_0 = \Pi_{i=1}^r a^i + 1.
	// See https://en.wikipedia.org/wiki/Divisor_function#Properties

	// So if we're looking for an upper bound on the minimum possible number with 500 divisors,
	// we set a_i = 1 regardless of i and have:-

	// 500 = \Pi_{i=1}^r 2 = 2^r

	// Hence r = lg(500) ~ 9, sets an upper bound of 223,092,870 (2*3*5*...*23).

	// After some fiddling around in a spreadsheet + then using a bruteforce to confirm found
	// the lowest possible number with 500 or more divisors is... 14,414,400 (2^6*3^2*5^2*7*11*13)

	// Now we have our lower bound we can solve 1/2n*(n+1) = 14,415,765 to get n = 5368.75, so
	// closest lowerbound triangle number is 1/2*5368*5369 = 14,410,396 and start iterating
	// from there :) :-

	n := int64(5368)
	for {
		triangle := n * (n + int64(1)) / int64(2)

		divisors := countDivisors(triangle)
		if divisors > 500 {
			fmt.Println(triangle, divisors)
			return
		}
		n++
	}
}
