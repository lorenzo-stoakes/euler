package main

import (
	"fmt"
	"math/big"
)

func main() {
	n := big.NewInt(1)

	two := big.NewInt(2)

	for i := 0; i < 1000; i++ {
		n.Mul(n, two)
	}

	chrToInt := func(chr rune) int {
		return int(uint8(chr) - '0')
	}

	sum := 0

	for _, chr := range n.String() {
		sum += chrToInt(chr)
	}

	fmt.Println(sum)
}
