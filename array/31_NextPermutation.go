package main

import "fmt"

func reverseSort(nums []int) {
	n := len(nums) - 1
	fmt.Println("n => ", n)
	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[n-i] = nums[n-i], nums[i]
	}
}

func nextPermutation(nums []int) {
	i := 0
	n := len(nums)
	for i = n - 1; i > 0; i-- {
		if nums[i] > nums[i-1] {
			break
		}
	}
	fmt.Println(i)
	if i == 0 {
		reverseSort(nums)
		return
	}

	//从nums[i+1:] 里找比nums[i]大的第一个数
	var j int
	for j = n - 1; j >= i; j-- {
		if nums[j] > nums[i-1] {
			break
		}
	}
	nums[i-1], nums[j] = nums[j], nums[i-1]
	reverseSort(nums[i:])
}

func main() {
	//nums := []int{1, 2, 3}
	nums := []int{7, 2, 5, 4, 3, 1}
	nextPermutation(nums)
	fmt.Println(nums)
}
