package main

import "fmt"

// Taken from 21.go, memoisation stripped.
func sumProperDivisors(n int) int {
	ret := 0
	for i := 1; i < n; i++ {
		if n%i == 0 {
			ret += i
		}
	}

	return ret
}

func main() {
	abundant := make(map[int]bool)
	abundantSums := make(map[int]bool)

	sum := 0

	for i := 1; i <= 28123; i++ {
		if sumProperDivisors(i) > i {
			abundant[i] = true
			for n, _ := range abundant {
				abundantSums[i+n] = true
			}
		}

		if !abundantSums[i] {
			sum += i
		}
	}

	fmt.Println(sum)
}
