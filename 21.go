package main

import "fmt"

var memo = make(map[int]int)

func sumProperDivisors(n int) int {
	if memory, has := memo[n]; has {
		return memory
	}

	ret := 0
	for i := 1; i < n; i++ {
		if n%i == 0 {
			ret += i
		}
	}

	memo[n] = ret

	return ret
}

func main() {
	amicable := make(map[int]bool)

	for i := 2; i < 10000; i++ {
		sum := sumProperDivisors(i)
		if sumProperDivisors(sum) == i && sum != i {
			amicable[i] = true
			amicable[sum] = true
		}
	}

	sum := 0
	for n, _ := range amicable {
		sum += n
	}
	fmt.Println(sum)
}
