package main

import "fmt"

func main() {
	a := []int{3, 4, 5, 1, 2}
	res := search(a, 1)
	fmt.Println(res)
}

func search(nums []int, target int) int {
	left := 0
	right := len(nums)

	for left < right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}

		if nums[left] < nums[mid] {
			if nums[left] <= target && target < nums[mid] {
				right = mid
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[right-1] {
				left = mid
			} else {
				right = mid - 1
			}

		}
	}
	return -1
}
