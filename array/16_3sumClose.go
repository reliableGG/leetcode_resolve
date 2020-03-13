package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	ca := []int{-1, 2, 1, -4}
	ans := threeSumClosest(ca, 1)
	fmt.Println(ans)
}

func threeSumClosest(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)
	min := nums[0] + nums[1] + nums[len(nums)-1]

	for i := 0; i < len(nums)-2; i++ {
		j := i + 1
		k := len(nums) - 1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum > target {
				k--
			} else {
				j++
			}

			if math.Abs(float64(sum-target)) < math.Abs(float64(min-target)) {
				min = sum
			}
		}
	}

	return min
}
