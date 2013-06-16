package main

import "fmt"

func main() {
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19}

	product := 1
	for _, prime := range primes {
		product *= prime
	}

	for i := 2; i <= 20; i++ {
		n := i
		for _, prime := range primes {
			for n%prime == 0 && product%i != 0 {
				product *= prime
				n /= prime
			}
		}
	}

	fmt.Println(product)
}
