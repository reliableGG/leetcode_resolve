package main

import "fmt"

func main() {
	//cases := []int{1, 2, 3, 4, 5}
	//cases := []int{1, 2, 3}
	cases := []int{1, 2, 3, 1, 2}
	nextPermutation(cases)
	fmt.Println(cases)
}

func nextPermutation(nums []int) {
	n := len(nums)
	i := 0
	for i = n - 1; i > 0; i-- {
		if nums[i] > nums[i-1] {
			break
		}
	}
	fmt.Println(i)
	if i == 0 {
		revertList(nums)
	} else {
		for j := n - 1; j > i-1; j-- {
			if nums[j] > nums[i-1] {
				nums[j], nums[i-1] = nums[i-1], nums[j]
				break
			}
		}
		revertList(nums[i:])
	}
}

func revertList(nums []int) {
	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[len(nums)-i-1] = nums[len(nums)-i-1], nums[i]
	}
}
