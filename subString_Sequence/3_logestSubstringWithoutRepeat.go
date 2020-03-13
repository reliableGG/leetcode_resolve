package main

import "fmt"

func main() {
	//str := "abcdeafkgkcc"
	str := "tmmzuxt"
	for i, v := range str {
		fmt.Println("i", i, "v", string(v), v)
	}
	res := lengthOfLongestSubstring(str)
	fmt.Println(res)
}

func lengthOfLongestSubstring(s string) int {
	m := map[rune]int{}

	start := 0
	max := 0

	for idx, c := range s {
		_, exist := m[c]
		if exist && start <= m[rune(s[idx])] {
			start = m[rune(s[idx])] + 1
		} else {
			max = Max(max, idx-start+1)
		}
		m[c] = idx
	}
	return max

}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
