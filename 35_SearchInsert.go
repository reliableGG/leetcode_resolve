package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	//ca := []int{1, 2, 3, 4, 6}
	//ca := []int{1, 3, 5, 6}
	ca := []int{1, 3}
	res := searchInsert(ca, 4)
	fmt.Println(res)
}

func searchInsert(nums []int, target int) int {
	var res int
	if len(nums) == 0 || target < nums[0] {
		return 0
	}
	if nums[0] == target {
		return 1
	}
	if nums[len(nums)-1] < target {
		return len(nums) + 1
	}
	for i := 1; i < len(nums); i++ {
		if nums[i-1] < target && target <= nums[i] {
			return i
		}

	}
	return res
}
