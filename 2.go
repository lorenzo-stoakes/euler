package main

import "fmt"

func main() {
	sum := 0

	prev := 1
	curr := 1

	for curr <= 4e6 {
		prev, curr = curr, prev+curr

		if curr%2 == 0 {
			sum += curr
		}
	}

	fmt.Println(sum)
}
