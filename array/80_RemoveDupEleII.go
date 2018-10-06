package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	index := 2
	for i := 2; i < len(nums); i++ {
		if nums[i] != nums[index-2] {
			nums[index] = nums[i]
			index++
		}
	}
	fmt.Println(nums)
	return index
}

func main() {
	nums := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	fmt.Println(removeDuplicates(nums))
}
