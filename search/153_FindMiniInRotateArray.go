package main

import "fmt"

func main() {
	cas := []int{4, 5, 6, 7, 0, 1, 2}

	fmt.Println(findMin(cas))
}

func findMin(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	left, right := 0, len(nums)-1

	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < nums[right] {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return nums[right]
}
