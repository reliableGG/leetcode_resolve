package main

import (
	"fmt"
)

func maxArea(height []int) int {
	water := 0
	length := len(height)
	i := 0
	j := length - 1
	for i <= length-1 && j >= 0 && j >= i {
		water = Max(water, (j-i)*Min(height[i], height[j]))
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return water
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	in := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	res := maxArea(in)
	fmt.Println(res)
}
