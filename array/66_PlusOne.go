package main

import "fmt"

func plusOne(digits []int) []int {
	digits = append([]int{0}, digits)
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i] += 1
			break
		} else {
			digits[i] = 0
		}
	}
	if digits[0] == 0 {
		return digits[1:]
	}
	return digits
}
