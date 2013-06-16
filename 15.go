package main

import "fmt"

// Sloooow brute forcing. Roughly 10 minutes on my system.

const GRID_SIZE = 20

func countRoutes(x, y int) int {
	ret := 1

	for ; x < GRID_SIZE && y < GRID_SIZE; x++ {
		ret += countRoutes(x, y+1)
	}

	return ret
}

func main() {
	fmt.Println(countRoutes(0, 0))
}
