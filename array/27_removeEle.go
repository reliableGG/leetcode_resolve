package main

import "fmt"

func removeElement(nums []int, val int) int {
	res := len(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			nums[i:] = nums[i+1:]
			res--
		}
	}
	return res
}

func main() {
	fmt.Println("vim-go")
}
