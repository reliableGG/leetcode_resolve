package main

import "fmt"

func removeDuplicates(nums []int) int {
	if nums == nil {
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
	fmt.Println("vim-go")
	nums := []int{1, 1, 1, 2, 3}
	fmt.Println(removeDuplicates(nums))

}
