package main

import "fmt"

var translation = map[int]string{
	1:    "one",
	2:    "two",
	3:    "three",
	4:    "four",
	5:    "five",
	6:    "six",
	7:    "seven",
	8:    "eight",
	9:    "nine",
	10:   "ten",
	11:   "eleven",
	12:   "twelve",
	13:   "thirteen",
	14:   "fourteen",
	15:   "fifteen",
	16:   "sixteen",
	17:   "seventeen",
	18:   "eighteen",
	19:   "nineteen",
	20:   "twenty",
	30:   "thirty",
	40:   "forty",
	50:   "fifty",
	60:   "sixty",
	70:   "seventy",
	80:   "eighty",
	90:   "ninety",
	100:  "onehundred",
	1000: "onethousand",
}

func main() {
	// Useful for debugging.
	var strs []string

	length := 0

	for i := 1; i <= 1000; i++ {
		n := i
		val := ""

		if n >= 100 && n < 1000 {
			hundred := n / 100
			val += translation[hundred] + "hundred"

			n -= hundred * 100

			if n > 0 {
				val += "and"
			}
		}

		if n <= 20 || n == 1000 {
			val += translation[n]
		} else if n > 0 && n < 100 {
			tenMul := n % 10
			val += translation[n-tenMul]

			if tenMul > 0 {
				val += translation[tenMul]
			}
		}

		length += len(val)
		strs = append(strs, val)
	}

	fmt.Println(length)
}
