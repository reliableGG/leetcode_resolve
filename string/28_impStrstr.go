package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	s := "hello"
	need := "ll"
	id := strStr(s, need)
	fmt.Println(id)

}
func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}

	if len(haystack) < len(needle) {
		return -1
	}

	for i := 0; i < len(haystack)-len(needle); i++ {
		if string([]rune(haystack)[i:i+len(needle)]) == needle {
			return i
		}
	}
	return -1

}
