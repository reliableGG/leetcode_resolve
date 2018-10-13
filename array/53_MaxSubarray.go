package main

import "fmt"

func maxSubarray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	sum := nums[0]
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if sum < 0 {
			sum = nums[i]
		} else {
			sum = sum + nums[i]
		}
		if sum > max {
			max = sum
		}
	}
	return max
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubarray(nums))
}
