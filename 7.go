package main

import (
	"fmt"
	"math"
)

const N = 10001

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

func main() {
	fmt.Println(primes(N)[N-1])
}
