package main

import (
	"fmt"
)

func main() {
	//ca := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	//ca := []int{1, 1}
	ca := []int{2, 3, 4, 5, 18, 17, 6}
	ans := maxArea(ca)
	fmt.Println(ans)
}
func maxArea(height []int) int {
	if len(height) < 2 {
		return 0
	}
	max := 0
	i := 0
	j := len(height) - 1
	for i < j {
		max = Max(max, (j-i)*Min(height[i], height[j]))
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return max
}

func Max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func Min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
