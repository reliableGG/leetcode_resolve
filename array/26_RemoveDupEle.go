package main

import "fmt"

func removeDuplicates(nums []int) int {
	if nums == nil {
		return 0
	}
	index := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[index-1] {
			nums[index] = nums[i]
			index++
		}
	}
	fmt.Println(nums)
	return index
}
func main() {
	fmt.Println("vim-go")
	nums := []int{1, 2}
	fmt.Println(removeDuplicates(nums))

}
