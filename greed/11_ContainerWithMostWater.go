package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	input := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	output := maxArea(input)
	fmt.Println(output)
}

func maxArea(height []int) int {
	i := 0
	j := len(height) - 1
	maxArea := 0
	for i < j {
		area := min(height[i], height[j]) * (j - i)
		maxArea = max(maxArea, area)
		if height[i] > height[j] {
			j--
		} else {
			i++
		}
	}
	return maxArea
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
