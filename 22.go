package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func score(str string) int {
	ret := 0

	for _, chr := range str {
		ret += int(uint8(chr)-'A') + 1
	}

	return ret
}

func main() {
	var str string

	if bytes, err := ioutil.ReadFile("data/names.txt"); err != nil {
		panic(err)
	} else {
		str = string(bytes)
	}

	names := strings.Split(str, ",")
	sort.Strings(names)

	sum := 0
	for i, name := range names {
		// Strip double quotes.
		name = name[1 : len(name)-1]

		sum += score(name) * (i + 1)
	}
	fmt.Println(sum)
}
