package main

import (
	"fmt"
)

func isSubsequence(s string, t string) bool {

	sIdx := 0
	for _, c := range t {
		if sIdx == len(s) {
			return true
		}
		if c == rune(s[sIdx]) {
			sIdx++
		}
	}

	if sIdx == len(t) {
		return true
	}
	return false
}

func main() {

	t := "abccc"
	s := "bc"
	fmt.Println(byte(t[0]))
	x := t[2:4]
	fmt.Println(x)

	fmt.Println(isSubsequence(s, t))
}
