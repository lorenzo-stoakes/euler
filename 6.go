package main

import "fmt"

func main() {
	// 1/2*n*(n+1) as made famous by Gauss.
	sum := 100 * 101 / 2
	sqSum := sum * sum

	sumSq := 0
	for i := 1; i <= 100; i++ {
		sumSq += i * i
	}

	fmt.Println(sqSum - sumSq)
}
