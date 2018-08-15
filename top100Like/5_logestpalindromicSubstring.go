package main

import "fmt"

func logestPalindrome(s string) string {
	if s == nil {
		return nil
	}

	max := s[0]
	for i := 0; i < len(s); i++ {
		max = checkPalindrome(s, i, i, max)
		max = checkPalindrome(s, i, i+1, max)
	}
	return max
}

func checkPalindrome(s string, i int, j int, max string) string {
	var sub string
	for i >= 0 && j <= len(s) && s[i] == s[j] {
		sub = s[i : j+1]
		i--
		j++
	}
	if len(max) < len(sub) {
		return sub
	}
}

func main() {
	fmt.Println("vim-go")
}
