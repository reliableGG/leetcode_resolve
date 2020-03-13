package main

import "fmt"

func main() {
	ca := []int{1, 2, 3, 4}
	res := productExceptSelf(ca)
	fmt.Println(res)
}

func productExceptSelf(nums []int) []int {
	left := map[int]int{}
	right := map[int]int{}
	res := make([]int, len(nums))

	left[0] = 1
	right[len(nums)-1] = 1

	for i := 1; i < len(nums); i++ {
		left[i] = left[i-1] * nums[i-1]
	}
	fmt.Println(left)
	for i := len(nums) - 2; i >= 0; i-- {
		right[i] = right[i+1] * nums[i+1]
	}
	fmt.Println(right)

	for i := 0; i < len(nums); i++ {
		res[i] = left[i] * right[i]
	}
	return res

}
