package main

import "fmt"

func longestConsecutive(nums []int) int {
	count := make(map[int]int)
	max := 0

	for _, i := range nums {
		if _, exist := count[i]; !exist {
			left := count[i-1]
			right := count[i+1]
			sum := left + right + 1

			if sum > max {
				max = sum
			}
			count[i] = sum
			count[i-left] = sum
			count[i+right] = sum
		}
	}
	return max
}

func main() {
	a := []int{100, 4, 200, 1, 3, 2}
	res := longestConsecutive(a)
	fmt.Println(res)
}
