package main

import (
	"fmt"
	"sort"
)

func main() {
	//a := []int{-1, 0, 1, 2, -1, -4}
	//a := []int{0, 0, 0}
	a := []int{-2, 0, 1, 1, 2}
	ans := threeSum(a)
	fmt.Println(ans)
}

func threeSum(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)

	for i := 0; i <= len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j := i + 1
		k := len(nums) - 1
		for j < k {
			if nums[i]+nums[j]+nums[k] == 0 {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				j++
				for ; j < k && nums[j] == nums[j-1]; j++ {
				}
			} else if nums[i]+nums[j]+nums[k] > 0 {
				k--
			} else {
				j++
			}
		}
	}
	return res
}
