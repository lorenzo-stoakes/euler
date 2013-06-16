package main

import (
	"fmt"
	"math/big"
)

func main() {
	n := big.NewInt(1)

	for i := 2; i <= 100; i++ {
		n.Mul(n, big.NewInt(int64(i)))
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
