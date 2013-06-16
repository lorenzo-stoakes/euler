package main

import (
	"fmt"
)

func main() {
	// a < b < c, a + b + c = 1000
	// Largest values 332 < 333 < 335, hence we don't have to iterate everything:-
	for a := 1; a < 333; a++ {
		for b := a + 1; a+b < 666; b++ {
			c := 1000 - a - b

			if a*a+b*b == c*c {
				fmt.Println(a * b * c)
				return
			}
		}
	}
}
