package main

import "fmt"

var rows = [][]int{
	[]int{75},
	[]int{95, 64},
	[]int{17, 47, 82},
	[]int{18, 35, 87, 10},
	[]int{20, 4, 82, 47, 65},
	[]int{19, 1, 23, 75, 3, 34},
	[]int{88, 2, 77, 73, 7, 63, 67},
	[]int{99, 65, 4, 28, 6, 16, 70, 92},
	[]int{41, 41, 26, 56, 83, 40, 80, 70, 33},
	[]int{41, 48, 72, 33, 47, 32, 37, 16, 94, 29},
	[]int{53, 71, 44, 65, 25, 43, 91, 52, 97, 51, 14},
	[]int{70, 11, 33, 28, 77, 73, 17, 78, 39, 68, 17, 57},
	[]int{91, 71, 52, 38, 17, 14, 91, 43, 58, 50, 27, 29, 48},
	[]int{63, 66, 4, 68, 89, 53, 67, 30, 73, 16, 69, 87, 40, 31},
	[]int{4, 62, 98, 27, 23, 9, 70, 98, 73, 93, 38, 53, 60, 4, 23},
}

// Taken from 11.go.
// Math library has max for floats. Avoid irritating typecasts by rolling our own.
func max(m int, ns ...int) int {
	for _, n := range ns {
		if n > m {
			m = n
		}
	}

	return m
}

// Just brute force it for now.
func walk(rowIndex, colIndex int) int {
	if rowIndex >= len(rows) {
		return 0
	}

	val := rows[rowIndex][colIndex]
	return val + max(walk(rowIndex+1, colIndex), walk(rowIndex+1, colIndex+1))
}

func main() {
	fmt.Println(walk(0, 0))
}
