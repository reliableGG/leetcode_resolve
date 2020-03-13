package main

import "fmt"

func main() {
	ca := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(ca, 3)
	fmt.Println(ca)
	k := 30
	k %= 10
	k = k % 10
	fmt.Println(k)
}

// first: open O(K) tmp space
func rotate(nums []int, k int) {
	if len(nums) == 0 {
		return
	}
	length := len(nums)
	k %= length

	tmp := make([]int, k)
	for i, v := range nums[len(nums)-k:] {
		tmp[i] = v
	}
	for i := length - 1; i >= k; i-- {
		nums[i] = nums[i-k]
	}
	for i, v := range tmp {
		nums[i] = v
	}
}

// second: reverse O(n) time complex O(1) space

func rotate(nums []int, k int) {
	if len(nums) == 0 || k > len(nums) {
		return
	}
	k %= len(nums)
	reverse(nums,

}

func reverse(nums []int, i, j  int) {
	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[len(nums)-i-1] = nums[len(nums)-i-1], nums[i]
	}
}
