package main

import "fmt"

func trap(h []int) int {
	if len(h) == 0 {
		return 0
	}

	left, right := make([]int, len(h)), make([]int, len(h))
	max := h[0]
	for i := 1; i < len(h); i++ {
		left[i] = max
		if h[i] > max {
			max = h[i]
		}
		fmt.Println(left[i])
	}
	max = h[len(h)-1]
	for i := len(h) - 2; i >= 0; i-- {
		right[i] = max
		if h[i] > max {
			max = h[i]
		}
	}

	fmt.Println(left, right)
	var res int
	for i := 1; i < len(h)-1; i++ {
		res += Max(Min(left[i], right[i])-h[i], 0)
	}
	return res
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	//num := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	num := []int{2, 0, 2}
	fmt.Println(trap(num))
}
