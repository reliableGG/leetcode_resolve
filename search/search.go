package main

import "fmt"

func main() {
	a := []int{3, 4, 5, 1, 2}
	res := search(a, 2)
	fmt.Println(res)
}

func search(nums []int, target int) int {
	first := 0
	last := len(nums)

	for first != last {
		mid := first + (last-first)/2
		if nums[mid] == target {
			return mid
		}

		if nums[first] <= nums[mid] {
			if nums[first] <= target && target < nums[mid] {
				last = mid
			} else {
				first = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[last-1] {
				first = mid + 1
			} else {
				last = mid
			}
		}
	}
	return -1
}
