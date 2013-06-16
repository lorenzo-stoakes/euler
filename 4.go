package main

import (
	"fmt"
	"strconv"
)

func isPalin(n int) bool {
	str := strconv.Itoa(n)
	length := len(str)

	for i := 0; i < length/2; i++ {
		if str[i] != str[length-1-i] {
			return false
		}
	}

	return true
}

func main() {
	max := 0

	for i := 100; i <= 999; i++ {
		for j := i; j <= 999; j++ {
			product := i * j
			if isPalin(product) && product > max {
				max = product
			}
		}
	}

	fmt.Println(max)
}
