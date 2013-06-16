package main

import "fmt"

func main() {
	// Presumably it'd be bad form to just use a library :-) so roll our own.

	dayCount := make([]int, 13)
	for i := range dayCount {
		switch i {
		// September, April, June and November.
		case 9, 4, 6, 11:
			dayCount[i] = 30
		case 2:
			// Deal with leap years separately.
			dayCount[i] = 28
		default:
			dayCount[i] = 31
		}
	}

	getDayCount := func(year, month int) int {
		ret := dayCount[month]

		// Leap years.
		if month == 2 && year%4 == 0 && (year%100 > 0 || year%400 == 0) {
			ret++
		}

		return ret
	}

	weekDay := 1 // Monday.
	count := 0

	for year := 1900; year <= 2000; year++ {
		for month := 1; month <= 12; month++ {
			if year > 1900 && weekDay == 0 {
				count++
			}

			weekDay += getDayCount(year, month)
			weekDay %= 7
		}
	}

	fmt.Println(count)
}
