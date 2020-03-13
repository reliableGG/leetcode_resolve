package main

import "fmt"

func main() {
	s := "abxxxdxxx"
	res := longestPalindrome(s)
	fmt.Println(res)
}

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	max := string(s[0])
	for i := 0; i < len(s); i++ {
		max = checkPalindromic(s, i, i, max)
		max = checkPalindromic(s, i, i+1, max)
	}
	return max

}

func checkPalindromic(s string, i, j int, max string) string {
	var sub string
	for i >= 0 && j < len(s)-1 && s[i] == s[j] {
		sub = s[i : j+1]
		i--
		j++
	}
	if len(max) < len(sub) {
		return sub
	}
	return max
}
