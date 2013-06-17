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

var (
	maxPrimeBound = 0
	primes        []int
	primeSet      = make(map[int]bool)
)

func isPrime(n int) bool {
	if (n > 2 && n%2 == 0) || (n > 3 && n%3 == 0) {
		return false
	}

	// At least one prime factor has to be <= sqrt(N).
	maxFactor := int(math.Sqrt(float64(n)))

	if maxFactor > maxPrimeBound {
		maxPrimeBound = maxFactor

		primes = sieve(maxFactor)

		for _, prime := range primes {
			primeSet[prime] = true
		}
	}

	// Since we are generating a set which might well comprise primes greater than the input
	// value for lower inputs than initial, we need to check to ensure that n % prime == 0
	// because n == prime. (Probably) quicker to use a hash then to check for n == prime below.
	if primeSet[n] {
		return true
	}

	for _, prime := range primes {
		if n%prime == 0 {
			return false
		}
	}

	return true
}

func main() {
	best := 0
	product := 0

	// Run backwards so we populate our cached primes early.
	for a := 999; a > -1000; a-- {
		for b := 999; b > -1000; b-- {
			n := 0

			for ; isPrime(n*n + a*n + b); n++ {
			}

			if n > best {
				best = n
				product = a * b
			}
		}
	}

	fmt.Println(best, product)
}
