package main

import "fmt"

const size = 1001

type direction int

const (
	right direction = iota
	down
	left
	up
)

var grid [size][size]int

func main() {
	// Start in centre of grid.
	x := size / 2
	y := x

	dir := right
	step := 1
	n := 1

	for x < size {
		var xDelta, yDelta int

		switch dir {
		case right:
			xDelta, yDelta = 1, 0
		case down:
			xDelta, yDelta = 0, 1
		case left:
			xDelta, yDelta = -1, 0
		case up:
			xDelta, yDelta = 0, -1
		}

		for i := 0; i < step; i++ {
			grid[y][x] = n

			x += xDelta
			y += yDelta
			n++
		}

		// Rotate through directions.
		dir = (dir + 1) % 4

		// Every 2 directions increase step by 2.
		if dir%2 == 0 {
			step++
		}
	}

	sum := 0
	for i, row := range grid {
		sum += row[i]
		if i != size/2 {
			sum += row[size-i-1]
		}
	}

	fmt.Println(sum)
}
