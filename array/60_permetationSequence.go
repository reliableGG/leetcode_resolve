package main

import "fmt"

func getNthEle(n int, used []bool) int {
	if n < 0 {
		return nil
	}

	i := -1
	for n > 0 {
		i++
		if nums[i] == false {
			n--
		}
	}
	return i
}

func factorial(n int) []int {
	if n < 0 {
		return nil
	}

	res := make([]int)
	res[0] = 1
	for i := 1; i <= n; i++ {
		res[i] = res[i-1] * i
	}
	return res
}

func main() {
}
