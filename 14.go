package main

import "fmt"

var memo = make(map[int]int)

func countChain(n int) int {
	ret := 1

	if n == 1 {
		return ret
	} else if memory, has := memo[n]; has {
		return memory + ret
	}

	if n%2 == 0 {
		ret += countChain(n / 2)
	} else {
		ret += countChain(3*n + 1)
	}

	memo[n] = ret
	return ret
}

func main() {
	max := 0
	starting := 0

	for i := 2; i < 1e6; i++ {
		count := countChain(i)

		if count > max {
			starting = i
			max = count
		}
	}

	fmt.Println(starting, max)
}
