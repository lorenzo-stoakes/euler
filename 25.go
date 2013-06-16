package main

import (
	"fmt"
	"math/big"
)

var (
	// If we memoise Fibonacci suddenly it's O(n) and *not* the poster child of poor recursive
	// algorithms! ZOMG!!!!111
	memo = make(map[int]*big.Int)

	one = big.NewInt(1)
)

func fib(n int) *big.Int {
	if n <= 2 {
		return one
	}

	if memory, has := memo[n]; has {
		return memory
	}

	ret := &big.Int{}
	ret.Add(fib(n-1), fib(n-2))

	memo[n] = ret
	return ret
}

func main() {
	for i := 1; ; i++ {
		n := fib(i)
		if len(n.String()) == 1000 {
			fmt.Println(i)
			return
		}
	}
}
