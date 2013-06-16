package main

import (
	"fmt"
	"math/big"
)

const (
	maxLength = 1e4
)

func cycleLen(n *big.Rat) int {
	str := n.FloatString(maxLength)[2:]

	length := len(str)

outer:
	for i := 1; i < length/2; i++ {
		cycle := str[:i]

		for j := i; j < length/2; j += i {
			if str[j:j+i] != cycle {
				continue outer
			}
		}
		return i
	}

	return -1
}

func main() {
	max := 0
	best := int64(0)

	for d := int64(2); d < 1000; d++ {
		frac := big.NewRat(1, d)
		cycles := cycleLen(frac)

		if cycles > max {
			max = cycles
			best = d
		}
	}

	fmt.Println(best, max)
}
