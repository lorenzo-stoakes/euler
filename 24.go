package main

import "fmt"

func permute(str string) []string {
	if len(str) == 1 {
		return []string{str}
	}

	var ret []string
	for i, chr := range str {
		// Since we always maintain order as input, if input is initially in lexicographical
		// order, output will maintain it.
		rest := str[:i] + str[i+1:]
		for _, perm := range permute(rest) {
			ret = append(ret, string(chr)+perm)
		}
	}
	return ret
}

func main() {
	fmt.Println(permute("0123456789")[1e6-1])
}
